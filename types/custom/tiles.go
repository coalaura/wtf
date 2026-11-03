package custom

import (
	"encoding/binary"

	"github.com/coalaura/wtf/types"
)

func detectTileHeader(b types.Buffer) *types.Metadata {
	if b.Len() < 16 || binary.LittleEndian.Uint32(b[4:]) != 1 {
		return nil
	}

	total := uint64(binary.LittleEndian.Uint32(b[8:]))

	if b.Has(0, []byte("cmpt")) {
		count := uint64(binary.LittleEndian.Uint32(b[12:]))
		if total >= 16 && count > 0 && count <= (total-16)/12 {
			return &types.Metadata{Kind: types.KindCompositeTile}
		}

		return nil
	}

	headerSize := uint64(28)

	if b.Has(0, []byte("i3dm")) {
		headerSize = 32
	}

	if uint64(b.Len()) < headerSize || total < headerSize {
		return nil
	}

	tablesSize := uint64(0)

	for offset := 12; offset < 28; offset += 4 {
		tablesSize += uint64(binary.LittleEndian.Uint32(b[offset:]))
	}

	if tablesSize > total-headerSize {
		return nil
	}

	switch string(b[:4]) {
	case "b3dm":
		return &types.Metadata{Kind: types.KindBatched3DModel}
	case "i3dm":
		if binary.LittleEndian.Uint32(b[28:]) <= 1 {
			return &types.Metadata{Kind: types.KindInstanced3DModel}
		}
	case "pnts":
		return &types.Metadata{Kind: types.KindPointCloudTile}
	}

	return nil
}
