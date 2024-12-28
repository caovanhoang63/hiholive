import {DefaultEventsMap, Server, Socket} from "socket.io";
import {chatSkio} from "../module/chat/transport/chatSkio";
import {userSkio} from "../module/user/transport/userSkio";
import {streamSkio} from "../module/stream/transport/streamSkio";


export const socketSetup = (
    io:  Server<DefaultEventsMap, DefaultEventsMap, DefaultEventsMap, any>,
    socket:  Socket<DefaultEventsMap, DefaultEventsMap, DefaultEventsMap, any>) => {
    streamSkio(io,socket)
    userSkio(socket)
    chatSkio(socket)
    socket.on("disconnect",()=> {
        console.log("Client disconnect")
    })
}