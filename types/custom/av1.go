package custom

import "github.com/coalaura/wtf/types"

func DetectAV1(b types.Buffer) *types.Metadata {
	if b.Len() < 4 {
		return nil
	}

	// AV1 OBU streams typically start with Temporal Delimiter (0x12)
	// followed by size byte (usually 0x00), then Sequence Header (0x0a)
	if b[0] != 0x12 || b[1] != 0x00 {
		return nil
	}

	if !isValidAV1OBUHeader(b[2], 1) {
		return nil
	}

	return &types.Metadata{Kind: types.KindAV1Video, Confidence: types.ConfidenceMedium}
}

func isValidAV1OBUHeader(header, expectedType byte) bool {
	if header&0x80 != 0 {
		return false
	}

	if header&0x01 != 0 {
		return false
	}

	if header&0x02 == 0 {
		return false
	}

	obuType := (header >> 3) & 0x0F

	return obuType == expectedType
}
