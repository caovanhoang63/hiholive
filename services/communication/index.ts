import "reflect-metadata";
import express, {Express} from "express";
import dotenv from "dotenv";
import cors from "cors"
import helmet from "helmet";
import bodyParser from "body-parser";
import logger from "morgan"
import cookieParser from "cookie-parser";
import {createServer} from "node:http";
import {socketSetup} from "./socketio/setupSocket";
import {Server} from "socket.io";
import amqplib from "amqplib";
import {container} from "./container";
import TYPES from "./types";
import {IPubSub} from "./component/pubsub/IPubsub";
import {SubscriberSetup} from "./subcriber";
import {RedisClientType, RedisDefaultModules, RedisFunctions, RedisModules, RedisScripts} from "redis";
import {createAdapter} from "@socket.io/redis-adapter";
import {setupCronJobs} from "./cron";
dotenv.config();
const app: Express = express();
const port = process.env.EXPRESS_PORT || 3000;
const httpServer = createServer(app);

const rdClient = container.get
    <RedisClientType>(TYPES.RedisClient);
const pub = rdClient.duplicate();
const sub = pub.duplicate();

(async () => {
    await Promise.all([
        pub.connect(),
        sub.connect()
    ]);
})();



export const io = new Server(httpServer,{
    connectionStateRecovery: {},
    adapter: createAdapter(pub, sub)
});
(BigInt.prototype as any).toJSON = function () {
    return this.toString();
};

(async ()=>{
    await SubscriberSetup();
})()


setupCronJobs()
app.use(logger('dev'));
app.use(cors());
app.use(helmet());
app.use(bodyParser.json());
app.use(express.urlencoded({extended: false}));
app.use(cookieParser());
io.on("connection", (socket)  => socketSetup(io,socket));

app.get("/ping",(req, res) => {
    res.status(200).json("pong")
});

httpServer.listen(port, () => {
    console.log(`[server]: Server is running at http://localhost:${port}`);
});

