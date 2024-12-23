import {ResultAsync} from "neverthrow";
import {Paging} from "../../../libs/paging";
import {Filter} from "../model/filter";
import {ChatMessage, ChatMessageCreate, ChatMessageTableName} from "../model/model";
import {IChatRepo} from "./IRepository";
import {dynamoClient} from "../../../dynamoClient";
import {PutItemCommandInput} from "@aws-sdk/client-dynamodb";

export class ChatDynamoRepo implements IChatRepo {


    create(create: ChatMessageCreate): ResultAsync<void, Error> {
        const command : PutItemCommandInput = {
            Item: {

            },
            TableName : ChatMessageTableName
        }
        dynamoClient.putItem(command)

        throw Error("")
    }

    list(filter: Filter, paging: Paging): ResultAsync<ChatMessage[], Error> {
        throw new Error("Method not implemented.");
    }

    delete(streamId: number, messageId: string): ResultAsync<void, Error> {
        throw new Error("Method not implemented.");
    }
}