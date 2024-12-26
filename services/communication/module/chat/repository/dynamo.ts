import {err, errAsync, fromPromise, okAsync, ResultAsync} from "neverthrow";
import {Paging} from "../../../libs/paging";
import {ChatMessageFilter} from "../model/chatMessageFilter";
import {ChatMessage, ChatMessageCreate, ChatMessageTableName} from "../model/model";
import {IChatRepo} from "./IRepository";
import {dynamoClient} from "../../../dynamoClient";
import {DeleteCommand, PutCommand, QueryCommand} from "@aws-sdk/lib-dynamodb";
import {createInternalError} from "../../../libs/errors";

export class ChatDynamoRepo implements IChatRepo {


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
                    paging.nextCursor = r.LastEvaluatedKey;
                    return okAsync(r.Items as ChatMessage[]);
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