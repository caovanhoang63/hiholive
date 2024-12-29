import {errAsync, fromPromise, okAsync, ResultAsync} from "neverthrow";
import {Paging} from "../../../libs/paging";
import {ChatMessageFilter} from "../model/chatMessageFilter";
import {ChatMessage, ChatMessageCreate, ChatMessageTableName} from "../model/model";
import {IChatRepo} from "./IRepository";
import {DeleteCommand, PutCommand, QueryCommand} from "@aws-sdk/lib-dynamodb";
import {createInternalError} from "../../../libs/errors";
import {UID} from "../../../libs/uid";
import {DbTypeUser} from "../../../libs/dbType";
import {IUserRepo} from "../../user/repository/IUserRepo";
import {inject, injectable} from "inversify";
import TYPES from "../../../types";
import * as AWS from "@aws-sdk/client-dynamodb";

@injectable()
export class ChatDynamoRepo implements IChatRepo {
    constructor(@inject(TYPES.IUserRepository) private _userRepo : IUserRepo ,
                @inject(TYPES.DynamoDBClient) private _dynamoClient : AWS.DynamoDB) {
    }
    create(create: ChatMessageCreate): ResultAsync<void, Error> {
        return fromPromise(this._dynamoClient.send(new PutCommand({
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
        return fromPromise(this._dynamoClient.send(new QueryCommand({
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
                                avatar: user.avatar,
                                systemRole: ""
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
        return fromPromise(this._dynamoClient.send(new DeleteCommand({
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