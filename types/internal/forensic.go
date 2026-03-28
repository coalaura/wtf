package internal

import "github.com/coalaura/wtf/types"

func init() {
	types.RegisterSignature(types.KindAccessDataFTK, types.TypeNone, 0, []byte{0xa9, 0x0d, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
	types.RegisterSignature(types.KindAdvancedForensicFormat, types.TypeNone, 0, []byte("AFF\x00"))
	types.RegisterSignature(types.KindBluetoothSnoop, types.TypeNone, 0, []byte("btsnoop\x00"))
	types.RegisterSignature(types.KindEnCaseCaseFile, types.TypeNone, 0, []byte("_CASE_"))
	types.RegisterSignature(types.KindEnCaseEvidenceV2, types.TypeNone, 0, []byte("EVF2\r\n\x81"))
	types.RegisterSignature(types.KindEnCaseImage, types.TypeNone, 0, []byte("EVF\x09\x0d\x0a\xff\x00"))
	types.RegisterSignature(types.KindExpertWitnessFormat, types.TypeNone, 0, []byte("EVF\x09\x0d\x0a\x81"))
	types.RegisterSignature(types.KindLogicalFileEvidence, types.TypeNone, 0, []byte("LVF\x09\x0d\x0a\xff\x00"))
	types.RegisterSignature(types.KindMCAPCapture, types.TypeNone, 0, []byte{0x89, 'M', 'C', 'A', 'P', '0', '\r', '\n'})
	types.RegisterSignature(types.KindMediaDescriptor, types.TypeNone, 0, []byte("MEDIA DESCRIPTOR"))
	types.RegisterSignature(types.KindMicrosoftNetworkMonitor, types.TypeNone, 0, []byte("MACROSOFT\x00"))
	types.RegisterSignature(types.KindMicrosoftNetworkMonitor, types.TypeNone, 0, []byte("NetMon\x00\x00"))
	types.RegisterSignature(types.KindPowerISODAA, types.TypeNone, 0, []byte("DAA\x00\x00\x00\x00\x00"))
	types.RegisterSignature(types.KindPowerISODAA, types.TypeNone, 0, []byte("DAA\x1a"))
	types.RegisterSignature(types.KindROSBag, types.TypeNone, 0, []byte("#ROSBAG V2.0\n"))
	types.RegisterSignature(types.KindSnoopCapture, types.TypeNone, 0, []byte("snoop\x00\x00\x00"))
	types.RegisterSignature(types.KindWinPharoahCapture, types.TypeNone, 0, []byte{0x1a, 0x35, 0x01, 0x00})
	types.RegisterSignature(types.KindWinPharoahFilter, types.TypeNone, 0, []byte{0xd2, 0x0a, 0x00, 0x00})

	types.RegisterMaskedSignature(types.KindPCAPNGCapture, types.TypeNone, 0, []byte{0x0a, 0x0d, 0x0d, 0x0a, 0x00, 0x00, 0x00, 0x00, 0x1a, 0x2b, 0x3c, 0x4d}, []byte{0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff})
	types.RegisterMaskedSignature(types.KindPCAPNGCapture, types.TypeNone, 0, []byte{0x0a, 0x0d, 0x0d, 0x0a, 0x00, 0x00, 0x00, 0x00, 0x4d, 0x3c, 0x2b, 0x1a}, []byte{0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff})
}
