# HGSS Converter

A command-line tool for converting between HGSS and various geospatial formats (GeoJSON, KML, GML, TopoJSON).

## Features

- Convert GeoJSON to HGSS
- Convert HGSS to GeoJSON
- Command-line interface
- Cross-platform builds (Linux, macOS, Windows)
- Comprehensive tests

## Installation

### From Source

```bash
git clone https://github.com/sulacolab/hgss-converter.git
cd hgss-converter
go mod tidy
make build
```

### Pre-built Binaries

Download from the [releases page](https://github.com/sulacolab/hgss-converter/releases).

## Usage

```bash
hgss-converter -i input.json -o output.json -f geojson -t hgss
```

### Options

- `-i`: Input file path
- `-o`: Output file path
- `-f`: Input format (geojson, hgss)
- `-t`: Output format (geojson, hgss)

### Examples

Convert GeoJSON to HGSS:
```bash
hgss-converter -i sample.geojson -o sample.hgss.json -f geojson -t hgss
```

Convert HGSS to GeoJSON:
```bash
hgss-converter -i sample.hgss.json -o sample.geojson -f hgss -t geojson
```

## Building

### Using Make (Linux/macOS)

```bash
make build          # Build for current platform
make build-all      # Build for all platforms
make test           # Run tests
make clean          # Clean build artifacts
```

### Using Batch Script (Windows)

```cmd
scripts\build.bat
```

## Testing

Run unit tests:
```bash
go test ./...
```

Run integration tests:
```bash
make test-integration
```

## Project Structure

```
hgss-converter/
├── main.go                 # CLI entry point
├── internal/
│   ├── converter/          # Conversion functions
│   │   ├── geojson.go      # GeoJSON conversions
│   │   └── geojson_test.go # Tests
│   └── models/             # Data models
│       ├── models.go       # HGSS/GeoJSON structs
│       └── models_test.go  # Tests
├── testdata/               # Test data files
├── scripts/                # Build scripts
├── Makefile                # Build configuration
└── README.md               # This file
```

## Supported Formats

- **HGSS**: Hierarchy Geometry Style Structure (JSON)
- **GeoJSON**: GeoJSON FeatureCollection (JSON)
- **KML**: Keyhole Markup Language (XML) - Folders, Styles, Placemarks

*Note: Support for TopoJSON, GML, Mapbox Style, SLD, CartoCSS, FlatGeobuf, and GeoPackage planned for future releases.*

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make changes and add tests
4. Run `make test`
5. Submit a pull request

## License

CC-BY-SA 4.0 - See LICENSE file for details.