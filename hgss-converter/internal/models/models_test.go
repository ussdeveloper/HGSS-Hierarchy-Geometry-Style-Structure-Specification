package models

import (
	"testing"
)

func TestHGSSDocument_Validate(t *testing.T) {
	tests := []struct {
		name    string
		doc     HGSSDocument
		wantErr bool
	}{
		{
			name: "valid document",
			doc: HGSSDocument{
				Type:    "HGSS",
				Version: "1.0",
				Root:    &Node{ID: "root", Type: "Group"},
			},
			wantErr: false,
		},
		{
			name: "invalid type",
			doc: HGSSDocument{
				Type:    "Invalid",
				Version: "1.0",
				Root:    &Node{ID: "root", Type: "Group"},
			},
			wantErr: true,
		},
		{
			name: "missing version",
			doc: HGSSDocument{
				Type: "HGSS",
				Root: &Node{ID: "root", Type: "Group"},
			},
			wantErr: true,
		},
		{
			name: "missing root",
			doc: HGSSDocument{
				Type:    "HGSS",
				Version: "1.0",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.doc.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}