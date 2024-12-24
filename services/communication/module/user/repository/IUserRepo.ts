import {ResultAsync} from "neverthrow";
import {User} from "../model/user";
import {Nullable} from "../../../libs/nullable";

export interface IUserRepo {
    getUserById(id : number) : ResultAsync<Nullable<User>, Error>
    getUserByIds(ids : number[]) : ResultAsync<User[], Error>
}