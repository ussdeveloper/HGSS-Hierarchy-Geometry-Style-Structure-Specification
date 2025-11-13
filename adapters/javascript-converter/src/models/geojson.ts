// GeoJSON Type Definitions

export interface GeoJSONFeatureCollection {
  type: "FeatureCollection";
  features: GeoJSONFeature[];
}

export interface GeoJSONFeature {
  type: "Feature";
  id?: string | number;
  geometry: GeoJSONGeometry;
  properties: Record<string, any>;
}

export interface GeoJSONGeometry {
  type: string;
  coordinates: any;
}