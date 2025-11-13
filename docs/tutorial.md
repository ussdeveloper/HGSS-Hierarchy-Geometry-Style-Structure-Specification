# HGSS Tutorial

This tutorial will guide you through creating, styling, and converting HGSS documents.

## 1. Creating a Basic HGSS Document

Start with a simple document containing a point and a polygon:

```json
{
  "type": "HGSS",
  "version": "1.0",
  "name": "My First Map",
  "styles": {
    "{Point}": { "icon": "marker", "size": 12 },
    "{Polygon}": { "fill": "#blue", "stroke": "#000", "stroke-width": 2 }
  },
  "root": {
    "id": "root",
    "type": "Group",
    "children": [
      {
        "id": "point1",
        "type": "Point",
        "name": "Location A",
        "geometry": {
          "type": "Point",
          "coordinates": [0, 0]
        }
      },
      {
        "id": "poly1",
        "type": "Polygon",
        "name": "Area B",
        "geometry": {
          "type": "Polygon",
          "coordinates": [[[0,0],[1,0],[1,1],[0,1],[0,0]]]
        }
      }
    ]
  }
}
```

## 2. Adding Hierarchy

Group features into logical layers:

```json
{
  "root": {
    "id": "root",
    "type": "Group",
    "name": "Layers",
    "children": [
      {
        "id": "points_layer",
        "type": "Group",
        "name": "Points of Interest",
        "children": [
          {
            "id": "poi1",
            "type": "Point",
            "name": "Restaurant",
            "geometry": { "type": "Point", "coordinates": [10, 20] }
          }
        ]
      },
      {
        "id": "areas_layer",
        "type": "Group",
        "name": "Zones",
        "children": [
          {
            "id": "zone1",
            "type": "Polygon",
            "name": "Park",
            "geometry": { /* polygon coords */ }
          }
        ]
      }
    ]
  }
}
```

## 3. Styling Features

Apply styles globally and override specifically:

```json
{
  "styles": {
    "{Polygon}": { "fill": "#green", "fill-opacity": 0.5 },
    "#zone1": { "fill": "#red" }
  },
  "root": {
    "children": [
      {
        "id": "zone1",
        "type": "Polygon",
        "style": { "stroke-width": 3 }
      }
    ]
  }
}
```

The final style for `zone1` will be: green fill with 0.5 opacity (from {Polygon}), red fill (from #zone1), and 3px stroke (from inline).

## 4. Adding Labels

Specify label positions:

```json
{
  "root": {
    "children": [
      {
        "id": "feature1",
        "type": "Polygon",
        "name": "My Polygon",
        "coordinates": [0.5, 0.5],  // Label at center
        "geometry": { /* polygon */ }
      }
    ]
  }
}
```

If `coordinates` is omitted, the label will be placed at the geometry's centroid.

## 5. Converting from GeoJSON

Use the conversion function:

```js
const geojson = {
  type: "FeatureCollection",
  features: [
    {
      type: "Feature",
      id: "f1",
      properties: { name: "Point A" },
      geometry: { type: "Point", coordinates: [1, 2] }
    }
  ]
};

const hgss = convertGeoJSONToHGSS(geojson);
```

## 6. Rendering HGSS

Traverse the tree and render each feature:

```js
function renderHGSS(node, styles, parentStyle = {}) {
  const style = resolveStyle(node, styles, parentStyle);

  if (node.geometry) {
    // Render geometry with style
    renderGeometry(node.geometry, style);
  }

  if (node.children) {
    for (const child of node.children) {
      renderHGSS(child, styles, style);
    }
  }
}

// Usage
renderHGSS(hgss.root, hgss.styles);
```

## 7. Validation

Always validate your HGSS documents:

```js
const isValid = validateHGSS(myHGSSDocument);
if (!isValid) {
  console.error("Invalid HGSS document");
}
```

## Next Steps

- Explore advanced styling options
- Learn about conversion to KML/GML
- Implement custom renderers
- Contribute to the HGSS ecosystem