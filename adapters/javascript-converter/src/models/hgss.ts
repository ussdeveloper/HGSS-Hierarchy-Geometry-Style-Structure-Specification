// HGSS Type Definitions

export interface HGSSDocument {
  type: "HGSS";
  version: string;
  name?: string;
  description?: string;
  styles: Record<string, any>;
  root: Node;
}

export interface Node {
  id: string;
  type: string; // "Group" or GeoJSON geometry type
  name?: string;
  description?: string;
  coordinates?: [number, number] | [number, number, number];
  geometry?: Geometry;
  style?: Record<string, any>;
  children?: Node[];
}

export interface Geometry {
  type: string;
  coordinates: any;
}

export interface StyleProperties {
  [key: string]: any;
}