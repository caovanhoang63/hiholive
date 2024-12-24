import {createClient} from "redis";
import dotenv from "dotenv";

dotenv.config();


const redisConnStr = process.env.REDIS_DSN || "";

export const redisClient = await  createClient({
    url: redisConnStr
});

