import {DefaultEventsMap, Socket} from "socket.io";
import {container} from "../../../container";
import {IStreamBusiness} from "../business/IStreamBusiness";
import TYPES from "../../../types";
import {UID} from "../../../libs/uid";
import {AppResponse} from "../../../libs/response";

export const streamSkio = (socket:  Socket<DefaultEventsMap, DefaultEventsMap, DefaultEventsMap, any>) => {
    const streamBiz = container.get<IStreamBusiness>(TYPES.IStreamBusiness);
    const onViewStream = async (streamId : string, callBack : (data: any) => void  ) => {
        await UID.FromBase58(streamId).match(
            async r => {
                const  stream = await streamBiz.findStreamById(r.localID)
                if (stream.isErr()) {
                    console.log(stream.error)
                    socket.emit("joinStream", stream.error);
                }
                socket.data.streamId = streamId
                socket.join(streamId)
                callBack(AppResponse.SimpleResponse(true))
            },
            e => {
                callBack(AppResponse.ErrorResponse(e))
            }
        )

    }
    socket.on("stream:view",onViewStream)
}