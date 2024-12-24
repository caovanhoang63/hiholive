import {IChatBusiness} from "../business/IBusiness";
import {DefaultEventsMap, Socket} from "socket.io";
import {ChatMessage, ChatMessageCreate, ChatMessageResponse} from "../model/model";
import {IRequester} from "../../../libs/IRequester";
import {createUnauthorizedError} from "../../../libs/errors";
import {UID} from "../../../libs/uid";
import {DbTypeStream} from "../../../libs/dbType";
import {User} from "../../user/model/user";


export class ChatSkio {
    private readonly _chatBusiness : IChatBusiness ;
    constructor(chatBusiness : IChatBusiness) {
        this._chatBusiness = chatBusiness
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
            },
            e => {
                console.log(e)
                socket.emit("sendMessage",e)
            }
        )
    }
}