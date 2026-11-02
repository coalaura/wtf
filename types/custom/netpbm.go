package custom

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/coalaura/wtf/types"
)

func DetectNetpbm(b types.Buffer) *types.Metadata {
	if b.Len() < 3 {
		return nil
	}

	if b[0] != 'P' {
		return nil
	}

	if !isNetpbmDelimiter(b[2]) {
		return nil
	}

	if b[1] == '7' {
		if !validPAMHeader(b[3:]) {
			return nil
		}
	} else {
		offset := 2

		width, ok := netpbmNumber(b, &offset)
		if !ok || width == 0 {
			return nil
		}

		height, ok := netpbmNumber(b, &offset)
		if !ok || height == 0 {
			return nil
		}

		if b[1] != '1' && b[1] != '4' {
			maximum, ok := netpbmNumber(b, &offset)
			if !ok || maximum == 0 || maximum > 65535 {
				return nil
			}
		}
	}

	switch b[1] {
	case '1':
		return &types.Metadata{
			Kind: types.KindNetpbmImage,
			Type: types.TypePBMASCII,
		}
	case '2':
		return &types.Metadata{
			Kind: types.KindNetpbmImage,
			Type: types.TypePGMASCII,
		}
	case '3':
		return &types.Metadata{
			Kind: types.KindNetpbmImage,
			Type: types.TypePPMASCII,
		}
	case '4':
		return &types.Metadata{
			Kind: types.KindNetpbmImage,
			Type: types.TypePBMBinary,
		}
	case '5':
		return &types.Metadata{
			Kind: types.KindNetpbmImage,
			Type: types.TypePGMBinary,
		}
	case '6':
		return &types.Metadata{
			Kind: types.KindNetpbmImage,
			Type: types.TypePPMBinary,
		}
	case '7':
		return &types.Metadata{
			Kind: types.KindNetpbmImage,
			Type: types.TypePAM,
		}
	default:
		return nil
	}
}

func isNetpbmDelimiter(c byte) bool {
	switch c {
	case ' ', '\t', '\r', '\n', '\f':
		return true
	default:
		return false
	}
}

func netpbmNumber(b types.Buffer, offset *int) (uint64, bool) {
	for *offset < b.Len() {
		if isNetpbmDelimiter(b[*offset]) {
			*offset++

			continue
		}

		if b[*offset] == '#' {
			end := bytes.IndexByte(b[*offset:], '\n')
			if end < 0 {
				return 0, false
			}

			*offset += end + 1

			continue
		}

		break
	}

	start := *offset

	for *offset < b.Len() && b[*offset] >= '0' && b[*offset] <= '9' {
		*offset++
	}

	if start == *offset || *offset >= b.Len() || !isNetpbmDelimiter(b[*offset]) {
		return 0, false
	}

	value, err := strconv.ParseUint(string(b[start:*offset]), 10, 32)
	return value, err == nil
}

func validPAMHeader(data []byte) bool {
	var (
		width   bool
		height  bool
		depth   bool
		maximum bool
	)

	for line := range strings.SplitSeq(string(data), "\n") {
		fields := strings.Fields(line)
		if len(fields) == 1 && fields[0] == "ENDHDR" {
			return width && height && depth && maximum
		}

		if len(fields) != 2 {
			continue
		}

		value, err := strconv.ParseUint(fields[1], 10, 32)
		if err != nil || value == 0 {
			continue
		}

		switch fields[0] {
		case "WIDTH":
			width = true
		case "HEIGHT":
			height = true
		case "DEPTH":
			depth = true
		case "MAXVAL":
			maximum = value <= 65535
		}
	}

	return false
}
