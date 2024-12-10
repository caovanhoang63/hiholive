package composer

import (
	"github.com/caovanhoang63/hiholive/services/auth/common"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/caovanhoang63/hiholive/shared/go/proto/pb"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func ComposeUserRPCClient(serviceCtx srvctx.ServiceContext) pb.UserServiceClient {
	configComp := serviceCtx.MustGet(core.KeyCompConf).(common.Config)
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.Dial(configComp.GetGRPCUserAddress(), opts)

	if err != nil {
		log.Fatal(err)
	}

	return pb.NewUserServiceClient(clientConn)
}
