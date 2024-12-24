import {errAsync, fromPromise, okAsync, ResultAsync} from "neverthrow";
import {IUserRepo} from "./IUserRepo";
import { Nullable } from "../../../libs/nullable";
import {User} from "../model/user";
import {userService} from "../../../userGRPCClient";
import {pb} from "../../../proto/pb/user";
import GetUserByIdReq = pb.GetUserByIdReq;
import {createEntityNotFoundError, createInternalError} from "../../../libs/errors";
import { UID } from "../../../libs/uid";
import {DbTypeUser} from "../../../libs/dbType";
import GetUsersByIdsReq = pb.GetUsersByIdsReq;

export class UserGRPCRepo implements IUserRepo {
    constructor() {
    }
    getUserById(id: number): ResultAsync<Nullable<User>, Error> {
       return fromPromise(new Promise<ResultAsync<Nullable<User>, Error>>((resolve,reject) => {
               userService.GetUserById(new GetUserByIdReq({id : id}), (e,r ) => {

                   if (e) {
                       return reject(errAsync(createInternalError(e)))
                   }
                   if (!r) return reject(createEntityNotFoundError("user"))
                   if (r) {
                        return resolve(okAsync({
                           id : r.user.id,
                           uid : new UID(id, DbTypeUser,1),
                           avatar : r.user.avatar,
                           firstName : r.user.first_name,
                           lastName : r.user.last_name
                       } as User))
                   }
               })
       }
       ), e=> e as Error).andThen(r => r );
    }
    getUserByIds(ids: number[]): ResultAsync<User[], Error> {
        return fromPromise(new Promise<ResultAsync<User[], Error>>((resolve,reject) =>
            userService.GetUsersByIds(new GetUsersByIdsReq(ids), (e,r ) => {
                let users: User[] = []
                if (e) return  reject(errAsync(e));

                if (r) {
                    users = r?.users.map((user,i) => {
                        return {
                            id : user.id,
                            uid : new UID(user.id, DbTypeUser,1),
                            avatar : user.avatar,
                            firstName : user.first_name,
                            lastName : user.last_name
                        } as User
                    })
                }

                return resolve(okAsync(users))
            })
        ), e => e as Error ).andThen(r=> r )
    }
}