import {UID} from "../../../libs/uid";
import {Image} from "../../../libs/image";


export interface User {
    id : number,
    uid : UID,
    firstName : string,
    lastName: string,
    avatar : Image
    systemRole : string
}

export interface UserRes {
    id : string ,
    firstName : string,
    lastName: string,
    avatar : Image
}