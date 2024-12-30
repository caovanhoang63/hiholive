export interface SendEmailMessage {
    source : string,
    toAddresses : string[],
    ccAddress : string[],
    bccAddress : string[],
    subject : string,
}