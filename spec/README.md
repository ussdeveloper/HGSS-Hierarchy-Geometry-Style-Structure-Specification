# HGSS Specification v1.0
**Hierarchy Geometry Style Structure**

**Status:** Draft Standard  
**Author:** Przemysław Lusina (SulacoLab)  
**Date:** 2025-11-13  
**License:** CC‑BY‑SA 4.0  
**Document ID:** HGSS‑RFC‑001

---

## Overview
**HGSS (Hierarchy Geometry Style Structure)** to otwarty standard JSON do reprezentowania danych przestrzennych w strukturze **hierarchicznej**, z pełnym wsparciem dla:
- stylizacji (CSS‑like style cascade),
- pozycji etykiet (`coordinates`),
- konwersji z/do popularnych formatów GIS (GeoJSON, KML, GML, TopoJSON).

---

## Core Principles
| Cecha | Opis |
|---|---|
| Hierarchia | zagnieżdżone `children` umożliwiają grupowanie warstw i obiektów |
| Kompatybilność | konwersja z GeoJSON, KML, GML, TopoJSON |
| Stylizacja | style globalne `{Type}` oraz selektory `#id`, inline overrides `node.style` |
| Wsparcie etykiet | każde `node` może mieć `coordinates` (lon/lat/alt) |
| JSON‑native | łatwy do parsowania, kompatybilny z JS/TS |
| Otwartość | licencja CC‑BY‑SA 4.0 |

---

## Document Structure
```jsonc
{
  "type": "HGSS",
  "version": "1.0",
  "name": "Optional title",
  "description": "Optional description",

  "styles": {
    "{Polygon}": { "stroke": "#444", "stroke-width": 1, "fill": "#ccc", "fill-opacity": 0.3 },
    "{Point}":   { "icon": "poi-default", "size": 10 },
    "#zone_a":     { "fill": "#ff0000", "fill-opacity": 0.4 }
  },

  "root": {
    "id": "root",
    "type": "Group",
    "name": "Warstwy",
    "coordinates": [19.945, 50.005],
    "children": [
      {
        "id": "districts",
        "type": "Group",
        "name": "Strefy",
        "coordinates": [19.949, 50.0075],
        "children": [
          {
            "id": "zone_a",
            "type": "Polygon",
            "name": "Strefa A",
            "description": "Strefa ścisłej ochrony.",
            "coordinates": [19.9445, 50.006],
            "geometry": {
              "type": "Polygon",
              "coordinates": [
                [[19.938,50.000],[19.950,50.000],[19.950,50.010],[19.938,50.010],[19.938,50.000]]
              ]
            },
            "style": { "fill-opacity": 0.5 }
          },
          {
            "id": "poi_center",
            "type": "Point",
            "name": "Punkt informacyjny",
            "geometry": { "type": "Point", "coordinates": [19.945, 50.005] }
          }
        ]
      }
    ]
  }
}
```

### Node Schema (semantic)
| Klucz | Typ | Opis |
|---|---|---|
| `id` | string | unikalny identyfikator węzła |
| `type` | string | `"Group"` lub typ geometrii GeoJSON (`Point`, `Polygon`, …) |
| `name` | string | (opcjonalnie) etykieta |
| `description` | string | (opcjonalnie) opis |
| `coordinates` | array | `[lon, lat, (alt)]` – pozycja etykiety |
| `geometry` | object | obiekt GeoJSON (opcjonalny) |
| `children` | array | zagnieżdżone węzły (opcjonalne) |
| `style` | object | lokalne nadpisania stylu (opcjonalne) |

---

## Styling
### Registry
- Klucze stylu:
  - `{Type}` – globalnie dla danego typu, np. `{Polygon}`, `{Point}`, `{Group}`,
  - `#id` – specyficznie dla danego elementu.
- Wartości: płaski obiekt właściwości (CSS‑like lub Mapbox‑like).

### Cascade Priority
1. `node.style` (inline)
2. `styles["#"+id]`
3. `styles["{"+Type+"}"]`
4. Dziedziczenie z rodzica

