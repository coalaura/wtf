package custom

import (
	"encoding/binary"

	"github.com/coalaura/wtf/types"
)

func DetectDirectX(buffer types.Buffer) *types.Metadata {
	if buffer.Len() < 32 || !buffer.Has(0, []byte("DXBC")) || binary.LittleEndian.Uint16(buffer[20:]) != 1 {
		return nil
	}

	size := uint64(binary.LittleEndian.Uint32(buffer[24:]))
	count := uint64(binary.LittleEndian.Uint32(buffer[28:]))
	tableEnd := 32 + count*4

	if count > 256 || tableEnd > uint64(buffer.Len()) || tableEnd > size {
		return nil
	}

	shader := types.TypeNone
	complete := true
	next := tableEnd
	previous := tableEnd - 8

	for index := range count {
		offset := uint64(binary.LittleEndian.Uint32(buffer[32+index*4:]))
		if offset < tableEnd || offset > size || size-offset < 8 {
			return nil
		}

		if offset < previous+8 || complete && offset != next {
			return nil
		}

		previous = offset

		if offset+8 > uint64(buffer.Len()) {
			complete = false

			continue
		}

		length := uint64(binary.LittleEndian.Uint32(buffer[offset+4:]))
		if length > size-offset-8 {
			return nil
		}

		next = offset + 8 + length

		part := string(buffer[offset : offset+4])
		if part == "DXIL" {
			shader = types.TypeDXIL
		} else if (part == "SHDR" || part == "SHEX") && shader == types.TypeNone {
			shader = types.TypeDXBC
		}
	}

	if !complete {
		shader = types.TypeNone
	} else if next != size {
		return nil
	}

	return &types.Metadata{Kind: types.KindDirectXShaderContainer, Type: shader}
}
