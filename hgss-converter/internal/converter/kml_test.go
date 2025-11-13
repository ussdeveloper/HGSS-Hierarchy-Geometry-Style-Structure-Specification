package converter

import (
	"os"
	"testing"
)

func TestConvertKMLToHGSS(t *testing.T) {
	// Load test data
	data, err := os.ReadFile("../../testdata/sample.kml")
	if err != nil {
		t.Fatalf("Failed to read test data: %v", err)
	}

	// Convert
	hgss, err := ConvertKMLToHGSS(data)
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
	if len(hgss.Root.Children) != 1 {
		t.Errorf("Expected 1 child, got %d", len(hgss.Root.Children))
	}

	// Check folder
	folder := hgss.Root.Children[0]
	if folder.Type != "Group" {
		t.Errorf("Expected Group type, got %s", folder.Type)
	}
	if folder.Name != "Zones" {
		t.Errorf("Expected 'Zones', got %s", folder.Name)
	}
	if len(folder.Children) != 2 {
		t.Errorf("Expected 2 children in folder, got %d", len(folder.Children))
	}

	// Check polygon
	polygon := folder.Children[0]
	if polygon.Type != "Polygon" {
		t.Errorf("Expected Polygon type, got %s", polygon.Type)
	}
	if polygon.Name != "Zone A" {
		t.Errorf("Expected 'Zone A', got %s", polygon.Name)
	}

	// Check point
	point := folder.Children[1]
	if point.Type != "Point" {
		t.Errorf("Expected Point type, got %s", point.Type)
	}
	if point.Name != "Point of Interest" {
		t.Errorf("Expected 'Point of Interest', got %s", point.Name)
	}
}

func TestParseKMLCoordinates(t *testing.T) {
	coordStr := "19.938,50.000 19.950,50.000 19.950,50.010"
	coords := parseKMLCoordinates(coordStr)
	expected := [][]float64{
		{19.938, 50.000},
		{19.950, 50.000},
		{19.950, 50.010},
	}
	if len(coords) != len(expected) {
		t.Fatalf("Expected %d coords, got %d", len(expected), len(coords))
	}
	for i, c := range coords {
		if c[0] != expected[i][0] || c[1] != expected[i][1] {
			t.Errorf("Coord %d: expected %v, got %v", i, expected[i], c)
		}
	}
}

func TestKMLColorToHex(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"ffcccccc", "#cccccc"},
		{"ff444444", "#444444"},
		{"invalid", "#000000"},
	}
	for _, tt := range tests {
		result := kmlColorToHex(tt.input)
		if result != tt.expected {
			t.Errorf("kmlColorToHex(%s) = %s, expected %s", tt.input, result, tt.expected)
		}
	}
}