package custom

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io"
	"strings"

	"github.com/coalaura/wtf/types"
)

func DetectXMLSubtypes(b types.Buffer) *types.Metadata {
	data := bytes.TrimPrefix(b, []byte{0xef, 0xbb, 0xbf})

	if b.Has(0, []byte{0xff, 0xfe}) || b.Has(0, []byte{0xfe, 0xff}) {
		decoded, _, ok := TryDecodeUTF16(b)
		if !ok {
			return nil
		}

		data = decoded
	}

	data = bytes.TrimSpace(data)
	if len(data) == 0 || data[0] != '<' {
		return nil
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = xmlCharset

	if b.Has(0, []byte{0xff, 0xfe}) || b.Has(0, []byte{0xfe, 0xff}) {
		decoder.CharsetReader = decodedXMLCharset
	}

	for {
		token, err := decoder.Token()
		if err != nil {
			return nil
		}

		if text, ok := token.(xml.CharData); ok && len(bytes.TrimSpace(text)) != 0 {
			return nil
		}

		if root, ok := token.(xml.StartElement); ok {
			return xmlRootMetadata(root)
		}
	}
}

func decodedXMLCharset(charset string, input io.Reader) (io.Reader, error) {
	switch strings.ToLower(charset) {
	case "utf-16", "utf-16le", "utf-16be":
		return input, nil
	default:
		return nil, errors.New("unsupported XML encoding")
	}
}

func xmlCharset(charset string, input io.Reader) (io.Reader, error) {
	if !strings.EqualFold(charset, "iso-8859-1") {
		return nil, errors.New("unsupported XML encoding")
	}

	data, err := io.ReadAll(input)
	if err != nil {
		return nil, err
	}

	var decoded strings.Builder

	decoded.Grow(len(data) * 2)

	for _, character := range data {
		decoded.WriteRune(rune(character))
	}

	return strings.NewReader(decoded.String()), nil
}

func xmlRootMetadata(root xml.StartElement) *types.Metadata {
	name := root.Name.Local
	namespace := root.Name.Space

	switch {
	case name == "svg" && namespace == "http://www.w3.org/2000/svg":
		return &types.Metadata{Kind: types.KindSVGImage}
	case name == "plist" && namespace == "":
		return &types.Metadata{Kind: types.KindAppleXMLPropertyList}
	case name == "kml" && (namespace == "http://www.opengis.net/kml/2.2" || namespace == "http://earth.google.com/kml/2.1"):
		return &types.Metadata{Kind: types.KindKeyholeMarkupLanguage}
	case name == "gpx" && (namespace == "http://www.topografix.com/GPX/1/1" || namespace == "http://www.topografix.com/GPX/1/0"):
		return &types.Metadata{Kind: types.KindGPSExchangeFormat}
	case name == "rss" && namespace == "":
		return &types.Metadata{Kind: types.KindRSSFeed}
	case name == "feed" && namespace == "http://www.w3.org/2005/Atom":
		return &types.Metadata{Kind: types.KindAtomFeed}
	case (name == "stylesheet" || name == "transform") && namespace == "http://www.w3.org/1999/XSL/Transform":
		return &types.Metadata{Kind: types.KindXMLDocument, Type: types.TypeXSLTStylesheet}
	case name == "Envelope" && (namespace == "http://schemas.xmlsoap.org/soap/envelope/" || namespace == "http://www.w3.org/2003/05/soap-envelope"):
		return &types.Metadata{Kind: types.KindSOAPMessage}
	case name == "COLLADA" && namespace == "http://www.collada.org/2005/11/COLLADASchema":
		return &types.Metadata{Kind: types.KindColladaModel}
	case name == "X3D" && (namespace == "" || namespace == "http://www.web3d.org/specifications/x3d-namespace"):
		return &types.Metadata{Kind: types.KindX3DModel}
	case namespace == "http://www.opengis.net/gml" || namespace == "http://www.opengis.net/gml/3.2":
		return &types.Metadata{Kind: types.KindGeographyMarkupLanguage}
	case name == "timedtext" && namespace == "":
		for _, attribute := range root.Attr {
			if attribute.Name.Local == "format" && attribute.Value == "3" {
				return &types.Metadata{Kind: types.KindYouTubeTimedText}
			}
		}
	case name == "assets" && namespace == "":
		return &types.Metadata{Kind: types.KindGameMakerStudioProject, Type: types.TypeGameMakerStudio, Confidence: types.ConfidenceMedium}
	}

	return &types.Metadata{Kind: types.KindXMLDocument, Confidence: types.ConfidenceMedium}
}
