package models

// HGSSDocument represents the root HGSS document
type HGSSDocument struct {
	Type        string                 `json:"type"`
	Version     string                 `json:"version"`
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	Styles      map[string]interface{} `json:"styles"`
	Root        *Node                  `json:"root"`
}

// Node represents a node in the HGSS hierarchy
type Node struct {
	ID          string                 `json:"id"`
	Type        string                 `json:"type"`
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	Coordinates []float64              `json:"coordinates,omitempty"`
	Geometry    *Geometry              `json:"geometry,omitempty"`
	Style       map[string]interface{} `json:"style,omitempty"`
	Children    []*Node                `json:"children,omitempty"`
}

// Geometry represents GeoJSON geometry
type Geometry struct {
	Type        string        `json:"type"`
	Coordinates interface{}   `json:"coordinates"`
}

// GeoJSONFeatureCollection represents GeoJSON FeatureCollection
type GeoJSONFeatureCollection struct {
	Type     string            `json:"type"`
	Features []GeoJSONFeature  `json:"features"`
}

// GeoJSONFeature represents a GeoJSON Feature
type GeoJSONFeature struct {
	Type       string                 `json:"type"`
	ID         interface{}            `json:"id,omitempty"`
	Geometry   *Geometry              `json:"geometry"`
	Properties map[string]interface{} `json:"properties"`
}

// Validate checks if the document is a valid HGSS
func (doc *HGSSDocument) Validate() error {
	if doc.Type != "HGSS" {
		return ErrInvalidType
	}
	if doc.Version == "" {
		return ErrMissingVersion
	}
	if doc.Root == nil {
		return ErrMissingRoot
	}
	return nil
}

// Errors
var (
	ErrInvalidType     = &ValidationError{"invalid type, must be HGSS"}
	ErrMissingVersion  = &ValidationError{"missing version"}
	ErrMissingRoot     = &ValidationError{"missing root node"}
)

// ValidationError represents a validation error
type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}