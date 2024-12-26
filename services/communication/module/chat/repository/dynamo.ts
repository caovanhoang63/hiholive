import {err, errAsync, fromPromise, okAsync, ResultAsync} from "neverthrow";
import {Paging} from "../../../libs/paging";
import {ChatMessageFilter} from "../model/chatMessageFilter";
import {ChatMessage, ChatMessageCreate, ChatMessageTableName} from "../model/model";
import {IChatRepo} from "./IRepository";
import {dynamoClient} from "../../../dynamoClient";
import {DeleteCommand, PutCommand, QueryCommand} from "@aws-sdk/lib-dynamodb";
import {createInternalError} from "../../../libs/errors";
import {userService} from "../../../userGRPCClient";
import {pb} from "../../../proto/pb/user";
import GetUsersByIdsReq = pb.GetUsersByIdsReq;
import {User} from "../../user/model/user";
import {UID} from "../../../libs/uid";
import {DbTypeUser} from "../../../libs/dbType";
import {IUserRepo} from "../../user/repository/IUserRepo";
import {UserGRPCRepo} from "../../user/repository/userGRPCRepo";

export class ChatDynamoRepo implements IChatRepo {
    private _userRepo : IUserRepo = new UserGRPCRepo()
    constructor() {
    }
    create(create: ChatMessageCreate): ResultAsync<void, Error> {
        return fromPromise(dynamoClient.send(new PutCommand({
                TableName: ChatMessageTableName,
                Item: {
                    streamId : create.streamId,
                    messageId: create.messageId,
                    userId : create.userId,
                    createdAt: create.createdAt.toISOString(),
                    updatedAt: create.updatedAt.toISOString(),
                    message: create.message,
                },
            })),
        e => createInternalError(e))
            .andThen(r=> {
                if (r.$metadata.httpStatusCode != 200) {
                    return errAsync(createInternalError())
                }
                return okAsync(undefined);
            })
    }

    list(filter: ChatMessageFilter, paging: Paging): ResultAsync<ChatMessage[], Error> {
        let response: ChatMessage[] = []
        return fromPromise(dynamoClient.send(new QueryCommand({
            ScanIndexForward: false,
            KeyConditionExpression :"streamId = :streamId",
            ExpressionAttributeValues: {
                ":streamId" : filter.streamId
            },
            Limit: paging.limit,
            ExclusiveStartKey : paging.cursor ,
            TableName: ChatMessageTableName
        })),
            e => createInternalError(e)).andThen(
                 r=> {
                    if (r.$metadata.httpStatusCode != 200) {
                        return errAsync(createInternalError())
                    }
                    const ids = r.Items?.map(v => (v["userId"]) as number) || []
                     response = r.Items as ChatMessage[]
                     paging.nextCursor = r.LastEvaluatedKey;
                    return  this._userRepo.getUserByIds(ids)
            }
        ).andThen(
            r => {
                if (r) {
                    response = response.map((item) => {
                        const user = r.find(user => user.id === item.userId);
                        if (user) {
                            item.user = {
                                id: user.id,
                                uid: new UID(user.id, DbTypeUser, 1),
                                firstName: user.firstName,
                                lastName: user.lastName,
                                avatar: user.avatar
                            };
                        }
                        return item;
                    });
                }


                return okAsync(response);

            }
        )
    }

    delete(streamId: number, messageId: string): ResultAsync<void, Error> {
        return fromPromise(dynamoClient.send(new DeleteCommand({
                Key: {
                    "streamId" : streamId,
                    "messageId" : messageId
                },
                TableName: ChatMessageTableName
            })),
            e => createInternalError(e)).andThen(
            r=> {
                if (r.$metadata.httpStatusCode != 200) {
                    return errAsync(createInternalError())
                }
                return okAsync(undefined);
            }
        )
    }
}