import {DefaultEventsMap, Socket} from "socket.io";
import {chatSkio} from "../module/chat/transport/chatSkio";
import {userSkio} from "../module/user/transport/userSkio";
import {streamSkio} from "../module/stream/transport/streamSkio";


export const socketSetup = (socket:  Socket<DefaultEventsMap, DefaultEventsMap, DefaultEventsMap, any>) => {
    streamSkio(socket)
    userSkio(socket)
    chatSkio(socket)
    socket.on("disconnect",()=> {
        console.log("Client disconnect")
    })
}