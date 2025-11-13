package converter

import (
	"fmt"

	"hgss-converter/internal/models"
)

// ConvertGeoJSONToHGSS converts a GeoJSON FeatureCollection to HGSS
func ConvertGeoJSONToHGSS(geojson *models.GeoJSONFeatureCollection) (*models.HGSSDocument, error) {
	doc := &models.HGSSDocument{
		Type:    "HGSS",
		Version: "1.0",
		Styles:  make(map[string]interface{}),
		Root: &models.Node{
			ID:   "root",
			Type: "Group",
			Children: make([]*models.Node, 0, len(geojson.Features)),
		},
	}

	for i, feature := range geojson.Features {
		node := &models.Node{
			ID:       fmt.Sprintf("feature_%d", i),
			Type:     getHGSSNodeType(feature.Geometry.Type),
			Geometry: feature.Geometry,
		}

		if feature.ID != nil {
			if id, ok := feature.ID.(string); ok {
				node.ID = id
			}
		}

		if feature.Properties != nil {
			props := feature.Properties
			if name, ok := props["name"].(string); ok {
				node.Name = name
			}
			if desc, ok := props["description"].(string); ok {
				node.Description = desc
			}
			if coords, ok := props["label_coordinates"].([]interface{}); ok {
				node.Coordinates = make([]float64, len(coords))
				for j, c := range coords {
					if f, ok := c.(float64); ok {
						node.Coordinates[j] = f
					}
				}
			}
		}

		doc.Root.Children = append(doc.Root.Children, node)
	}

	return doc, nil
}

// ConvertHGSStoGeoJSON converts HGSS to GeoJSON FeatureCollection
func ConvertHGSStoGeoJSON(hgss *models.HGSSDocument) (*models.GeoJSONFeatureCollection, error) {
	features := make([]models.GeoJSONFeature, 0)
	flattenHGSS(hgss.Root, &features)

	return &models.GeoJSONFeatureCollection{
		Type:     "FeatureCollection",
		Features: features,
	}, nil
}

// flattenHGSS recursively flattens HGSS tree into GeoJSON features
func flattenHGSS(node *models.Node, features *[]models.GeoJSONFeature) {
	if node.Geometry != nil {
		feature := models.GeoJSONFeature{
			Type:     "Feature",
			Geometry: node.Geometry,
			Properties: map[string]interface{}{
				"name":        node.Name,
				"description": node.Description,
			},
		}

		if node.ID != "" {
			feature.ID = node.ID
		}

		if node.Coordinates != nil {
			feature.Properties["label_coordinates"] = node.Coordinates
		}

		*features = append(*features, feature)
	}

	if node.Children != nil {
		for _, child := range node.Children {
			flattenHGSS(child, features)
		}
	}
}

// getHGSSNodeType maps GeoJSON geometry type to HGSS node type
func getHGSSNodeType(geoType string) string {
	switch geoType {
	case "Point", "MultiPoint", "LineString", "MultiLineString", "Polygon", "MultiPolygon", "GeometryCollection":
		return geoType
	default:
		return "Group"
	}
}