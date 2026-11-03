package internal

import "github.com/coalaura/wtf/types"

func init() {
	types.RegisterSignature(types.KindGoCoverageCounters, types.TypeNone, 0, []byte("\x00cwm\x01\x00\x00\x00"))
	types.RegisterSignature(types.KindGoCoverageMetadata, types.TypeNone, 0, []byte("\x00cvm\x01\x00\x00\x00"))
	types.RegisterSignature(types.KindGoExecutionTrace, types.TypeNone, 0, []byte("go 1.11 trace\x00\x00\x00"))
	types.RegisterSignature(types.KindGoExecutionTrace, types.TypeNone, 0, []byte("go 1.19 trace\x00\x00\x00"))
	types.RegisterSignature(types.KindGoExecutionTrace, types.TypeNone, 0, []byte("go 1.21 trace\x00\x00\x00"))
	types.RegisterSignature(types.KindGoExecutionTrace, types.TypeNone, 0, []byte("go 1.22 trace\x00\x00\x00"))
	types.RegisterSignature(types.KindGoExecutionTrace, types.TypeNone, 0, []byte("go 1.23 trace\x00\x00\x00"))
	types.RegisterSignature(types.KindGoExecutionTrace, types.TypeNone, 0, []byte("go 1.25 trace\x00\x00\x00"))
	types.RegisterSignature(types.KindGoExecutionTrace, types.TypeNone, 0, []byte("go 1.26 trace\x00\x00\x00"))
	types.RegisterSignature(types.KindGoExecutionTrace, types.TypeNone, 0, []byte("go 1.5 trace\x00\x00\x00"))
	types.RegisterSignature(types.KindGoExecutionTrace, types.TypeNone, 0, []byte("go 1.7 trace\x00\x00\x00"))
	types.RegisterSignature(types.KindGoExecutionTrace, types.TypeNone, 0, []byte("go 1.9 trace\x00\x00\x00"))
	types.RegisterSignature(types.KindLLVMIndexedProfile, types.TypeNone, 0, []byte("\xfflprofi\x81"))
	types.RegisterSignature(types.KindLLVMRawProfile, types.Type32BitBigEndian, 0, []byte("\xfflprofR\x81"))
	types.RegisterSignature(types.KindLLVMRawProfile, types.Type32BitLittleEndian, 0, []byte("\x81Rforpl\xff"))
	types.RegisterSignature(types.KindLLVMRawProfile, types.Type64BitBigEndian, 0, []byte("\xfflprofr\x81"))
	types.RegisterSignature(types.KindLLVMRawProfile, types.Type64BitLittleEndian, 0, []byte("\x81rforpl\xff"))
}
