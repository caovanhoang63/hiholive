import {DefaultEventsMap, Socket} from "socket.io";
import {container} from "../../../container";
import TYPES from "../../../types";
import {IChatBusiness} from "../business/IBusiness";
import {createInternalError, createInvalidRequestError, createUnauthorizedError} from "../../../libs/errors";
import {UID} from "../../../libs/uid";
import {ChatMessageCreate, ChatMessageResponse} from "../model/model";
import {User} from "../../user/model/user";
import {AppResponse} from "../../../libs/response";
import {ChatMessageFilter} from "../model/chatMessageFilter";
import {Paging} from "../../../libs/paging";
import {Filter} from "@grpc/grpc-js/build/src/filter";
import {DbTypeStream} from "../../../libs/dbType";


export const chatSkio = (socket : Socket<DefaultEventsMap, DefaultEventsMap, DefaultEventsMap, any>) =>  {
    const chatBiz =  container.get<IChatBusiness>(TYPES.IChatBusiness)
    const onCreateMessage = async (message : any, callback?:(data : any) => void ) => {
        const user = socket.data.user as User
        const streamId = socket.data.streamId
        if (!user ) {
            callback?.(AppResponse.ErrorResponse(createUnauthorizedError()))
            return
        }
        if (!streamId) {
            callback?.(AppResponse.ErrorResponse(createInvalidRequestError(new Error("You are not viewing any stream"))))
            return
        }
        const id = UID.FromBase58(streamId)
        if (id.isErr()) {
            console.log(id.error)
            callback?.(createInternalError(id.isErr));
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
                callback?.(AppResponse.SimpleResponse(true))
            },
            e => {
                console.log(e)
                callback?.(createInternalError(e))
            }
        )
    }
    const onListMessage = async ({filter, paging} : {filter: ChatMessageFilter, paging: Paging}, callback?:(data : any) => void) => {
        console.log(filter,paging)
        if (!filter?.streamId) {
            callback?.(AppResponse.ErrorResponse(createInvalidRequestError(new Error("streamId is required"))))
            return
        }
        let oldCursor = paging.cursor
        paging = new Paging(paging?.limit || 0,paging?.page || 0)
        paging.default()
        paging.cursor = oldCursor
        const streamId = UID.FromBase58(filter.streamId.toString())
        if (streamId.isErr()) {
            callback?.(AppResponse.ErrorResponse(streamId.error))
            return
        }
        if (paging?.cursor?.streamId && paging?.cursor?.messageId ) {
            oldCursor = paging?.cursor
            paging.cursor = {
                streamId :  streamId.value.localID,
                messageId : paging.cursor.messageId
            }
        } else {
            paging.cursor = undefined
        }

        filter.streamId= streamId.value.localID
        const r = await chatBiz.list(filter,paging)
        r.match(
            list => {
                const res = list.map(v => {
                    return ({
                        streamId: streamId.value.toString(),
                        messageId: v.messageId.toString(),
                        user: {
                            id: v.user.uid.toString(),
                            firstName: v.user.firstName,
                            lastName:v.user.lastName,
                            avatar: v.user.avatar
                        },
                        message: v.message,
                        createdAt: v.createdAt,
                        updatedAt: v.updatedAt,
                    } as ChatMessageResponse)
                })
                if (paging.nextCursor) {
                    paging.nextCursor.streamId = new UID(paging.nextCursor.streamId,DbTypeStream,1).toString();
                }
                paging.cursor = oldCursor
                callback?.(AppResponse.SuccessResponse(res,paging,{}))
            },
            err => {
                console.log(err)
                callback?.(AppResponse.ErrorResponse(err))
            }
        )
    }
    socket.on("chat:create",onCreateMessage)
    socket.on("chat:list", onListMessage)
    socket.on("chat:delete", () => {})
}