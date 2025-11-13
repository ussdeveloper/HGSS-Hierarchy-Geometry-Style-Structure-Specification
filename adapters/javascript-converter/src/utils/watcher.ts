// Reactive Watcher using Proxy

export type ChangeCallback = (path: string[], oldValue: any, newValue: any) => void;

export function createReactiveProxy<T extends object>(
  target: T,
  callback: ChangeCallback,
  path: string[] = []
): T {
  return new Proxy(target, {
    get(target, property) {
      const value = (target as any)[property];
      if (typeof value === 'object' && value !== null) {
        return createReactiveProxy(value, callback, [...path, property as string]);
      }
      return value;
    },
    set(target, property, value) {
      const oldValue = (target as any)[property];
      (target as any)[property] = value;
      callback([...path, property as string], oldValue, value);
      return true;
    },
    deleteProperty(target, property) {
      const oldValue = (target as any)[property];
      delete (target as any)[property];
      callback([...path, property as string], oldValue, undefined);
      return true;
    }
  });
}