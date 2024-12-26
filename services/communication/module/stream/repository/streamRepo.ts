import {Stream} from "../model/stream";
import {IStreamRepo} from "./IStreamRepo";
import {fromPromise, okAsync, ResultAsync} from "neverthrow";
import {createEntityNotFoundError, createInternalError} from "../../../libs/errors";
import {pb} from "../../../proto/pb/stream";
import {inject, injectable} from "inversify";
import TYPES from "../../../types";
import FindStreamReq = pb.FindStreamReq;
import StreamServiceClient = pb.StreamServiceClient;


@injectable()
export class StreamRepo implements IStreamRepo {
    constructor(@inject(TYPES.StreamGRPCClient) private _videoService : StreamServiceClient) {
    }
    findStreamById(id: number): ResultAsync<Stream, Error> {
        return fromPromise(new Promise<ResultAsync<Stream, Error>>((resolve, reject) => {
            this._videoService.FindStreamById(new FindStreamReq({id: id}), (e, r) => {

                if (e) {
                    return reject(createInternalError(e))
                }

                if (!r) return reject(createEntityNotFoundError("stream"))

                return resolve(
                    okAsync(
                        {
                            state: r.state,
                            status: r.status,
                            title: r.title,
                            channelId: r.channel_id
                        } as Stream
                    )
                )
            });
        }),e => e as Error ).andThen(r=>r);
    }
}