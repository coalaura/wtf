package custom

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"io"
	"slices"
	"unicode/utf8"

	"github.com/coalaura/wtf/types"
)

type tensorExtent struct {
	start uint64
	end   uint64
}

func DetectSafeTensors(buffer types.Buffer) *types.Metadata {
	return detectSafeTensorsSized(buffer, int64(buffer.Len()))
}

func detectSafeTensorsSized(buffer types.Buffer, size int64) *types.Metadata {
	if buffer.Len() < 10 || buffer[8] != '{' {
		return nil
	}

	headerSize := binary.LittleEndian.Uint64(buffer)
	if headerSize < 2 || headerSize > maxScanSize-8 || headerSize > uint64(buffer.Len()-8) || uint64(size) < 8+headerSize {
		return nil
	}

	header := buffer[8 : 8+headerSize]
	if !utf8.Valid(header) {
		return nil
	}

	fields, ok := tensorJSONObject(header)
	if !ok {
		return nil
	}

	extents := make([]tensorExtent, 0, len(fields))

	for name, raw := range fields {
		if name == "__metadata__" {
			metadata, valid := tensorJSONObject(raw)
			if !valid {
				return nil
			}

			for _, value := range metadata {
				var text string

				err := json.Unmarshal(value, &text)
				if err != nil || len(value) == 0 || value[0] != '"' {
					return nil
				}
			}

			continue
		}

		descriptor, valid := tensorJSONObject(raw)
		if !valid {
			return nil
		}

		var (
			dtype   string
			shape   []uint64
			offsets []uint64
		)

		err := json.Unmarshal(descriptor["dtype"], &dtype)
		if err != nil {
			return nil
		}

		err = json.Unmarshal(descriptor["shape"], &shape)
		if err != nil || shape == nil || bytes.Contains(descriptor["shape"], []byte("null")) {
			return nil
		}

		err = json.Unmarshal(descriptor["data_offsets"], &offsets)
		if err != nil || len(offsets) != 2 || bytes.Contains(descriptor["data_offsets"], []byte("null")) || offsets[1] < offsets[0] {
			return nil
		}

		bits := tensorDtypeBits(dtype)
		if bits == 0 {
			return nil
		}

		empty := false

		for _, dimension := range shape {
			if dimension == 0 {
				empty = true

				continue
			}

			if dimension > ^uint64(0)/bits {
				return nil
			}

			bits *= dimension
		}

		if empty {
			bits = 0
		}

		if bits%8 != 0 || bits/8 != offsets[1]-offsets[0] {
			return nil
		}

		extents = append(extents, tensorExtent{start: offsets[0], end: offsets[1]})
	}

	slices.SortFunc(extents, compareTensorExtents)
	end := uint64(0)

	for _, extent := range extents {
		if extent.start != end {
			return nil
		}

		end = extent.end
	}

	if end != uint64(size)-8-headerSize {
		return nil
	}

	return &types.Metadata{Kind: types.KindSafeTensors}
}

func tensorJSONObject(data []byte) (map[string]json.RawMessage, bool) {
	decoder := json.NewDecoder(bytes.NewReader(data))

	token, err := decoder.Token()
	if err != nil || token != json.Delim('{') {
		return nil, false
	}

	fields := make(map[string]json.RawMessage)

	for decoder.More() {
		token, err = decoder.Token()
		if err != nil {
			return nil, false
		}

		name, ok := token.(string)
		if !ok {
			return nil, false
		}

		if _, exists := fields[name]; exists {
			return nil, false
		}

		var value json.RawMessage

		err = decoder.Decode(&value)
		if err != nil {
			return nil, false
		}

		fields[name] = value
	}

	token, err = decoder.Token()
	if err != nil || token != json.Delim('}') {
		return nil, false
	}

	token, err = decoder.Token()
	return fields, errors.Is(err, io.EOF) && token == nil
}

func tensorDtypeBits(dtype string) uint64 {
	switch dtype {
	case "F4":
		return 4
	case "F6_E3M2", "F6_E2M3":
		return 6
	case "BOOL", "U8", "I8", "F8_E5M2", "F8_E4M3", "F8_E8M0", "F8_E4M3FNUZ", "F8_E5M2FNUZ":
		return 8
	case "I16", "U16", "F16", "BF16":
		return 16
	case "I32", "U32", "F32":
		return 32
	case "I64", "U64", "F64", "C64":
		return 64
	}

	return 0
}

func compareTensorExtents(left, right tensorExtent) int {
	if left.start < right.start || left.start == right.start && left.end < right.end {
		return -1
	}

	if left == right {
		return 0
	}

	return 1
}
