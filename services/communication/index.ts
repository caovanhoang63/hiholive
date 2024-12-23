import "reflect-metadata";
import express, {Express} from "express";
import dotenv from "dotenv";
import cors from "cors"
import helmet from "helmet";
import bodyParser from "body-parser";
import logger from "morgan"
import cookieParser from "cookie-parser";
import {createServer} from "node:http";
import {Server} from "socket.io";
import * as AWS from "@aws-sdk/client-dynamodb";
import {socketSetup} from "./setupSocket";
dotenv.config();

const accessKey = process.env.S3_API_KEY
const secretAccessKey = process.env.S3_SECRET

export const client = new AWS.DynamoDB({
    region: "ap-southeast-1",
    credentials : {
        accessKeyId : accessKey!,
        secretAccessKey : secretAccessKey!
    }
});


(async () => {
    const r =await client.listTables()
    console.log(r.TableNames)
})()

const app: Express = express();
const port = process.env.EXPRESS_PORT || 3000;



const httpServer = createServer(app);
export const io = new Server(httpServer);

io.on("connection", socketSetup);



(BigInt.prototype as any).toJSON = function () {
    return this.toString();
};




app.use(logger('dev'));
app.use(cors());
app.use(helmet());
app.use(bodyParser.json());
app.use(express.urlencoded({extended: false}));
app.use(cookieParser());

app.get("/ping",(req, res) => {
    res.status(200).json("pong")
})


httpServer.listen(port, () => {
    console.log(`[server]: Server is running at http://localhost:${port}`);
});