---

## Label Coordinates Logic
1. Jeśli `node.coordinates` → użyj.
2. W przeciwnym razie centroid geometrii:
   - `Point` → współrzędne punktu,
   - `LineString` → środek długości,
   - `Polygon` → centroid (shoelace formula),
   - `Multi*` → centroid ważony,
3. Jeśli `Group` bez geometrii → średnia centroidów dzieci.

---

## Conversion
### GeoJSON → HGSS
| GeoJSON | HGSS | Reguła |
|---|---|---|
| `FeatureCollection` | `root:{type:"Group"}` | każdy Feature → child |
| `Feature.id` | `node.id` | kopiuj/generuj |
| `Feature.properties.name` | `node.name` | |
| `Feature.properties.description` | `node.description` | |
| `Feature.geometry` | `node.geometry` | |
| `Feature.properties.label_coordinates` | `node.coordinates` | |
| (style brak) | `styles` | opcjonalne auto‑generowanie |

### HGSS → GeoJSON
| HGSS | GeoJSON | Reguła |
|---|---|---|
| `root` (tree) | `FeatureCollection` | spłaszczone |
| `node.id` | `Feature.id` | |
| `name`, `description` | `Feature.properties.*` | |
| `geometry` | `Feature.geometry` | |
| `coordinates` | `Feature.properties.label_coordinates` | |
| scalone style | `Feature.properties.style` | |

### KML → HGSS
| KML | HGSS | Reguła |
|---|---|---|
| `<Document>`/`<Folder>` | `type:"Group"` | hierarchia → children |
| `<Placemark>` | `Feature node` | typ wg geometrii |
| `<name>` | `name` | |
| `<description>` | `description` | |
| `<coordinates>` | `geometry.coordinates` / `node.coordinates` | |
| `<Style>` | `styles` | konwersja (fill/stroke/width/opacity/icon) |
| folder path | HGSS path | mapowanie na drzewo |

### HGSS → KML
- `Group` → `<Folder>`,
- Feature → `<Placemark>` z geometrią,
- `styles` → `<Style>` + `styleUrl`.

### GML / WFS → HGSS
| GML | HGSS |
|---|---|
| `<featureMember>` | `children[]` |
| `<gml:name>` | `name` |
| `<gml:description>` | `description` |
| `gml:*Geometry*` | `geometry` |
| `<gml:id>` | `id` |

---

## Geometry Type Mapping
| HGSS/GeoJSON | KML | GML |
|---|---|---|
| Point | `<Point>` | `gml:Point` |
| MultiPoint | `<MultiGeometry><Point>` | `gml:MultiPoint` |
| LineString | `<LineString>` | `gml:LineString` |
| MultiLineString | `<MultiGeometry><LineString>` | `gml:MultiCurve` |
| Polygon | `<Polygon>` | `gml:Polygon` |
| MultiPolygon | `<MultiGeometry><Polygon>` | `gml:MultiSurface` |
| GeometryCollection | `<MultiGeometry>` | `gml:CompositeGeometry` |
| Group | `<Folder>/<Document>` | logiczne grupowanie |

---

## JSON Schema
Zobacz: [`spec/hgss.schema.json`](hgss.schema.json)

Walidacja (AJV):
```bash
ajv validate -s spec/hgss.schema.json -d examples/demo.hgss.json
```

# HGSS Specification v1.0
**Hierarchy Geometry Style Structure**

**Status:** Draft Standard  
**Author:** Przemysław Lusina (SulacoLab)  
**Date:** 2025-11-13  
**License:** CC‑BY‑SA 4.0  
**Document ID:** HGSS‑RFC‑001

---

## Overview
**HGSS (Hierarchy Geometry Style Structure)** to otwarty standard JSON do reprezentowania danych przestrzennych w strukturze **hierarchicznej**, z pełnym wsparciem dla:
- stylizacji (CSS‑like style cascade),
- pozycji etykiet (`coordinates`),
- konwersji z/do popularnych formatów GIS (GeoJSON, KML, GML, TopoJSON).

---

