import "reflect-metadata";
import express, {Express} from "express";
import dotenv from "dotenv";
import cors from "cors"
import helmet from "helmet";
import bodyParser from "body-parser";
import logger from "morgan"
import cookieParser from "cookie-parser";
import {createServer} from "node:http";
import {socketSetup} from "./setupSocket";
import {Server} from "socket.io";
dotenv.config();
const app: Express = express();
const port = process.env.EXPRESS_PORT || 3000;
const httpServer = createServer(app);
const io = new Server(httpServer);
(BigInt.prototype as any).toJSON = function () {
    return this.toString();
};

app.use(logger('dev'));
app.use(cors());
app.use(helmet());
app.use(bodyParser.json());
app.use(express.urlencoded({extended: false}));
app.use(cookieParser());
io.on("connection", socketSetup);

app.get("/ping",(req, res) => {
    res.status(200).json("pong")
});


(async () => {

})()

httpServer.listen(port, () => {
    console.log(`[server]: Server is running at http://localhost:${port}`);
});

