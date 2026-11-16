package custom

import (
	"bytes"
	"encoding/binary"
	"math/bits"

	"github.com/coalaura/wtf/types"
)

type pdbStream struct {
	name   string
	offset uint64
	size   uint64
}

func DetectPortablePDB(buffer types.Buffer) *types.Metadata {
	return detectPortablePDBSized(buffer, int64(buffer.Len()))
}

func detectPortablePDBSized(buffer types.Buffer, size int64) *types.Metadata {
	if buffer.Len() < 24 || !buffer.Has(0, []byte("BSJB")) || !buffer.Has(4, []byte{1, 0, 1, 0}) {
		return nil
	}

	versionSize := uint64(binary.LittleEndian.Uint32(buffer[12:]))
	if versionSize == 0 || versionSize > 256 || versionSize%4 != 0 || 20+versionSize > uint64(buffer.Len()) {
		return nil
	}

	position := 16 + versionSize

	count := int(binary.LittleEndian.Uint16(buffer[position+2:]))
	if count < 2 || count > 16 {
		return nil
	}

	position += 4

	var streams [16]pdbStream

	for index := range count {
		if position+12 > uint64(buffer.Len()) {
			return nil
		}

		stream := &streams[index]
		stream.offset = uint64(binary.LittleEndian.Uint32(buffer[position:]))
		stream.size = uint64(binary.LittleEndian.Uint32(buffer[position+4:]))
		position += 8
		nameLimit := min(position+32, uint64(buffer.Len()))

		nameLength := bytes.IndexByte(buffer[position:nameLimit], 0)
		if nameLength < 0 {
			return nil
		}

		stream.name = string(buffer[position : position+uint64(nameLength)])
		position += (uint64(nameLength) + 4) &^ 3
	}

	var (
		debug  bool
		tables bool
	)

	for index := range count {
		stream := &streams[index]
		if stream.offset < position || stream.offset%4 != 0 || stream.offset+stream.size > uint64(size) {
			return nil
		}

		for previous := range index {
			other := &streams[previous]
			overlap := stream.offset < other.offset+other.size && other.offset < stream.offset+stream.size

			if stream.name == other.name || overlap {
				return nil
			}
		}

		switch stream.name {
		case "#Pdb":
			if stream.size < 32 || stream.offset+32 > uint64(buffer.Len()) {
				return nil
			}

			mask := binary.LittleEndian.Uint64(buffer[stream.offset+24:])
			debug = stream.size == 32+4*uint64(bits.OnesCount64(mask))
		case "#~":
			if stream.size < 24 || stream.offset+24 > uint64(buffer.Len()) {
				return nil
			}

			header := buffer[stream.offset:]
			mask := binary.LittleEndian.Uint64(header[8:])
			tables = header[4] == 2 && header[5] == 0 && stream.size >= 24+4*uint64(bits.OnesCount64(mask))
		}
	}

	if !debug || !tables {
		return nil
	}

	return &types.Metadata{Kind: types.KindPortablePDB}
}
