package custom

import (
	"bytes"
	"unicode/utf16"
	"unicode/utf8"

	"github.com/coalaura/wtf/types"
)

func DetectTextEncoding(b types.Buffer) (types.TypeID, bool) {
	if b.Len() < 2 {
		return types.TypeNone, false
	}

	// Check for BOMs
	if b.Has(0, []byte{0xff, 0xfe, 0x00, 0x00}) {
		return types.TypeUTF32LE, true
	}

	if b.Has(0, []byte{0x00, 0x00, 0xfe, 0xff}) {
		return types.TypeUTF32BE, true
	}

	if b.Has(0, []byte{0xff, 0xfe}) {
		return types.TypeUTF16LE, true
	}

	if b.Has(0, []byte{0xfe, 0xff}) {
		return types.TypeUTF16BE, true
	}

	if b.Has(0, []byte{0xef, 0xbb, 0xbf}) {
		return types.TypeUTF8, true
	}

	// Try to detect UTF-16 without BOM
	if isLikelyUTF16LE(b) {
		return types.TypeUTF16LE, true
	}

	if isLikelyUTF16BE(b) {
		return types.TypeUTF16BE, true
	}

	return types.TypeNone, false
}

func isLikelyUTF16LE(b types.Buffer) bool {
	limit := min(b.Len(), 4096)
	if limit < 4 {
		return false
	}

	data := b[:limit]

	var (
		asciiPatterns int
		nullInOdd     int
		totalChecked  int
	)

	for i := 0; i+1 < len(data); i += 2 {
		totalChecked++

		if data[i] >= 0x20 && data[i] <= 0x7e && data[i+1] == 0x00 {
			asciiPatterns++
		}

		if data[i+1] == 0x00 && data[i] != 0x00 {
			nullInOdd++
		}
	}

	if totalChecked == 0 {
		return false
	}

	if nullInOdd == 0 {
		return false
	}

	nullRatio := float64(nullInOdd) / float64(totalChecked)
	asciiRatio := float64(asciiPatterns) / float64(totalChecked)

	if asciiRatio > 0.7 {
		return true
	}

	if nullRatio > 0.4 && asciiRatio > 0.3 {
		return true
	}

	return false
}

func isLikelyUTF16BE(b types.Buffer) bool {
	limit := min(b.Len(), 4096)
	if limit < 4 {
		return false
	}

	data := b[:limit]

	var (
		asciiPatterns int
		nullInEven    int
		totalChecked  int
	)

	for i := 0; i+1 < len(data); i += 2 {
		totalChecked++

		if data[i] == 0x00 && data[i+1] >= 0x20 && data[i+1] <= 0x7e {
			asciiPatterns++
		}

		if data[i] == 0x00 && data[i+1] != 0x00 {
			nullInEven++
		}
	}

	if totalChecked == 0 {
		return false
	}

	if nullInEven == 0 {
		return false
	}

	nullRatio := float64(nullInEven) / float64(totalChecked)
	asciiRatio := float64(asciiPatterns) / float64(totalChecked)

	if asciiRatio > 0.7 {
		return true
	}

	if nullRatio > 0.4 && asciiRatio > 0.3 {
		return true
	}

	return false
}

func isLikelyLatin1(b types.Buffer) bool {
	limit := min(b.Len(), 4096)
	if limit == 0 {
		return false
	}

	data := b[:limit]

	if utf8.Valid(data) {
		return false
	}

	var (
		highBytes int
		printable int
	)

	for _, c := range data {
		if c == 0 {
			return false
		}

		if c >= 0x80 {
			highBytes++
		}

		if c == '\n' || c == '\r' || c == '\t' || (c >= 0x20 && c <= 0x7e) || c >= 0xa0 {
			printable++
		}
	}

	if highBytes == 0 {
		return false
	}

	highRatio := float64(highBytes) / float64(limit)
	printableRatio := float64(printable) / float64(limit)

	// Latin-1 typically has some high bytes but is mostly printable
	if highRatio > 0.01 && highRatio < 0.5 && printableRatio > 0.9 {
		return true
	}

	return false
}

func TryDecodeUTF16(b types.Buffer) ([]byte, types.TypeID, bool) {
	if b.Len() < 2 {
		return nil, types.TypeNone, false
	}

	var (
		data []byte
		typ  types.TypeID
	)

	if b.Has(0, []byte{0xff, 0xfe}) {
		data = b[2:]
		typ = types.TypeUTF16LE
	} else if b.Has(0, []byte{0xfe, 0xff}) {
		data = b[2:]
		typ = types.TypeUTF16BE
	} else if isLikelyUTF16LE(b) {
		data = b
		typ = types.TypeUTF16LE
	} else if isLikelyUTF16BE(b) {
		data = b
		typ = types.TypeUTF16BE
	} else {
		return nil, types.TypeNone, false
	}

	if len(data)%2 != 0 {
		data = data[:len(data)-1]
	}

	if len(data) == 0 {
		return nil, typ, false
	}

	u16 := make([]uint16, len(data)/2)

	for i := range u16 {
		if typ == types.TypeUTF16LE {
			u16[i] = uint16(data[i*2]) | uint16(data[i*2+1])<<8
		} else {
			u16[i] = uint16(data[i*2])<<8 | uint16(data[i*2+1])
		}
	}

	runes := utf16.Decode(u16)

	var buf bytes.Buffer

	buf.Grow(len(runes) * 4)

	for _, r := range runes {
		buf.WriteRune(r)
	}

	return buf.Bytes(), typ, true
}
