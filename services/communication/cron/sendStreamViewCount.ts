import {CronJob} from "cron";
import {RedisClientType} from "redis";
import {container} from "../container";
import TYPES from "../types";
import {io} from "../index";
import {createMessage, IPubSub} from "../component/pubsub/IPubsub";
import {TopicUpdateStreamViewCount} from "../libs/topic";

export const jobUpdateStreamViewCount = new CronJob(
    '*/10 * * * * *', // cronTime
    function () {
        const rdClient = container.get<RedisClientType>(TYPES.RedisClient)
        const ps = container.get<IPubSub>(TYPES.PubSub)
        rdClient.SMEMBERS("active_stream").then(
            r => {
                for (let i = 0; i < r.length; i++ ){
                    if (io.sockets.adapter.rooms.has(r[i])) {
                        // ps.publish(TopicUpdateStreamViewCount,createMessage({
                        //
                        // }))
                        console.log(io.sockets.adapter.rooms.get(r[i])?.size)
                    }
                }
            }
        )

        // console.log('You will see this message every second');
    }, // onTick
    null, // onComplete
    false, // start
);
