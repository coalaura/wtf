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
		{file: "3g2.3g2", kind: types.KindISOBaseMedia, typ: types.Type3G2Media},
		{file: "3gp.3gp", kind: types.KindISOBaseMedia, typ: types.Type3GPPMedia},
		{file: "7z.7z", kind: types.Kind7ZipArchive, typ: types.TypeNone},
		{file: "aifc.aif", kind: types.KindIFFContainer, typ: types.TypeAIFCAudio},
		{file: "android.apk", kind: types.KindZIPArchive, typ: types.TypeAndroidPackageAPK},
		{file: "arw.arw", kind: types.KindTIFFImage, typ: types.TypeSonyRAWARW},
		{file: "avi.avi", kind: types.KindRIFFContainer, typ: types.TypeAVIVideo},
		{file: "avif.avif", kind: types.KindISOBaseMedia, typ: types.TypeAVIFImageSequence},
		{file: "bigtiff.tif", kind: types.KindTIFFImage, typ: types.TypeBigTIFF},
		{file: "cr2.cr2", kind: types.KindCanonRAWImage, typ: types.TypeCanonRAWHE},
		{file: "cursor.cur", kind: types.KindICOCURImage, typ: types.TypeWindowsCursor},
		{file: "dicom.dcm", kind: types.KindDICOMMedicalImage, typ: types.TypeNone},
		{file: "dng.dng", kind: types.KindTIFFImage, typ: types.TypeAdobeDNGDNG},
		{file: "document.doc", kind: types.KindOLECompoundDocument, typ: types.TypeMicrosoftWordDocumentDOC},
		{file: "dp3_icc.icc", kind: types.KindICCProfile, typ: types.TypeNone},
		{file: "elf_x64", kind: types.KindExecutableAndLinkableFormat, typ: types.TypeELF64LittleEndian},
		{file: "epub.epub", kind: types.KindZIPArchive, typ: types.TypeEPUBDocument},
		{file: "excel.xls", kind: types.KindOLECompoundDocument, typ: types.TypeMicrosoftExcelWorkbookXLS},
		{file: "excel.xlsx", kind: types.KindZIPArchive, typ: types.TypeMicrosoftExcelSpreadsheetXLSX},
		{file: "gif87a.gif", kind: types.KindGIFImage, typ: types.TypeGIF87a},
		{file: "gif89a.gif", kind: types.KindGIFImage, typ: types.TypeGIF89a},
		{file: "gz.gz", kind: types.KindGzipArchive, typ: types.TypeNone},
		{file: "heif.heif", kind: types.KindISOBaseMedia, typ: types.TypeHEIFImage},
		{file: "ico.ico", kind: types.KindICOCURImage, typ: types.TypeWindowsIcon},
		{file: "jar.jar", kind: types.KindZIPArchive, typ: types.TypeJavaArchiveJAR},
		{file: "java.class", kind: types.KindJavaClass, typ: types.TypeNone},
		{file: "jpeg.jpeg", kind: types.KindJPEGImage, typ: types.TypeNone},
		{file: "jpegxl.jxl", kind: types.KindJPEGXLImage, typ: types.TypeCodestream},
		{file: "m4a.m4a", kind: types.KindISOBaseMedia, typ: types.TypeMPEG4AudioM4AFamily},
		{file: "macho", kind: types.KindMachOBinary, typ: types.Type64BitLittleEndian},
		{file: "mkv.mkv", kind: types.KindEBMLContainer, typ: types.TypeMatroska},
		{file: "mp4.mp4", kind: types.KindISOBaseMedia, typ: types.TypeMP4Video},
		{file: "mpeg2.ts", kind: types.KindMPEGTransportStream, typ: types.TypeTS},
		{file: "mpeg3.mp3", kind: types.KindMPEGAudio, typ: types.TypeMP3ID3Tagged},
		{file: "nef.nef", kind: types.KindTIFFImage, typ: types.TypeNikonRAWNEF},
		{file: "nifti.nii", kind: types.KindNIfTIMedicalImage, typ: types.TypeNone},
		{file: "ogg.ogg", kind: types.KindOggContainer, typ: types.TypeVorbisAudio},
		{file: "open-doc.odt", kind: types.KindZIPArchive, typ: types.TypeOpenDocumentTextODT},
		{file: "open-spread.ods", kind: types.KindZIPArchive, typ: types.TypeOpenDocumentSpreadsheetODS},
		{file: "opus.opus", kind: types.KindOggContainer, typ: types.TypeOpusAudio},
		{file: "pcap.pcap", kind: types.KindPCAPCapture, typ: types.TypeLittleEndian},
		{file: "pcapng.pcapng", kind: types.KindPCAPNGCapture, typ: types.TypeNone},
		{file: "pcx.pcx", kind: types.KindPCXImage, typ: types.TypeNone},
		{file: "pdf.pdf", kind: types.KindPDFDocument, typ: types.TypeNone},
		{file: "pe_x64", kind: types.KindPortableExecutable, typ: types.TypePE32PlusX8664},
		{file: "png.png", kind: types.KindPNGImage, typ: types.TypeNone},
		{file: "powerpoint.ppt", kind: types.KindOLECompoundDocument, typ: types.TypeMicrosoftPowerPointPresentationPPT},
		{file: "powerpoint.pptx", kind: types.KindZIPArchive, typ: types.TypeMicrosoftPowerPointPresentationPPTX},
		{file: "pptm.pptm", kind: types.KindZIPArchive, typ: types.TypeMicrosoftPowerPointMacroEnabledPresentationPPTM},
		{file: "quicktime.mov", kind: types.KindISOBaseMedia, typ: types.TypeQuickTimeMovie},
		{file: "rar4.rar", kind: types.KindRARArchive, typ: types.TypeRARLegacy},
		{file: "rar5.rar", kind: types.KindRARArchive, typ: types.TypeRAR5},
		{file: "sqlite.db-shm", kind: types.KindSQLite3SharedMemory, typ: types.TypeLittleEndian},
		{file: "sqlite.db-wal", kind: types.KindSQLite3WriteAheadLog, typ: types.TypeBigEndian},
		{file: "sqlite.db", kind: types.KindSQLiteDatabase, typ: types.TypeNone},
		{file: "svg.svg", kind: types.KindSVGImage, typ: types.TypeNone},
		{file: "tar.tar", kind: types.KindTARArchive, typ: types.TypeNone},
		{file: "torrent.torrent", kind: types.KindTorrentFile, typ: types.TypeNone},
		{file: "war.war", kind: types.KindZIPArchive, typ: types.TypeJavaWebArchiveWAR},
		{file: "wav.wav", kind: types.KindRIFFContainer, typ: types.TypeWAVAudio},
		{file: "webm.webm", kind: types.KindEBMLContainer, typ: types.TypeWebM},
		{file: "webp.webp", kind: types.KindRIFFContainer, typ: types.TypeWebPImage},
		{file: "word.docx", kind: types.KindZIPArchive, typ: types.TypeMicrosoftWordDocumentDOCX},
		{file: "xlsm.xlsm", kind: types.KindZIPArchive, typ: types.TypeMicrosoftExcelMacroEnabledWorkbookXLSM},
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
