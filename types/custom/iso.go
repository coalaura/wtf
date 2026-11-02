package custom

import (
	"io"
	"slices"

	"github.com/coalaura/wtf/types"
)

func DetectISOBaseMedia(b types.Buffer) *types.Metadata {
	if b.Has(0, []byte{0, 0, 0, 12, 'j', 'P', ' ', ' ', 13, 10, 135, 10}) {
		b = b[12:]
	}

	if b.Len() < 16 {
		return nil
	}

	boxSize32, ok := b.U32BE(0)
	if !ok || !b.Has(4, []byte("ftyp")) {
		return nil
	}

	brandOffset := 8
	compatibleOffset := 16
	boxEnd := uint64(boxSize32)

	switch boxSize32 {
	case 1:
		high, ok := b.U32BE(8)
		if !ok {
			return nil
		}

		low, ok := b.U32BE(12)
		if !ok {
			return nil
		}

		boxEnd = uint64(high)<<32 | uint64(low)
		brandOffset = 16
		compatibleOffset = 24
	case 0:
		boxEnd = uint64(b.Len())
	}

	if boxEnd < uint64(compatibleOffset) || b.Len() < compatibleOffset {
		return nil
	}

	if boxEnd > uint64(b.Len()) {
		return nil
	}

	boxEnd = min(boxEnd, uint64(maxScanSize))

	if !hasISOBrand(b, brandOffset, compatibleOffset, int(boxEnd), "isom", "iso2", "iso3", "iso4", "iso5", "iso6", "mp41", "mp42", "dash", "avif", "avis", "heic", "heix", "hevc", "hevx", "mif1", "msf1", "mjp2", "M4A ", "M4B ", "M4P ", "M4V ", "f4v ", "qt  ", "3gp4", "3gp5", "3gp6", "3gs7", "3ge6", "3gg6", "3gp1", "3gp2", "3g2a", "3g2b", "crx ", "braw", "MSNV") {
		return nil
	}

	boxLimit := int(boxEnd)

	// Get major brand
	var majorBrand string

	if brandOffset+4 <= b.Len() {
		majorBrand = string(b[brandOffset : brandOffset+4])
	}

	// Check major brand first
	switch majorBrand {
	case "qt  ":
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeQuickTimeMovie}
	case "M4V ":
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeM4VVideo}
	case "f4v ":
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeF4VVideo}
	case "M4A ", "M4B ", "M4P ":
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeMPEG4Audio}
	case "avif":
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeAVIFImage}
	case "avis":
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeAVIFImageSequence}
	case "heic", "heix", "hevc", "hevx", "mif1", "msf1":
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeHEIFImage}
	case "mjp2":
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeMotionJPEG2000Video}
	case "crx ":
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeCanonRAW3Image}
	case "braw":
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeBlackmagicRAWVideo}
	case "3g2a", "3g2b":
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.Type3G2Video}
	}

	// Check for 3GPP prefix in major brand (3gp1-3gp6, 3gs7, 3ge6, 3gg6)
	if len(majorBrand) == 4 && (majorBrand[:3] == "3gp" || majorBrand[:3] == "3g2") {
		if majorBrand[3] >= '1' && majorBrand[3] <= '9' {
			return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.Type3GPPVideo}
		}
	}

	// For generic major brands (isom, iso2, mp41, etc.), check compatible brands
	if hasISOBrand(b, brandOffset, compatibleOffset, boxLimit, "avif") {
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeAVIFImage}
	}

	if hasISOBrand(b, brandOffset, compatibleOffset, boxLimit, "avis") {
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeAVIFImageSequence}
	}

	if hasISOBrand(b, brandOffset, compatibleOffset, boxLimit, "heic", "heix", "hevc", "hevx", "mif1", "msf1") {
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeHEIFImage}
	}

	if hasISOBrand(b, brandOffset, compatibleOffset, boxLimit, "M4V ") {
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeM4VVideo}
	}

	if hasISOBrand(b, brandOffset, compatibleOffset, boxLimit, "f4v ") {
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeF4VVideo}
	}

	if hasISOBrand(b, brandOffset, compatibleOffset, boxLimit, "3g2a", "3g2b") {
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.Type3G2Video}
	}

	if hasISOBrandPrefix(b, brandOffset, compatibleOffset, boxLimit, "3gp", "3g2") {
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.Type3GPPVideo}
	}

	if hasISOBrand(b, brandOffset, compatibleOffset, boxLimit, "M4A ", "M4B ", "M4P ") {
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeMPEG4Audio}
	}

	if hasISOBrand(b, brandOffset, compatibleOffset, boxLimit, "qt  ") {
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeQuickTimeMovie}
	}

	if hasISOBrand(b, brandOffset, compatibleOffset, boxLimit, "mjp2") {
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeMotionJPEG2000Video}
	}

	if hasISOBrand(b, brandOffset, compatibleOffset, boxLimit, "crx ") {
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeCanonRAW3Image}
	}

	if hasISOBrand(b, brandOffset, compatibleOffset, boxLimit, "braw") {
		return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeBlackmagicRAWVideo}
	}

	if hasISOBrand(b, brandOffset, compatibleOffset, boxLimit, "isom", "iso2", "iso3", "iso4", "iso5", "iso6", "mp41", "mp42", "dash") {
		audio, video := isoTrackKinds(b, 0)
		if video {
			return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeMP4Video}
		}

		if audio {
			return &types.Metadata{Kind: types.KindISOBaseMedia, Type: types.TypeMPEG4Audio}
		}
	}

	return &types.Metadata{
		Kind: types.KindISOBaseMedia,
	}
}

