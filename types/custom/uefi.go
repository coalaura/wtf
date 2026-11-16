package custom

import (
	"encoding/binary"

	"github.com/coalaura/wtf/types"
)

func DetectUEFIVolume(buffer types.Buffer) *types.Metadata {
	if buffer.Len() < 64 || !buffer.Has(40, []byte("_FVH")) || buffer[54] != 0 || buffer[55] != 2 {
		return nil
	}

	length := binary.LittleEndian.Uint64(buffer[32:])
	headerSize := int(binary.LittleEndian.Uint16(buffer[48:]))
	extension := int(binary.LittleEndian.Uint16(buffer[52:]))

	if headerSize < 64 || headerSize%8 != 0 || headerSize > buffer.Len() || uint64(headerSize) > length || extension != 0 && (extension < 64 || uint64(extension)+20 > length) {
		return nil
	}

	checksum := uint16(0)

	for offset := 0; offset < headerSize; offset += 2 {
		checksum += binary.LittleEndian.Uint16(buffer[offset:])
	}

	if checksum != 0 {
		return nil
	}

	blocks := uint64(0)

	for offset := 56; offset+8 <= headerSize; offset += 8 {
		count := uint64(binary.LittleEndian.Uint32(buffer[offset:]))
		blockSize := uint64(binary.LittleEndian.Uint32(buffer[offset+4:]))

		if count == 0 && blockSize == 0 {
			if blocks == length {
				return &types.Metadata{Kind: types.KindUEFIFirmwareVolume}
			}

			return nil
		}

		if count == 0 || blockSize == 0 || count*blockSize > length-blocks {
			return nil
		}

		blocks += count * blockSize
	}

	return nil
}
