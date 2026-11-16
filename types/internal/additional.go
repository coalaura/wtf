package internal

import "github.com/coalaura/wtf/types"

func init() {
	types.RegisterSignature(types.KindBAMIndex, types.TypeNone, 0, []byte{'B', 'A', 'I', 1})
	types.RegisterSignature(types.KindDwarFSImage, types.TypeNone, 0, []byte("DWARFS\x02"))
	types.RegisterSignature(types.KindJavaFlightRecorder, types.TypeNone, 0, []byte("FLR\x00"))
	types.RegisterSignature(types.KindJavaHeapDump, types.TypeNone, 0, []byte("JAVA PROFILE 1.0.1\x00"))
	types.RegisterSignature(types.KindJavaHeapDump, types.TypeNone, 0, []byte("JAVA PROFILE 1.0.2\x00"))
	types.RegisterSignature(types.KindLLVMOptimizationRemarks, types.TypeNone, 0, []byte("RMRK"))
	types.RegisterSignature(types.KindLuaJITBytecode, types.TypeNone, 0, []byte{0x1b, 'L', 'J', 1})
	types.RegisterSignature(types.KindLuaJITBytecode, types.TypeNone, 0, []byte{0x1b, 'L', 'J', 2})
	types.RegisterSignature(types.KindNASACDF, types.TypeNone, 0, []byte{0xcd, 0xf2, 0x60, 2, 0, 0, 0xff, 0xff})
	types.RegisterSignature(types.KindNASACDF, types.TypeNone, 0, []byte{0xcd, 0xf2, 0x60, 2, 0xcc, 0xcc, 0, 1})
	types.RegisterSignature(types.KindNASACDF, types.TypeNone, 0, []byte{0xcd, 0xf3, 0, 1, 0, 0, 0xff, 0xff})
	types.RegisterSignature(types.KindNASACDF, types.TypeNone, 0, []byte{0xcd, 0xf3, 0, 1, 0xcc, 0xcc, 0, 1})
	types.RegisterSignature(types.KindNixArchive, types.TypeNone, 0, []byte("\x0d\x00\x00\x00\x00\x00\x00\x00nix-archive-1\x00\x00\x00"))
	types.RegisterSignature(types.KindOpenZIMArchive, types.TypeNone, 0, []byte{'Z', 'I', 'M', 4})
	types.RegisterSignature(types.KindPAR2Recovery, types.TypeNone, 0, []byte("PAR2\x00PKT"))
	types.RegisterSignature(types.KindPostgreSQLBinaryCopy, types.TypeNone, 0, []byte("PGCOPY\n\xff\r\n\x00"))
	types.RegisterSignature(types.KindRDFHDTDataset, types.TypeNone, 0, []byte("$HDT\x01"))
	types.RegisterSignature(types.KindStandardFlowgram, types.TypeNone, 0, []byte{'.', 's', 'f', 'f', 0, 0, 0, 1})
	types.RegisterSignature(types.KindUBIImage, types.TypeNone, 0, []byte("UBI#\x01"))
	types.RegisterSignature(types.KindWebAssemblyComponent, types.TypeNone, 0, []byte{0, 'a', 's', 'm', 0x0d, 0, 1, 0})
}
