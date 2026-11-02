package custom

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
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

// The caller validates and bounds the central directory before opening the ZIP.
func detectOOXML(reader io.ReaderAt, size int64, central []byte) *types.Metadata {
	if len(central) > maxCentralSize {
		return nil
	}

	archive, err := zip.NewReader(reader, size)
	if err != nil {
		return nil
	}

	manifest := zipPart(archive, "[Content_Types].xml")
	relations := zipPart(archive, "_rels/.rels")

	if manifest == nil || relations == nil {
		return nil
	}

	var relationships packageRelationships

	err = xml.Unmarshal(relations, &relationships)
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

	if mainPart == "" || !hasZIPPart(archive, strings.TrimPrefix(mainPart, "/")) {
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

func hasZIPPart(archive *zip.Reader, name string) bool {
	count := 0

	for _, file := range archive.File {
		if file.Name == name {
			count++
		}
	}

	return count == 1
}

func zipPart(archive *zip.Reader, name string) []byte {
	if !hasZIPPart(archive, name) {
		return nil
	}

	for _, file := range archive.File {
		if file.Name != name {
			continue
		}

		if file.CompressedSize64 > maxScanSize || file.UncompressedSize64 > maxScanSize || file.Flags&1 != 0 {
			return nil
		}

		reader, err := file.Open()
		if err != nil {
			return nil
		}

		var output bytes.Buffer

		output.Grow(int(file.UncompressedSize64))
		_, err = io.Copy(&output, io.LimitReader(reader, maxScanSize+1))
		reader.Close()

		if err != nil || output.Len() > maxScanSize {
			return nil
		}

		return output.Bytes()
	}

	return nil
}
