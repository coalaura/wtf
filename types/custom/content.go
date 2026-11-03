package custom

import "github.com/coalaura/wtf/types"

func DetectContentHeaders(b types.Buffer) *types.Metadata {
	if b.Len() < 4 {
		return nil
	}

	// Dispatch before inspecting fields so unrelated files retain a cheap path.
	switch string(b[:4]) {
	case "MIDX", "REFT", "MTME", "gcno", "oncg", "gcda", "adcg":
		return detectDevelopmentHeaders(b)
	case "\xd0\x0d\xfe\xed", "\xd7\xb7\xab\x1e", "AVB0", "CrAU", "UF2\n", "\x9f\xeb\x01\x00", "\xeb\x9f\x01\x00":
		return detectFirmwareHeaders(b)
	case "b3dm", "i3dm", "pnts", "cmpt":
		return detectTileHeader(b)
	case "\x26\xfc\x8f\x88", "\x88\x8f\xfc\x26", "\xeb\xf2\x89\x87", "\x87\x89\xf2\xeb", "\x43\x27\x41\x1a", "\x1a\x41\x27\x43":
		return detectUCSCHeader(b)
	case "qoaf", "DRAC", "----":
		return detectMediaHeaders(b)
	}

	if b.Has(4, []byte{0xe0, 0xa5}) {
		return detectAsepriteHeader(b)
	}

	return nil
}
