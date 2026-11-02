package custom

import (
	"encoding/json"

	"github.com/coalaura/wtf/types"
)

type jsonIdentity struct {
	Type            string          `json:"type"`
	ResourceType    string          `json:"resourceType"`
	ResourceVersion string          `json:"resourceVersion"`
	Features        json.RawMessage `json:"features"`
	Geometry        json.RawMessage `json:"geometry"`
	Coordinates     json.RawMessage `json:"coordinates"`
	Objects         json.RawMessage `json:"objects"`
	Arcs            json.RawMessage `json:"arcs"`
	Resources       json.RawMessage `json:"resources"`
}

func DetectJSON(b types.Buffer) *types.Metadata {
	if !json.Valid(b) {
		return nil
	}

	var identity jsonIdentity

	err := json.Unmarshal(b, &identity)
	if err == nil {
		switch identity.Type {
		case "FeatureCollection":
			if len(identity.Features) > 0 {
				return &types.Metadata{Kind: types.KindJSONDocument, Type: types.TypeGeoJSON}
			}
		case "Feature":
			if len(identity.Geometry) > 0 {
				return &types.Metadata{Kind: types.KindJSONDocument, Type: types.TypeGeoJSON}
			}
		case "Point", "MultiPoint", "LineString", "MultiLineString", "Polygon", "MultiPolygon":
			if len(identity.Coordinates) > 0 {
				return &types.Metadata{Kind: types.KindJSONDocument, Type: types.TypeGeoJSON}
			}
		case "Topology":
			if len(identity.Objects) > 0 && len(identity.Arcs) > 0 {
				return &types.Metadata{Kind: types.KindJSONDocument, Type: types.TypeTopoJSON}
			}
		}

		if identity.ResourceType == "GMProject" && identity.ResourceVersion != "" && len(identity.Resources) > 0 {
			return &types.Metadata{Kind: types.KindGameMakerStudioProject, Type: types.TypeGameMakerStudio2}
		}
	}

	return &types.Metadata{Kind: types.KindJSONDocument}
}
