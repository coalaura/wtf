package custom

import "github.com/coalaura/wtf/types"

func DetectPCAPNG(b types.Buffer) *types.Metadata {
	if b.Len() < 12 {
		return nil
	}

	if !b.Has(0, []byte{0x0a, 0x0d, 0x0d, 0x0a}) {
		return nil
	}

	if !b.Has(8, []byte{0x1a, 0x2b, 0x3c, 0x4d}) && !b.Has(8, []byte{0x4d, 0x3c, 0x2b, 0x1a}) {
		return nil
	}

	return &types.Metadata{
		Kind: types.KindPCAPNGCapture,
	}
}
