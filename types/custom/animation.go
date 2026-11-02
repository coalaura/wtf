package custom

import (
	"math"

	"github.com/coalaura/wtf/types"
)

func DetectFLIC(b types.Buffer) *types.Metadata {
	if b.Len() < 128 {
		return nil
	}

	size, _ := b.U32LE(0)
	magic, _ := b.U16LE(4)
	frames, _ := b.U16LE(6)
	width, _ := b.U16LE(8)
	height, _ := b.U16LE(10)
	depth, _ := b.U16LE(12)

	validMagic := magic == 0xaf11 || magic == 0xaf12 || magic == 0xaf44
	validDepth := depth == 8 || depth == 15 || depth == 16 || depth == 24 || depth == 32

	if validMagic && validDepth && size >= 128 && frames > 0 && width > 0 && height > 0 {
		return &types.Metadata{Kind: types.KindFLICAnimation}
	}

	return nil
}

func DetectSpine(b types.Buffer) *types.Metadata {
	if b.Len() < 32 {
		return nil
	}

	// Spine 4.x starts with an eight-byte hash and a length-prefixed version, not a magic string.
	length := int(b[8]) - 1
	if length < 5 || length > 13 || 9+length+16 > b.Len() || !b.Has(9, []byte("4.")) {
		return nil
	}

	version := b[9 : 9+length]
	if version[2] < '0' || version[2] > '2' || version[3] != '.' {
		return nil
	}

	for _, character := range version[4:] {
		if character < '0' || character > '9' {
			return nil
		}
	}

	for index := range 4 {
		bits, _ := b.U32BE(9 + length + index*4)

		value := math.Float32frombits(bits)
		if math.IsNaN(float64(value)) || math.IsInf(float64(value), 0) || index >= 2 && value < 0 {
			return nil
		}
	}

	return &types.Metadata{Kind: types.KindSpineAnimation, Type: types.TypeSpineBinary, Confidence: types.ConfidenceMedium}
}
