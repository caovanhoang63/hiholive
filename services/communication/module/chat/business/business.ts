import {err, errAsync, fromPromise, ok, okAsync, ResultAsync} from "neverthrow";
import { IRequester } from "../../../libs/IRequester";
import { Paging } from "../../../libs/paging";
import { Filter } from "../model/filter";
import { ChatMessageCreate, ChatMessage } from "../model/model";
import {IChatBusiness} from "./IBusiness";
import {IChatRepo} from "../repository/IRepository";
import {createUnauthorizedError} from "../../../libs/errors";


export class ChatBusiness implements IChatBusiness {
    private chatRepo : IChatRepo;

    constructor(chatRepo : IChatRepo) {
        this.chatRepo = chatRepo;
    }

    create(requester: IRequester, create: ChatMessageCreate): ResultAsync<void, Error> {
        return fromPromise((async () => {
            if (!requester.userId) return err(createUnauthorizedError())
            create.userId = requester.userId
            const r =  await this.chatRepo.create(create)
            if (r.isErr()) {
                return err(r.error)
            }else {
                return ok(undefined)
            }
        })(), e => e as Error)
            .andThen(r=>r);
    }
    list(filter: Filter, paging: Paging): ResultAsync<ChatMessage[], Error> {
        return fromPromise((async () => {
            const r =  await this.chatRepo.list(filter,paging)
            if (r.isErr()) {
                return err(r.error);
            }else {
                return ok(r.value);
            }
        })(), e => e as Error)
            .andThen(r=>r);
    }
    delete(requester: IRequester, streamId: number, messageId: string): ResultAsync<void, Error> {

        throw new Error("Method not implemented.");
    }

}