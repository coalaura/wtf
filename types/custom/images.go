package custom

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/coalaura/wtf/types"
)

func DetectJPEG(b types.Buffer) *types.Metadata {
	if !b.Has(0, []byte{0xff, 0xd8, 0xff}) {
		return nil
	}

	for offset := 2; offset+4 <= b.Len(); {
		if b[offset] != 0xff {
			break
		}

		marker := b[offset+1]
		if marker == 0xff {
			offset++

			continue
		}

		length, _ := b.U16BE(offset + 2)
		if length < 2 || int(length) > b.Len()-offset-2 {
			break
		}

		if marker == 0xf7 && length >= 11 {
			return &types.Metadata{Kind: types.KindJPEGLSImage}
		}

		if marker >= 0xc0 && marker <= 0xcf && marker != 0xc4 && marker != 0xc8 && marker != 0xcc && length >= 11 {
			return &types.Metadata{Kind: types.KindJPEGImage}
		}

		if marker == 0xda || marker == 0xd9 {
			break
		}

		offset += 2 + int(length)
	}

	return &types.Metadata{Kind: types.KindJPEGImage, Confidence: types.ConfidenceMedium}
}

func DetectRasterHeaders(b types.Buffer) *types.Metadata {
	if b.Has(0, []byte("BM")) && b.Len() >= 26 {
		fileSize, _ := b.U32LE(2)
		reserved, _ := b.U32LE(6)
		pixels, _ := b.U32LE(10)
		dibSize, _ := b.U32LE(14)
		validDIB := dibSize == 12 || dibSize == 40 || dibSize == 52 || dibSize == 56 || dibSize == 64 || dibSize == 108 || dibSize == 124

		if validDIB && reserved == 0 && pixels >= 14+dibSize && fileSize > pixels && uint64(dibSize)+14 <= uint64(b.Len()) {
			return &types.Metadata{Kind: types.KindBMPImage}
		}
	}

	if b.Has(0, []byte{0, 0, 2, 0}) && b.Len() >= 22 {
		count, _ := b.U16LE(4)
		if count == 0 || int(count) > (b.Len()-6)/16 {
			return nil
		}

		for index := range int(count) {
			offset := 6 + index*16
			length, _ := b.U32LE(offset + 8)
			position, _ := b.U32LE(offset + 12)

			if b[offset+3] != 0 || length < 8 || position < uint32(6+int(count)*16) {
				return nil
			}
		}

		return &types.Metadata{Kind: types.KindWindowsCursorImage}
	}

	if b.Has(0, []byte("8BPS")) && b.Len() >= 26 {
		version, _ := b.U16BE(4)

		reserved, _ := b.U32BE(6)
		reservedEnd, _ := b.U16BE(10)
		channels, _ := b.U16BE(12)
		height, _ := b.U32BE(14)
		width, _ := b.U32BE(18)

		if reserved != 0 || reservedEnd != 0 || channels == 0 || channels > 56 || height == 0 || width == 0 {
			return nil
		}

		switch version {
		case 1:
			return &types.Metadata{Kind: types.KindPhotoshopDocument}
		case 2:
			return &types.Metadata{Kind: types.KindPhotoshopDocument, Type: types.TypeLargeDocument}
		}
	}

	if b.Len() >= 28 {
		headerSize, _ := b.U32BE(0)
		version, _ := b.U32BE(4)

		width, _ := b.U32BE(8)
		height, _ := b.U32BE(12)
		channels, _ := b.U32BE(16)

		if width == 0 || height == 0 || channels == 0 || channels > 4 || headerSize > uint32(b.Len()) {
			return nil
		}

		if b.Has(20, []byte("GIMP")) && version == 2 && headerSize > 28 && (channels == 1 || channels == 4) {
			return &types.Metadata{Kind: types.KindGIMPBrush}
		}

		if b.Has(20, []byte("GPAT")) && version == 1 && headerSize > 24 {
			return &types.Metadata{Kind: types.KindGIMPPattern}
		}
	}

	return nil
}

func DetectTrueType(b types.Buffer) *types.Metadata {
	if !b.Has(0, []byte("true")) && !b.Has(0, []byte{0, 1, 0, 0}) {
		return nil
	}

	count, ok := b.U16BE(4)
	if !ok || count == 0 || int(count) > (b.Len()-12)/16 {
		return nil
	}

	var (
		hasHead   bool
		hasCmap   bool
		hasGlyphs bool
	)

	for index := range int(count) {
		offset := 12 + index*16

		position, _ := b.U32BE(offset + 8)
		length, _ := b.U32BE(offset + 12)

		if position < uint32(12+int(count)*16) || position%4 != 0 || length == 0 {
			return nil
		}

		switch string(b[offset : offset+4]) {
		case "head":
			hasHead = true
		case "cmap":
			hasCmap = true
		case "glyf":
			hasGlyphs = true
		}
	}

	if hasHead && hasCmap && hasGlyphs {
		return &types.Metadata{Kind: types.KindTrueTypeFont}
	}

	return nil
}

func DetectXBM(b types.Buffer) *types.Metadata {
	if !isLikelyASCIIText(b) || !bytes.Contains(b, []byte("#define")) {
		return nil
	}

	var (
		name   string
		width  bool
		height bool
	)

	for line := range strings.SplitSeq(string(b), "\n") {
		fields := strings.Fields(line)
		if len(fields) != 3 || fields[0] != "#define" {
			continue
		}

		value, err := strconv.ParseUint(fields[2], 10, 32)
		if err != nil || value == 0 {
			continue
		}

		prefix, found := strings.CutSuffix(fields[1], "_width")
		if found {
			name = prefix
			width = true
		}

		if name != "" && fields[1] == name+"_height" {
			height = true
		}
	}

	if width && height && bytes.Contains(b, []byte(name+"_bits")) && bytes.Contains(b, []byte("static")) && bytes.Contains(b, []byte("{")) && bytes.Contains(b, []byte("0x")) {
		return &types.Metadata{Kind: types.KindXBitMapImage, Confidence: types.ConfidenceMedium}
	}

	return nil
}
