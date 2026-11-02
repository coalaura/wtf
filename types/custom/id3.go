package custom

import (
	"bytes"
	"io"

	"github.com/coalaura/wtf/types"
)

func SkipID3(b types.Buffer) (types.Buffer, bool) {
	return SkipID3ReaderAt(b, bytes.NewReader(b), int64(b.Len()))
}

func RefineID3Metadata(meta *types.Metadata) *types.Metadata {
	if meta != nil && meta.Kind == types.KindMPEGAudio && meta.Type == types.TypeMPEGLayer3 {
		return &types.Metadata{Kind: types.KindMPEGAudio, Type: types.TypeMP3ID3Tagged, Confidence: meta.Confidence}
	}

	return meta
}

func SkipID3ReaderAt(b types.Buffer, reader io.ReaderAt, size int64) (types.Buffer, bool) {
	var offset int64

	for tags := 0; b.Has(0, []byte("ID3")); tags++ {
		if tags >= 16 || b.Len() < 10 || b[3] < 2 || b[3] > 4 || b[4] == 0xff || b[6]|b[7]|b[8]|b[9] >= 0x80 {
			return nil, false
		}

		allowed := byte(0xc0)

		switch b[3] {
		case 3:
			allowed = 0xe0
		case 4:
			allowed = 0xf0
		}

		if b[5] & ^allowed != 0 {
			return nil, false
		}

		length := int64(b[6])<<21 | int64(b[7])<<14 | int64(b[8])<<7 | int64(b[9])
		length += 10

		footer := b[3] == 4 && b[5]&0x10 != 0
		if footer {
			length += 10
		}

		if length > size-offset {
			return nil, false
		}

		if footer {
			var trailer [10]byte

			if !readAtExactly(reader, trailer[:], offset+length-10) || string(trailer[:3]) != "3DI" || !bytes.Equal(trailer[3:], b[3:10]) {
				return nil, false
			}
		}

		offset += length

		if length <= int64(b.Len()) {
			b = b[int(length):]
		} else {
			b = nil
		}

		remaining := int(min(size-offset, maxScanSize))
		if b.Len() < remaining {
			b = make([]byte, remaining)
			if !readAtExactly(reader, b, offset) {
				return nil, false
			}
		}
	}

	return b, offset > 0
}
