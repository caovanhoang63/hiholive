import {Stream} from "../model/stream";
import {ResultAsync} from "neverthrow";


export interface IStreamRepo {
    findStreamById(id : number ) : ResultAsync<Stream,Error>
}