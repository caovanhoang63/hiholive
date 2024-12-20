package grpc

import (
	"github.com/caovanhoang63/hiholive/services/user/module/user/usermodel"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/caovanhoang63/hiholive/shared/go/proto/pb"
	"golang.org/x/net/context"
)

type Business interface {
	CreateNewUser(ctx context.Context, data *usermodel.UserCreate) error
	GetUserRole(ctx context.Context, userId int) (string, error)
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

func (s *grpcService) GetUserRole(ctx context.Context, req *pb.GetUserRoleReq) (*pb.GetUserRoleReps, error) {
	role, err := s.business.GetUserRole(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.GetUserRoleReps{Role: role}, nil
}
