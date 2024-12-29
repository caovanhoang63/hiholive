import {container} from "./container";
import {IPubSub} from "./component/pubsub/IPubsub";
import TYPES from "./types";
import {RedisClientType} from "redis";
import {StreamSubscriber} from "./module/stream/transport/streamSubscriber";
import {TopicStreamStart} from "./libs/topic";

export const SubscriberSetup = async () => {
    const ps = await container.getAsync<IPubSub>(TYPES.PubSub);
    await ps.start()

    const rdClient  =  container.get<RedisClientType>(TYPES.RedisClient);
    const streamService = new StreamSubscriber(rdClient)


    ps.subscribe(TopicStreamStart,[streamService.activeStream()])

}