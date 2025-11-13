package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"hgss-converter/internal/converter"
	"hgss-converter/internal/models"
)

func main() {
	var (
		inputFile  = flag.String("i", "", "Input file path")
		outputFile = flag.String("o", "", "Output file path")
		fromFormat = flag.String("f", "geojson", "Input format (geojson, hgss, kml)")
		toFormat   = flag.String("t", "hgss", "Output format (geojson, hgss)")
	)
	flag.Parse()

	if *inputFile == "" || *outputFile == "" {
		fmt.Println("Usage: hgss-converter -i input.json -o output.json -f geojson -t hgss")
		os.Exit(1)
	}

	// Read input file
	inputData, err := os.ReadFile(*inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		os.Exit(1)
	}

	var outputData []byte

	switch *fromFormat {
	case "geojson":
		if *toFormat != "hgss" {
			fmt.Println("Unsupported conversion")
			os.Exit(1)
		}
		var geojson models.GeoJSONFeatureCollection
		if err := json.Unmarshal(inputData, &geojson); err != nil {
			fmt.Printf("Error parsing GeoJSON: %v\n", err)
			os.Exit(1)
		}
		hgss, err := converter.ConvertGeoJSONToHGSS(&geojson)
		if err != nil {
			fmt.Printf("Error converting to HGSS: %v\n", err)
			os.Exit(1)
		}
		outputData, err = json.MarshalIndent(hgss, "", "  ")
		if err != nil {
			fmt.Printf("Error marshaling HGSS: %v\n", err)
			os.Exit(1)
		}

	case "kml":
		if *toFormat != "hgss" {
			fmt.Println("Unsupported conversion")
			os.Exit(1)
		}
		hgss, err := converter.ConvertKMLToHGSS(inputData)
		if err != nil {
			fmt.Printf("Error converting KML to HGSS: %v\n", err)
			os.Exit(1)
		}
		outputData, err = json.MarshalIndent(hgss, "", "  ")
		if err != nil {
			fmt.Printf("Error marshaling HGSS: %v\n", err)
			os.Exit(1)
		}

	case "hgss":
		if *toFormat != "geojson" {
			fmt.Println("Unsupported conversion")
			os.Exit(1)
		}
		var hgss models.HGSSDocument
		if err := json.Unmarshal(inputData, &hgss); err != nil {
			fmt.Printf("Error parsing HGSS: %v\n", err)
			os.Exit(1)
		}
		if err := hgss.Validate(); err != nil {
			fmt.Printf("Invalid HGSS document: %v\n", err)
			os.Exit(1)
		}
		geojson, err := converter.ConvertHGSStoGeoJSON(&hgss)
		if err != nil {
			fmt.Printf("Error converting to GeoJSON: %v\n", err)
			os.Exit(1)
		}
		outputData, err = json.MarshalIndent(geojson, "", "  ")
		if err != nil {
			fmt.Printf("Error marshaling GeoJSON: %v\n", err)
			os.Exit(1)
		}

	default:
		fmt.Printf("Unsupported input format: %s\n", *fromFormat)
		os.Exit(1)
	}

	// Write output file
	if err := os.WriteFile(*outputFile, outputData, 0644); err != nil {
		fmt.Printf("Error writing output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Conversion completed: %s -> %s\n", *inputFile, *outputFile)
}