## Core Principles
| Cecha | Opis |
|---|---|
| Hierarchia | zagnieżdżone `children` umożliwiają grupowanie warstw i obiektów |
| Kompatybilność | konwersja z GeoJSON, KML, GML, TopoJSON |
| Stylizacja | style globalne `{Type}` oraz selektory `#id`, inline overrides `node.style` |
| Wsparcie etykiet | każde `node` może mieć `coordinates` (lon/lat/alt) |
| JSON‑native | łatwy do parsowania, kompatybilny z JS/TS |
| Otwartość | licencja CC‑BY‑SA 4.0 |

---

## Document Structure
```jsonc
{
  "type": "HGSS",
  "version": "1.0",
  "name": "Optional title",
  "description": "Optional description",

  "styles": {
    "{Polygon}": { "stroke": "#444", "stroke-width": 1, "fill": "#ccc", "fill-opacity": 0.3 },
    "{Point}":   { "icon": "poi-default", "size": 10 },
    "#zone_a":     { "fill": "#ff0000", "fill-opacity": 0.4 }
  },

  "root": {
    "id": "root",
    "type": "Group",
    "name": "Warstwy",
    "coordinates": [19.945, 50.005],
    "children": [
      {
        "id": "districts",
        "type": "Group",
        "name": "Strefy",
        "coordinates": [19.949, 50.0075],
        "children": [
          {
            "id": "zone_a",
            "type": "Polygon",
            "name": "Strefa A",
            "description": "Strefa ścisłej ochrony.",
            "coordinates": [19.9445, 50.006],
            "geometry": {
              "type": "Polygon",
              "coordinates": [
                [[19.938,50.000],[19.950,50.000],[19.950,50.010],[19.938,50.010],[19.938,50.000]]
              ]
            },
            "style": { "fill-opacity": 0.5 }
          },
          {
            "id": "poi_center",
            "type": "Point",
            "name": "Punkt informacyjny",
            "geometry": { "type": "Point", "coordinates": [19.945, 50.005] }
          }
        ]
      }
    ]
  }
}
```

### Node Schema (semantic)
| Klucz | Typ | Opis |
|---|---|---|
| `id` | string | unikalny identyfikator węzła |
| `type` | string | `"Group"` lub typ geometrii GeoJSON (`Point`, `Polygon`, …) |
| `name` | string | (opcjonalnie) etykieta |
| `description` | string | (opcjonalnie) opis |
| `coordinates` | array | `[lon, lat, (alt)]` – pozycja etykiety |
| `geometry` | object | obiekt GeoJSON (opcjonalny) |
| `children` | array | zagnieżdżone węzły (opcjonalne) |
| `style` | object | lokalne nadpisania stylu (opcjonalne) |

---

## Styling
### Registry
- Klucze stylu:
  - `{Type}` – globalnie dla danego typu, np. `{Polygon}`, `{Point}`, `{Group}`,
  - `#id` – specyficznie dla danego elementu.
- Wartości: płaski obiekt właściwości (CSS‑like lub Mapbox‑like).

### Cascade Priority
1. `node.style` (inline)
2. `styles["#"+id]`
3. `styles["{"+Type+"}"]`
4. Dziedziczenie z rodzica

---

## Label Coordinates Logic
1. Jeśli `node.coordinates` → użyj.
2. W przeciwnym razie centroid geometrii:
   - `Point` → współrzędne punktu,
   - `LineString` → środek długości,
   - `Polygon` → centroid (shoelace formula),
   - `Multi*` → centroid ważony,
3. Jeśli `Group` bez geometrii → średnia centroidów dzieci.

---

## Conversion
### GeoJSON → HGSS
| GeoJSON | HGSS | Reguła |
|---|---|---|
| `FeatureCollection` | `root:{type:"Group"}` | każdy Feature → child |
| `Feature.id` | `node.id` | kopiuj/generuj |
| `Feature.properties.name` | `node.name` | |
| `Feature.properties.description` | `node.description` | |
| `Feature.geometry` | `node.geometry` | |
| `Feature.properties.label_coordinates` | `node.coordinates` | |
| (style brak) | `styles` | opcjonalne auto‑generowanie |

