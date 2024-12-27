import {fromPromise, ResultAsync} from "neverthrow";
import {IPubSub, Message} from "./IPubsub";
import {inject, injectable} from "inversify";
import TYPES from "../../types";
import {Connection} from "amqplib";
import {createInternalError} from "../../libs/errors";


@injectable()
export class RabbitPubSub implements IPubSub {
    constructor(@inject(TYPES.RabbitMQClient) private client : Connection) {
    }



    publish(topic: string, message: Message): ResultAsync<void, Error> {
        throw new Error("Method not implemented.");
    }
    subscribe(topic: string): ResultAsync<[Message[], () => void], Error> {
        throw new Error("Method not implemented.");
    }
}