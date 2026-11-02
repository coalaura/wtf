package main

import (
	"archive/zip"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/coalaura/wtf/types"
)

type testCase struct {
	file string
	kind types.KindID
	typ  types.TypeID
}

type collisionCase struct {
	name      string
	data      []byte
	kind      types.KindID
	typ       types.TypeID
	forbidden types.KindID
}

func TestConsistency(t *testing.T) {
	typeNames := getIdentifiers(t, "types/ids_type.go", "Type")
	kindNames := getIdentifiers(t, "types/ids_kind.go", "Kind")

	for i := types.TypeID(1); i < types.MaxTypes; i++ {
		if i.String() == "" {
			name := "Unknown Identifier"

			if int(i) < len(typeNames) {
				name = typeNames[i]
			}

			t.Errorf("Type %s (%d) is missing a label in typeNames array", name, i)
		}
	}

	for i := types.KindID(1); i < types.MaxKinds; i++ {
		if i.String() == "Unknown" {
			name := "Unknown Identifier"

			if int(i) < len(kindNames) {
				name = kindNames[i]
			}

			t.Errorf("Kind %s (%d) is missing a label in kindNames array", name, i)
		}
	}
}

func TestDetectFixtures(t *testing.T) {
	cases := []testCase{
		{file: "3g2.3g2", kind: types.KindISOBaseMedia, typ: types.Type3G2Video},
		{file: "3gp.3gp", kind: types.KindISOBaseMedia, typ: types.Type3GPPVideo},
		{file: "7z.7z", kind: types.Kind7ZipArchive, typ: types.TypeNone},
		{file: "aifc.aif", kind: types.KindIFFContainer, typ: types.TypeAIFCAudio},
		{file: "android.apk", kind: types.KindZIPArchive, typ: types.TypeAndroidPackage},
		{file: "archive.zst", kind: types.KindZstandardArchive, typ: types.TypeNone},
		{file: "arw.arw", kind: types.KindTIFFImage, typ: types.TypeSonyRAW},
		{file: "audio.flac", kind: types.KindFLACAudio, typ: types.TypeNone},
		{file: "avi.avi", kind: types.KindRIFFContainer, typ: types.TypeAVIVideo},
		{file: "avif.avif", kind: types.KindISOBaseMedia, typ: types.TypeAVIFImageSequence},
		{file: "big_endian.tif", kind: types.KindTIFFImage, typ: types.TypeBigEndian},
		{file: "bigtiff.tif", kind: types.KindTIFFImage, typ: types.TypeBigTIFF},
		{file: "bundle.aab", kind: types.KindZIPArchive, typ: types.TypeAndroidAppBundle},
		{file: "comic.cbz", kind: types.KindZIPArchive, typ: types.TypeComicBook},
		{file: "cr2.cr2", kind: types.KindCanonRAWImage, typ: types.TypeCanonRAW2},
		{file: "cursor.cur", kind: types.KindWindowsCursorImage, typ: types.TypeNone},
		{file: "data.json", kind: types.KindJSONDocument, typ: types.TypeNone},
		{file: "dicom.dcm", kind: types.KindDICOMMedicalImage, typ: types.TypeNone},
		{file: "dng.dng", kind: types.KindTIFFImage, typ: types.TypeAdobeDNG},
		{file: "document.doc", kind: types.KindOLECompoundDocument, typ: types.TypeMicrosoftWordDocument},
		{file: "dp3_icc.icc", kind: types.KindICCProfile, typ: types.TypeNone},
		{file: "elf_x64", kind: types.KindExecutableAndLinkableFormat, typ: types.Type64BitLittleEndian},
		{file: "elf_x86", kind: types.KindExecutableAndLinkableFormat, typ: types.Type32BitLittleEndian},
		{file: "epub.epub", kind: types.KindZIPArchive, typ: types.TypeEPUBDocument},
		{file: "excel.xls", kind: types.KindOLECompoundDocument, typ: types.TypeMicrosoftExcelWorkbook},
		{file: "excel.xlsx", kind: types.KindZIPArchive, typ: types.TypeMicrosoftExcelWorkbook},
		{file: "gif87a.gif", kind: types.KindGIFImage, typ: types.TypeGIF87a},
		{file: "gif89a.gif", kind: types.KindGIFImage, typ: types.TypeGIF89a},
		{file: "gz.gz", kind: types.KindGzipArchive, typ: types.TypeNone},
		{file: "heif.heif", kind: types.KindISOBaseMedia, typ: types.TypeHEIFImage},
		{file: "ico.ico", kind: types.KindWindowsIconImage, typ: types.TypeNone},
		{file: "installer.msi", kind: types.KindOLECompoundDocument, typ: types.TypeMicrosoftInstaller},
		{file: "jar.jar", kind: types.KindZIPArchive, typ: types.TypeJavaArchive},
		{file: "java.class", kind: types.KindJavaClassBytecode, typ: types.TypeNone},
		{file: "jpeg.jpeg", kind: types.KindJPEGImage, typ: types.TypeNone},
		{file: "jpegxl.jxl", kind: types.KindJPEGXLImage, typ: types.TypeCodestream},
		{file: "little_endian.tif", kind: types.KindTIFFImage, typ: types.TypeLittleEndian},
		{file: "m4a.m4a", kind: types.KindISOBaseMedia, typ: types.TypeMPEG4Audio},
		{file: "macho", kind: types.KindMachOBinary, typ: types.Type64BitLittleEndian},
		{file: "mkv.mkv", kind: types.KindEBMLContainer, typ: types.TypeMatroskaVideo},
		{file: "module.wasm", kind: types.KindWebAssemblyModule, typ: types.TypeNone},
		{file: "mp4.mp4", kind: types.KindISOBaseMedia, typ: types.TypeMP4Video},
		{file: "mpeg2.ts", kind: types.KindMPEGTransportStream, typ: types.TypeNone},
		{file: "mpeg3.mp3", kind: types.KindMPEGAudio, typ: types.TypeMP3ID3Tagged},
		{file: "nef.nef", kind: types.KindTIFFImage, typ: types.TypeNikonRAW},
		{file: "nifti.nii", kind: types.KindNIfTIMedicalImage, typ: types.TypeNone},
		{file: "ogg.ogg", kind: types.KindOggContainer, typ: types.TypeVorbisAudio},
		{file: "open-doc.odt", kind: types.KindZIPArchive, typ: types.TypeOpenDocumentText},
		{file: "open-spread.ods", kind: types.KindZIPArchive, typ: types.TypeOpenDocumentSpreadsheet},
		{file: "opus.opus", kind: types.KindOggContainer, typ: types.TypeOpusAudio},
		{file: "package.deb", kind: types.KindDebianPackage, typ: types.TypeNone},
		{file: "pcap.pcap", kind: types.KindPCAPCapture, typ: types.TypeLittleEndian},
		{file: "pcapng.pcapng", kind: types.KindPCAPNGCapture, typ: types.TypeNone},
		{file: "pcx.pcx", kind: types.KindPCXImage, typ: types.TypeNone},
		{file: "pdf.pdf", kind: types.KindPDFDocument, typ: types.TypeNone},
		{file: "pe_x64", kind: types.KindPortableExecutable, typ: types.TypePE32PlusX8664},
		{file: "pe_x86", kind: types.KindPortableExecutable, typ: types.TypePE32X86},
		{file: "png.png", kind: types.KindPNGImage, typ: types.TypeNone},
		{file: "powerpoint.ppt", kind: types.KindOLECompoundDocument, typ: types.TypeMicrosoftPowerPointPresentation},
		{file: "powerpoint.pptx", kind: types.KindZIPArchive, typ: types.TypeMicrosoftPowerPointPresentation},
		{file: "pptm.pptm", kind: types.KindZIPArchive, typ: types.TypeMicrosoftPowerPointMacroEnabledPresentation},
		{file: "quicktime.mov", kind: types.KindISOBaseMedia, typ: types.TypeQuickTimeMovie},
		{file: "rar4.rar", kind: types.KindRARArchive, typ: types.TypeLegacy},
		{file: "rar5.rar", kind: types.KindRARArchive, typ: types.TypeVersion5},
		{file: "sqlite.db-shm", kind: types.KindSQLite3SharedMemory, typ: types.TypeLittleEndian},
		{file: "sqlite.db-wal", kind: types.KindSQLite3WriteAheadLog, typ: types.TypeLittleEndian},
		{file: "sqlite.db", kind: types.KindSQLiteDatabase, typ: types.TypeNone},
		{file: "svg.svg", kind: types.KindSVGImage, typ: types.TypeNone},
		{file: "tar.tar", kind: types.KindTARArchive, typ: types.TypeNone},
		{file: "torrent.torrent", kind: types.KindTorrentDocument, typ: types.TypeNone},
		{file: "universal.macho", kind: types.KindMachOUniversalBinary, typ: types.TypeNone},
		{file: "war.war", kind: types.KindZIPArchive, typ: types.TypeJavaWebArchive},
		{file: "wav.wav", kind: types.KindRIFFContainer, typ: types.TypeWAVAudio},
		{file: "webm.webm", kind: types.KindEBMLContainer, typ: types.TypeWebMVideo},
		{file: "webp.webp", kind: types.KindRIFFContainer, typ: types.TypeWebPImage},
		{file: "word.docx", kind: types.KindZIPArchive, typ: types.TypeMicrosoftWordDocument},
		{file: "xlsm.xlsm", kind: types.KindZIPArchive, typ: types.TypeMicrosoftExcelMacroEnabledWorkbook},
		{file: "zip.zip", kind: types.KindZIPArchive, typ: types.TypeNone},
	}

	for _, tc := range cases {
		t.Run(tc.file, func(t *testing.T) {
			path := filepath.Join("test", tc.file)

			data, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("read fixture: %v", err)
			}

			meta, err := types.Detect(tc.file, data)
			if err != nil {
				t.Fatalf("detect: %v", err)
			}

			if meta == nil {
				t.Fatal("detect returned nil metadata")
			}

			if meta.Kind != tc.kind {
				t.Fatalf("kind mismatch: got %v (%q), want %v (%q)", meta.Kind, meta.Kind.String(), tc.kind, tc.kind.String())
			}

			if meta.Type != tc.typ {
				t.Fatalf("type mismatch: got %v (%q), want %v (%q)", meta.Type, meta.Type.String(), tc.typ, tc.typ.String())
			}

			streamed, err := types.DetectReaderAt(tc.file, data[:min(len(data), 2048)], bytes.NewReader(data), int64(len(data)))
			if err != nil || streamed == nil || *streamed != *meta {
				t.Fatalf("ReaderAt mismatch: got %+v, %v; want %+v", streamed, err, meta)
			}
		})
	}
}

