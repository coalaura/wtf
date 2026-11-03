package custom

import (
	"encoding/binary"

	"github.com/coalaura/wtf/types"
)

func detectFirmwareHeaders(b types.Buffer) *types.Metadata {
	switch string(b[:4]) {
	case "\xd0\x0d\xfe\xed":
		if validDevicetreeHeader(b) {
			return &types.Metadata{Kind: types.KindFlattenedDevicetree}
		}
	case "\xd7\xb7\xab\x1e":
		if validAndroidDTHeader(b) {
			return &types.Metadata{Kind: types.KindAndroidDTTable}
		}
	case "AVB0":
		if b.Len() >= 256 && binary.BigEndian.Uint32(b[4:]) == 1 {
			authSize := binary.BigEndian.Uint64(b[12:])
			auxSize := binary.BigEndian.Uint64(b[20:])

			if authSize%64 == 0 && auxSize%64 == 0 {
				return &types.Metadata{Kind: types.KindAndroidVBMeta}
			}
		}
	case "CrAU":
		if b.Len() >= 20 {
			version := binary.BigEndian.Uint64(b[4:])
			manifestSize := binary.BigEndian.Uint64(b[12:])

			if manifestSize > 0 && (version == 1 || version == 2 && b.Len() >= 24) {
				return &types.Metadata{Kind: types.KindUpdateEnginePayload}
			}
		}
	case "UF2\n":
		if b.Has(4, []byte{0x57, 0x51, 0x5d, 0x9e}) && b.Has(508, []byte{0x30, 0x6f, 0xb1, 0x0a}) {
			payloadSize := binary.LittleEndian.Uint32(b[16:])
			block := binary.LittleEndian.Uint32(b[20:])
			blocks := binary.LittleEndian.Uint32(b[24:])

			if payloadSize <= 476 && (blocks == 0 || block < blocks) {
				return &types.Metadata{Kind: types.KindUF2Firmware}
			}
		}
	case "\x9f\xeb\x01\x00":
		if validBTFHeader(b, binary.LittleEndian) {
			return &types.Metadata{Kind: types.KindBPFTypeFormat, Type: types.TypeLittleEndian}
		}
	case "\xeb\x9f\x01\x00":
		if validBTFHeader(b, binary.BigEndian) {
			return &types.Metadata{Kind: types.KindBPFTypeFormat, Type: types.TypeBigEndian}
		}
	}

	return nil
}

func validDevicetreeHeader(b types.Buffer) bool {
	if b.Len() < 28 {
		return false
	}

	total := uint64(binary.BigEndian.Uint32(b[4:]))
	structure := uint64(binary.BigEndian.Uint32(b[8:]))
	strings := uint64(binary.BigEndian.Uint32(b[12:]))
	reserve := uint64(binary.BigEndian.Uint32(b[16:]))
	version := binary.BigEndian.Uint32(b[20:])
	compatible := binary.BigEndian.Uint32(b[24:])
	headerSize := uint64(28)

	if version >= 17 {
		headerSize = 40
	} else if version >= 3 {
		headerSize = 36
	} else if version == 2 {
		headerSize = 32
	}

	if version < 1 || version > 17 || compatible > version || uint64(b.Len()) < headerSize || total < headerSize {
		return false
	}

	if structure < headerSize || structure >= total || structure%4 != 0 || strings < headerSize || strings > total || reserve < headerSize || reserve%8 != 0 || reserve+16 > total {
		return false
	}

	if version >= 3 && strings+uint64(binary.BigEndian.Uint32(b[32:])) > total {
		return false
	}

	return version < 17 || structure+uint64(binary.BigEndian.Uint32(b[36:])) <= total
}

func validAndroidDTHeader(b types.Buffer) bool {
	if b.Len() < 32 {
		return false
	}

	total := uint64(binary.BigEndian.Uint32(b[4:]))
	headerSize := uint64(binary.BigEndian.Uint32(b[8:]))
	entrySize := uint64(binary.BigEndian.Uint32(b[12:]))
	count := uint64(binary.BigEndian.Uint32(b[16:]))
	offset := uint64(binary.BigEndian.Uint32(b[20:]))
	version := binary.BigEndian.Uint32(b[28:])

	return version <= 1 && headerSize >= 32 && entrySize >= 32 && offset >= headerSize && offset <= total && count <= (total-offset)/entrySize
}

func validBTFHeader(b types.Buffer, order binary.ByteOrder) bool {
	if b.Len() < 24 {
		return false
	}

	headerSize := uint64(order.Uint32(b[4:]))
	typeOffset := uint64(order.Uint32(b[8:]))
	typeSize := uint64(order.Uint32(b[12:]))
	stringOffset := uint64(order.Uint32(b[16:]))
	stringSize := order.Uint32(b[20:])

	if headerSize < 24 || headerSize%4 != 0 || typeOffset%4 != 0 || typeSize%4 != 0 || stringOffset < typeOffset+typeSize || stringSize == 0 {
		return false
	}

	stringStart := headerSize + stringOffset

	return stringStart >= uint64(b.Len()) || b[stringStart] == 0
}
