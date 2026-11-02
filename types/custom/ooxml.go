package custom

import (
	"bytes"
	"compress/flate"
	"encoding/xml"
	"hash/crc32"
	"io"
	"path"
	"strings"

	"github.com/coalaura/wtf/types"
)

type packageRelationships struct {
	XMLName xml.Name              `xml:"http://schemas.openxmlformats.org/package/2006/relationships Relationships"`
	Entries []packageRelationship `xml:"http://schemas.openxmlformats.org/package/2006/relationships Relationship"`
}

type packageRelationship struct {
	Type       string `xml:"Type,attr"`
	Target     string `xml:"Target,attr"`
	TargetMode string `xml:"TargetMode,attr"`
}

type packageContentTypes struct {
	XMLName   xml.Name          `xml:"http://schemas.openxmlformats.org/package/2006/content-types Types"`
	Overrides []packageOverride `xml:"http://schemas.openxmlformats.org/package/2006/content-types Override"`
}

type packageOverride struct {
	PartName    string `xml:"PartName,attr"`
	ContentType string `xml:"ContentType,attr"`
}

// Resolve parts only through the validated directory, never a second EOCD search.
func detectOOXML(reader io.ReaderAt, centralOffset int64, central []byte) *types.Metadata {
	manifest := zipPart(reader, centralOffset, central, "[Content_Types].xml")
	relations := zipPart(reader, centralOffset, central, "_rels/.rels")

	if manifest == nil || relations == nil {
		return nil
	}

	var relationships packageRelationships

	err := xml.Unmarshal(relations, &relationships)
	if err != nil {
		return nil
	}

	mainPart := ""

	for _, relationship := range relationships.Entries {
		if relationship.Type != "http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument" && relationship.Type != "http://purl.oclc.org/ooxml/officeDocument/relationships/officeDocument" && relationship.Type != "http://schemas.microsoft.com/xps/2005/06/fixedrepresentation" {
			continue
		}

		if mainPart != "" || relationship.TargetMode == "External" || strings.ContainsAny(relationship.Target, "\\?#:") {
			return nil
		}

		mainPart = path.Clean("/" + relationship.Target)
	}

	if mainPart == "" || findZIPPart(central, strings.TrimPrefix(mainPart, "/")) == nil {
		return nil
	}

	var contentTypes packageContentTypes

	err = xml.Unmarshal(manifest, &contentTypes)
	if err != nil {
		return nil
	}

	var subtype types.TypeID

	for _, override := range contentTypes.Overrides {
		if override.PartName != mainPart {
			continue
		}

		if subtype != types.TypeNone {
			return nil
		}

		switch override.ContentType {
		case "application/vnd.ms-package.xps-fixeddocumentsequence+xml":
			subtype = types.TypeXMLPaperSpecification
		case "application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml":
			subtype = types.TypeMicrosoftWordDocument
		case "application/vnd.openxmlformats-officedocument.wordprocessingml.template.main+xml":
			subtype = types.TypeMicrosoftWordTemplate
		case "application/vnd.ms-word.document.macroEnabled.main+xml":
			subtype = types.TypeMicrosoftWordMacroEnabledDocument
		case "application/vnd.ms-word.template.macroEnabledTemplate.main+xml":
			subtype = types.TypeMicrosoftWordMacroEnabledTemplate
		case "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet.main+xml":
			subtype = types.TypeMicrosoftExcelWorkbook
		case "application/vnd.openxmlformats-officedocument.spreadsheetml.template.main+xml":
			subtype = types.TypeMicrosoftExcelTemplate
		case "application/vnd.ms-excel.sheet.macroEnabled.main+xml":
			subtype = types.TypeMicrosoftExcelMacroEnabledWorkbook
		case "application/vnd.ms-excel.template.macroEnabled.main+xml":
			subtype = types.TypeMicrosoftExcelMacroEnabledTemplate
		case "application/vnd.ms-excel.addin.macroEnabled.main+xml":
			subtype = types.TypeMicrosoftExcelAddIn
		case "application/vnd.openxmlformats-officedocument.presentationml.presentation.main+xml":
			subtype = types.TypeMicrosoftPowerPointPresentation
		case "application/vnd.openxmlformats-officedocument.presentationml.template.main+xml":
			subtype = types.TypeMicrosoftPowerPointTemplate
		case "application/vnd.openxmlformats-officedocument.presentationml.slideshow.main+xml":
			subtype = types.TypeMicrosoftPowerPointSlideshow
		case "application/vnd.ms-powerpoint.presentation.macroEnabled.main+xml":
			subtype = types.TypeMicrosoftPowerPointMacroEnabledPresentation
		case "application/vnd.ms-powerpoint.template.macroEnabled.main+xml":
			subtype = types.TypeMicrosoftPowerPointMacroEnabledTemplate
		case "application/vnd.ms-powerpoint.slideshow.macroEnabled.main+xml":
			subtype = types.TypeMicrosoftPowerPointMacroEnabledSlideshow
		case "application/vnd.ms-powerpoint.addin.macroEnabled.main+xml":
			subtype = types.TypeMicrosoftPowerPointAddIn
		}
	}

	if subtype == types.TypeNone {
		return nil
	}

	return &types.Metadata{Kind: types.KindZIPArchive, Type: subtype}
}