### HGSS → GeoJSON
| HGSS | GeoJSON | Reguła |
|---|---|---|
| `root` (tree) | `FeatureCollection` | spłaszczone |
| `node.id` | `Feature.id` | |
| `name`, `description` | `Feature.properties.*` | |
| `geometry` | `Feature.geometry` | |
| `coordinates` | `Feature.properties.label_coordinates` | |
| scalone style | `Feature.properties.style` | |

### KML → HGSS
| KML | HGSS | Reguła |
|---|---|---|
| `<Document>`/`<Folder>` | `type:"Group"` | hierarchia → children |
| `<Placemark>` | `Feature node` | typ wg geometrii |
| `<name>` | `name` | |
| `<description>` | `description` | |
| `<coordinates>` | `geometry.coordinates` / `node.coordinates` | |
| `<Style>` | `styles` | konwersja (fill/stroke/width/opacity/icon) |
| folder path | HGSS path | mapowanie na drzewo |

### HGSS → KML
- `Group` → `<Folder>`,
- Feature → `<Placemark>` z geometrią,
- `styles` → `<Style>` + `styleUrl`.

### GML / WFS → HGSS
| GML | HGSS |
|---|---|
| `<featureMember>` | `children[]` |
| `<gml:name>` | `name` |
| `<gml:description>` | `description` |
| `gml:*Geometry*` | `geometry` |
| `<gml:id>` | `id` |

---

## Geometry Type Mapping
| HGSS/GeoJSON | KML | GML |
|---|---|---|
| Point | `<Point>` | `gml:Point` |
| MultiPoint | `<MultiGeometry><Point>` | `gml:MultiPoint` |
| LineString | `<LineString>` | `gml:LineString` |
| MultiLineString | `<MultiGeometry><LineString>` | `gml:MultiCurve` |
| Polygon | `<Polygon>` | `gml:Polygon` |
| MultiPolygon | `<MultiGeometry><Polygon>` | `gml:MultiSurface` |
| GeometryCollection | `<MultiGeometry>` | `gml:CompositeGeometry` |
| Group | `<Folder>/<Document>` | logiczne grupowanie |

---

## JSON Schema
Zobacz: [`spec/hgss.schema.json`](hgss.schema.json)

Walidacja (AJV):
```bash
ajv validate -s spec/hgss.schema.json -d examples/demo.hgss.json
```

---

## Reference Snippets (JS)
```js
function resolveStyle(node, styles, parentStyle = {}) {
  const s = { ...parentStyle };
  if (styles[`{${node.type}}`]) Object.assign(s, styles[`{${node.type}}`]);
  if (styles[`#${node.id}`]) Object.assign(s, styles[`#${node.id}`]);
  if (node.style) Object.assign(s, node.style);
  return s;
}

function flattenHGSS(root, path="root", acc=[]) {
  const currentPath = `${path}/${root.id}`;
  if (root.geometry) {
    acc.push({
      type: "Feature",
      id: root.id,
      properties: {
        name: root.name,
        description: root.description,
        path: currentPath,
        style: root.style
      },
      geometry: root.geometry
    });
  }
  for (const c of (root.children || [])) flattenHGSS(c, currentPath, acc);
  return acc;
}
```

---

## Implementation Guidelines

### Parsing HGSS
- Use standard JSON parsers.
- Validate against `hgss.schema.json`.
- Resolve styles using the cascade priority.

### Rendering
- Traverse the tree recursively.
- For each node with geometry, apply resolved style.
- Position labels using `coordinates` or computed centroid.

### Extensions
- Custom properties in `node` or `style` are allowed via `additionalProperties`.
- New geometry types should follow GeoJSON conventions.

### Best Practices
- Use unique `id`s across the document.
- Prefer `{Type}` styles for consistency.
- Include `coordinates` for important labels.
- Keep hierarchy shallow for performance.

---

## Changelog
- **v1.0 (2025-11-13)**: Initial draft specification.
