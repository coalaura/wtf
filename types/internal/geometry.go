package internal

import "github.com/coalaura/wtf/types"

func init() {
	types.RegisterSignature(types.KindAC3DModel, types.TypeAC3DVersionB, 0, []byte("AC3Db"))
	types.RegisterSignature(types.KindAC3DModel, types.TypeAC3DVersionC, 0, []byte("AC3Dc"))
	types.RegisterSignature(types.KindE57PointCloud, types.TypeNone, 0, []byte("ASTM-E57\x01\x00\x00\x00\x00\x00\x00\x00"))
	types.RegisterSignature(types.KindNanoVDB, types.TypeNanoVDBFile, 0, []byte("NanoVDB2"))
	types.RegisterSignature(types.KindNanoVDB, types.TypeNanoVDBGrid, 0, []byte("NanoVDB1"))
	types.RegisterSignature(types.KindNanoVDB, types.TypeNanoVDBLegacy, 0, []byte("NanoVDB0"))
	types.RegisterSignature(types.KindOpenCTMMesh, types.TypeNone, 0, []byte("OCTM\x05\x00\x00\x00"))
	types.RegisterSignature(types.KindPrusaBinaryGCode, types.TypeNone, 0, []byte("GCDE\x01\x00\x00\x00"))
	types.RegisterSignature(types.KindPtexTexture, types.TypeNone, 0, []byte("Ptex\x01\x00\x00\x00"))
	types.RegisterSignature(types.KindXaraDrawing, types.TypeNone, 0, []byte("XARA\xa3\xa3\r\n"))

	types.RegisterMaskedSignature(types.KindIFFContainer, types.TypeLightWaveLWO2, 0, []byte("FORM\x00\x00\x00\x00LWO2"), []byte("\xff\xff\xff\xff\x00\x00\x00\x00\xff\xff\xff\xff"))
	types.RegisterMaskedSignature(types.KindIFFContainer, types.TypeLightWaveLWO3, 0, []byte("FORM\x00\x00\x00\x00LWO3"), []byte("\xff\xff\xff\xff\x00\x00\x00\x00\xff\xff\xff\xff"))
	types.RegisterMaskedSignature(types.KindIFFContainer, types.TypeLightWaveLWOB, 0, []byte("FORM\x00\x00\x00\x00LWOB"), []byte("\xff\xff\xff\xff\x00\x00\x00\x00\xff\xff\xff\xff"))
	types.RegisterMaskedSignature(types.KindIFFContainer, types.TypeModoLXO, 0, []byte("FORM\x00\x00\x00\x00LXOB"), []byte("\xff\xff\xff\xff\x00\x00\x00\x00\xff\xff\xff\xff"))
	types.RegisterMaskedSignature(types.KindX11Cursor, types.TypeNone, 0, []byte("Xcur\x00\x00\x00\x00\x00\x00\x01\x00"), []byte("\xff\xff\xff\xff\x00\x00\x00\x00\xff\xff\xff\xff"))
}
