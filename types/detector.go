//go:generate go run ../gen/optimizer

package types

import (
	"bytes"
	"io"
)

type Detector interface {
	Detect(Buffer) *Metadata
}

type DetectFunc func(Buffer) *Metadata

func (df DetectFunc) Detect(b Buffer) *Metadata {
	return df(b)
}

func Detect(name string, data []byte) (*Metadata, error) {
	return DetectReaderAt(name, data, bytes.NewReader(data), int64(len(data)))
}

// DetectReaderAt detects a file from a bounded prefix and refines containers
// using metadata read at their real file offsets.
func DetectReaderAt(name string, data []byte, r io.ReaderAt, size int64) (*Metadata, error) {
	if size < 0 {
		return nil, io.ErrUnexpectedEOF
	}

	var meta *Metadata

	prefixSize := int(min(size, int64(maxScanSize)))
	if len(data) < prefixSize {
		data = make([]byte, prefixSize)
		if !readAtExactly(r, data, 0) {
			return nil, io.ErrUnexpectedEOF
		}
	}

	buffer := Buffer(data[:prefixSize])
	if buffer.Has(0, []byte("PK\x03\x04")) || buffer.Has(0, []byte("PK\x05\x06")) {
		meta = detectZIPReaderAt(r, size)
		if meta != nil && meta.Kind == KindZIPArchive && meta.Type == TypeNone {
			meta = detectZIPEntries(buffer)
		}
	} else if buffer.Has(0, []byte{0xd0, 0xcf, 0x11, 0xe0, 0xa1, 0xb1, 0x1a, 0xe1}) {
		meta = detectOLEReaderAt(r, size)
	}

	if meta == nil && size >= 48 {
		tailSize := min(size, 512)

		var tail []byte

		if int64(len(data)) >= size {
			tail = data[size-tailSize : size]
		} else {
			tail = make([]byte, tailSize)
			if !readAtExactly(r, tail, size-tailSize) {
				return nil, io.ErrUnexpectedEOF
			}
		}

		meta = detectFileTrailerSized(Buffer(tail), r, size)
	}

	if meta == nil {
		meta = detectPKCS12Sized(buffer, size)
	}

	if meta == nil {
		meta = DetectTIFFReaderAt(buffer, r, size)
	}

	if meta == nil {
		meta = DetectISOReaderAt(buffer, r, size)
	}

	if meta == nil {
		meta = detectStrong(buffer)
	}

	if meta == nil {
		meta = DetectHDF5ReaderAt(buffer, r, size)
	}

	if meta == nil {
		codec, tagged := SkipID3ReaderAt(buffer, r, size)
		if tagged {
			meta = RefineID3Metadata(detectOptimized(codec))
		} else {
			meta = detectWeak(buffer)
		}
	}

	// A complete-looking prefix does not validate the unread remainder of JSON.
	if meta != nil && size > int64(buffer.Len()) && (meta.Kind == KindJSONDocument || meta.Kind == KindGameMakerStudioProject && meta.Type == TypeGameMakerStudio2) {
		meta = &Metadata{Kind: KindText, Type: TypeUTF8, Confidence: ConfidenceLow}
	}

	if meta != nil {
		meta.File = name

		return meta, nil
	}

	return nil, ErrUnknownFormat
}
