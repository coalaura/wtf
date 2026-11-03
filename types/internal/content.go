package internal

import (
	"github.com/coalaura/wtf/types"
	"github.com/coalaura/wtf/types/custom"
)

func init() {
	types.RegisterSignature(types.KindAgeEncryptedFile, types.TypeNone, 0, []byte("age-encryption.org/v1\n"))
	types.RegisterSignature(types.KindBSDIFFPatch, types.TypeNone, 0, []byte("BSDIFF40"))
	types.RegisterSignature(types.KindBtrfsSendStream, types.TypeNone, 0, []byte("btrfs-stream\x00\x01\x00\x00\x00"))
	types.RegisterSignature(types.KindBtrfsSendStream, types.TypeNone, 0, []byte("btrfs-stream\x00\x02\x00\x00\x00"))
	types.RegisterSignature(types.KindBtrfsSendStream, types.TypeNone, 0, []byte("btrfs-stream\x00\x03\x00\x00\x00"))
	types.RegisterSignature(types.KindSQLiteRollbackJournal, types.TypeNone, 0, []byte("\xd9\xd5\x05\xf9\x20\xa1\x63\xd7"))

	types.RegisterMaskedSignature(types.KindFlatGeobuf, types.TypeNone, 0, []byte("fgb\x03fgb\x00"), []byte("\xff\xff\xff\xff\xff\xff\xff\x00"))
	types.RegisterMaskedSignature(types.KindVCDIFFDelta, types.TypeNone, 0, []byte("\xd6\xc3\xc4\x00\x00"), []byte("\xff\xff\xff\xff\xf8"))
	types.RegisterMaskedSignature(types.KindVCDIFFDelta, types.TypeNone, 0, []byte("\xd6\xc3\xc4S\x00"), []byte("\xff\xff\xff\xff\xf8"))

	types.Register(types.DetectFunc(custom.DetectContentHeaders))
}
