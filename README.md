# HGSS - Hierarchy Geometry Style Structure

[![License: CC BY-SA 4.0](https://img.shields.io/badge/License-CC%20BY--SA%204.0-lightgrey.svg)](https://creativecommons.org/licenses/by-sa/4.0/)
[![Version](https://img.shields.io/badge/version-1.0-blue.svg)](https://github.com/sulacolab/hgss/releases)

**HGSS (Hierarchy Geometry Style Structure)** is an open JSON standard for representing spatial data in a hierarchical structure, with full support for styling (CSS-like style cascade), label positioning, and conversions to/from popular GIS formats like GeoJSON, KML, GML, and TopoJSON.

## Table of Contents

- [Overview](#overview)
- [Key Features](#key-features)
- [Quick Start](#quick-start)
- [Specification](#specification)
- [Examples](#examples)
- [API Reference](#api-reference)
- [Contributing](#contributing)
- [License](#license)

## Overview

HGSS enables developers and GIS professionals to structure spatial data hierarchically, apply cascading styles, and easily convert between formats. It's designed to be JSON-native, making it ideal for web applications, data visualization, and spatial analysis.

### Use Cases
- Web mapping applications
- GIS data interchange
- Spatial data visualization with styling
- Hierarchical organization of geographic features

## Key Features

- **Hierarchical Structure**: Nested `children` for grouping layers and objects
- **Styling System**: CSS-like cascade with global styles, ID-specific overrides, and inline properties
- **Label Positioning**: Explicit `coordinates` for feature labels
- **Format Compatibility**: Bidirectional conversion with GeoJSON, KML, GML, TopoJSON
- **JSON-Native**: Easy parsing and manipulation in JavaScript/TypeScript
- **Extensible**: Open standard with CC-BY-SA 4.0 license

## Quick Start

### Installation

Clone the repository:
```bash
git clone https://github.com/sulacolab/hgss.git
cd hgss
```

### Basic Example

```json
{
  "type": "HGSS",
  "version": "1.0",
  "styles": {
    "{Polygon}": { "fill": "#ccc", "stroke": "#444" }
  },
  "root": {
    "id": "root",
    "type": "Group",
    "children": [
      {
        "id": "feature1",
        "type": "Polygon",
        "geometry": {
          "type": "Polygon",
          "coordinates": [[[0,0],[1,0],[1,1],[0,1],[0,0]]]
        }
      }
    ]
  }
}
```

### Validation

Use AJV to validate HGSS documents:
```bash
npm install -g ajv-cli
ajv validate -s spec/hgss.schema.json -d your-file.hgss.json
```

## Specification

Detailed specification available in [spec/README.md](spec/README.md).

Key components:
- Document structure
- Node schema
- Styling cascade
- Label coordinates logic
- Conversion mappings

## Examples

See [examples/](examples/) for sample HGSS files.

- [demo.hgss.json](examples/demo.hgss.json) - Basic example with zones and points

## API Reference

### JavaScript Utilities

See [docs/api.md](docs/api.md) for JavaScript functions to work with HGSS.

### Conversion Tools

Planned: Command-line tools for format conversion.

## Contributing

We welcome contributions! See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

### Development Setup

1. Fork the repository
2. Create a feature branch
3. Make changes
4. Run tests: `npm test`
5. Submit a pull request

## License

This project is licensed under the Creative Commons Attribution-ShareAlike 4.0 International License - see the [LICENSE](LICENSE) file for details.

## Authors

- **Przemysław Lusina** (SulacoLab) - Lead Author

## Changelog

See [CHANGELOG.md](CHANGELOG.md) for version history.