# HGSS API Reference

This document provides reference for JavaScript/TypeScript utilities to work with HGSS documents.

## Functions

### resolveStyle(node, styles, parentStyle = {})

Resolves the final style for a node by applying the cascade priority.

**Parameters:**
- `node` (object): HGSS node object
- `styles` (object): Global styles registry
- `parentStyle` (object): Inherited style from parent (default: {})

**Returns:** Object with resolved style properties

**Example:**
```js
const styles = {
  "{Polygon}": { fill: "#ccc" },
  "#zone_a": { fill: "#ff0000" }
};
const node = {
  id: "zone_a",
  type: "Polygon",
  style: { fillOpacity: 0.5 }
};
const resolved = resolveStyle(node, styles);
// { fill: "#ff0000", fillOpacity: 0.5 }
```

### flattenHGSS(root, path = "root", acc = [])

Flattens the HGSS tree into a GeoJSON FeatureCollection array.

**Parameters:**
- `root` (object): HGSS root node
- `path` (string): Current path in tree (default: "root")
- `acc` (array): Accumulator for features (default: [])

**Returns:** Array of GeoJSON Feature objects

**Example:**
```js
const hgss = { /* HGSS document */ };
const features = flattenHGSS(hgss.root);
console.log(features); // [{ type: "Feature", ... }, ...]
```

### computeCentroid(geometry)

Computes the centroid of a GeoJSON geometry.

**Parameters:**
- `geometry` (object): GeoJSON geometry object

**Returns:** Array [lon, lat] or null if unsupported

**Supported types:** Point, LineString, Polygon, MultiPoint, MultiLineString, MultiPolygon

### validateHGSS(document)

Validates an HGSS document against the JSON schema.

**Parameters:**
- `document` (object): HGSS document

**Returns:** Boolean - true if valid, throws error if invalid

### convertGeoJSONToHGSS(geojson)

Converts a GeoJSON FeatureCollection to HGSS.

**Parameters:**
- `geojson` (object): GeoJSON FeatureCollection

**Returns:** HGSS document object

### convertHGSStoGeoJSON(hgss)

Converts HGSS to GeoJSON FeatureCollection.

**Parameters:**
- `hgss` (object): HGSS document

**Returns:** GeoJSON FeatureCollection object

## Classes

### HGSSDocument

Class representing an HGSS document.

**Methods:**
- `constructor(data)`: Initialize with HGSS object
- `validate()`: Validate document
- `flatten()`: Return flattened GeoJSON features
- `resolveStyle(nodeId)`: Get resolved style for node by ID
- `getNodeById(id)`: Find node by ID
- `addNode(parentId, node)`: Add child node
- `removeNode(id)`: Remove node by ID

**Example:**
```js
const doc = new HGSSDocument(hgssData);
doc.validate();
const features = doc.flatten();
```

## TypeScript Types

```ts
interface HGSSDocument {
  type: "HGSS";
  version: string;
  name?: string;
  description?: string;
  styles: Record<string, StyleProperties>;
  root: Node;
}

interface Node {
  id: string;
  type: string;
  name?: string;
  description?: string;
  coordinates?: [number, number] | [number, number, number];
  geometry?: GeoJSON.Geometry;
  style?: StyleProperties;
  children?: Node[];
}

type StyleProperties = Record<string, any>;
```