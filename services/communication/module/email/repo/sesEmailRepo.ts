import {errAsync, fromPromise, okAsync, ResultAsync} from "neverthrow";
import { Paging } from "../../../libs/paging";
import { EmailTemplate } from "../model/emailTemplate";
import { EmailFilter } from "../model/filter";
import { SendEmailMessage } from "../model/sendEmailMessage";
import {IEmailRepo} from "./IEmailRepo";
import {inject, injectable} from "inversify";
import TYPES from "../../../types";
import {
    CreateTemplateCommand,
    GetTemplateCommand,
    ListTemplatesCommand,
    SendTemplatedEmailCommand,
    SESClient, UpdateTemplateCommand
} from "@aws-sdk/client-ses";
import {createInternalError} from "../../../libs/errors";


@injectable()
export class SesEmailRepo implements IEmailRepo {
    constructor(@inject(TYPES.SESClient)private readonly SESClient : SESClient) {
    }

    updateEmailTemplate(email: EmailTemplate): ResultAsync<void, Error> {
        const command = new UpdateTemplateCommand({
            Template: {
                TemplateName: email.templateName,
                TextPart: email.text,
                HtmlPart: email.html,
                SubjectPart: email.subject
            }
        })
        return fromPromise(this.SESClient.send(command),
            e => createInternalError(e)).
        andThen(r => {
            if (r.$metadata.httpStatusCode != 200) {
                return errAsync(createInternalError())
            }
            return okAsync(undefined)
        })
    }
    createEmailTemplate(email: EmailTemplate): ResultAsync<void, Error> {
        const command = new CreateTemplateCommand({
            Template: {
                TemplateName: email.templateName,
                TextPart: email.text,
                HtmlPart: email.html,
                SubjectPart: email.subject
            }
        })
        return fromPromise(this.SESClient.send(command),
                e => createInternalError(e)).
        andThen(r => {
            if (r.$metadata.httpStatusCode != 200) {
                return errAsync(createInternalError())
            }
            return okAsync(undefined)
        })
    }
    listEmailTemplate(filter: EmailFilter, paging: Paging): ResultAsync<EmailTemplate[], Error> {
        const command = new ListTemplatesCommand({
            NextToken : paging.cursor,
            MaxItems : paging.limit
        })
        return fromPromise(this.SESClient.send(command),
            e => createInternalError(e)).
        andThen(r => {
            if (r.$metadata.httpStatusCode != 200) {
                return errAsync(createInternalError())
            }
            if( !r.TemplatesMetadata ){
                return okAsync([])
            }
            paging.nextCursor = r.NextToken
            return fromPromise(Promise.all(
                r.TemplatesMetadata.map( t => {
                        const getTemplateCommand = new GetTemplateCommand({
                            TemplateName: t.Name
                        })
                        return this.SESClient.send(getTemplateCommand)
                    }
                )
            ), e => createInternalError(e))
        }).andThen(r=> {
            const res : EmailTemplate[] = []
            r.forEach(t => {
                if (t.Template &&t.$metadata.httpStatusCode)
                    res.push({
                        templateName : t.Template.TemplateName,
                        text : t.Template.TextPart,
                        html : t.Template.HtmlPart,
                        subject : t.Template.SubjectPart
                    } as EmailTemplate)
            })
            return okAsync(res)
        })

    }
    sendEmail(message: SendEmailMessage): ResultAsync<void, Error> {
        const command = new SendTemplatedEmailCommand({
            Destination: {
                ToAddresses : message.toAddresses,
                CcAddresses : message.ccAddress,
                BccAddresses :message.bccAddress,
            },
            Source: message.source,
            Template: message.template,
            TemplateData: message.templateData
        })
        return fromPromise(this.SESClient.send(command),
                e => createInternalError(e)).
        andThen( r => {
            if (r.$metadata.httpStatusCode == 200) {
                return okAsync(undefined)
            }
            return errAsync(createInternalError(new Error("Fail to send Email")))
        })
    }

}