func TestTarSizeOverflow(t *testing.T) {
	data := make([]byte, 512)
	copy(data, "member")
	copy(data[257:], "ustar\x00")
	data[124] = 0x80
	binary.BigEndian.PutUint64(data[128:], ^uint64(0)-511)
	checksum := 8 * int(' ')

	for _, value := range data {
		checksum += int(value)
	}

	copy(data[148:156], fmt.Sprintf("%06o\x00 ", checksum))

	meta, err := types.Detect("overflow.tar", data)
	if err != nil || meta.Kind != types.KindTARArchive {
		t.Fatalf("detect: %v, %v", meta, err)
	}
}

func TestTIFFCountOverflow(t *testing.T) {
	data := make([]byte, 32)
	copy(data, "II\x2a\x00\x08\x00\x00\x00")
	binary.LittleEndian.PutUint16(data[8:], 1)
	binary.LittleEndian.PutUint16(data[10:], 0x010f)
	binary.LittleEndian.PutUint16(data[12:], 2)
	binary.LittleEndian.PutUint32(data[14:], 0xffffffff)
	binary.LittleEndian.PutUint32(data[18:], 16)

	meta, err := types.Detect("overflow.tif", data)
	if err != nil || meta.Kind != types.KindTIFFImage {
		t.Fatalf("detect: %v, %v", meta, err)
	}
}

