package internal

import "github.com/coalaura/wtf/types"

func init() {
	types.RegisterSignature(types.KindAccessDataFTKEvidence, types.TypeNone, 0, []byte{0xa9, 0x0d, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
	types.RegisterSignature(types.KindAdvancedForensicFormat, types.TypeNone, 0, []byte("AFF\x00"))
	types.RegisterSignature(types.KindBluetoothSnoopCapture, types.TypeNone, 0, []byte("btsnoop\x00"))
	types.RegisterSignature(types.KindEnCaseCase, types.TypeNone, 0, []byte("_CASE_"))
	types.RegisterSignature(types.KindEnCaseEvidenceV2, types.TypeNone, 0, []byte("EVF2\r\n\x81"))
	types.RegisterSignature(types.KindEnCaseImage, types.TypeNone, 0, []byte("EVF\x09\x0d\x0a\xff\x00"))
	types.RegisterSignature(types.KindExpertWitnessDiskImage, types.TypeNone, 0, []byte("EVF\x09\x0d\x0a\x81"))
	types.RegisterSignature(types.KindLogicalFileEvidence, types.TypeNone, 0, []byte("LVF\x09\x0d\x0a\xff\x00"))
	types.RegisterSignature(types.KindMacOSFSEventsLog, types.TypeNone, 0, []byte("1SLD"))
	types.RegisterSignature(types.KindMacOSFSEventsLog, types.TypeNone, 0, []byte("2SLD"))
	types.RegisterSignature(types.KindMCAPCapture, types.TypeNone, 0, []byte{0x89, 'M', 'C', 'A', 'P', '0', '\r', '\n'})
	types.RegisterSignature(types.KindMediaDescriptor, types.TypeNone, 0, []byte("MEDIA DESCRIPTOR"))
	types.RegisterSignature(types.KindMicrosoftNetworkMonitorCapture, types.TypeNone, 0, []byte("MACROSOFT\x00"))
	types.RegisterSignature(types.KindMicrosoftNetworkMonitorCapture, types.TypeNone, 0, []byte("NetMon\x00\x00"))
	types.RegisterSignature(types.KindPacketSnifferXCPCapture, types.TypeNone, 0, []byte("XCP\x00"))
	types.RegisterSignature(types.KindROSBag, types.TypeNone, 0, []byte("#ROSBAG V2.0\n"))
	types.RegisterSignature(types.KindSnoopCapture, types.TypeNone, 0, []byte("snoop\x00\x00\x00"))
	types.RegisterSignature(types.KindWinPharaohCapture, types.TypeNone, 0, []byte{0x1a, 0x35, 0x01, 0x00})
	types.RegisterSignature(types.KindWinPharaohFilter, types.TypeNone, 0, []byte{0xd2, 0x0a, 0x00, 0x00})

	types.RegisterMaskedSignature(types.KindPCAPNGCapture, types.TypeNone, 0, []byte{0x0a, 0x0d, 0x0d, 0x0a, 0x00, 0x00, 0x00, 0x00, 0x1a, 0x2b, 0x3c, 0x4d}, []byte{0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff})
	types.RegisterMaskedSignature(types.KindPCAPNGCapture, types.TypeNone, 0, []byte{0x0a, 0x0d, 0x0d, 0x0a, 0x00, 0x00, 0x00, 0x00, 0x4d, 0x3c, 0x2b, 0x1a}, []byte{0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff})

	types.RegisterWeakSignature(types.KindCellebriteUFED, types.TypeNone, 0, []byte("[Cellebrite]"))
	types.RegisterWeakSignature(types.KindCellebriteUFED, types.TypeNone, 0, []byte("[Dump]"))
}
