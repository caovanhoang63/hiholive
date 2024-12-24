import {createClient} from "redis";
import dotenv from "dotenv";

dotenv.config();


const redisConnStr = process.env.REDIS_DSN || "";

export const redisClient = await  createClient({
    url: redisConnStr
}).on('error', err => console.log('Redis Client Error', err)).connect();

