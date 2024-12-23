import dotenv from "dotenv";
import * as AWS from "@aws-sdk/client-dynamodb";

dotenv.config();

const accessKey = process.env.S3_API_KEY
const secretAccessKey = process.env.S3_SECRET

export const dynamoClient = new AWS.DynamoDB({
    region: "ap-southeast-1",
    credentials : {
        accessKeyId : accessKey!,
        secretAccessKey : secretAccessKey!
    }
});
