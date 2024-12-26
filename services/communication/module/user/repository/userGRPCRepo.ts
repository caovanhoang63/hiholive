import {errAsync, fromPromise, okAsync, ResultAsync} from "neverthrow";
import {IUserRepo} from "./IUserRepo";
import { Nullable } from "../../../libs/nullable";
import {User} from "../model/user";
import {createEntityNotFoundError, createInternalError} from "../../../libs/errors";
import { UID } from "../../../libs/uid";
import {DbTypeUser} from "../../../libs/dbType";
import {inject, injectable} from "inversify";
import TYPES from "../../../types";
import {pb} from "../../../proto/pb/user";
import UserServiceClient = pb.UserServiceClient;
import GetUserByIdReq = pb.GetUserByIdReq;
import GetUsersByIdsReq = pb.GetUsersByIdsReq;

@injectable()
export class UserGRPCRepo implements IUserRepo {
    constructor(@inject(TYPES.UserGRPCClient) private _userService : UserServiceClient ) {
    }
    getUserById(id: number): ResultAsync<Nullable<User>, Error> {
       return fromPromise(new Promise<ResultAsync<Nullable<User>, Error>>((resolve,reject) => {
               this._userService.GetUserById(new GetUserByIdReq({id : id}), (e,r ) => {

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
            this._userService.GetUsersByIds(new GetUsersByIdsReq({ids: ids}), (e,r ) => {
                let users: User[] = []
                if (e) {
                    console.log(e)
                    return reject(errAsync(createInternalError(e)))
                }

                if (r) {
                    users = r?.users.map((user) => {
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