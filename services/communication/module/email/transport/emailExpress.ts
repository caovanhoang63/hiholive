import {inject, injectable} from "inversify";
import TYPES from "../../../types";
import {IEmailBusiness} from "../business/IEmailBusiness";
import express from "express";
import {EmailTemplate} from "../model/emailTemplate";
import {AppResponse} from "../../../libs/response";
import {writeErrorResponse} from "../../../libs/writeErrorResponse";
import {ReqHelper} from "../../../libs/reqHelper";


@injectable()
export class EmailExpress {
    constructor(@inject(TYPES.IEmailBusiness) private readonly EmailBusiness : IEmailBusiness) {
    }
    base64UrlEncode(data : string) {
        return Buffer.from(data).toString('base64').replace(/\+/g, '-').replace(/\//g, '_').replace(/=+$/, '');
    }

    base64UrlDecode(data :string) {
        const paddedData = data.replace(/-/g, '+').replace(/_/g, '/');
        return Buffer.from(paddedData, 'base64').toString();
    }
    create() : express.Handler {
        return async (req , res)=> {
            const create = req.body as EmailTemplate

            (await this.EmailBusiness.createEmailTemplate(create)).match(
                r => {
                   res.status(200).json(AppResponse.SimpleResponse(true))
                },
                e => {
                    writeErrorResponse(res,e)
                }
            )

        }
    }

    list() : express.Handler {
        return async (req,res) =>{
            const paging = ReqHelper.getPaging(req.query);
            if (paging.cursor) {
                paging.cursor = this.base64UrlDecode(paging.cursor)
            }
            (await this.EmailBusiness.listEmailTemplate({},paging)).match(
                r => {
                    if (paging.nextCursor) {
                        paging.nextCursor = this.base64UrlEncode(paging.nextCursor)
                    }
                    if (paging.cursor) {
                        paging.cursor = this.base64UrlEncode(paging.cursor)
                    }
                    res.status(200).json(AppResponse.SuccessResponse(r,paging,{}))
                },
                e => {
                    writeErrorResponse(res,e)
                }
            )
        }
    }


}