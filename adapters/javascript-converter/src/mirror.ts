import { HGSSDocument } from './models/hgss';
import { createReactiveProxy, ChangeCallback } from './utils/watcher';

export interface Converter {
  toHGSS(data: any): HGSSDocument;
  fromHGSS(hgss: HGSSDocument): any;
}

export class HGSSMirror {
  private hgss: HGSSDocument;
  private converters: Record<string, Converter>;
  private mirrors: Record<string, any> = {};
  private isUpdating = false;

  constructor(hgss: HGSSDocument, converters: Record<string, Converter>) {
    this.hgss = hgss;
    this.converters = converters;

    // Create reactive proxy for HGSS
    this.hgss = createReactiveProxy(hgss, this.onHGSSChange.bind(this));

    // Initialize mirrors
    this.updateAllMirrors();
  }

  get(format: string): any {
    return this.mirrors[format];
  }

  setHGSS(hgss: HGSSDocument): void {
    this.hgss = createReactiveProxy(hgss, this.onHGSSChange.bind(this));
    this.updateAllMirrors();
  }

  addConverter(name: string, converter: Converter): void {
    this.converters[name] = converter;
    this.updateMirror(name);
  }

  private onHGSSChange(path: string[], oldValue: any, newValue: any): void {
    if (this.isUpdating) return;
    this.isUpdating = true;

    // Update all mirrors
    this.updateAllMirrors();

    this.isUpdating = false;
  }

  private updateAllMirrors(): void {
    for (const format in this.converters) {
      this.updateMirror(format);
    }
  }

  private updateMirror(format: string): void {
    const converter = this.converters[format];
    if (converter) {
      this.mirrors[format] = converter.fromHGSS(this.hgss);
    }
  }
}