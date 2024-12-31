import {inject, injectable} from "inversify";
import TYPES from "../../../types";
import {IEmailBusiness} from "../business/IEmailBusiness";
import {ConsumerJob, Message} from "../../../component/pubsub/IPubsub";
import {errAsync, fromPromise, okAsync, ResultAsync} from "neverthrow";
import {UID} from "../../../libs/uid";
import {createInternalError} from "../../../libs/errors";


@injectable()
export class EmailSub  {
    constructor(@inject(TYPES.IEmailBusiness) private readonly EmailBusiness : IEmailBusiness) {
    }
    sendForgotPasswordEmail() : ConsumerJob {
        const handler = (message: Message): ResultAsync<void, Error> => {
            return fromPromise((async ()  => {
                    console.log(message.data.email)
                    console.log(message.data.pin);
                    (await this.EmailBusiness.sendEmail({
                        template: "RESET_PASSWORD",
                        bccAddress: [],
                        ccAddress: [],
                        source: "no-reply@hiholive.fun",
                        toAddresses: [message.data.email],
                        templateData: JSON.stringify(message.data) || "{}"
                    })).match(
                        r => {

                        },
                        e => {
                            console.log(e)
                        }
                    )
                    return okAsync(undefined)
                })(),
                e => createInternalError(e) as Error
            ).andThen(r => r)
        }
        return {
            Title: "Send forgot password email",
            Handler :handler
        }
    }


}