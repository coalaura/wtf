package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/coalaura/wtf/types"
)

type testCase struct {
	file string
	kind types.KindID
	typ  types.TypeID
}

func TestDetectFixtures(t *testing.T) {
	cases := []testCase{
		{file: "3gp.3gp", kind: types.KindISOBaseMedia, typ: types.Type3GPPMedia},
		{file: "aifc.aif", kind: types.KindIFFContainer, typ: types.TypeAIFCAudio},
		{file: "android.apk", kind: types.KindZIPArchive, typ: types.TypeAndroidPackageAPK},
		{file: "avif.avif", kind: types.KindISOBaseMedia, typ: types.TypeAVIFImageSequence},
		{file: "dicom.dcm", kind: types.KindDICOMMedicalImage, typ: types.TypeNone},
		{file: "document.doc", kind: types.KindOLECompoundDocument, typ: types.TypeMicrosoftWordDocumentDOC},
		{file: "dng.dng", kind: types.KindTIFFImage, typ: types.TypeAdobeDNGDNG},
		{file: "dp3_icc.icc", kind: types.KindICCProfile, typ: types.TypeNone},
		{file: "elf_x64", kind: types.KindExecutableAndLinkableFormat, typ: types.TypeELF64LittleEndian},
		{file: "epub.epub", kind: types.KindZIPArchive, typ: types.TypeEPUBDocument},
		{file: "excel.xls", kind: types.KindOLECompoundDocument, typ: types.TypeMicrosoftExcelWorkbookXLS},
		{file: "excel.xlsx", kind: types.KindZIPArchive, typ: types.TypeMicrosoftExcelSpreadsheetXLSX},
		{file: "gif87a.gif", kind: types.KindGIFImage, typ: types.TypeGIF87a},
		{file: "gif89a.gif", kind: types.KindGIFImage, typ: types.TypeGIF89a},
		{file: "heif.heif", kind: types.KindISOBaseMedia, typ: types.TypeHEIFImage},
		{file: "jar.jar", kind: types.KindZIPArchive, typ: types.TypeJavaArchiveJAR},
		{file: "java.class", kind: types.KindJavaClass, typ: types.TypeNone},
		{file: "jpeg.jpeg", kind: types.KindJPEGImage, typ: types.TypeNone},
		{file: "jpegxl.jxl", kind: types.KindJPEGXLImage, typ: types.TypeCodestream},
		{file: "m4a.m4a", kind: types.KindISOBaseMedia, typ: types.TypeMPEG4AudioM4AFamily},
		{file: "macho", kind: types.KindMachOBinary, typ: types.Type64BitLittleEndian},
		{file: "mp4.mp4", kind: types.KindISOBaseMedia, typ: types.TypeMP4Video},
		{file: "mpeg2.ts", kind: types.KindMPEGTransportStream, typ: types.TypeTS},
		{file: "mpeg3.mp3", kind: types.KindMPEGAudio, typ: types.TypeMP3ID3Tagged},
		{file: "nef.nef", kind: types.KindTIFFImage, typ: types.TypeNikonRAWNEF},
		{file: "nifti.nii", kind: types.KindNIfTIMedicalImage, typ: types.TypeNone},
		{file: "ogg.ogg", kind: types.KindOggContainer, typ: types.TypeVorbisAudio},
		{file: "open-doc.odt", kind: types.KindZIPArchive, typ: types.TypeOpenDocumentTextODT},
		{file: "open-spread.ods", kind: types.KindZIPArchive, typ: types.TypeOpenDocumentSpreadsheetODS},
		{file: "pcx.pcx", kind: types.KindPCXImage, typ: types.TypeNone},
		{file: "pdf.pdf", kind: types.KindPDFDocument, typ: types.TypeNone},
		{file: "pe_x64", kind: types.KindPortableExecutable, typ: types.TypePE32PlusX8664},
		{file: "powerpoint.ppt", kind: types.KindOLECompoundDocument, typ: types.TypeMicrosoftPowerPointPresentationPPT},
		{file: "powerpoint.pptx", kind: types.KindZIPArchive, typ: types.TypeMicrosoftPowerPointPresentationPPTX},
		{file: "sqlite.db", kind: types.KindSQLiteDatabase, typ: types.TypeNone},
		{file: "sqlite.db-shm", kind: types.KindSQLite3SharedMemory, typ: types.TypeLittleEndian},
		{file: "sqlite.db-wal", kind: types.KindSQLite3WriteAheadLog, typ: types.TypeBigEndian},
		{file: "svg.svg", kind: types.KindSVGImage, typ: types.TypeNone},
		{file: "torrent.torrent", kind: types.KindTorrentFile, typ: types.TypeNone},
		{file: "war.war", kind: types.KindZIPArchive, typ: types.TypeJavaWebArchiveWAR},
		{file: "webm.webm", kind: types.KindEBMLContainer, typ: types.TypeWebM},
		{file: "webp.webp", kind: types.KindRIFFContainer, typ: types.TypeWebPImage},
		{file: "word.docx", kind: types.KindZIPArchive, typ: types.TypeMicrosoftWordDocumentDOCX},
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
		})
	}
}
