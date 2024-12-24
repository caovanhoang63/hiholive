import {DefaultEventsMap, Socket} from "socket.io";
import {Authentication} from "./libs/rthandler/authentication";
import {IRequester} from "./libs/IRequester";
import {ChatDynamoRepo} from "./module/chat/repository/dynamo";
import {v7 as uuidv7} from 'uuid';
import {Paging} from "./libs/paging";

export const socketSetup = (socket:  Socket<DefaultEventsMap, DefaultEventsMap, DefaultEventsMap, any>) => {
    console.log("Client connected");
    const repo = new ChatDynamoRepo()

    socket.on("authentication",async message => {
        let requester : IRequester | null = null
        if (!message.token) {
            socket.emit("authentication","unauthorized")
            return
        }
        const result   = await  Authentication(message.token)
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

    socket.on("joinStream",async message => {
        socket.join(message.streamId)



        const r = await repo.create({
            createdAt: new Date(),
            message: "asdasdas",
            messageId: uuidv7(),
            streamId: 123,
            updatedAt: new Date(),
            userId: 123
        })

        if (r.isErr()) {
            console.log(r.error)
        }


    })



    socket.on("disconnect",()=> {
        console.log("Client disconnect")
    })
}