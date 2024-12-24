import {err, ok, Result} from "neverthrow";
import base58 from "bs58";
export class UID   {

    private readonly _localID: number;
    private readonly _objectType: number;
    private readonly _shardID: number;
    constructor(localID: number, objectType: number, shardID: number) {
        this._localID = localID;
        this._objectType = objectType;
        this._shardID = shardID;
    }
    public get shardID(): number {
        return this._shardID;
    }
    public get objectType(): number {
        return this._objectType;
    }
    public  get localID(): number {
        return this._localID;
    }
    static FromBase58(s : string)  {
        const decodedBuffer = base58.decode(s);

        const decodedStr = String.fromCharCode(...decodedBuffer);
            return UID.DecomposeUID(decodedStr);
    }
    static DecomposeUID(s:  string) : Result<UID,Error>  {
        const uid = BigInt(s)
        if (!uid) return err(new Error("Invalid UID"))
        console.log(uid )
        const u  = new UID (
            Number (uid >> BigInt(28)),
            Number((uid >> BigInt(18)) & BigInt(0x3FF)),
            Number((uid >> BigInt(0)) & BigInt(0x3FFF))
        )
        console.log(u.localID)
        return ok(u)
    }
    toString(): string {
        const val =
            ((this._localID) << 28) |
            ((this._objectType) << 18) |
            (this._shardID);
        return base58.encode(Buffer.from(val.toString()));
    }
}

