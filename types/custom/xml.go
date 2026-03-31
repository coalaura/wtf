package custom

import (
	"bytes"

	"github.com/coalaura/wtf/types"
)

func DetectXMLSubtypes(b types.Buffer) *types.Metadata {
	limit := min(b.Len(), 1024)
	data := b[:limit]

	trimmed := bytes.TrimSpace(data)
	if len(trimmed) < 5 || trimmed[0] != '<' {
		return nil
	}

	isXML := bytes.HasPrefix(trimmed, []byte("<?xml")) || bytes.HasPrefix(trimmed, []byte("<!DOCTYPE"))

	if isXML || bytes.HasPrefix(trimmed, []byte("<svg")) {
		if bytes.Contains(data, []byte("<svg")) {
			return &types.Metadata{Kind: types.KindSVGImage, Confidence: types.ConfidenceMedium}
		}
	}

	if isXML {
		if bytes.Contains(data, []byte("<plist")) || bytes.Contains(data, []byte("<!DOCTYPE plist")) {
			return &types.Metadata{Kind: types.KindAppleXMLPropertyList, Confidence: types.ConfidenceMedium}
		}

		if bytes.Contains(data, []byte("<kml")) {
			return &types.Metadata{Kind: types.KindKeyholeMarkupLanguage, Confidence: types.ConfidenceMedium}
		}

		if bytes.Contains(data, []byte("<gpx")) {
			return &types.Metadata{Kind: types.KindGPSExchangeFormat, Confidence: types.ConfidenceMedium}
		}

		if bytes.Contains(data, []byte("<rss")) {
			return &types.Metadata{Kind: types.KindRSSFeed, Confidence: types.ConfidenceMedium}
		}

		if bytes.Contains(data, []byte("<feed")) {
			return &types.Metadata{Kind: types.KindAtomFeed, Confidence: types.ConfidenceMedium}
		}

		if bytes.Contains(data, []byte("<xsl:stylesheet")) || bytes.Contains(data, []byte("<xsl:transform")) {
			return &types.Metadata{Kind: types.KindXMLDocument, Type: types.TypeXSLTStylesheet, Confidence: types.ConfidenceMedium}
		}

		if bytes.Contains(data, []byte("<soap:Envelope")) {
			return &types.Metadata{Kind: types.KindSOAPMessage, Confidence: types.ConfidenceMedium}
		}

		if bytes.Contains(data, []byte("<COLLADA")) {
			return &types.Metadata{Kind: types.KindColladaModel, Confidence: types.ConfidenceMedium}
		}

		if bytes.Contains(data, []byte("<X3D")) {
			return &types.Metadata{Kind: types.KindX3DModel, Confidence: types.ConfidenceMedium}
		}

		if bytes.Contains(data, []byte("<gml:")) {
			return &types.Metadata{Kind: types.KindGeographyMarkupLanguage, Confidence: types.ConfidenceMedium}
		}

		return &types.Metadata{Kind: types.KindXMLDocument, Confidence: types.ConfidenceMedium}
	}

	return nil
}