func TestDetectCollisions(t *testing.T) {
	payload := make([]byte, 2048)
	copy(payload[1080-30-len("payload"):], []byte{0x53, 0xef})

	zipPayload := collisionZIP(t, map[string][]byte{"payload": payload}, zip.Store)
	quotedMIME := collisionZIP(t, map[string][]byte{"notes.txt": []byte("application/vnd.ms-word.document.macroEnabled.main+xml")}, zip.Store)
	project := collisionZIP(t, map[string][]byte{"project.json": []byte("{}"), "payload/data.txt": []byte("data")}, zip.Deflate)

	template := collisionZIP(t, map[string][]byte{
		"[Content_Types].xml": []byte(`<Types xmlns="http://schemas.openxmlformats.org/package/2006/content-types"><Override PartName="/word/document.xml" ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.template.main+xml"/></Types>`),
		"_rels/.rels":         []byte(`<Relationships xmlns="http://schemas.openxmlformats.org/package/2006/relationships"><Relationship Id="rId1" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument" Target="word/document.xml"/></Relationships>`),
		"word/document.xml":   []byte(`<w:document xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main"/>`),
	}, zip.Deflate)

	cases := []collisionCase{
		{name: "zip-ext", data: zipPayload, kind: types.KindZIPArchive},
		{name: "quoted-mime", data: quotedMIME, kind: types.KindZIPArchive},
		{name: "generic-project", data: project, kind: types.KindZIPArchive},
		{name: "deflated-template", data: template, kind: types.KindZIPArchive, typ: types.TypeMicrosoftWordTemplate},
		{name: "boolean-not-font", data: []byte("true"), kind: types.KindJSONDocument},
		{name: "json-array", data: []byte(" \n[1, null, {}]\t"), kind: types.KindJSONDocument},
		{name: "json-empty", data: []byte("{}"), kind: types.KindJSONDocument},
		{name: "json-malformed", data: []byte(`{"type":"FeatureCollection","features":}`), forbidden: types.KindJSONDocument},
		{name: "resources-key", data: []byte(`{"resources":[]}`), kind: types.KindJSONDocument},
		{name: "xml-comment", data: []byte("\xef\xbb\xbf<?xml version=\"1.0\"?><!-- <svg --> <root/>"), kind: types.KindXMLDocument},
		{name: "bitmap-prose", data: []byte("BM is an abbreviation"), forbidden: types.KindBMPImage},
		{name: "netpbm-prose", data: []byte("P6 is a processor family"), forbidden: types.KindNetpbmImage},
		{name: "c-header", data: []byte("#define VALUE 42\n"), forbidden: types.KindXBitMapImage},
		{name: "tagged-flac", data: []byte("ID3\x04\x00\x00\x00\x00\x00\x00fLaC"), kind: types.KindFLACAudio},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			full, fullErr := types.Detect(test.name, test.data)
			streamed, streamErr := types.DetectReaderAt(test.name, test.data[:min(32, len(test.data))], bytes.NewReader(test.data), int64(len(test.data)))

			if !errors.Is(fullErr, streamErr) || (full == nil) != (streamed == nil) || full != nil && *full != *streamed {
				t.Fatalf("API mismatch: %v, %v / %v, %v", full, fullErr, streamed, streamErr)
			}

			if test.forbidden != types.KindUnknown {
				if full != nil && full.Kind == test.forbidden {
					t.Fatalf("false positive: %+v", full)
				}

				return
			}

			if fullErr != nil || full == nil || full.Kind != test.kind || full.Type != test.typ {
				t.Fatalf("got %+v, %v; want %v / %v", full, fullErr, test.kind, test.typ)
			}
		})
	}
}

func collisionZIP(t *testing.T, files map[string][]byte, method uint16) []byte {
	t.Helper()

	var output bytes.Buffer

	writer := zip.NewWriter(&output)

	for name, data := range files {
		entry, err := writer.CreateHeader(&zip.FileHeader{Name: name, Method: method})
		if err != nil {
			t.Fatal(err)
		}

		_, err = entry.Write(data)
		if err != nil {
			t.Fatal(err)
		}
	}

	err := writer.Close()
	if err != nil {
		t.Fatal(err)
	}

	return output.Bytes()
}

func getIdentifiers(t *testing.T, filename, prefix string) []string {
	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		t.Fatalf("Failed to parse %s: %v", filename, err)
	}

	var names []string

	for _, decl := range node.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.CONST {
			continue
		}

		for _, spec := range genDecl.Specs {
			vs, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}

			for _, name := range vs.Names {
				if strings.HasPrefix(name.Name, prefix) {
					names = append(names, name.Name)
				}
			}
		}
	}

	return names
}
