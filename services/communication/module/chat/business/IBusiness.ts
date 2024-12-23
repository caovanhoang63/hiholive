import {ChatMessage, ChatMessageCreate} from "../model/model";
import {ResultAsync} from "neverthrow";
import {Filter} from "../model/filter";
import {Paging} from "../../../libs/paging";
import {IRequester} from "../../../libs/IRequester";

export interface IChatBusiness {
    create(requester : IRequester,create : ChatMessageCreate) : ResultAsync<void, Error>
    list(filter: Filter, paging: Paging) : ResultAsync<ChatMessage[], Error>
    delete(requester : IRequester, streamId : number, messageId: string) : ResultAsync<void, Error>
}