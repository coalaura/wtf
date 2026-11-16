package custom

import (
	"bytes"
	"strconv"

	"github.com/coalaura/wtf/types"
)

func DetectKiCad(buffer types.Buffer) *types.Metadata {
	position := 0
	input := buffer[:min(buffer.Len(), 4096)]

	if !kicadSkipSpace(input, &position) || input[position] != '(' {
		return nil
	}

	position++

	document := kicadToken(input, &position)
	if document != "kicad_pcb" && document != "kicad_sch" {
		return nil
	}

	if kicadToken(input, &position) != "(" || kicadToken(input, &position) != "version" {
		return nil
	}

	version, err := strconv.ParseUint(kicadToken(input, &position), 10, 32)
	if err != nil || version == 0 || kicadToken(input, &position) != ")" || kicadToken(input, &position) != "(" {
		return nil
	}

	field := kicadToken(input, &position)
	if field != "generator" && (field != "host" || document != "kicad_pcb") {
		return nil
	}

	if !kicadHeaderValue(kicadToken(input, &position)) {
		return nil
	}

	if field == "host" && !kicadHeaderValue(kicadToken(input, &position)) {
		return nil
	}

	if kicadToken(input, &position) != ")" {
		return nil
	}

	if document == "kicad_pcb" {
		return &types.Metadata{Kind: types.KindKiCadPCB}
	}

	return &types.Metadata{Kind: types.KindKiCadSchematic}
}

func kicadSkipSpace(input []byte, position *int) bool {
	for *position < len(input) {
		value := input[*position]
		if value == ' ' || value == '\t' || value == '\n' || value == '\r' || value == 0 {
			*position++

			continue
		}

		if value == '#' {
			end := bytes.IndexByte(input[*position:], '\n')
			if end < 0 {
				return false
			}

			*position += end + 1

			continue
		}

		break
	}

	return *position < len(input)
}

func kicadToken(input []byte, position *int) string {
	if !kicadSkipSpace(input, position) {
		return ""
	}

	start := *position
	value := input[start]
	*position++

	if value == '(' || value == ')' {
		return string(input[start:*position])
	}

	if value == '"' {
		for *position < len(input) {
			value = input[*position]
			*position++

			if value == '"' {
				return string(input[start:*position])
			}

			if value == '\\' && *position < len(input) {
				*position++
			}
		}

		return ""
	}

	for *position < len(input) {
		value = input[*position]
		if value == '(' || value == ')' || value == ' ' || value == '\t' || value == '\r' || value == '\n' || value == 0 {
			return string(input[start:*position])
		}

		*position++
	}

	return ""
}

func kicadHeaderValue(value string) bool {
	return value != "" && value != "(" && value != ")"
}
