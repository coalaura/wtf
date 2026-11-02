package custom

import (
	"bytes"
	"io"

	"github.com/coalaura/wtf/types"
)

const (
	tiffLittleEndian tiffByteOrder = iota
	tiffBigEndian
)

type tiffByteOrder uint8

func DetectTIFFSubtypes(b types.Buffer) *types.Metadata {
	return DetectTIFFReaderAt(b, bytes.NewReader(b), int64(b.Len()))
}

func DetectTIFFReaderAt(b types.Buffer, reader io.ReaderAt, size int64) *types.Metadata {
	if b.Has(0, []byte{'I', 'I', 'U', 0x00, 0x18, 0x00, 0x00, 0x00}) {
		return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypePanasonicRAW}
	}

	if !isTIFFHeader(b) {
		if b.Has(0, []byte{'I', 'I', 'R', 'O', 0x08, 0x00}) || b.Has(0, []byte{'M', 'M', 'O', 'R', 0x00, 0x00}) {
			return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypeOlympusRAW}
		}

		return nil
	}

	if b.Has(0, []byte{'I', 'I', 0x2a, 0x00}) && b.Has(8, []byte{'C', 'R', 2, 0}) && b.Len() >= 16 {
		return &types.Metadata{Kind: types.KindCanonRAWImage, Type: types.TypeCanonRAW2}
	}

	order, ifd0, ok := tiffHeaderInfo(b)
	if !ok {
		if b.Has(0, []byte{'I', 'I'}) {
			return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypeLittleEndian}
		}

		return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypeBigEndian}
	}

	if tiffHasTag(b, order, ifd0, 0xc612) {
		return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypeAdobeDNG}
	}

	makeValue, ok := tiffASCIIValueForTag(b, order, ifd0, 0x010f)
	budget := 32

	if ok && tiffHasRawImage(reader, size, order, ifd0, 0, &budget) {
		if bytes.HasPrefix(makeValue, []byte("NIKON")) || bytes.HasPrefix(makeValue, []byte("Nikon")) {
			return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypeNikonRAW}
		}

		if bytes.HasPrefix(makeValue, []byte("OLYMPUS")) || bytes.HasPrefix(makeValue, []byte("Olympus")) {
			return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypeOlympusRAW}
		}

		if bytes.HasPrefix(makeValue, []byte("PENTAX")) {
			return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypePentaxRAW}
		}

		if bytes.HasPrefix(makeValue, []byte("SONY")) {
			return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypeSonyRAW}
		}
	}

	if b.Has(0, []byte{'I', 'I'}) {
		return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypeLittleEndian}
	}

	return &types.Metadata{Kind: types.KindTIFFImage, Type: types.TypeBigEndian}
}

func tiffHasRawImage(reader io.ReaderAt, size int64, order tiffByteOrder, ifdOffset uint32, depth int, budget *int) bool {
	if depth > 4 || *budget == 0 || int64(ifdOffset)+2 > size {
		return false
	}

	*budget--

	var header [2]byte

	if !readAtExactly(reader, header[:], int64(ifdOffset)) {
		return false
	}

	count, _ := tiffU16(header[:], order, 0)

	length := int(count) * 12
	if length > 4096 || int64(length) > size-int64(ifdOffset)-2 {
		return false
	}

	b := make(types.Buffer, length)
	if !readAtExactly(reader, b, int64(ifdOffset)+2) {
		return false
	}

	for index := range int(count) {
		entry := index * 12

		tag, _ := tiffU16(b, order, entry)
		fieldType, _ := tiffU16(b, order, entry+2)
		items, _ := tiffU32(b, order, entry+4)

		if tag == 0x0106 && fieldType == 3 && items == 1 {
			value, _ := tiffU16(b, order, entry+8)
			if value == 32803 || value == 34892 {
				return true
			}
		}

		if tag == 0x014a && fieldType == 4 && items > 0 && items <= 16 {
			var offsets [64]byte

			values := b[entry+8 : entry+12]

			if items > 1 {
				position, _ := tiffU32(b, order, entry+8)
				values = offsets[:int(items)*4]

				if int64(position)+int64(len(values)) > size || !readAtExactly(reader, values, int64(position)) {
					continue
				}
			}

			for child := range int(items) {
				childIFD, _ := tiffU32(values, order, child*4)
				if childIFD > ifdOffset && tiffHasRawImage(reader, size, order, childIFD, depth+1, budget) {
					return true
				}
			}
		}
	}

	return false
}

