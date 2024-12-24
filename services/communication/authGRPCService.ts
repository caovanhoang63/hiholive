import dotenv from "dotenv";
import {credentials} from "@grpc/grpc-js";
import {pb} from "./proto/pb/auth";
import AuthServiceClient = pb.AuthServiceClient;
dotenv.config()
const authAddress = process.env.GRPC_AUTH_ADDRESS!


export const authService = new AuthServiceClient(authAddress,credentials.createInsecure());

