import {container} from "./container";
import {IPubSub} from "./component/pubsub/IPubsub";
import TYPES from "./types";

export const SubscriberSetup = async () => {
    const ps = await container.getAsync<IPubSub>(TYPES.PubSub);
    await ps.start()
}