export namespace main {
	
	export class FileRecord {
	    filename: string;
	    type: string;
	    old_path: string;
	    new_path: string;
	    timestamp: string;
	
	    static createFrom(source: any = {}) {
	        return new FileRecord(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.filename = source["filename"];
	        this.type = source["type"];
	        this.old_path = source["old_path"];
	        this.new_path = source["new_path"];
	        this.timestamp = source["timestamp"];
	    }
	}
	export class OrganizeResult {
	    success: boolean;
	    error?: string;
	    moved: number;
	    files: FileRecord[];
	
	    static createFrom(source: any = {}) {
	        return new OrganizeResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.success = source["success"];
	        this.error = source["error"];
	        this.moved = source["moved"];
	        this.files = this.convertValues(source["files"], FileRecord);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class StatsResponse {
	    images: number;
	    videos: number;
	    docs: number;
	    music: number;
	    others: number;
	
	    static createFrom(source: any = {}) {
	        return new StatsResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.images = source["images"];
	        this.videos = source["videos"];
	        this.docs = source["docs"];
	        this.music = source["music"];
	        this.others = source["others"];
	    }
	}

}

