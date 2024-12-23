import dotenv from "dotenv";
import {pb} from "./proto/pb/auth";
import AuthServiceClient = pb.AuthServiceClient;
import {credentials} from "@grpc/grpc-js";
dotenv.config()
const authAddress = process.env.GRPC_AUTH_ADDRESS!

export const authService = new AuthServiceClient(authAddress,credentials.createInsecure());


