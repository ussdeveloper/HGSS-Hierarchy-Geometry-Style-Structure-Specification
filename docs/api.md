# HGSS API Reference

This document provides reference for the JavaScript/TypeScript HGSS library with reactive mirrors.

## Installation

```bash
npm install @hgss/javascript-converter
```

## Core Classes

### HGSSMirror

Main class for reactive data mirroring between HGSS and other formats.

```typescript
class HGSSMirror {
  constructor(hgss: HGSSDocument, converters: Record<string, Converter>);
  get(format: string): any;
  setHGSS(hgss: HGSSDocument): void;
  addConverter(name: string, converter: Converter): void;
}
```

**Parameters:**
- `hgss`: HGSS document object
- `converters`: Object mapping format names to converter instances

**Methods:**
- `get(format)`: Returns the mirrored data in specified format
- `setHGSS(hgss)`: Updates the HGSS document and refreshes all mirrors
- `addConverter(name, converter)`: Adds a new format converter

### Converter Interface

Interface for format converters.

```typescript
interface Converter {
  toHGSS(data: any): HGSSDocument;
  fromHGSS(hgss: HGSSDocument): any;
}
```

## Built-in Converters

### GeoJSONConverter

Converts between HGSS and GeoJSON FeatureCollection.

```typescript
class GeoJSONConverter implements Converter {
  toHGSS(data: GeoJSONFeatureCollection): HGSSDocument;
  fromHGSS(hgss: HGSSDocument): GeoJSONFeatureCollection;
}
```

### KMLConverter

Converts between HGSS and KML (XML string).

```typescript
class KMLConverter implements Converter {
  toHGSS(data: string): HGSSDocument; // KML as XML string
  fromHGSS(hgss: HGSSDocument): string; // Returns KML XML string
}
```

## Type Definitions

### HGSSDocument

```typescript
interface HGSSDocument {
  type: "HGSS";
  version: string;
  name?: string;
  description?: string;
  styles: Record<string, any>;
  root: Node;
}
```

### Node

```typescript
interface Node {
  id: string;
  type: string; // "Group" or GeoJSON geometry type
  name?: string;
  description?: string;
  coordinates?: [number, number] | [number, number, number];
  geometry?: Geometry;
  style?: Record<string, any>;
  children?: Node[];
}
```

### Geometry

```typescript
interface Geometry {
  type: string;
  coordinates: any;
}
```

### GeoJSON Types

```typescript
interface GeoJSONFeatureCollection {
  type: "FeatureCollection";
  features: GeoJSONFeature[];
}

interface GeoJSONFeature {
  type: "Feature";
  id?: string | number;
  geometry: GeoJSONGeometry;
  properties: Record<string, any>;
}

interface GeoJSONGeometry {
  type: string;
  coordinates: any;
}
```

## Reactive Behavior

The HGSSMirror uses JavaScript Proxies to automatically detect changes to the HGSS document:

```typescript
const mirror = new HGSSMirror(hgssDoc, { geojson: new GeoJSONConverter() });

// Changes to hgssDoc automatically update mirrors
hgssDoc.root.children.push({
  id: "newPoint",
  type: "Point",
  geometry: { type: "Point", coordinates: [1, 1] }
});

const geojson = mirror.get('geojson'); // Updated automatically
```

## Custom Converters

Create custom converters by implementing the Converter interface:

```typescript
class CustomConverter implements Converter {
  toHGSS(data: any): HGSSDocument {
    // Implement conversion logic
    return {
      type: "HGSS",
      version: "1.0",
      styles: {},
      root: {
        id: "root",
        type: "Group",
        children: [] // Convert data to HGSS nodes
      }
    };
  }

  fromHGSS(hgss: HGSSDocument): any {
    // Implement conversion logic
    return {}; // Convert HGSS to target format
  }
}

// Add to mirror
const mirror = new HGSSMirror(hgssDoc, {
  custom: new CustomConverter()
});
```

## Error Handling

Converters may throw errors for invalid input data. Always wrap operations in try-catch:

```typescript
try {
  const geojson = mirror.get('geojson');
  // Process geojson
} catch (error) {
  console.error('Conversion failed:', error);
}
```

## Performance Notes

- Mirrors are updated lazily when `get()` is called
- Deep object changes are tracked efficiently using Proxies
- Large documents may benefit from selective mirroring (only needed formats)