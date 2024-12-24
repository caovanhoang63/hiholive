import {DefaultEventsMap, Socket} from "socket.io";
import {Authentication} from "./module/rthandler/authentication";
import {IRequester} from "./libs/IRequester";
import {ChatDynamoRepo} from "./module/chat/repository/dynamo";
import {Paging} from "./libs/paging";
import {StreamRepo} from "./module/stream/repository/streamRepo";
import {StreamBusiness} from "./module/stream/business/streamBusiness";
import {UID} from "./libs/uid";
import {ChatBusiness} from "./module/chat/business/business";
import {ChatSkio} from "./module/chat/transport/chatskio";
import {ChatMessageCreate} from "./module/chat/model/model";

export const socketSetup = (socket:  Socket<DefaultEventsMap, DefaultEventsMap, DefaultEventsMap, any>) => {
    let requester : IRequester | null = null

    socket.on("authentication",async message => {
        if (!message.token) {
            socket.emit("authentication","unauthorized")
            return
        }
        const result = await  Authentication(message.token)
        result.match(
            ok => {
                requester = ok
                socket.emit("authentication", "ok")
            },
            e => {
                console.log(e)
                socket.emit("authentication", "unauthorized")
            }
        )
    })

    socket.on("joinStream",async (message) => {
        if (!message.streamId) {
            socket.emit("joinStream", "error")
            return
        }
        const uid = message.streamId
        await UID.FromBase58(uid).match(
            async r => {
                const repo = new StreamRepo()
                const biz = new StreamBusiness(repo)
                const  stream = await biz.findStreamById(r.localID)
                if (stream.isErr()) {
                    console.log(stream.error)
                    socket.emit("joinStream", stream.error);
                }
                socket.join(message.streamId)

                const chatRepo = new ChatDynamoRepo()
                const chatBiz = new ChatBusiness(chatRepo)
                const chatSkio = new ChatSkio(chatBiz)

                socket.on("sendMessage",async (message) => {
                    message.streamId = r.localID
                    await chatSkio.sendMessage(socket,requester,message as ChatMessageCreate)
                })
            },
            e => {
                console.log(e)
                socket.emit("joinStream", e)
            }
        )
    })

    socket.on("disconnect",()=> {
        console.log("Client disconnect")
    })
}