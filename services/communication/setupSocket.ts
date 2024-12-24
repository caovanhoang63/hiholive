import {DefaultEventsMap, Socket} from "socket.io";
import {Authentication} from "./module/rthandler/authentication";
import {IRequester} from "./libs/IRequester";
import {ChatDynamoRepo} from "./module/chat/repository/dynamo";
import {Paging} from "./libs/paging";
import {StreamRepo} from "./module/stream/repository/streamRepo";
import {StreamBusiness} from "./module/stream/business/streamBusiness";
import {UID} from "./libs/uid";

export const socketSetup = (socket:  Socket<DefaultEventsMap, DefaultEventsMap, DefaultEventsMap, any>) => {
    const repo = new ChatDynamoRepo()

    socket.on("authentication",async message => {
        let requester : IRequester | null = null
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

    socket.on("get", async message => {
        const r = await repo.list({streamId: 123},new Paging(1,20))
        if (r.isErr()) {
            console.log(r.error)
            return
        }
        console.log(r.value)

    })


    socket.on("delete", async message => {
        const r = await repo.delete(123,message.messageId)
        if (r.isErr()) {
            console.log(r.error)
            return
        }
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