func DetectISOReaderAt(b types.Buffer, reader io.ReaderAt, size int64) *types.Metadata {
	meta := DetectISOBaseMedia(b)
	if meta == nil || meta.Type != types.TypeNone {
		return meta
	}

	var (
		header [16]byte
		offset int64
	)

	for range 256 {
		if size-offset < 8 || !readAtExactly(reader, header[:8], offset) {
			break
		}

		length32, _ := types.Buffer(header[:]).U32BE(0)
		length := uint64(length32)
		headerSize := 8

		if length == 1 {
			if size-offset < 16 || !readAtExactly(reader, header[8:], offset+8) {
				break
			}

			high, _ := types.Buffer(header[:]).U32BE(8)
			low, _ := types.Buffer(header[:]).U32BE(12)
			length = uint64(high)<<32 | uint64(low)
			headerSize = 16
		}

		if length == 0 {
			length = uint64(size - offset)
		}

		if length < uint64(headerSize) || length > uint64(size-offset) {
			break
		}

		if string(header[4:8]) == "moov" && length <= maxScanSize {
			box := make(types.Buffer, int(length))
			if !readAtExactly(reader, box, offset) {
				break
			}

			audio, video := isoTrackKinds(box, 0)
			if video {
				meta.Type = types.TypeMP4Video
			} else if audio {
				meta.Type = types.TypeMPEG4Audio
			}

			break
		}

		offset += int64(length)
	}

	return meta
}

func isoTrackKinds(b types.Buffer, depth int) (bool, bool) {
	if depth > 4 {
		return false, false
	}

	var (
		audio bool
		video bool
	)

	for offset := 0; offset+8 <= b.Len(); {
		length32, _ := b.U32BE(offset)
		length := uint64(length32)
		header := 8

		if length == 1 {
			high, _ := b.U32BE(offset + 8)
			low, _ := b.U32BE(offset + 12)
			length = uint64(high)<<32 | uint64(low)
			header = 16
		}

		if length == 0 {
			length = uint64(b.Len() - offset)
		}

		if length < uint64(header) || length > uint64(b.Len()-offset) {
			break
		}

		payload := b[offset+header : offset+int(length)]

		switch string(b[offset+4 : offset+8]) {
		case "moov", "trak", "mdia":
			childAudio, childVideo := isoTrackKinds(payload, depth+1)
			audio = audio || childAudio
			video = video || childVideo
		case "hdlr":
			if depth == 3 && payload.Len() >= 12 {
				audio = audio || payload.Has(8, []byte("soun"))
				video = video || payload.Has(8, []byte("vide"))
			}
		}

		offset += int(length)
	}

	return audio, video
}

func hasISOBrand(b types.Buffer, majorOffset int, compatOffset int, boxEnd int, brands ...string) bool {
	if majorOffset+4 <= b.Len() {
		if slices.Contains(brands, string(b[majorOffset:majorOffset+4])) {
			return true
		}
	}

	for off := compatOffset; off+4 <= boxEnd && off+4 <= b.Len(); off += 4 {
		if slices.Contains(brands, string(b[off:off+4])) {
			return true
		}
	}

	return false
}

func hasISOBrandPrefix(b types.Buffer, majorOffset int, compatOffset int, boxEnd int, prefixes ...string) bool {
	if majorOffset+3 <= b.Len() {
		if slices.Contains(prefixes, string(b[majorOffset:majorOffset+3])) {
			return true
		}
	}

	for off := compatOffset; off+4 <= boxEnd && off+4 <= b.Len(); off += 4 {
		if slices.Contains(prefixes, string(b[off:off+3])) {
			return true
		}
	}

	return false
}
