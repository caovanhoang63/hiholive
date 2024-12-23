import dotenv from "dotenv";
import * as AWS from "@aws-sdk/client-dynamodb";

dotenv.config();

const accessKey = process.env.S3_API_KEY
const secretAccessKey = process.env.S3_SECRET

const marshallOptions = {
    // Whether to automatically convert empty strings, blobs, and sets to `null`.
    // convertEmptyValues: false, // false, by default.
    // Whether to remove undefined values while marshalling.
    removeUndefinedValues: true, // false, by default.
    // Whether to convert typeof object to map attribute.
    convertClassInstanceToMap: true // false, by default. <---- Set this flag
}
const unmarshallOptions = {
    // Whether to return numbers as a string instead of converting them to native JavaScript numbers.
    // wrapNumbers: false, // false, by default.
}
export const dynamoClient = new AWS.DynamoDB({
    region: "ap-southeast-1",
    credentials : {
        accessKeyId : accessKey!,
        secretAccessKey : secretAccessKey!
    },
},);
