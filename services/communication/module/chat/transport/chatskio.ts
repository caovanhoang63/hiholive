import {IChatBusiness} from "../business/IBusiness";
import {DefaultEventsMap, Socket} from "socket.io";
import {ChatMessage, ChatMessageCreate, ChatMessageResponse} from "../model/model";
import {IRequester} from "../../../libs/IRequester";
import {createInvalidRequestError, createUnauthorizedError} from "../../../libs/errors";
import {UID} from "../../../libs/uid";
import {DbTypeStream} from "../../../libs/dbType";
import {User} from "../../user/model/user";
import {ChatMessageFilter} from "../model/chatMessageFilter";
import {Paging} from "../../../libs/paging";
import {AppResponse} from "../../../libs/response";


export class ChatSkio {
    private readonly _chatBusiness : IChatBusiness ;
    constructor(chatBusiness : IChatBusiness) {
        this._chatBusiness = chatBusiness
    }
    async listChatMessage(socket:  Socket<DefaultEventsMap, DefaultEventsMap, DefaultEventsMap, any>,
                          message : any,
                          ) {
        if (!message.filter) {
            socket.emit("listChat",createInvalidRequestError(new Error("filter required")))
            return
        }
        const streamId = UID.FromBase58(message.filter.streamId.toString())
        if (streamId.isErr()) {
            socket.emit("listChat",streamId.error)
            return
        }

        const filter : ChatMessageFilter = {
            streamId: streamId.value.localID
        }
        const paging : Paging = new Paging(message?.paging?.page || 0,message?.paging?.limit || 0)
        paging.default()

        let  oldCursor = undefined
        if (message.paging?.cursor) {
            oldCursor = message.paging?.cursor
            paging.cursor = {
                streamId :  streamId.value.localID,
                messageId : message.paging.cursor.messageId
            }
        }


        const r = await this._chatBusiness.list(filter,paging)
        r.match(
            list => {
                console.log(list)
                const res = list.map(v => {
                    return ({
                        streamId: streamId.value.toString(),
                        messageId: v.messageId.toString(),
                        user: {},
                        message: v.message,
                        createdAt: v.createdAt,
                        updatedAt: v.updatedAt,
                    } as ChatMessageResponse)
                })
                if (paging.nextCursor) {
                    paging.nextCursor.streamId = new UID(paging.nextCursor.streamId,DbTypeStream,1).toString();
                }
                paging.cursor = oldCursor
                const response = AppResponse.SuccessResponse(res,paging,{})
                socket.emit("listChat",response)
            },
            err => {
                console.log(err)
                socket.emit("listChat",err)
            }
        )
    }
    async sendMessage(socket:  Socket<DefaultEventsMap, DefaultEventsMap, DefaultEventsMap, any>,
                      requester : IRequester | null,
                      user: User,
                      message : ChatMessageCreate) {
        if (!requester) {
            socket.emit("sendMessage",createUnauthorizedError())
            return
        }
        const r = await this._chatBusiness.create(requester,message)
        r.match(
            r => {
                const roomId =new UID(message.streamId,DbTypeStream,1).toString()

                const mesRes : ChatMessageResponse= {
                    streamId: roomId,
                    messageId: message.messageId,
                    user: {
                        id: user.uid.toString(),
                        firstName: user.firstName,
                        lastName:user.lastName,
                        avatar: user.avatar
                    },
                    message:message.message,
                    createdAt: message.createdAt,
                    updatedAt: message.updatedAt,
                }
                socket.to(roomId).emit("newMessage",mesRes )
                socket.emit("sendMessage",true)

            },
            e => {
                console.log(e)
                socket.emit("sendMessage",e)
            }
        )
    }
}