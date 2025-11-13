# HGSS JavaScript Converter

A professional TypeScript library for converting between HGSS and various geospatial formats with reactive data mirroring.

## Features

- **Bidirectional Conversions**: HGSS ↔ GeoJSON, HGSS ↔ KML
- **Reactive Mirrors**: Automatic synchronization between HGSS and other formats using watchers and optimized references
- **TypeScript Support**: Full type definitions and IntelliSense
- **Extensible**: Easy to add new format converters
- **Lightweight**: No external dependencies for core functionality

## Installation

```bash
npm install @hgss/javascript-converter
```

## Quick Start

```typescript
import { HGSSMirror, GeoJSONConverter, KMLConverter } from '@hgss/javascript-converter';

// Create HGSS document
const hgssDoc = {
  type: "HGSS",
  version: "1.0",
  styles: {},
  root: {
    id: "root",
    type: "Group",
    children: [{
      id: "point1",
      type: "Point",
      geometry: { type: "Point", coordinates: [0, 0] }
    }]
  }
};

// Create mirror with GeoJSON and KML
const mirror = new HGSSMirror(hgssDoc, {
  geojson: new GeoJSONConverter(),
  kml: new KMLConverter()
});

// Access mirrored formats
console.log(mirror.get('geojson')); // GeoJSON FeatureCollection
console.log(mirror.get('kml')); // KML string

// Modify HGSS - mirrors update automatically
hgssDoc.root.children[0].geometry.coordinates = [1, 1];
console.log(mirror.get('geojson')); // Updated coordinates
```

## API Reference

### HGSSMirror

Main class for reactive data mirroring.

```typescript
class HGSSMirror {
  constructor(hgss: HGSSDocument, converters: Record<string, Converter>);
  get(format: string): any;
  setHGSS(hgss: HGSSDocument): void;
  addConverter(name: string, converter: Converter): void;
}
```

### Converters

#### GeoJSONConverter

```typescript
class GeoJSONConverter implements Converter {
  toHGSS(data: GeoJSONFeatureCollection): HGSSDocument;
  fromHGSS(hgss: HGSSDocument): GeoJSONFeatureCollection;
}
```

#### KMLConverter

```typescript
class KMLConverter implements Converter {
  toHGSS(data: string): HGSSDocument; // KML as string
  fromHGSS(hgss: HGSSDocument): string; // Returns KML string
}
```

## Reactive Updates

The library uses JavaScript Proxies to watch for changes in the HGSS document. When you modify the HGSS object, all mirrored formats are automatically updated:

```typescript
const mirror = new HGSSMirror(hgssDoc, { geojson: new GeoJSONConverter() });

// Direct modification
hgssDoc.root.children.push({
  id: "point2",
  type: "Point",
  geometry: { type: "Point", coordinates: [2, 2] }
});

// Mirror updates instantly
const geojson = mirror.get('geojson');
console.log(geojson.features.length); // 2
```

## Adding Custom Converters

Implement the `Converter` interface:

```typescript
interface Converter {
  toHGSS(data: any): HGSSDocument;
  fromHGSS(hgss: HGSSDocument): any;
}

class CustomConverter implements Converter {
  toHGSS(data: any): HGSSDocument {
    // Your conversion logic
  }

  fromHGSS(hgss: HGSSDocument): any {
    // Your conversion logic
  }
}
```

## Building

```bash
npm run build
```

## Testing

```bash
npm test
```

## License

CC-BY-SA-4.0