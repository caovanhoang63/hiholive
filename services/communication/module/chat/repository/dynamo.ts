import { ResultAsync } from "neverthrow";
import { Paging } from "../../../libs/paging";
import { Filter } from "../model/filter";
import { ChatMessageCreate, ChatMessage } from "../model/model";
import {IChatRepo} from "./IRepository";
import {client} from "../../../index";

export class ChatDynamoRepo implements IChatRepo {


    create(create: ChatMessageCreate): ResultAsync<void, Error> {
        throw new Error("Method not implemented.");
    }

    list(filter: Filter, paging: Paging): ResultAsync<ChatMessage[], Error> {
        throw new Error("Method not implemented.");
    }

    delete(streamId: number, messageId: string): ResultAsync<void, Error> {
        throw new Error("Method not implemented.");
    }
}