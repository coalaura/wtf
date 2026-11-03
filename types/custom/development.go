package custom

import (
	"encoding/binary"

	"github.com/coalaura/wtf/types"
)

func detectDevelopmentHeaders(b types.Buffer) *types.Metadata {
	if b.Len() < 12 {
		return nil
	}

	switch string(b[:4]) {
	case "MIDX":
		if (b[4] == 1 || b[4] == 2) && (b[5] == 1 || b[5] == 2) && b[6] >= 4 {
			return &types.Metadata{Kind: types.KindGitMultiPackIndex}
		}
	case "REFT":
		if b.Len() < 24 || binary.BigEndian.Uint64(b[8:]) > binary.BigEndian.Uint64(b[16:]) {
			return nil
		}

		if b[4] == 1 || b[4] == 2 && (b.Has(24, []byte("sha1")) || b.Has(24, []byte("s256"))) {
			return &types.Metadata{Kind: types.KindGitReftable}
		}
	case "MTME":
		version := binary.BigEndian.Uint32(b[4:])
		hash := binary.BigEndian.Uint32(b[8:])

		if version == 1 && (hash == 1 || hash == 2) {
			return &types.Metadata{Kind: types.KindGitCruftPackMtimes}
		}
	case "gcno":
		if validGCCVersion(binary.BigEndian.Uint32(b[4:])) {
			return &types.Metadata{Kind: types.KindGCCCoverageNotes, Type: types.TypeBigEndian}
		}
	case "oncg":
		if validGCCVersion(binary.LittleEndian.Uint32(b[4:])) {
			return &types.Metadata{Kind: types.KindGCCCoverageNotes, Type: types.TypeLittleEndian}
		}
	case "gcda":
		if validGCCVersion(binary.BigEndian.Uint32(b[4:])) {
			return &types.Metadata{Kind: types.KindGCCCoverageData, Type: types.TypeBigEndian}
		}
	case "adcg":
		if validGCCVersion(binary.LittleEndian.Uint32(b[4:])) {
			return &types.Metadata{Kind: types.KindGCCCoverageData, Type: types.TypeLittleEndian}
		}
	}

	return nil
}

func validGCCVersion(version uint32) bool {
	major := byte(version >> 24)
	minor := byte(version >> 16)
	patch := byte(version >> 8)
	status := byte(version)

	validMajor := major >= '3' && major <= '9' || major >= 'A' && major <= 'Z'
	validStatus := status == '*' || status == 'e' || status == 'p' || status == 'r'

	return validMajor && minor >= '0' && minor <= '9' && patch >= '0' && patch <= '9' && validStatus
}