func tiffHeaderInfo(b types.Buffer) (tiffByteOrder, uint32, bool) {
	if b.Has(0, []byte{'I', 'I', 42, 0}) {
		ifd0, ok := b.U32LE(4)
		if !ok || ifd0 >= uint32(b.Len()) {
			return 0, 0, false
		}

		return tiffLittleEndian, ifd0, true
	}

	if b.Has(0, []byte{'M', 'M', 0, 42}) {
		ifd0, ok := b.U32BE(4)
		if !ok || ifd0 >= uint32(b.Len()) {
			return 0, 0, false
		}

		return tiffBigEndian, ifd0, true
	}

	return 0, 0, false
}

func tiffHasTag(b types.Buffer, order tiffByteOrder, ifdOffset uint32, tag uint16) bool {
	count, ok := tiffU16(b, order, int(ifdOffset))
	if !ok {
		return false
	}

	entryBase := int(ifdOffset) + 2
	entryEnd := entryBase + int(count)*12

	if entryBase < 0 || entryEnd > b.Len() {
		return false
	}

	for i := range int(count) {
		entry := entryBase + i*12

		currentTag, ok := tiffU16(b, order, entry)
		if !ok {
			return false
		}

		if currentTag == tag {
			return true
		}
	}

	return false
}

func tiffASCIIValueForTag(b types.Buffer, order tiffByteOrder, ifdOffset uint32, tag uint16) ([]byte, bool) {
	count, ok := tiffU16(b, order, int(ifdOffset))
	if !ok {
		return nil, false
	}

	entryBase := int(ifdOffset) + 2
	entryEnd := entryBase + int(count)*12

	if entryBase < 0 || entryEnd > b.Len() {
		return nil, false
	}

	for i := range int(count) {
		entry := entryBase + i*12

		currentTag, ok := tiffU16(b, order, entry)
		if !ok || currentTag != tag {
			continue
		}

		fieldType, ok := tiffU16(b, order, entry+2)
		if !ok || fieldType != 2 {
			return nil, false
		}

		valueCount, ok := tiffU32(b, order, entry+4)
		if !ok || valueCount == 0 {
			return nil, false
		}

		if valueCount <= 4 {
			end := entry + 8 + int(valueCount)
			if end > b.Len() {
				return nil, false
			}

			return bytes.TrimRight(b[entry+8:end], "\x00"), true
		}

		valueOffset, ok := tiffU32(b, order, entry+8)
		if !ok {
			return nil, false
		}

		if uint64(valueOffset) > uint64(b.Len()) || uint64(valueCount) > uint64(b.Len())-uint64(valueOffset) {
			return nil, false
		}

		start := int(valueOffset)
		end := start + int(valueCount)

		return bytes.TrimRight(b[start:end], "\x00"), true
	}

	return nil, false
}

func tiffU16(b types.Buffer, order tiffByteOrder, offset int) (uint16, bool) {
	if order == tiffLittleEndian {
		return b.U16LE(offset)
	}

	return b.U16BE(offset)
}

func tiffU32(b types.Buffer, order tiffByteOrder, offset int) (uint32, bool) {
	if order == tiffLittleEndian {
		return b.U32LE(offset)
	}

	return b.U32BE(offset)
}

func isTIFFHeader(b types.Buffer) bool {
	return b.Has(0, []byte{'I', 'I', 42, 0}) || b.Has(0, []byte{'M', 'M', 0, 42})
}
