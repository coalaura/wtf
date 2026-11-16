package custom

import (
	"bytes"
	"encoding/binary"
	"math"

	"github.com/coalaura/wtf/types"
)

func DetectScientificHeaders(buffer types.Buffer) *types.Metadata {
	if buffer.Len() >= 24 && buffer.Has(0, []byte("RRD\x00")) && buffer.Has(4, []byte("000")) && buffer[7] >= '1' && buffer[7] <= '5' && buffer[8] == 0 {
		cookie := math.Float64bits(8.642135e130)

		for offset := 12; offset <= 16; offset += 4 {
			if binary.LittleEndian.Uint64(buffer[offset:]) == cookie {
				return &types.Metadata{Kind: types.KindRRDtoolDatabase, Type: types.TypeLittleEndian}
			}

			if binary.BigEndian.Uint64(buffer[offset:]) == cookie {
				return &types.Metadata{Kind: types.KindRRDtoolDatabase, Type: types.TypeBigEndian}
			}
		}
	}

	if buffer.Len() >= 7 && buffer.Has(0, []byte("#ASDF")) && (buffer[5] == ' ' || buffer[5] == '\t') {
		end := bytes.IndexByte(buffer[:min(buffer.Len(), 64)], '\n')
		if end < 0 {
			return nil
		}

		version := bytes.TrimSpace(buffer[6:end])

		var (
			components int
			digits     int
		)

		for _, value := range version {
			if value >= '0' && value <= '9' {
				digits++
			} else if value == '.' && digits > 0 {
				components++
				digits = 0
			} else {
				return nil
			}
		}

		if components == 2 && digits > 0 {
			return &types.Metadata{Kind: types.KindASDFDocument}
		}
	}

	if buffer.Len() < 1024 || !buffer.Has(208, []byte("MAP ")) {
		return nil
	}

	var (
		order   binary.ByteOrder
		subtype types.TypeID
	)

	if buffer.Has(212, []byte{0x44, 0x44}) || buffer.Has(212, []byte{0x44, 0x41}) {
		order = binary.LittleEndian
		subtype = types.TypeLittleEndian
	} else if buffer.Has(212, []byte{0x11, 0x11}) {
		order = binary.BigEndian
		subtype = types.TypeBigEndian
	} else {
		return nil
	}

	for offset := 0; offset < 12; offset += 4 {
		dimension := order.Uint32(buffer[offset:])
		if dimension == 0 || dimension > math.MaxInt32 {
			return nil
		}
	}

	mode := order.Uint32(buffer[12:])

	validMode := mode <= 4 || mode == 6 || mode == 12 || mode == 101
	version := order.Uint32(buffer[108:])

	if !validMode || version != 20140 && version != 20141 || order.Uint32(buffer[92:]) > math.MaxInt32 {
		return nil
	}

	return &types.Metadata{Kind: types.KindMRCImage, Type: subtype}
}
