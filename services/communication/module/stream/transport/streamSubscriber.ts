import {id, inject, injectable} from "inversify";
import TYPES from "../../../types";
import {RedisClientType} from "redis";
import {errAsync, fromPromise, okAsync, ResultAsync} from "neverthrow";
import {createInternalError} from "../../../libs/errors";
import {ConsumerJob, Message} from "../../../component/pubsub/IPubsub";
import {UID} from "../../../libs/uid";
import {io} from "../../../index";
import {AppResponse} from "../../../libs/response";

@injectable()
export class StreamSubscriber   {
    constructor(@inject(TYPES.RedisClient) private redisClient : RedisClientType) {
    }
    activeStream (): ConsumerJob {
        const handler = (message: Message): ResultAsync<void, Error> => {
            return fromPromise((async ()  => {
                const id =  UID.FromBase58(message?.data)
                if (id.isErr())
                    return errAsync(id.error)
                console.log(id.value.localID)
                const r =  await this.redisClient.SADD("active_stream",id.value.localID.toString())
                console.log(r)
                return okAsync(undefined)
                })(),
                e => createInternalError(e) as Error
            ).andThen(r => r)
        }
        return {
            Title: "",
            Handler :handler
        }
    }

    endStream ():  ConsumerJob  {
        const handler =(message: Message): ResultAsync<void, Error> => {
            return fromPromise((async ()  => {
                    UID.FromBase58(message?.data?.stream_id).match(
                        async r => {
                            io.to(r.localID.toString()).emit("stream:end",AppResponse.SimpleResponse(true))
                            await this.redisClient.sRem("active_stream",r.localID.toString())
                        },
                        e => {
                            console.log(e)
                        }
                    )


                })(),
                e => createInternalError(e) as Error)
        }

        return {
            Title : "",
            Handler : handler
        }
    }

}
