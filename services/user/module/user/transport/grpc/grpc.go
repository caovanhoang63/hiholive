package grpc

import (
	"github.com/caovanhoang63/hiholive/services/user/module/user/usermodel"
	"github.com/caovanhoang63/hiholive/shared/golang/core"
	"github.com/caovanhoang63/hiholive/shared/golang/proto/pb"
	"golang.org/x/net/context"
)

type Business interface {
	CreateNewUser(ctx context.Context, data *usermodel.UserCreate) error
	GetUserRole(ctx context.Context, userId int) (string, error)
	FindUserById(ctx context.Context, id int) (*usermodel.User, error)
	FindUserByIds(ctx context.Context, ids []int) ([]usermodel.User, error)
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

func (s *grpcService) GetUserById(ctx context.Context, req *pb.GetUserByIdReq) (*pb.PublicUserInfoResp, error) {
	user, err := s.business.FindUserById(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	data := &pb.PublicUserInfoResp{
		User: &pb.PublicUserInfo{
			Id:        int32(user.Id),
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
	}

	if user.Avatar != nil {
		data.User.Avatar = &pb.Image{
			Id:        int64(user.Avatar.Id),
			Url:       user.Avatar.Url,
			Width:     int32(user.Avatar.Width),
			Height:    int32(user.Avatar.Height),
			CloudName: user.Avatar.CloudName,
			Extension: user.Avatar.Extension,
		}
	}
	return data, nil
}

func (s *grpcService) GetUsersByIds(ctx context.Context, req *pb.GetUsersByIdsReq) (*pb.PublicUsersInfoResp, error) {
	ids := make([]int, len(req.Ids))

	for i := range req.Ids {
		ids[i] = int(req.Ids[i])
	}

	users, err := s.business.FindUserByIds(ctx, ids)
	if err != nil {
		return nil, err
	}
	publicUserInfo := make([]*pb.PublicUserInfo, len(users))

	for i := range users {
		publicUserInfo[i] = &pb.PublicUserInfo{
			Id:        int32(users[i].Id),
			FirstName: users[i].FirstName,
			LastName:  users[i].LastName,
		}
		if users[i].Avatar != nil {
			publicUserInfo[i].Avatar = &pb.Image{
				Id:        int64(users[i].Avatar.Id),
				Url:       users[i].Avatar.Url,
				Width:     int32(users[i].Avatar.Width),
				Height:    int32(users[i].Avatar.Height),
				CloudName: users[i].Avatar.CloudName,
				Extension: users[i].Avatar.Extension,
			}
		}

	}
	return &pb.PublicUsersInfoResp{Users: publicUserInfo}, nil
}
