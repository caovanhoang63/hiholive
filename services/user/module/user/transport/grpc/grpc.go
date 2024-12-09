package grpc

import (
	"github.com/caovanhoang63/hiholive/services/user/module/user/usermodel"
	"github.com/caovanhoang63/hiholive/services/user/proto/pb"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"golang.org/x/net/context"
)

type Business interface {
	CreateNewUser(ctx context.Context, data *usermodel.UserCreate) error
}

type grpcService struct {
	business Business
}

func NewService(business Business) *grpcService {
	return &grpcService{business: business}
}

func (s *grpcService) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.NewUserIdResp, error) {
	newUserData := usermodel.NewUserForCreation(req.FirstName, req.LastName, req.Email)

	if err := s.business.CreateNewUser(ctx, &newUserData); err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	return &pb.NewUserIdResp{Id: int32(newUserData.Id)}, nil
}
