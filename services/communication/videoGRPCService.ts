import dotenv from "dotenv";

import {credentials} from "@grpc/grpc-js";
import {pb} from "./proto/pb/stream";
import StreamServiceClient = pb.StreamServiceClient;

dotenv.config()


const videoAddress = process.env.GRPC_VIDEO_ADDRESS!
export const videoService = new StreamServiceClient(videoAddress,credentials.createInsecure())
