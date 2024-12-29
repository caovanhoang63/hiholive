import {DefaultEventsMap, Server, Socket} from "socket.io";
import {container} from "../../../container";
import {IStreamBusiness} from "../business/IStreamBusiness";
import TYPES from "../../../types";
import {UID} from "../../../libs/uid";
import {AppResponse} from "../../../libs/response";
import {createInvalidRequestError} from "../../../libs/errors";

export const streamSkio = (
    io:  Server<DefaultEventsMap, DefaultEventsMap, DefaultEventsMap, any>,
    socket:  Socket<DefaultEventsMap, DefaultEventsMap, DefaultEventsMap, any>) => {
    const streamBiz = container.get<IStreamBusiness>(TYPES.IStreamBusiness);
    const onViewStream = async (streamId : string, callBack? : (data: any) => void  ) => {
        await UID.FromBase58(streamId).match(
            async r => {
                const  stream = await streamBiz.findStreamById(r.localID)
                if (stream.isErr()) {
                    console.log(stream.error)
                    callBack?.(stream.error);
                }
                socket.data.streamId = r.localID.toString()
                socket.join(r.localID.toString())
                callBack?.(AppResponse.SimpleResponse(true))
            },
            e => {
                callBack?.(AppResponse.ErrorResponse(createInvalidRequestError(e)))
            }
        )
    }
    const onLeaveStream = async (message : any, callBack? : (data: any) => void ) => {
        const streamId = socket.data.streamId
        if (!streamId) return
        if (socket.rooms.has(streamId)) {
            socket.data.streamId = undefined
            socket.leave(streamId)
        }
    }
    const onGetView = async (streamId : string,callBack? : (data: any) => void) => {
        if (!streamId) return
        UID.FromBase58(streamId).match(
            r => {
                const id = r.localID.toString()
                if (io.sockets.adapter.rooms.has(id)) {
                    callBack?.(AppResponse.SimpleResponse(io.sockets.adapter.rooms.get(id)?.size))
                    return
                }
                callBack?.(AppResponse.ErrorResponse(createInvalidRequestError(new Error("stream not found or not have any viewer"))))
            },
            e => {
                callBack?.(AppResponse.ErrorResponse(createInvalidRequestError(e)))
            }
        )
    }
    socket.on("stream:view",onViewStream)
    socket.on("stream:leave",onLeaveStream)
    socket.on("stream:getView",onGetView)
    socket.on("disconnect",onLeaveStream)
}