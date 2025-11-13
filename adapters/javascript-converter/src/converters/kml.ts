import { HGSSDocument, Node } from '../models/hgss';
import { KMLDocument } from '../models/kml';

export class KMLConverter {
  toHGSS(kml: KMLDocument): HGSSDocument {
    // Simplified KML parsing - in real implementation, use a proper XML parser
    // For now, return a basic HGSS document
    const doc: HGSSDocument = {
      type: "HGSS",
      version: "1.0",
      styles: {},
      root: {
        id: "root",
        type: "Group",
        name: "KML Import",
        children: []
      }
    };
    // TODO: Implement proper KML parsing
    return doc;
  }

  fromHGSS(hgss: HGSSDocument): KMLDocument {
    let kml = '<?xml version="1.0" encoding="UTF-8"?>\n';
    kml += '<kml xmlns="http://www.opengis.net/kml/2.2">\n';
    kml += '<Document>\n';
    kml += `<name>${hgss.name || 'HGSS Export'}</name>\n`;

    // Add styles
    for (const [key, style] of Object.entries(hgss.styles)) {
      if (key.startsWith('#')) {
        kml += '<Style id="' + key.slice(1) + '">\n';
        // Add style details (simplified)
        kml += '</Style>\n';
      }
    }

    // Convert nodes
    kml += this.nodeToKML(hgss.root);

    kml += '</Document>\n';
    kml += '</kml>\n';
    return kml;
  }

  private nodeToKML(node: Node): string {
    let kml = '';

    if (node.type === 'Group') {
      kml += '<Folder>\n';
      if (node.name) kml += `<name>${node.name}</name>\n`;
      if (node.children) {
        node.children.forEach(child => {
          kml += this.nodeToKML(child);
        });
      }
      kml += '</Folder>\n';
    } else if (node.geometry) {
      kml += '<Placemark>\n';
      if (node.name) kml += `<name>${node.name}</name>\n`;
      if (node.description) kml += `<description>${node.description}</description>\n`;

      // Add geometry
      if (node.geometry.type === 'Point') {
        kml += '<Point>\n';
        kml += `<coordinates>${node.geometry.coordinates.join(',')}</coordinates>\n`;
        kml += '</Point>\n';
      } else if (node.geometry.type === 'Polygon') {
        kml += '<Polygon>\n';
        kml += '<outerBoundaryIs>\n';
        kml += '<LinearRing>\n';
        const coords = (node.geometry.coordinates as any[][])[0];
        kml += `<coordinates>${coords.map(c => c.join(',')).join(' ')}</coordinates>\n`;
        kml += '</LinearRing>\n';
        kml += '</outerBoundaryIs>\n';
        kml += '</Polygon>\n';
      }

      kml += '</Placemark>\n';
    }

    return kml;
  }
}