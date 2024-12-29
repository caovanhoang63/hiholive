import {errAsync, fromPromise, okAsync, ResultAsync} from "neverthrow";
import {ConsumerJob, IPubSub, Message} from "./IPubsub";
import {inject, injectable, interfaces} from "inversify";
import TYPES from "../../types";
import {Channel, Connection, ConsumeMessage} from "amqplib";
import {createInternalError} from "../../libs/errors";


interface IQueueMap {
    [topic: string]: Message[][];
}


@injectable()
export class RabbitPubSub implements IPubSub {
    private channel : Channel | null = null ;
    private readonly channelMap: IQueueMap = {};
    private readonly messageQueue: Message[] = [];
    constructor(@inject(TYPES.RabbitMQClient) private client : Promise<Connection>) {

    }

    async start() {
        const conn = await this.client
        this.channel = await conn.createChannel()
        console.log("Rabbit PubSub Started")
    }

    publish(topic: string, message: Message): ResultAsync<void, Error>
    {
        if (this.channel == null) return errAsync(new Error("Channel wasn't created"))
        message.channel = topic
        return fromPromise(
            this.channel!.assertExchange(topic,"fanout",{
                durable: true,
            autoDelete : false,
            internal: false}),
                e => createInternalError(e)
        ).andThen( (t) => {
            const  r=  this.channel?.publish(topic,'',Buffer.from(JSON.stringify(message)))
            if (!r) {
                return errAsync(createInternalError(new Error(`Fail to publish event ${topic}`)))
            }
            console.log(`Publish event ${topic} with data ${JSON.stringify(message.data)}`)
            return okAsync(undefined)
        })
    }
    subscribe(topic: string,fn : ConsumerJob[] ): ResultAsync<void, Error> {
        if (this.channel == null) return errAsync(new Error("Channel wasn't created"))
        return fromPromise((async () => {
            const queue = await this.channel!.assertQueue(topic)
            await this.channel?.bindQueue(queue.queue,topic,"")
            await this.channel?.consume(queue.queue, msg => {
                const data =JSON.parse(msg?.content.toString() ?? "") as Message
                for (let i = 0; i < fn.length; i++) {
                    fn[i].Handler?.(data)
                }
            },{
                noAck : true,
            })
        })(),e=> createInternalError(e),)
    }
}