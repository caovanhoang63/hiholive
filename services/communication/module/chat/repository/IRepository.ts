import {ChatMessage, ChatMessageCreate} from "../model/model";
import {ResultAsync} from "neverthrow";
import {ChatMessageFilter} from "../model/chatMessageFilter";
import {Paging} from "../../../libs/paging";

export interface IChatRepo {
    create(create : ChatMessageCreate) : ResultAsync<void, Error>
    list(filter: ChatMessageFilter, paging: Paging) : ResultAsync<ChatMessage[], Error>
    delete(streamId : number, messageId: string) : ResultAsync<void, Error>
}