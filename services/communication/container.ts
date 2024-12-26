import 'reflect-metadata';
import dotenv from "dotenv";
import {Container} from "inversify";
import {IChatRepo} from "./module/chat/repository/IRepository";
import TYPES from "./types";
import {ChatDynamoRepo} from "./module/chat/repository/dynamo";
import {IChatBusiness} from "./module/chat/business/IBusiness";
import {ChatBusiness} from "./module/chat/business/business";
import {createClient, RedisClientType, RedisDefaultModules, RedisFunctions, RedisModules, RedisScripts} from "redis";
import {pb as userPb} from "./proto/pb/user";
import UserServiceClient = userPb.UserServiceClient;
import {credentials} from "@grpc/grpc-js";
import {pb as streamPb} from "./proto/pb/stream";
import StreamServiceClient = streamPb.StreamServiceClient;
import {pb as authPb} from "./proto/pb/auth";
import AuthServiceClient = authPb.AuthServiceClient;
import * as AWS from "@aws-sdk/client-dynamodb";
import {IUserRepo} from "./module/user/repository/IUserRepo";
import {UserGRPCRepo} from "./module/user/repository/userGRPCRepo";
import { IStreamRepo } from './module/stream/repository/IStreamRepo';
import {StreamRepo} from "./module/stream/repository/streamRepo";
import {IStreamBusiness} from "./module/stream/business/IStreamBusiness";
import {StreamBusiness} from "./module/stream/business/streamBusiness";


dotenv.config();

const redisConnStr = process.env.REDIS_DSN || "";
const userAddress = process.env.GRPC_USER_ADDRESS || "";
const videoAddress = process.env.GRPC_VIDEO_ADDRESS || "";
const authAddress = process.env.GRPC_AUTH_ADDRESS || "";
const accessKey = process.env.S3_API_KEY || "";
const secretAccessKey = process.env.S3_SECRET || "";

const container = new Container();
// Repository
container.bind<IChatRepo>(TYPES.IChatRepository).to(ChatDynamoRepo).inRequestScope();
container.bind<IUserRepo>(TYPES.IUserRepository).to(UserGRPCRepo).inRequestScope();
container.bind<IStreamRepo>(TYPES.IStreamRepository).to(StreamRepo).inRequestScope();

// Business
container.bind<IChatBusiness>(TYPES.IChatBusiness).to(ChatBusiness).inRequestScope();
container.bind<IStreamBusiness>(TYPES.IStreamBusiness).to(StreamBusiness).inRequestScope();

container.bind<RedisClientType<RedisDefaultModules & RedisModules, RedisFunctions, RedisScripts>>("RedisClient").toDynamicValue( () => {
    const client = createClient({ url: redisConnStr });
    client.connect().then().catch().finally();
    return client;
}).inSingletonScope();

container.bind<UserServiceClient>(TYPES.UserGRPCClient).toDynamicValue(() => {
    return new UserServiceClient(userAddress,credentials.createInsecure())
}).inSingletonScope()

container.bind<StreamServiceClient>(TYPES.StreamGRPCClient).toDynamicValue(( ) => {
    return new StreamServiceClient(videoAddress,credentials.createInsecure())
}).inSingletonScope()

container.bind<AuthServiceClient>(TYPES.AuthGRPCClient).toDynamicValue(( ) => {
    return new AuthServiceClient(authAddress,credentials.createInsecure())
}).inSingletonScope()

container.bind<AWS.DynamoDB>(TYPES.DynamoDBClient).toDynamicValue(() => {
    return new AWS.DynamoDB({
        region: "ap-southeast-1",
        credentials : {
            accessKeyId : accessKey!,
            secretAccessKey : secretAccessKey!
        },
    },);
}).inSingletonScope();

export {container};
