import {errAsync, fromPromise, okAsync, ResultAsync} from "neverthrow";
import {IRequester} from "../../libs/IRequester";
import {createUnauthorizedError} from "../../libs/errors";
import {UID} from "../../libs/uid";
import {authService} from "../../authGRPCService";
import {pb} from "../../proto/pb/auth";
import IntrospectReq = pb.IntrospectReq;

export const Authentication = (token: string) : ResultAsync<IRequester, Error > => {
    return fromPromise(new Promise<ResultAsync<IRequester, Error>>((resolve, reject) => {
        const requester = {} as IRequester;
        authService.IntrospectToken(new IntrospectReq({ access_token: token }), (e, r) => {
            if (e) {
                return resolve(errAsync(createUnauthorizedError()));
            }
            if (!r?.sub) {
                return resolve(errAsync(createUnauthorizedError()));
            }
            UID.FromBase58(r.sub).match(
                (res) => {
                    requester.userId = res.localID;
                    resolve(okAsync(requester));
                },
                (er) => {
                    resolve(errAsync(er));
                }
            );
        });
    }),e => e as Error ).andThen(r=>r);
}