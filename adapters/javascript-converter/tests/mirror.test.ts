import { HGSSMirror, GeoJSONConverter } from '../src';

describe('HGSSMirror', () => {
  const sampleHGSS = {
    type: "HGSS" as const,
    version: "1.0",
    styles: {},
    root: {
      id: "root",
      type: "Group",
      children: [{
        id: "point1",
        type: "Point",
        name: "Test Point",
        geometry: {
          type: "Point",
          coordinates: [0, 0]
        }
      }]
    }
  };

  it('should create mirrors correctly', () => {
    const mirror = new HGSSMirror(sampleHGSS, {
      geojson: new GeoJSONConverter()
    });

    const geojson = mirror.get('geojson');
    expect(geojson.type).toBe('FeatureCollection');
    expect(geojson.features).toHaveLength(1);
    expect(geojson.features[0].properties.name).toBe('Test Point');
  });

  it('should update mirrors when HGSS changes', () => {
    const mirror = new HGSSMirror(sampleHGSS, {
      geojson: new GeoJSONConverter()
    });

    // Modify HGSS
    sampleHGSS.root.children[0].geometry.coordinates = [1, 1];

    const geojson = mirror.get('geojson');
    expect(geojson.features[0].geometry.coordinates).toEqual([1, 1]);
  });
});