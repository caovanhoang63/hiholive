import {DefaultEventsMap, Socket} from "socket.io";


export const socketSetup = (socket:  Socket<DefaultEventsMap, DefaultEventsMap, DefaultEventsMap, any>) => {
    console.log("Client connected");
    socket.on("disconnect",()=> {
        console.log("Client disconnect")
    })
}