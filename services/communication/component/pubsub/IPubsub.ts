import {v7} from "uuid";
import {ResultAsync} from "neverthrow";

export interface Message {
    id: string;
    data: any;
    channel : string,
    createdAt : Date
}

export function createMessage(data: any) : Message {
    return {
        channel: "",
        createdAt: new Date(),
        data: data,
        id: v7()
    }
}

export interface IPubSub {
    publish(topic: string, message: Message): ResultAsync<void, Error>;
    subscribe(topic: string): ResultAsync<[Message[], () => void], Error>;
}
