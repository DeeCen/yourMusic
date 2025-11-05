export namespace api {
	
	export class FileHash {
	    Hash: string;
	
	    static createFrom(source: any = {}) {
	        return new FileHash(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Hash = source["Hash"];
	    }
	}
	export class Song {
	    SingerName?: string;
	    Image?: string;
	    FileHash?: string;
	    AlbumID?: string;
	    FileName?: string;
	    SQ?: FileHash;
	    HQ?: FileHash;
	
	    static createFrom(source: any = {}) {
	        return new Song(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.SingerName = source["SingerName"];
	        this.Image = source["Image"];
	        this.FileHash = source["FileHash"];
	        this.AlbumID = source["AlbumID"];
	        this.FileName = source["FileName"];
	        this.SQ = this.convertValues(source["SQ"], FileHash);
	        this.HQ = this.convertValues(source["HQ"], FileHash);
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
	export class RespAPISearchSongData {
	    total?: number;
	    lists?: Song[];
	
	    static createFrom(source: any = {}) {
	        return new RespAPISearchSongData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.total = source["total"];
	        this.lists = this.convertValues(source["lists"], Song);
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

}

export namespace main {
	
	export class RespGetSongURL {
	    size: number;
	    errMsg: string;
	    data: string[];
	    lyric: string;
	
	    static createFrom(source: any = {}) {
	        return new RespGetSongURL(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.size = source["size"];
	        this.errMsg = source["errMsg"];
	        this.data = source["data"];
	        this.lyric = source["lyric"];
	    }
	}
	export class RespLogin {
	    errMsg: string;
	    // Go type: struct { Userid int "json:\"userid,omitempty\""; Token string "json:\"token,omitempty\""; Pic string "json:\"pic,omitempty\""; Dfid string "json:\"dfid,omitempty\"" }
	    data: any;
	
	    static createFrom(source: any = {}) {
	        return new RespLogin(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.errMsg = source["errMsg"];
	        this.data = this.convertValues(source["data"], Object);
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
	export class RespMsg {
	    is_success: boolean;
	    msg: string;
	
	    static createFrom(source: any = {}) {
	        return new RespMsg(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.is_success = source["is_success"];
	        this.msg = source["msg"];
	    }
	}
	export class RespSearch {
	    errMsg: string;
	    data: api.RespAPISearchSongData;
	
	    static createFrom(source: any = {}) {
	        return new RespSearch(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.errMsg = source["errMsg"];
	        this.data = this.convertValues(source["data"], api.RespAPISearchSongData);
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

}

