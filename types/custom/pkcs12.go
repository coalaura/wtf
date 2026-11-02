package custom

import "github.com/coalaura/wtf/types"

func DetectPKCS12(buffer types.Buffer) *types.Metadata {
	return detectPKCS12Sized(buffer, int64(buffer.Len()))
}

func detectPKCS12Sized(buffer types.Buffer, size int64) *types.Metadata {
	offset, end, ok := asn1Header(buffer, 0, 0x30, size)
	if !ok || end != size {
		return nil
	}

	version, versionEnd, ok := asn1Header(buffer, offset, 2, end)
	if !ok || versionEnd != int64(version)+1 || !buffer.Has(version, []byte{3}) {
		return nil
	}

	content, contentEnd, ok := asn1Header(buffer, version+1, 0x30, end)
	if !ok {
		return nil
	}

	oid, oidEnd, ok := asn1Header(buffer, content, 6, contentEnd)
	if !ok || oidEnd != int64(oid)+9 || !buffer.Has(oid, []byte{0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d, 1, 7, 1}) {
		return nil
	}

	_, _, ok = asn1Header(buffer, oid+9, 0xa0, contentEnd)
	if !ok {
		return nil
	}

	return &types.Metadata{Kind: types.KindPKCS12}
}

func asn1Header(buffer types.Buffer, offset int, tag byte, limit int64) (int, int64, bool) {
	if offset < 0 || offset > buffer.Len()-2 || buffer[offset] != tag {
		return 0, 0, false
	}

	length := int64(buffer[offset+1])
	header := 2

	if length&0x80 != 0 {
		count := int(length & 0x7f)
		if count == 0 || count > 4 || count > buffer.Len()-offset-2 || buffer[offset+2] == 0 {
			return 0, 0, false
		}

		length = 0

		for index := range count {
			length = length<<8 | int64(buffer[offset+2+index])
		}

		if length < 128 {
			return 0, 0, false
		}

		header += count
	}

	start := offset + header
	if int64(start) > limit || length > limit-int64(start) {
		return 0, 0, false
	}

	return start, int64(start) + length, true
}
