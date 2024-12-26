import {ChatMessage, ChatMessageCreate} from "../model/model";
import {ResultAsync} from "neverthrow";
import {ChatMessageFilter} from "../model/chatMessageFilter";
import {Paging} from "../../../libs/paging";
import {IRequester} from "../../../libs/IRequester";


export interface IChatBusiness {
    create(requester : IRequester,create : ChatMessageCreate) : ResultAsync<void, Error>
    list(filter: ChatMessageFilter, paging: Paging) : ResultAsync<ChatMessage[], Error>
    delete(requester : IRequester, streamId : number, messageId: string) : ResultAsync<void, Error>
}