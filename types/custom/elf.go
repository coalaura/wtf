package custom

import "github.com/coalaura/wtf/types"

func DetectELF(b types.Buffer) *types.Metadata {
	if !b.Has(0, []byte{0x7f, 'E', 'L', 'F'}) {
		return nil
	}

	if b.Has(8, []byte{'A', 'I', 0x01}) || b.Has(8, []byte{'A', 'I', 0x02}) {
		return &types.Metadata{
			Kind: types.KindAppImage,
		}
	}

	class := byte(0)
	data := byte(0)

	if b.Len() > 4 {
		class = b[4]
	}

	if b.Len() > 5 {
		data = b[5]
	}

	typ := elfType(class, data)

	return &types.Metadata{
		Kind: types.KindExecutableAndLinkableFormat,
		Type: typ,
	}
}

func elfType(class byte, data byte) types.TypeID {
	switch class {
	case 1:
		switch data {
		case 1:
			return types.Type32BitLittleEndian
		case 2:
			return types.Type32BitBigEndian
		default:
			return types.Type32Bit
		}
	case 2:
		switch data {
		case 1:
			return types.Type64BitLittleEndian
		case 2:
			return types.Type64BitBigEndian
		default:
			return types.Type64Bit
		}
	default:
		return types.TypeNone
	}
}
