package composer

import (
	"github.com/caovanhoang63/hiholive/shared/golang/core"
	"github.com/caovanhoang63/hiholive/shared/golang/proto/pb"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type hlsClient struct {
	grpcHlsClient pb.HlsServiceClient
}

func (ac *hlsClient) NewHlsStream(ctx context.Context, streamId string, serverUrl, streamKey string, fps, resolution int) (err error) {
	_, err = ac.grpcHlsClient.NewHlsStream(ctx, &pb.NewHlsStreamReq{
		StreamId:   streamId,
		StreamKey:  streamKey,
		ServerUrl:  serverUrl,
		Fps:        int32(fps),
		Resolution: int32(resolution),
	})
	if err != nil {
		return err
	}

	return nil
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
