import {errAsync, fromPromise, okAsync, ResultAsync} from "neverthrow";
import { Stream } from "../model/stream";
import {IStreamRepo} from "../repository/IStreamRepo";
import {IStreamBusiness} from "./IStreamBusiness";
import {inject, injectable} from "inversify";
import TYPES from "../../../types";


@injectable()
export class StreamBusiness implements IStreamBusiness {
    constructor(@inject(TYPES.IStreamRepository)private readonly _streamRepo: IStreamRepo) {
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