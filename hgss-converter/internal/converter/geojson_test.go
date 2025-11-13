package converter

import (
	"encoding/json"
	"os"
	"testing"

	"hgss-converter/internal/models"
)

func TestConvertGeoJSONToHGSS(t *testing.T) {
	// Load test data
	data, err := os.ReadFile("../../testdata/sample.geojson")
	if err != nil {
		t.Fatalf("Failed to read test data: %v", err)
	}

	var geojson models.GeoJSONFeatureCollection
	if err := json.Unmarshal(data, &geojson); err != nil {
		t.Fatalf("Failed to unmarshal GeoJSON: %v", err)
	}

	// Convert
	hgss, err := ConvertGeoJSONToHGSS(&geojson)
	if err != nil {
		t.Fatalf("Conversion failed: %v", err)
	}

	// Validate
	if hgss.Type != "HGSS" {
		t.Errorf("Expected type HGSS, got %s", hgss.Type)
	}
	if hgss.Version != "1.0" {
		t.Errorf("Expected version 1.0, got %s", hgss.Version)
	}
	if len(hgss.Root.Children) != 2 {
		t.Errorf("Expected 2 children, got %d", len(hgss.Root.Children))
	}

	// Check first child
	child1 := hgss.Root.Children[0]
	if child1.Type != "Point" {
		t.Errorf("Expected Point type, got %s", child1.Type)
	}
	if child1.Name != "Sample Point" {
		t.Errorf("Expected 'Sample Point', got %s", child1.Name)
	}
}

func TestConvertHGSStoGeoJSON(t *testing.T) {
	// Load test data
	data, err := os.ReadFile("../../testdata/sample.hgss.json")
	if err != nil {
		t.Fatalf("Failed to read test data: %v", err)
	}

	var hgss models.HGSSDocument
	if err := json.Unmarshal(data, &hgss); err != nil {
		t.Fatalf("Failed to unmarshal HGSS: %v", err)
	}

	// Convert
	geojson, err := ConvertHGSStoGeoJSON(&hgss)
	if err != nil {
		t.Fatalf("Conversion failed: %v", err)
	}

	// Validate
	if geojson.Type != "FeatureCollection" {
		t.Errorf("Expected FeatureCollection, got %s", geojson.Type)
	}
	if len(geojson.Features) != 2 {
		t.Errorf("Expected 2 features, got %d", len(geojson.Features))
	}

	// Check first feature
	feature1 := geojson.Features[0]
	if feature1.Geometry.Type != "Point" {
		t.Errorf("Expected Point geometry, got %s", feature1.Geometry.Type)
	}
	if feature1.Properties != nil {
		props := feature1.Properties
		if name, ok := props["name"].(string); !ok || name != "Sample Point" {
			t.Errorf("Expected name 'Sample Point', got %v", props["name"])
		}
	}
}

func TestRoundTripConversion(t *testing.T) {
	// Load HGSS
	data, err := os.ReadFile("../../testdata/sample.hgss.json")
	if err != nil {
		t.Fatalf("Failed to read test data: %v", err)
	}

	var original models.HGSSDocument
	if err := json.Unmarshal(data, &original); err != nil {
		t.Fatalf("Failed to unmarshal HGSS: %v", err)
	}

	// Convert HGSS -> GeoJSON -> HGSS
	geojson, err := ConvertHGSStoGeoJSON(&original)
	if err != nil {
		t.Fatalf("HGSS to GeoJSON failed: %v", err)
	}

	converted, err := ConvertGeoJSONToHGSS(geojson)
	if err != nil {
		t.Fatalf("GeoJSON to HGSS failed: %v", err)
	}

	// Basic checks
	if converted.Type != original.Type {
		t.Errorf("Type mismatch: %s vs %s", converted.Type, original.Type)
	}
	if len(converted.Root.Children) != len(original.Root.Children) {
		t.Errorf("Children count mismatch: %d vs %d", len(converted.Root.Children), len(original.Root.Children))
	}
}