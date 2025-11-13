# HGSS Tutorial

This tutorial will guide you through creating, styling, and converting HGSS documents using the JavaScript/TypeScript library.

## Installation

```bash
npm install @hgss/javascript-converter
```

## 1. Creating a Basic HGSS Document

Start with a simple document containing a point and a polygon:

```typescript
import { HGSSMirror, GeoJSONConverter } from '@hgss/javascript-converter';

const hgssDoc = {
  type: "HGSS" as const,
  version: "1.0",
  name: "My First Map",
  styles: {
    "{Point}": { "icon": "marker", "size": 12 },
    "{Polygon}": { "fill": "#blue", "stroke": "#000", "stroke-width": 2 }
  },
  root: {
    id: "root",
    type: "Group",
    children: [
      {
        id: "point1",
        type: "Point",
        name: "Location A",
        geometry: {
          type: "Point",
          coordinates: [0, 0]
        }
      },
      {
        id: "poly1",
        type: "Polygon",
        name: "Area B",
        geometry: {
          type: "Polygon",
          coordinates: [[[0,0],[1,0],[1,1],[0,1],[0,0]]]
        }
      }
    ]
  }
};
```

## 2. Adding Hierarchy

Group features into logical layers:

```typescript
const hgssDoc = {
  type: "HGSS" as const,
  version: "1.0",
  root: {
    id: "root",
    type: "Group",
    name: "Layers",
    children: [
      {
        id: "points_layer",
        type: "Group",
        name: "Points of Interest",
        children: [
          {
            id: "poi1",
            type: "Point",
            name: "Restaurant",
            geometry: { type: "Point", coordinates: [10, 20] }
          }
        ]
      },
      {
        id: "areas_layer",
        type: "Group",
        name: "Zones",
        children: [
          {
            id: "zone1",
            type: "Polygon",
            name: "Park",
            geometry: {
              type: "Polygon",
              coordinates: [[[0,0],[1,0],[1,1],[0,1],[0,0]]]
            }
          }
        ]
      }
    ]
  }
};
```

## 3. Styling Features

Apply styles globally and override specifically:

```typescript
const hgssDoc = {
  type: "HGSS" as const,
  version: "1.0",
  styles: {
    "{Polygon}": { "fill": "#green", "fill-opacity": 0.5 },
    "#zone1": { "fill": "#red" }
  },
  root: {
    id: "root",
    type: "Group",
    children: [
      {
        id: "zone1",
        type: "Polygon",
        style: { "stroke-width": 3 }
      }
    ]
  }
};

// The final style for zone1 will be:
// green fill with 0.5 opacity (from {Polygon}),
// red fill (from #zone1),
// and 3px stroke (from inline)
```

## 4. Adding Labels

Specify label positions:

```typescript
const hgssDoc = {
  type: "HGSS" as const,
  version: "1.0",
  root: {
    id: "root",
    type: "Group",
    children: [
      {
        id: "feature1",
        type: "Polygon",
        name: "My Polygon",
        coordinates: [0.5, 0.5],  // Label at center
        geometry: {
          type: "Polygon",
          coordinates: [[[0,0],[1,0],[1,1],[0,1],[0,0]]]
        }
      }
    ]
  }
};
```

If `coordinates` is omitted, the label will be placed at the geometry's centroid.

## 5. Creating Reactive Mirrors

Use HGSSMirror for automatic synchronization with other formats:

```typescript
import { HGSSMirror, GeoJSONConverter, KMLConverter } from '@hgss/javascript-converter';

// Create mirror with multiple formats
const mirror = new HGSSMirror(hgssDoc, {
  geojson: new GeoJSONConverter(),
  kml: new KMLConverter()
});

// Access mirrored formats
const geojson = mirror.get('geojson'); // GeoJSON FeatureCollection
const kml = mirror.get('kml');         // KML XML string

console.log(geojson.features.length); // Number of features
```

## 6. Reactive Updates

Changes to the HGSS document automatically update all mirrors:

```typescript
// Add a new point
hgssDoc.root.children.push({
  id: "point2",
  type: "Point",
  name: "New Location",
  geometry: { type: "Point", coordinates: [2, 2] }
});

// Mirrors update automatically
const updatedGeojson = mirror.get('geojson');
console.log(updatedGeojson.features.length); // 3 features now
```

## 7. Converting Between Formats

Direct conversion without reactive mirroring:

```typescript
import { GeoJSONConverter } from '@hgss/javascript-converter';

const converter = new GeoJSONConverter();

// Convert GeoJSON to HGSS
const geojsonData = {
  type: "FeatureCollection",
  features: [{
    type: "Feature",
    id: "f1",
    properties: { name: "Point A" },
    geometry: { type: "Point", coordinates: [1, 2] }
  }]
};

const hgss = converter.toHGSS(geojsonData);

// Convert HGSS back to GeoJSON
const backToGeojson = converter.fromHGSS(hgss);
```

## 8. Custom Converters

Create converters for other formats:

```typescript
import { HGSSMirror, Converter } from '@hgss/javascript-converter';

class TopoJSONConverter implements Converter {
  toHGSS(data: any): HGSSDocument {
    // Implement TopoJSON to HGSS conversion
    return {
      type: "HGSS",
      version: "1.0",
      styles: {},
      root: {
        id: "root",
        type: "Group",
        children: [] // Convert TopoJSON features
      }
    };
  }

  fromHGSS(hgss: HGSSDocument): any {
    // Implement HGSS to TopoJSON conversion
    return {}; // TopoJSON object
  }
}

// Add to mirror
const mirror = new HGSSMirror(hgssDoc, {
  geojson: new GeoJSONConverter(),
  topojson: new TopoJSONConverter()
});
```

## 9. Error Handling

Always handle conversion errors:

```typescript
try {
  const geojson = mirror.get('geojson');
  // Process data
} catch (error) {
  console.error('Conversion failed:', error.message);
}
```

## Next Steps

- Explore advanced styling options
- Learn about conversion to KML/GML/TopoJSON
- Implement custom renderers
- Contribute to the HGSS ecosystem