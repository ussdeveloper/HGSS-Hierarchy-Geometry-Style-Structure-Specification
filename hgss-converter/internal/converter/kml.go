package converter

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"

	"hgss-converter/internal/models"
)

// KML structures
type KML struct {
	XMLName  xml.Name `xml:"kml"`
	Document Document `xml:"Document"`
}

type Document struct {
	Name   string  `xml:"name"`
	Styles []Style `xml:"Style"`
	Folders []Folder `xml:"Folder"`
}

type Folder struct {
	Name        string       `xml:"name"`
	Placemarks  []Placemark  `xml:"Placemark"`
	SubFolders  []Folder     `xml:"Folder"`
}

type Placemark struct {
	Name        string      `xml:"name"`
	Description string      `xml:"description"`
	StyleURL    string      `xml:"styleUrl"`
	Point       *Point      `xml:"Point"`
	Polygon     *Polygon    `xml:"Polygon"`
}

type Point struct {
	Coordinates string `xml:"coordinates"`
}

type Polygon struct {
	OuterBoundary OuterBoundary `xml:"outerBoundaryIs"`
}

type OuterBoundary struct {
	LinearRing LinearRing `xml:"LinearRing"`
}

type LinearRing struct {
	Coordinates string `xml:"coordinates"`
}

type Style struct {
	ID         string     `xml:"id,attr"`
	PolyStyle  *PolyStyle `xml:"PolyStyle"`
	LineStyle   *LineStyle  `xml:"LineStyle"`
}

type PolyStyle struct {
	Color   string `xml:"color"`
	Fill    string `xml:"fill"`
	Outline string `xml:"outline"`
}

type LineStyle struct {
	Color string `xml:"color"`
	Width string `xml:"width"`
}

// ConvertKMLToHGSS converts KML to HGSS
func ConvertKMLToHGSS(kmlData []byte) (*models.HGSSDocument, error) {
	var kml KML
	if err := xml.Unmarshal(kmlData, &kml); err != nil {
		return nil, fmt.Errorf("failed to parse KML: %w", err)
	}

	doc := &models.HGSSDocument{
		Type:    "HGSS",
		Version: "1.0",
		Styles:  make(map[string]interface{}),
		Root: &models.Node{
			ID:       "root",
			Type:     "Group",
			Name:     kml.Document.Name,
			Children: make([]*models.Node, 0),
		},
	}

	// Process styles
	for _, style := range kml.Document.Styles {
		styleProps := make(map[string]interface{})
		if style.PolyStyle != nil {
			if style.PolyStyle.Color != "" {
				styleProps["fill"] = kmlColorToHex(style.PolyStyle.Color)
			}
			if style.PolyStyle.Fill == "1" {
				styleProps["fill-opacity"] = 1.0
			}
		}
		if style.LineStyle != nil {
			if style.LineStyle.Color != "" {
				styleProps["stroke"] = kmlColorToHex(style.LineStyle.Color)
			}
			if style.LineStyle.Width != "" {
				if w, err := strconv.ParseFloat(style.LineStyle.Width, 64); err == nil {
					styleProps["stroke-width"] = w
				}
			}
		}
		doc.Styles["#"+style.ID] = styleProps
	}

	// Process folders recursively
	for _, folder := range kml.Document.Folders {
		node := convertKMLFolderToNode(&folder)
		doc.Root.Children = append(doc.Root.Children, node)
	}

	return doc, nil
}

func convertKMLFolderToNode(folder *Folder) *models.Node {
	node := &models.Node{
		ID:       generateID(folder.Name),
		Type:     "Group",
		Name:     folder.Name,
		Children: make([]*models.Node, 0),
	}

	for _, pm := range folder.Placemarks {
		child := convertKMLPlacemarkToNode(&pm)
		node.Children = append(node.Children, child)
	}

	for _, subFolder := range folder.SubFolders {
		child := convertKMLFolderToNode(&subFolder)
		node.Children = append(node.Children, child)
	}

	return node
}

func convertKMLPlacemarkToNode(pm *Placemark) *models.Node {
	node := &models.Node{
		ID:          generateID(pm.Name),
		Name:        pm.Name,
		Description: pm.Description,
	}

	if pm.Point != nil {
		coords := parseKMLCoordinates(pm.Point.Coordinates)
		if len(coords) > 0 {
			node.Type = "Point"
			node.Geometry = &models.Geometry{
				Type:        "Point",
				Coordinates: []interface{}{coords[0][0], coords[0][1]},
			}
		}
	} else if pm.Polygon != nil {
		coords := parseKMLCoordinates(pm.Polygon.OuterBoundary.LinearRing.Coordinates)
		if len(coords) > 0 {
			node.Type = "Polygon"
			node.Geometry = &models.Geometry{
				Type:        "Polygon",
				Coordinates: []interface{}{coords},
			}
		}
	}

	if pm.StyleURL != "" {
		// Note: In full implementation, would resolve styleUrl to style properties
		node.Style = map[string]interface{}{"styleUrl": pm.StyleURL}
	}

	return node
}

func parseKMLCoordinates(coordStr string) [][]float64 {
	parts := strings.Fields(coordStr)
	coords := make([][]float64, 0, len(parts))
	for _, part := range parts {
		subParts := strings.Split(strings.TrimSpace(part), ",")
		if len(subParts) >= 2 {
			if lon, err := strconv.ParseFloat(subParts[0], 64); err == nil {
				if lat, err := strconv.ParseFloat(subParts[1], 64); err == nil {
					coords = append(coords, []float64{lon, lat})
				}
			}
		}
	}
	return coords
}

func kmlColorToHex(kmlColor string) string {
	if len(kmlColor) != 8 {
		return "#000000"
	}
	// KML color is AABBGGRR, convert to #RRGGBB
	return "#" + kmlColor[6:8] + kmlColor[4:6] + kmlColor[2:4]
}

func generateID(name string) string {
	return strings.ToLower(strings.ReplaceAll(name, " ", "_"))
}