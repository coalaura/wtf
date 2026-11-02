package custom

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/coalaura/wtf/types"
)

func DetectFileTrailer(buffer types.Buffer) *types.Metadata {
	return detectFileTrailerSized(buffer, bytes.NewReader(buffer), int64(buffer.Len()))
}

func detectFileTrailerSized(buffer types.Buffer, reader io.ReaderAt, size int64) *types.Metadata {
	meta := detectAppleDiskImageSized(buffer, reader, size)
	if meta != nil {
		return meta
	}

	if validLevelDBFooter(buffer, size) {
		return &types.Metadata{Kind: types.KindLevelDB}
	}

	if buffer.Len() < 512 {
		return nil
	}

	foot := buffer[buffer.Len()-512:]
	version, _ := foot.U32BE(12)
	diskType, _ := foot.U32BE(60)
	checksum, _ := foot.U32BE(64)

	if !foot.Has(0, []byte("conectix")) || version != 0x00010000 || diskType < 2 || diskType > 4 {
		return nil
	}

	var sum uint32

	for index, value := range foot {
		if index < 64 || index >= 68 {
			sum += uint32(value)
		}
	}

	if ^sum != checksum {
		return nil
	}

	return &types.Metadata{Kind: types.KindVHDDiskImage}
}

func validLevelDBFooter(buffer types.Buffer, size int64) bool {
	if buffer.Len() < 48 || size < 48 || !buffer.Has(buffer.Len()-8, []byte{0x57, 0xfb, 0x80, 0x8b, 0x24, 0x75, 0x47, 0xdb}) {
		return false
	}

	footer := buffer[buffer.Len()-48 : buffer.Len()-8]
	limit := uint64(size - 48)

	var previousEnd uint64

	for range 2 {
		offset, count := binary.Uvarint(footer)
		if count <= 0 {
			return false
		}

		footer = footer[count:]

		length, count := binary.Uvarint(footer)
		if count <= 0 || offset < previousEnd || offset > limit || limit-offset < 5 || length == 0 || length > limit-offset-5 {
			return false
		}

		footer = footer[count:]
		previousEnd = offset + length + 5
	}

	for _, value := range footer {
		if value != 0 {
			return false
		}
	}

	return true
}
