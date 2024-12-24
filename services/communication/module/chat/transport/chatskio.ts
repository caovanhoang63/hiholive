import {IChatBusiness} from "../business/IBusiness";
import {DefaultEventsMap, Socket} from "socket.io";
import {ChatMessage, ChatMessageCreate} from "../model/model";
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
                const mesRes = message as ChatMessage
                mesRes.user = user
                const roomId =new UID(message.streamId,DbTypeStream,1).toString()
                socket.to(roomId).emit("newMessage",mesRes )
            },
            e => {
                console.log(e)
                socket.emit("sendMessage",e)
            }
        )
    }
}