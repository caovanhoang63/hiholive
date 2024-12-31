import {container} from "./container";
import {IPubSub} from "./component/pubsub/IPubsub";
import TYPES from "./types";
import {RedisClientType} from "redis";
import {StreamSubscriber} from "./module/stream/transport/streamSubscriber";
import {TopicForgotPassword, TopicStreamEnded, TopicStreamStart} from "./libs/topic";
import {EmailSub} from "./module/email/transport/emailSub";

export const SubscriberSetup = async () => {
    const ps = await container.getAsync<IPubSub>(TYPES.PubSub);
    await ps.start()

    const streamService = container.get<StreamSubscriber>(TYPES.StreamSubscriber)
    const emailService = container.get<EmailSub>(TYPES.EmailSubscriber)

    ps.subscribe(TopicStreamStart,[streamService.activeStream()])
    ps.subscribe(TopicStreamEnded,[streamService.endStream()])
    ps.subscribe(TopicForgotPassword,[emailService.sendForgotPasswordEmail()])
}