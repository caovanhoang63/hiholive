package videocomposer

import (
	"github.com/caovanhoang63/hiholive/services/video/common"
	"github.com/caovanhoang63/hiholive/shared/golang/core"
	"github.com/caovanhoang63/hiholive/shared/golang/proto/pb"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type authClient struct {
	grpcAuthClient pb.AuthServiceClient
}

type userClient struct {
	grpcUserClient pb.UserServiceClient
}

func (ac *authClient) IntrospectToken(ctx context.Context, accessToken string) (sub string, tid string, err error) {
	resp, err := ac.grpcAuthClient.IntrospectToken(ctx, &pb.IntrospectReq{AccessToken: accessToken})

	if err != nil {
		return "", "", err
	}

	return resp.Sub, resp.Tid, nil
}

func (uc *userClient) GetUserRole(ctx context.Context, userId int) (string, error) {
	resp, err := uc.grpcUserClient.GetUserRole(ctx, &pb.GetUserRoleReq{Id: int32(userId)})

	if err != nil {
		return "", err
	}

	return resp.Role, nil
}

func (uc *userClient) GetUserById(ctx context.Context, id int) (*common.User, error) {
	data, err := uc.grpcUserClient.GetUserById(ctx, &pb.GetUserByIdReq{Id: int32(id)})
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}
	user := &common.User{
		Uid:         core.NewUIDP(uint32(data.User.Id), core.DbTypeUser, 1),
		FirstName:   data.User.FirstName,
		LastName:    data.User.LastName,
		UserName:    data.User.UserName,
		DisplayName: data.User.DisplayName,
		SystemRole:  data.User.SystemRole,
	}
	if data.User.Avatar != nil {
		user.Avatar = &core.Image{
			Id:        0,
			Url:       data.User.Avatar.Url,
			Width:     int(data.User.Avatar.Width),
			Height:    int(data.User.Avatar.Height),
			CloudName: data.User.Avatar.CloudName,
			Extension: data.User.Avatar.Extension,
		}
	}
	return user, nil
}

func ComposeUserRPCClient(serviceCtx srvctx.ServiceContext) *userClient {
	configComp := serviceCtx.MustGet(core.KeyCompConf).(core.Config)

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.Dial(configComp.GetGRPCUserAddress(), opts)

	if err != nil {
		log.Fatal(err)
	}

	return &userClient{pb.NewUserServiceClient(clientConn)}
}

// ComposeAuthRPCClient use only for middleware: get token info
func ComposeAuthRPCClient(serviceCtx srvctx.ServiceContext) *authClient {
	configComp := serviceCtx.MustGet(core.KeyCompConf).(core.Config)

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.Dial(configComp.GetGRPCAuthAddress(), opts)

	if err != nil {
		log.Fatal(err)
	}

	return &authClient{pb.NewAuthServiceClient(clientConn)}
}
