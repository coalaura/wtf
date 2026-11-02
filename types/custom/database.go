package custom

import (
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
	for offset := 0; offset+8 <= b.Len(); {
		if b.Has(offset, []byte{0x89, 'H', 'D', 'F', 13, 10, 26, 10}) {
			return &types.Metadata{Kind: types.KindHDF5Data}
		}

		if offset == 0 {
			offset = 512
		} else {
			offset *= 2
		}
	}

	return nil
}

func DetectHDF5ReaderAt(b types.Buffer, reader io.ReaderAt, size int64) *types.Metadata {
	meta := DetectHDF5(b)
	if meta != nil {
		return meta
	}

	var signature [8]byte

	for offset := int64(512); offset <= size-8; {
		if offset >= int64(b.Len()) && readAtExactly(reader, signature[:], offset) && bytesEqualHDF5(signature[:]) {
			return &types.Metadata{Kind: types.KindHDF5Data}
		}

		if offset > (size-8)/2 {
			break
		}

		offset *= 2
	}

	return nil
}

func bytesEqualHDF5(signature []byte) bool {
	return string(signature) == "\x89HDF\r\n\x1a\n"
}

func berkeleyMagic(magic uint32) bool {
	return magic == 0x00053162 || magic == 0x00061561 || magic == 0x00042253 || magic == 0x00040988
}

func validDatabasePageSize(size uint32) bool {
	return size >= 512 && size <= 65536 && size&(size-1) == 0
}
