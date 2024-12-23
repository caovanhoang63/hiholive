export interface ChatMessage {
    streamId : number,
    messageId: string,
    userId : number,
    message: string,
    createdAt: Date,
    updatedAt: Date,
    status: number
}

export interface ChatMessageCreate {
    streamId : number,
    messageId: string,
    timestamp: Date,
    userId : number,
    createdAt: Date,
    updatedAt: Date,
    message: string,
}