func findZIPPart(central []byte, name string) types.Buffer {
	var found types.Buffer

	for offset := 0; offset < len(central); {
		record := types.Buffer(central[offset:])
		nameLength, _ := record.U16LE(28)
		extraLength, _ := record.U16LE(30)
		commentLength, _ := record.U16LE(32)

		if string(record[centralHeaderSize:centralHeaderSize+int(nameLength)]) == name {
			if found != nil {
				return nil
			}

			found = record[:centralHeaderSize]
		}

		offset += centralHeaderSize + int(nameLength) + int(extraLength) + int(commentLength)
	}

	return found
}

func zipPart(reader io.ReaderAt, centralOffset int64, central []byte, name string) []byte {
	record := findZIPPart(central, name)
	if record == nil {
		return nil
	}

	flags, _ := record.U16LE(8)
	method, _ := record.U16LE(10)
	checksum, _ := record.U32LE(16)

	compressedSize, _ := record.U32LE(20)
	uncompressedSize, _ := record.U32LE(24)
	localOffset, _ := record.U32LE(42)

	if compressedSize > maxScanSize || uncompressedSize > maxScanSize || flags&0x41 != 0 || method != 0 && method != 8 || int64(localOffset)+localHeaderSize+int64(len(name)) > centralOffset {
		return nil
	}

	local := make(types.Buffer, localHeaderSize+len(name))
	if !readAtExactly(reader, local, int64(localOffset)) || !local.Has(0, []byte("PK\x03\x04")) {
		return nil
	}

	localFlags, _ := local.U16LE(6)
	localMethod, _ := local.U16LE(8)
	localChecksum, _ := local.U32LE(14)
	localCompressedSize, _ := local.U32LE(18)
	localUncompressedSize, _ := local.U32LE(22)
	nameLength, _ := local.U16LE(26)
	extraLength, _ := local.U16LE(28)
	dataOffset := int64(localOffset) + localHeaderSize + int64(nameLength) + int64(extraLength)

	if localFlags != flags || localMethod != method || int(nameLength) != len(name) || string(local[localHeaderSize:]) != name || dataOffset+int64(compressedSize) > centralOffset {
		return nil
	}

	if flags&8 == 0 && (localChecksum != checksum || localCompressedSize != compressedSize || localUncompressedSize != uncompressedSize) {
		return nil
	}

	compressed := make([]byte, int(compressedSize))
	if !readAtExactly(reader, compressed, dataOffset) {
		return nil
	}

	output := compressed

	if method == 8 {
		inflater := flate.NewReader(bytes.NewReader(compressed))

		var err error

		output, err = io.ReadAll(io.LimitReader(inflater, int64(uncompressedSize)+1))
		inflater.Close()

		if err != nil {
			return nil
		}
	}

	if len(output) != int(uncompressedSize) || crc32.ChecksumIEEE(output) != checksum {
		return nil
	}

	return output
}
