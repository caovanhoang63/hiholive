package composer

import (
	"context"
	"github.com/caovanhoang63/hiholive/services/user/proto/pb"
	"github.com/caovanhoang63/hiholive/shared/go/shared"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type authClient struct {
	grpcAuthClient pb.AuthServiceClient
}

func (ac *authClient) IntrospectToken(ctx context.Context, accessToken string) (sub string, tid string, err error) {
	resp, err := ac.grpcAuthClient.IntrospectToken(ctx, &pb.IntrospectReq{AccessToken: accessToken})

	if err != nil {
		return "", "", err
	}

	return resp.Sub, resp.Tid, nil
}

// ComposeAuthRPCClient use only for middleware: get token info
func ComposeAuthRPCClient(serviceCtx srvctx.ServiceContext) *authClient {
	configComp := serviceCtx.MustGet(shared.KeyCompConf).(shared.Config)

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.Dial(configComp.GetGRPCAuthServerAddress(), opts)

	if err != nil {
		log.Fatal(err)
	}

	return &authClient{pb.NewAuthServiceClient(clientConn)}
}
