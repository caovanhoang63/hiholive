import { Stream } from "../model/stream";
import {IStreamRepo} from "./IStreamRepo";
import {fromPromise, ok, okAsync, ResultAsync} from "neverthrow";
import {createEntityNotFoundError, createInternalError, createUnauthorizedError} from "../../../libs/errors";
import {videoService} from "../../../videoGRPCService";
import {pb} from "../../../proto/pb/stream";
import FindStreamReq = pb.FindStreamReq;


export class StreamRepo implements IStreamRepo {
    findStreamById(id: number): ResultAsync<Stream, Error> {
        return fromPromise(new Promise<ResultAsync<Stream, Error>>((resolve, reject) => {
            videoService.FindStreamById(new FindStreamReq({id: id}), (e, r) => {

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