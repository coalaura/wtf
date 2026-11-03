package custom

import (
	"encoding/binary"

	"github.com/coalaura/wtf/types"
)

func detectUCSCHeader(b types.Buffer) *types.Metadata {
	switch string(b[:4]) {
	case "\x26\xfc\x8f\x88":
		if validBBIHeader(b, binary.LittleEndian) {
			return &types.Metadata{Kind: types.KindUCSCBigWig, Type: types.TypeLittleEndian}
		}
	case "\x88\x8f\xfc\x26":
		if validBBIHeader(b, binary.BigEndian) {
			return &types.Metadata{Kind: types.KindUCSCBigWig, Type: types.TypeBigEndian}
		}
	case "\xeb\xf2\x89\x87":
		if validBBIHeader(b, binary.LittleEndian) {
			return &types.Metadata{Kind: types.KindUCSCBigBed, Type: types.TypeLittleEndian}
		}
	case "\x87\x89\xf2\xeb":
		if validBBIHeader(b, binary.BigEndian) {
			return &types.Metadata{Kind: types.KindUCSCBigBed, Type: types.TypeBigEndian}
		}
	case "\x43\x27\x41\x1a":
		if validTwoBitHeader(b, binary.LittleEndian) {
			return &types.Metadata{Kind: types.KindUCSCTwoBit, Type: types.TypeLittleEndian}
		}
	case "\x1a\x41\x27\x43":
		if validTwoBitHeader(b, binary.BigEndian) {
			return &types.Metadata{Kind: types.KindUCSCTwoBit, Type: types.TypeBigEndian}
		}
	}

	return nil
}

func validBBIHeader(b types.Buffer, order binary.ByteOrder) bool {
	if b.Len() < 64 {
		return false
	}

	version := order.Uint16(b[4:])
	headerSize := uint64(64) + uint64(order.Uint16(b[6:]))*24
	treeOffset := order.Uint64(b[8:])
	dataOffset := order.Uint64(b[16:])
	indexOffset := order.Uint64(b[24:])

	return version >= 1 && version <= 4 && treeOffset >= headerSize && dataOffset >= headerSize && indexOffset >= headerSize
}

func validTwoBitHeader(b types.Buffer, order binary.ByteOrder) bool {
	return b.Len() >= 16 && order.Uint32(b[4:]) <= 1 && order.Uint32(b[12:]) == 0
}
