import {Paging} from "./paging";
import {createInternalError, Err} from "./errors";

export interface IResponse {
    data: any,
    paging?: Paging
    extra?: any
}

export interface IErrorResponse {
    message: string,
    code: number,
    key: string,
    metadata?: Record<string, unknown>;
}

export class AppResponse {
    public static SimpleResponse(data: any): IResponse {
        return {
            data: data
        }
    }

    public static SuccessResponse(data: any, paging: Paging, extra: any): IResponse {
        return {
            data: data,
            extra: extra,
            paging: paging
        }
    }

    public static ErrorResponse(err: Error ): IErrorResponse {
        if (err instanceof Err) {
            return {
                code: err.code,
                message: err.message,
                key: err.key,
                metadata: err.metadata
            }
        }
        return createInternalError(err)
    }
}