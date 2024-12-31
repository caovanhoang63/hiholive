import {IStreamRepo} from "./module/stream/repository/IStreamRepo";
import {StreamSubscriber} from "./module/stream/transport/streamSubscriber";

const TYPES = {
    // Repository
    IChatRepository : Symbol.for("IChatRepository"),
    IUserRepository : Symbol.for("IUserRepository"),
    IStreamRepository: Symbol.for("IStreamRepository"),
    IEmailRepository : Symbol.for("IEmailRepository"),
    // GRPC Client
    UserGRPCClient : Symbol.for("IUserGRPCClient"),
    AuthGRPCClient : Symbol.for("IAuthGRPCClient"),
    StreamGRPCClient : Symbol.for("IStreamGRPCClient"),

    // Business
    IChatBusiness : Symbol.for("IChatBusiness"),
    IStreamBusiness: Symbol.for("IStreamBusiness"),
    IEmailBusiness : Symbol.for("IEmailBusiness"),
    // Rest

    // Controller
    EmailController : Symbol.for("EmailController"),


    StreamSubscriber :Symbol.for("StreamSubscriber"),
    EmailSubscriber : Symbol.for("EmailSubscriber"),

    // GRPC server


    // Infrastructure
    PubSub : Symbol.for("PubSub"),
    DynamoDBClient : Symbol.for("DynamoClient"),
    RabbitMQClient : Symbol.for("RabbitMQClient"),
    RedisClient : Symbol.for("RedisClient"),
    SESClient : Symbol.for("SESClient")

}


export default TYPES;