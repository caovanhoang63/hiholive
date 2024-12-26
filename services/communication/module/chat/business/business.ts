import {err, fromPromise, ok, ResultAsync} from "neverthrow";
import {IRequester} from "../../../libs/IRequester";
import {Paging} from "../../../libs/paging";
import {ChatMessageFilter} from "../model/chatMessageFilter";
import {ChatMessage, ChatMessageCreate} from "../model/model";
import {IChatBusiness} from "./IBusiness";
import {IChatRepo} from "../repository/IRepository";
import {createUnauthorizedError} from "../../../libs/errors";
import {v7} from "uuid";
import {inject} from "inversify";
import TYPES from "../../../types";

export class ChatBusiness implements IChatBusiness {
    constructor(@inject(TYPES.IChatRepository) private chatRepo: IChatRepo) {
    }

    create(requester: IRequester, create: ChatMessageCreate): ResultAsync<void, Error> {
        return fromPromise((async () => {
            if (!requester.userId) return err(createUnauthorizedError())
            create.userId = requester.userId
            create.createdAt = new Date()
            create.updatedAt = new Date()
            create.messageId = v7();
            const r = await this.chatRepo.create(create)
            if (r.isErr()) {
                return err(r.error)
            } else {
                return ok(undefined)
            }
        })(), e => e as Error)
            .andThen(r => r);
    }

    list(filter: ChatMessageFilter, paging: Paging): ResultAsync<ChatMessage[], Error> {
        return fromPromise((async () => {
            const r = await this.chatRepo.list(filter, paging)
            if (r.isErr()) {
                return err(r.error);
            } else {
                return ok(r.value);
            }
        })(), e => e as Error)
            .andThen(r => r);
    }

    delete(requester: IRequester, streamId: number, messageId: string): ResultAsync<void, Error> {

        throw new Error("Method not implemented.");
    }

}