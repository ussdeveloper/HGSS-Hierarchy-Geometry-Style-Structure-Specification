import { HGSSDocument, Node } from '../models/hgss';
import { GeoJSONFeatureCollection, GeoJSONFeature } from '../models/geojson';

export class GeoJSONConverter {
  toHGSS(geojson: GeoJSONFeatureCollection): HGSSDocument {
    const doc: HGSSDocument = {
      type: "HGSS",
      version: "1.0",
      styles: {},
      root: {
        id: "root",
        type: "Group",
        children: geojson.features.map((feature, i) => this.featureToNode(feature, i))
      }
    };
    return doc;
  }

  fromHGSS(hgss: HGSSDocument): GeoJSONFeatureCollection {
    const features: GeoJSONFeature[] = [];
    this.flattenHGSS(hgss.root, features);
    return {
      type: "FeatureCollection",
      features
    };
  }

  private featureToNode(feature: GeoJSONFeature, index: number): Node {
    const node: Node = {
      id: feature.id ? feature.id.toString() : `feature_${index}`,
      type: feature.geometry.type,
      geometry: feature.geometry
    };

    if (feature.properties) {
      if (feature.properties.name) node.name = feature.properties.name;
      if (feature.properties.description) node.description = feature.properties.description;
      if (feature.properties.label_coordinates) node.coordinates = feature.properties.label_coordinates;
    }

    return node;
  }

  private flattenHGSS(node: Node, features: GeoJSONFeature[]): void {
    if (node.geometry) {
      const feature: GeoJSONFeature = {
        type: "Feature",
        geometry: node.geometry,
        properties: {
          name: node.name,
          description: node.description
        }
      };

      if (node.id) feature.id = node.id;
      if (node.coordinates) feature.properties.label_coordinates = node.coordinates;

      features.push(feature);
    }

    if (node.children) {
      node.children.forEach(child => this.flattenHGSS(child, features));
    }
  }
}