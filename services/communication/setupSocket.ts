import {DefaultEventsMap, Socket} from "socket.io";
import {Authentication} from "./libs/rthandler/authentication";
import {IRequester} from "./libs/IRequester";


export const socketSetup = (socket:  Socket<DefaultEventsMap, DefaultEventsMap, DefaultEventsMap, any>) => {
    console.log("Client connected");
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

    socket.on("joinStream",async message => {




    })



    socket.on("disconnect",()=> {
        console.log("Client disconnect")
    })
}