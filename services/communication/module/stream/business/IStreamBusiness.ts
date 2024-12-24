import {ResultAsync} from "neverthrow";
import {Stream} from "../model/stream";


export interface IStreamBusiness {
    findStreamById(id : number) : ResultAsync<Stream,Error>
}



