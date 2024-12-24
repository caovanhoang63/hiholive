import {User, UserRes} from "../../user/model/user";

export interface ChatMessage {
    streamId : number,
    messageId: string,
    user: User,
    userId : number,
    message: string,
    createdAt: Date,
    updatedAt: Date,
    status: number
}

export interface ChatMessageCreate {
    streamId : number,
    messageId: string,
    userId : number,
    createdAt: Date,
    updatedAt: Date,
    message: string,
}
export interface ChatMessageResponse {
    streamId : string,
    messageId: string,
    user: UserRes,
    message: string,
    createdAt: Date,
    updatedAt: Date,
}


export const ChatMessageTableName = "chatMessages";