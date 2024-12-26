import {IStreamRepo} from "./module/stream/repository/IStreamRepo";

const TYPES = {
    // Repository
    IChatRepository : Symbol.for("IChatRepository"),
    IUserRepository : Symbol.for("IUserRepository"),
    IStreamRepository: Symbol.for("IStreamRepository"),
    // GRPC Client
    UserGRPCClient : Symbol.for("IUserGRPCClient"),
    AuthGRPCClient : Symbol.for("IAuthGRPCClient"),
    StreamGRPCClient : Symbol.for("IStreamGRPCClient"),

    // Business
    IChatBusiness : Symbol.for("IChatBusiness"),
    IStreamBusiness: Symbol.for("IStreamBusiness"),


    // Rest


    // GRPC server


    // Infrastructure
    DynamoDBClient : Symbol.for("DynamoClient"),




}


export default TYPES;