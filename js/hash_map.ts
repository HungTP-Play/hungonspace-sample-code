interface Hash{
    hash(key: any): number
}

class Hash32Shift implements Hash{
    hash(key: any): number {
        let hash = 0;
        for (let i = 0; i < key.length; i++) {
            hash = (hash << 5) - hash + key.charCodeAt(i);
            hash = hash & hash; // Convert to 32bit integer
        }
        return hash;
    }
}

class HashMap {
    private map: Map<number, any> = new Map<number, any>();
    constructor(private hash: Hash = new Hash32Shift()) {}

    public set(key: any, value: any): void {
        this.map.set(this.hash.hash(key), value);
    }

    public get(key: any): any {
        return this.map.get(this.hash.hash(key));
    }
}