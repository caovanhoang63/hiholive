package composer

import (
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/caovanhoang63/hiholive/shared/go/proto/pb"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type hlsClient struct {
	grpcHlsClient pb.HlsServiceClient
}

func (ac *hlsClient) NewHlsStream(ctx context.Context, serverUrl, streamKey string) (err error) {
	_, err = ac.grpcHlsClient.NewHlsStream(ctx, &pb.NewHlsStreamReq{StreamKey: streamKey, ServerUrl: serverUrl})
	if err != nil {
		panic(err)
	}

	return err
}

func ComposeHlsRPCClient(serviceCtx srvctx.ServiceContext) *hlsClient {
	configComp := serviceCtx.MustGet(core.KeyCompConf).(core.Config)

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.Dial(configComp.GetGRPCHlsAddress(), opts)

	if err != nil {
		log.Fatal(err)
	}

	return &hlsClient{pb.NewHlsServiceClient(clientConn)}
}
