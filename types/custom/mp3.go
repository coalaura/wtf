package custom

import "github.com/coalaura/wtf/types"

func DetectMP3(b types.Buffer) *types.Metadata {
	if b.Len() < 4 {
		return nil
	}

	if b[0] != 0xff || (b[1]&0xe0) != 0xe0 {
		return nil
	}

	version := (b[1] >> 3) & 0x03
	layer := (b[1] >> 1) & 0x03
	bitrate := (b[2] >> 4) & 0x0f
	sampleRate := (b[2] >> 2) & 0x03

	if version == 1 || layer == 0 || bitrate == 0 || bitrate == 0x0f || sampleRate == 0x03 {
		return nil
	}

	if layer == 1 {
		return &types.Metadata{Kind: types.KindMPEGAudio, Type: types.TypeMPEGLayer3, Confidence: types.ConfidenceMedium}
	}

	if layer == 2 {
		return &types.Metadata{Kind: types.KindMPEGAudio, Type: types.TypeMPEGLayer2, Confidence: types.ConfidenceMedium}
	}

	if layer == 3 {
		return &types.Metadata{Kind: types.KindMPEGAudio, Type: types.TypeMPEGLayer1, Confidence: types.ConfidenceMedium}
	}

	return nil
}
