import {DefaultEventsMap, Socket} from "socket.io";
import {container} from "../../../container";
import TYPES from "../../../types";
import {IChatBusiness} from "../business/IBusiness";
import {createInternalError, createInvalidRequestError, createUnauthorizedError} from "../../../libs/errors";
import {UID} from "../../../libs/uid";
import {ChatMessageCreate, ChatMessageResponse} from "../model/model";
import {User} from "../../user/model/user";
import {AppResponse} from "../../../libs/response";


export const chatSkio = (socket : Socket<DefaultEventsMap, DefaultEventsMap, DefaultEventsMap, any>) =>  {
    const chatBiz =  container.get<IChatBusiness>(TYPES.IChatBusiness)
    const onCreateMessage = async (message : any, callback:(data : any) => void ) => {
        const user = socket.data.user as User
        const streamId = socket.data.streamId
        if (!user || !streamId) {
            callback(AppResponse.ErrorResponse(createUnauthorizedError()))
            return
        }
        const id = UID.FromBase58(streamId)
        if (id.isErr()) {
            console.log(id.error)
            callback(createInternalError(id.isErr))
            return
        }
        const create = {message : message,streamId: id.value.localID}  as ChatMessageCreate

        const r = await chatBiz.create({userId : user.id, userRole: user.systemRole}, create)
        r.match(
            r => {
                const mesRes : ChatMessageResponse= {
                    streamId: streamId,
                    messageId: create.messageId,
                    user: {
                        id: user.uid.toString(),
                        firstName: user.firstName,
                        lastName:user.lastName,
                        avatar: user.avatar
                    },
                    message:create.message,
                    createdAt: create.createdAt,
                    updatedAt: create.updatedAt,
                }
                socket.to(streamId).emit("newMessage",mesRes )
                callback(AppResponse.SimpleResponse(true))
            },
            e => {
                console.log(e)
                callback(createInternalError(e))
            }
        )
    }

    const onListMessage = async (message : any, callback:(data : any) => void) => {

    }




    socket.on("chat:create",onCreateMessage)
    socket.on("chat:list", onListMessage)
    socket.on("chat:delete", () => {})
}