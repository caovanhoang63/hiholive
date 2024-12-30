import {EmailTemplate} from "../model/emailTemplate";
import {ResultAsync} from "neverthrow";
import {EmailFilter} from "../model/filter";
import {Paging} from "../../../libs/paging";
import {SendEmailMessage} from "../model/sendEmailMessage";


export interface IEmailRepo {
    createEmailTemplate(email:  EmailTemplate) : ResultAsync<void, Error>
    listEmailTemplate(filter : EmailFilter, paging : Paging) :  ResultAsync<EmailTemplate[], Error>
    sendEmail(message: SendEmailMessage):  ResultAsync<void, Error>
}