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
            (await this.EmailBusiness.listEmailTemplate({},paging)).match(
                r => {
                    res.status(200).json(AppResponse.SuccessResponse(r,paging,{}))
                },
                e => {
                    writeErrorResponse(res,e)
                }
            )
        }
    }


}