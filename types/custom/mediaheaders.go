package custom

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"

	"github.com/coalaura/wtf/types"
)

const ageArmorHeader = "-----BEGIN AGE ENCRYPTED FILE-----"

func detectMediaHeaders(b types.Buffer) *types.Metadata {
	switch string(b[:4]) {
	case "qoaf":
		if b.Len() < 16 {
			return nil
		}

		channels := uint32(b[8])
		rate := binary.BigEndian.Uint32(b[8:]) & 0xffffff
		samples := uint32(binary.BigEndian.Uint16(b[12:]))
		frameSize := uint32(binary.BigEndian.Uint16(b[14:]))
		expectedSize := 8 + 16*channels + 8*((samples+19)/20)*channels

		if channels > 0 && channels <= 8 && rate > 0 && samples > 0 && samples <= 5120 && frameSize == expectedSize {
			return &types.Metadata{Kind: types.KindQOAAudio}
		}
	case "DRAC":
		if b.Len() >= 11 && b[4] == 'O' && b[7] <= 1 && b[8] <= 1 {
			maxMinor := byte(3 - b[7])
			if b[5] == 1 || b[5] == 2 && b[6] <= maxMinor {
				return &types.Metadata{Kind: types.KindDracoGeometry}
			}
		}
	case "----":
		if validAgeArmorPrefix(b) {
			return &types.Metadata{Kind: types.KindAgeEncryptedFile, Type: types.TypeASCIIArmored}
		}
	}

	return nil
}

func detectAsepriteHeader(b types.Buffer) *types.Metadata {
	if b.Len() < 144 || !b.Has(132, []byte{0xfa, 0xf1}) {
		return nil
	}

	total := uint64(binary.LittleEndian.Uint32(b))
	frames := binary.LittleEndian.Uint16(b[6:])
	width := binary.LittleEndian.Uint16(b[8:])
	height := binary.LittleEndian.Uint16(b[10:])
	depth := binary.LittleEndian.Uint16(b[12:])
	frameSize := uint64(binary.LittleEndian.Uint32(b[128:]))

	if frames == 0 || width == 0 || height == 0 || depth != 8 && depth != 16 && depth != 32 || frameSize < 16 || 128+frameSize > total {
		return nil
	}

	chunks := uint64(binary.LittleEndian.Uint32(b[140:]))
	if chunks == 0 {
		chunks = uint64(binary.LittleEndian.Uint16(b[134:]))
	}

	if chunks > (frameSize-16)/6 || uint64(frames)*16 > total-128 {
		return nil
	}

	return &types.Metadata{Kind: types.KindAsepriteSprite}
}

func validAgeArmorPrefix(b types.Buffer) bool {
	if !b.Has(0, []byte(ageArmorHeader)) {
		return false
	}

	offset := len(ageArmorHeader)
	if offset < b.Len() && b[offset] == '\r' {
		offset++
	}

	if offset >= b.Len() || b[offset] != '\n' {
		return false
	}

	// Only decode enough armor to corroborate the native version line.
	var (
		encoded [32]byte
		decoded [24]byte
	)

	count := 0
	limit := min(b.Len(), offset+128)

	for offset++; offset < limit && count < len(encoded); offset++ {
		character := b[offset]
		if character == '\r' || character == '\n' {
			continue
		}

		encoded[count] = character
		count++
	}

	if count != len(encoded) {
		return false
	}

	length, err := base64.StdEncoding.Strict().Decode(decoded[:], encoded[:])

	return err == nil && bytes.HasPrefix(decoded[:length], []byte("age-encryption.org/v1\n"))
}
