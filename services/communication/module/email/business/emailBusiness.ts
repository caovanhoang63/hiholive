import {IEmailBusiness} from "./IEmailBusiness";
import {inject, injectable} from "inversify";
import TYPES from "../../../types";
import {IEmailRepo} from "../repo/IEmailRepo";
import { ResultAsync } from "neverthrow";
import { Paging } from "../../../libs/paging";
import { EmailTemplate } from "../model/emailTemplate";
import { EmailFilter } from "../model/filter";
import { SendEmailMessage } from "../model/sendEmailMessage";

@injectable()
export class EmailBusiness implements IEmailBusiness {
    constructor(@inject(TYPES.IEmailRepository) private readonly EmailRepo: IEmailRepo) {
    }

    createEmailTemplate(email: EmailTemplate): ResultAsync<void, Error> {
        return this.EmailRepo.createEmailTemplate(email)
    }
    listEmailTemplate(filter: EmailFilter, paging: Paging): ResultAsync<EmailTemplate[], Error> {
        return this.EmailRepo.listEmailTemplate(filter,paging)
    }
    sendEmail(message: SendEmailMessage): ResultAsync<void, Error> {
        return this.EmailRepo.sendEmail(message)
    }
}