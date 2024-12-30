import {DefaultEventsMap, Socket} from "socket.io";
import {container} from "../../../container";
import {IUserRepo} from "../repository/IUserRepo";
import TYPES from "../../../types";
import {Authentication} from "../../rthandler/authentication";
import {createInternalError, createUnauthorizedError} from "../../../libs/errors";
import {User} from "../model/user";
import {AppResponse} from "../../../libs/response";

export const userSkio = (socket:  Socket<DefaultEventsMap, DefaultEventsMap, DefaultEventsMap, any> ) =>{
    const userRepo = container.get<IUserRepo>(TYPES.IUserRepository)
    const onAuthentication = async (token : string, callBack? : (data :any) => void )=>  {
        const result = await  Authentication(token)
        result.match(
            async ok => {
                if (!ok) {
                    socket.emit("authentication", createUnauthorizedError())
                    return
                }
                const ur = await userRepo.getUserById(ok.userId!)
                ur.match(
                    r => {
                        if (r) {
                            socket.data.user = {
                                id : r.id,
                                uid : r.uid,
                                lastName : r.lastName,
                                firstName : r.firstName,
                                systemRole : ok.userRole,
                                avatar : r.avatar,
                                userName :r.userName,
                                displayName :r.displayName
                            } as User
                        }
                        callBack?.(AppResponse.SimpleResponse(true))
                    },
                    e => {
                        console.log(e)
                        callBack?.(AppResponse.ErrorResponse(createInternalError(e)))
                    }
                )
            },
            e => {
                callBack?.(AppResponse.ErrorResponse(createUnauthorizedError()))
            }
        )
    }
    socket.on("user:authentication",onAuthentication)
}