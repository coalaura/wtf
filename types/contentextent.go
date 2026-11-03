package types

import "encoding/binary"

// Header offsets refer to the whole file, not the bounded detection prefix.
func validContentExtent(b Buffer, kind KindID, size uint64) bool {
	switch kind {
	case KindGitMultiPackIndex:
		hashSize := uint64(20)

		if b[5] == 2 {
			hashSize = 32
		}

		return size >= 12+(uint64(b[6])+1)*12+hashSize
	case KindGitReftable:
		return size >= 92+uint64(b[4]-1)*8
	case KindGitCruftPackMtimes:
		hashSize := uint64(20)

		if binary.BigEndian.Uint32(b[8:]) == 2 {
			hashSize = 32
		}

		return size >= 12+2*hashSize && size%4 == 0
	case KindGoCoverageMetadata:
		if b.Len() < 56 {
			return false
		}

		length := binary.LittleEndian.Uint64(b[8:])

		return length >= 56 && length <= size
	case KindGoCoverageCounters:
		return b.Len() >= 32 && (b[24] == 1 || b[24] == 2) && b[25] <= 1
	case KindLLVMRawProfile:
		if b.Len() < 64 {
			return false
		}

		var order binary.ByteOrder = binary.LittleEndian

		if b[0] == 0xff {
			order = binary.BigEndian
		}

		// The high version bits are independently defined instrumentation flags.
		version := order.Uint64(b[8:]) & 0xffffffff

		return version >= 1 && version <= 11
	case KindLLVMIndexedProfile:
		if b.Len() < 40 {
			return false
		}

		version := binary.LittleEndian.Uint64(b[8:]) & 0xffffffff

		return version >= 1 && version <= 14
	case KindFlattenedDevicetree, KindAndroidDTTable:
		return uint64(binary.BigEndian.Uint32(b[4:])) <= size
	case KindAndroidVBMeta:
		authSize := binary.BigEndian.Uint64(b[12:])
		auxSize := binary.BigEndian.Uint64(b[20:])

		return contentExtentFits(256, authSize, size) && contentExtentFits(256+authSize, auxSize, size)
	case KindUpdateEnginePayload:
		headerSize := uint64(20)
		signatureSize := uint64(0)

		if binary.BigEndian.Uint64(b[4:]) == 2 {
			headerSize = 24
			signatureSize = uint64(binary.BigEndian.Uint32(b[20:]))
		}

		manifestSize := binary.BigEndian.Uint64(b[12:])

		return contentExtentFits(headerSize, manifestSize, size) && contentExtentFits(headerSize+manifestSize, signatureSize, size)
	case KindUF2Firmware:
		return size%512 == 0
	case KindBPFTypeFormat:
		var order binary.ByteOrder = binary.LittleEndian

		if b[0] == 0xeb {
			order = binary.BigEndian
		}

		headerSize := uint64(order.Uint32(b[4:]))
		typesEnd := headerSize + uint64(order.Uint32(b[8:])) + uint64(order.Uint32(b[12:]))
		stringsEnd := headerSize + uint64(order.Uint32(b[16:])) + uint64(order.Uint32(b[20:]))

		if headerSize >= 32 {
			if b.Len() < 32 || headerSize+uint64(order.Uint32(b[24:]))+uint64(order.Uint32(b[28:])) > size {
				return false
			}
		}

		return typesEnd <= size && stringsEnd <= size
	case KindSQLiteRollbackJournal:
		return validRollbackJournalExtent(b, size)
	case KindBSDIFFPatch:
		if b.Len() < 32 || b[15]&0x80 != 0 || b[23]&0x80 != 0 || b[31]&0x80 != 0 {
			return false
		}

		controlSize := binary.LittleEndian.Uint64(b[8:])
		diffSize := binary.LittleEndian.Uint64(b[16:])

		return contentExtentFits(32, controlSize, size) && contentExtentFits(32+controlSize, diffSize, size)
	case KindAsepriteSprite:
		return uint64(binary.LittleEndian.Uint32(b)) <= size
	case KindQOAAudio:
		return 8+uint64(binary.BigEndian.Uint16(b[14:])) <= size
	case KindFlatGeobuf:
		if b.Len() < 12 {
			return false
		}

		headerSize := uint64(binary.LittleEndian.Uint32(b[8:]))

		return headerSize >= 8 && contentExtentFits(12, headerSize, size)
	case KindUCSCBigWig, KindUCSCBigBed:
		var order binary.ByteOrder = binary.LittleEndian

		if b[0] == 0x88 || b[0] == 0x87 {
			order = binary.BigEndian
		}

		return order.Uint64(b[8:]) < size && order.Uint64(b[16:]) < size && order.Uint64(b[24:]) < size
	case KindUCSCTwoBit:
		var order binary.ByteOrder = binary.LittleEndian

		if b[0] == 0x1a {
			order = binary.BigEndian
		}

		entrySize := uint64(5 + 4*order.Uint32(b[4:]))

		return uint64(order.Uint32(b[8:])) <= (size-16)/entrySize
	case KindBatched3DModel, KindInstanced3DModel, KindPointCloudTile, KindCompositeTile:
		return uint64(binary.LittleEndian.Uint32(b[8:])) <= size
	}

	return true
}

func validRollbackJournalExtent(b Buffer, size uint64) bool {
	if b.Len() < 28 {
		return false
	}

	pages := uint64(binary.BigEndian.Uint32(b[8:]))

	sectorSize := uint64(binary.BigEndian.Uint32(b[20:]))
	pageSize := uint64(binary.BigEndian.Uint32(b[24:]))

	if sectorSize < 32 || sectorSize > 65536 || sectorSize&(sectorSize-1) != 0 || pageSize < 512 || pageSize > 65536 || pageSize&(pageSize-1) != 0 || sectorSize > size {
		return false
	}

	return pages == 0xffffffff || pages <= (size-sectorSize)/(pageSize+8)
}

func contentExtentFits(offset, length, size uint64) bool {
	return offset <= size && length <= size-offset
}
