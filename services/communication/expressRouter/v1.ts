import express from "express";
import {container} from "../container";
import {EmailExpress} from "../module/email/transport/emailExpress";
import TYPES from "../types";

export const v1Router = () => {
    const router = express.Router();
    const emailController = container.get<EmailExpress>(TYPES.EmailController);

    const emailRouter = express.Router();
    emailRouter.post("", emailController.create());
    emailRouter.get("", emailController.list());
    emailRouter.patch("", emailController.update());

    router.use("/email", emailRouter);

    return router;
};