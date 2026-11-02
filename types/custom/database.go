package custom

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/coalaura/wtf/types"
)

func DetectDatabaseHeaders(b types.Buffer) *types.Metadata {
	if b.Len() >= 28 {
		magic, _ := b.U32LE(12)
		version, _ := b.U32LE(16)
		pageSize, _ := b.U32LE(20)

		if berkeleyMagic(magic) && version > 0 && version <= 20 && validDatabasePageSize(pageSize) {
			return &types.Metadata{Kind: types.KindBerkeleyDatabase, Type: types.TypeLittleEndian}
		}

		magic, _ = b.U32BE(12)
		version, _ = b.U32BE(16)
		pageSize, _ = b.U32BE(20)

		if berkeleyMagic(magic) && version > 0 && version <= 20 && validDatabasePageSize(pageSize) {
			return &types.Metadata{Kind: types.KindBerkeleyDatabase, Type: types.TypeBigEndian}
		}

		magic, _ = b.U32LE(16)
		version, _ = b.U32LE(20)
		pageSize, _ = b.U32LE(24)
		flags, _ := b.U16LE(8)

		if magic == 0xed0cdaed && version == 2 && flags == 4 && validDatabasePageSize(pageSize) {
			return &types.Metadata{Kind: types.KindBoltDB}
		}
	}

	if b.Has(0, []byte{0x3f, 0xd7, 0x6c, 0x17}) && b.Len() >= 10 {
		nameLength := int(b[4])
		if nameLength > 0 && nameLength < 128 && nameLength <= b.Len()-9 {
			for _, character := range b[5 : 5+nameLength] {
				if character < 0x20 || character > 0x7e {
					return nil
				}
			}

			return &types.Metadata{Kind: types.KindLuceneIndex}
		}
	}

	return nil
}

func DetectHDF5(b types.Buffer) *types.Metadata {
	return DetectHDF5ReaderAt(b, bytes.NewReader(b), int64(b.Len()))
}

func DetectHDF5ReaderAt(b types.Buffer, reader io.ReaderAt, size int64) *types.Metadata {
	var header [96]byte

	for offset := int64(0); offset <= size-48; {
		length := min(int64(len(header)), size-offset)

		var superblock []byte

		if offset+length <= int64(b.Len()) {
			superblock = b[offset : offset+length]
		} else if readAtExactly(reader, header[:length], offset) {
			superblock = header[:length]
		}

		if validHDF5Superblock(superblock, uint64(offset), uint64(size)) {
			return &types.Metadata{Kind: types.KindHDF5Data}
		}

		if offset > (size-48)/2 {
			break
		}

		if offset == 0 {
			offset = 512
		} else {
			offset *= 2
		}
	}

	return nil
}

func validHDF5Superblock(header []byte, offset uint64, size uint64) bool {
	if len(header) < 48 || string(header[:8]) != "\x89HDF\r\n\x1a\n" {
		return false
	}

	version := header[8]
	addressStart := 12
	width := int(header[9])
	lengthWidth := int(header[10])
	rootStart := 0

	switch version {
	case 0, 1:
		width = int(header[13])
		lengthWidth = int(header[14])
		addressStart = 24

		if header[9] != 0 || header[10] != 0 || header[11] != 0 || header[12] != 0 || header[15] != 0 || binary.LittleEndian.Uint16(header[16:]) == 0 || binary.LittleEndian.Uint16(header[18:]) == 0 || binary.LittleEndian.Uint32(header[20:]) & ^uint32(3) != 0 {
			return false
		}

		if version == 1 {
			addressStart = 28

			if binary.LittleEndian.Uint16(header[24:]) == 0 || header[26] != 0 || header[27] != 0 {
				return false
			}
		}

		rootStart = addressStart + 4*width + lengthWidth
	case 2, 3:
		if header[11] & ^byte(5) != 0 {
			return false
		}

		rootStart = addressStart + 3*width
	default:
		return false
	}

	if !validHDF5Width(width) || !validHDF5Width(lengthWidth) || rootStart+width > len(header) {
		return false
	}

	base := hdf5Address(header[addressStart : addressStart+width])
	endStart := addressStart + 2*width
	end := hdf5Address(header[endStart : endStart+width])
	root := hdf5Address(header[rootStart : rootStart+width])

	return base == offset && end > offset+uint64(rootStart+width) && end <= size && root > 0 && root < end-base
}

func validHDF5Width(width int) bool {
	return width == 1 || width == 2 || width == 4 || width == 8
}

func hdf5Address(data []byte) uint64 {
	var value uint64

	for index, part := range data {
		value |= uint64(part) << (index * 8)
	}

	return value
}

func berkeleyMagic(magic uint32) bool {
	return magic == 0x00053162 || magic == 0x00061561 || magic == 0x00042253 || magic == 0x00040988
}

func validDatabasePageSize(size uint32) bool {
	return size >= 512 && size <= 65536 && size&(size-1) == 0
}
