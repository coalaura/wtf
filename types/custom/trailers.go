package custom

import "github.com/coalaura/wtf/types"

func DetectFileTrailer(buffer types.Buffer) *types.Metadata {
	meta := DetectAppleDiskImage(buffer)
	if meta != nil {
		return meta
	}

	if buffer.Len() >= 48 && buffer.Has(buffer.Len()-8, []byte{0x57, 0xfb, 0x80, 0x8b, 0x24, 0x75, 0x47, 0xdb}) {
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
