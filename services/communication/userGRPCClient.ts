import dotenv from "dotenv";

import {credentials} from "@grpc/grpc-js";
import {pb} from "./proto/pb/user";
import UserServiceClient = pb.UserServiceClient;

dotenv.config()

const userAddress = process.env.GRPC_USER_ADDRESS!
export const userService = new UserServiceClient(userAddress,credentials.createInsecure())
