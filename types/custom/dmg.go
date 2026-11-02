package custom

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/coalaura/wtf/types"
)

func DetectAppleDiskImage(b types.Buffer) *types.Metadata {
	return detectAppleDiskImageSized(b, bytes.NewReader(b), int64(b.Len()))
}

func detectAppleDiskImageSized(b types.Buffer, reader io.ReaderAt, size int64) *types.Metadata {
	if b.Len() < 512 {
		return nil
	}

	offset := b.Len() - 512

	version, versionOK := b.U32BE(offset + 4)
	headerSize, sizeOK := b.U32BE(offset + 8)
	segmentNumber, segmentOK := b.U32BE(offset + 56)
	segmentCount, countOK := b.U32BE(offset + 60)

	if !b.Has(offset, []byte("koly")) || !versionOK || version != 4 || !sizeOK || headerSize != 512 || !segmentOK || !countOK || segmentNumber > segmentCount {
		return nil
	}

	if segmentNumber > 0 && segmentCount > 0 {
		return &types.Metadata{Kind: types.KindAppleDiskImage}
	}

	if segmentNumber != 0 || segmentCount != 0 || size < 512 {
		return nil
	}

	dataOffset := binary.BigEndian.Uint64(b[offset+24:])
	dataLength := binary.BigEndian.Uint64(b[offset+32:])
	xmlOffset := binary.BigEndian.Uint64(b[offset+216:])
	xmlLength := binary.BigEndian.Uint64(b[offset+224:])
	sectorCount := binary.BigEndian.Uint64(b[offset+492:])
	limit := uint64(size - 512)

	if dataOffset > limit || dataLength == 0 || dataLength > limit-dataOffset || xmlOffset < dataOffset+dataLength || xmlOffset > limit || xmlLength < 32 || xmlLength > limit-xmlOffset || sectorCount == 0 {
		return nil
	}

	var resource [512]byte

	metadata := resource[:min(xmlLength, uint64(len(resource)))]
	if !readAtExactly(reader, metadata, int64(xmlOffset)) || !bytes.HasPrefix(metadata, []byte("<?xml")) || !bytes.Contains(metadata, []byte("<plist")) {
		return nil
	}

	return &types.Metadata{Kind: types.KindAppleDiskImage}
}
