import {errAsync, fromPromise, okAsync, ResultAsync} from "neverthrow";
import { Stream } from "../model/stream";
import {IStreamRepo} from "../repository/IStreamRepo";
import {IStreamBusiness} from "./IStreamBusiness";

export class StreamBusiness implements IStreamBusiness {
    private readonly _streamRepo: IStreamRepo;

    constructor(streamRepo: IStreamRepo) {
        this._streamRepo = streamRepo
    }

    findStreamById(id: number): ResultAsync<Stream, Error> {
        return fromPromise((async () => {
            const r = await this._streamRepo.findStreamById(id)
            if (r.isErr()) {
                return errAsync(r.error)
            }
            return okAsync(r.value)
        })(),e => e as Error ). andThen(r => r )
    }
}