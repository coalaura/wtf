package custom

import "github.com/coalaura/wtf/types"

func DetectSkippableFrames(b types.Buffer) *types.Metadata {
	var skipped bool

	for b.Len() >= 8 && b[0]&0xf0 == 0x50 && b.Has(1, []byte{0x2a, 0x4d, 0x18}) {
		skipped = true

		length, _ := b.U32LE(4)
		if uint64(length) > uint64(b.Len()-8) {
			return &types.Metadata{Kind: types.KindZstandardLZ4Stream, Type: types.TypeSkinnableFrame, Confidence: types.ConfidenceMedium}
		}

		b = b[8+int(length):]
	}

	if !skipped {
		return nil
	}

	if b.Has(0, []byte{0x28, 0xb5, 0x2f, 0xfd}) {
		return &types.Metadata{Kind: types.KindZstandardArchive}
	}

	if b.Has(0, []byte{4, 0x22, 0x4d, 0x18}) {
		return &types.Metadata{Kind: types.KindLZ4Frame}
	}

	return &types.Metadata{Kind: types.KindZstandardLZ4Stream, Type: types.TypeSkinnableFrame}
}
