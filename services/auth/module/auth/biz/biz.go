package biz

import (
	"github.com/caovanhoang63/hiholive/services/auth/common"
	"github.com/caovanhoang63/hiholive/services/auth/module/auth/entity"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type AuthBiz interface {
	Register(c context.Context, register *entity.AuthRegister) error
	Login(c context.Context, user *entity.AuthEmailPassword) (*entity.TokenResponse, error)
}

type AuthRepository interface {
	Create(ctx context.Context, register *entity.Auth) error
	FindByEmail(ctx context.Context, email string) (*entity.Auth, error)
}

type UserRepository interface {
	CreateUser(ctx context.Context, firstName, lastName, email string) (newId int, err error)
}

type biz struct {
	serviceContext srvctx.ServiceContext
	repo           AuthRepository
	userRepo       UserRepository
	jwtProvider    common.JWTProvider
}

func NewAuthBiz(serviceContext srvctx.ServiceContext, repo AuthRepository, userRepo UserRepository) *biz {
	return &biz{serviceContext: serviceContext, repo: repo, userRepo: userRepo}
}

func (b *biz) Register(x context.Context, register *entity.AuthRegister) error {
	field, err := core.Validator.ValidateField(register)
	if err != nil {
		return errors.Errorf("INVALID_%s", field)
	}

	old, err := b.repo.FindByEmail(x, register.Email)
	if old != nil {
		return errors.Errorf("Email existed!")
	}

	id, err := b.userRepo.CreateUser(x, register.FirstName, register.LastName, register.Email)
	if err != nil {
		return errors.New("ERROR")
	}

	salt := core.GenSalt(50)
	auth := entity.NewAuthWithEmailPassword(id, register.Email, salt, register.Password)
	return b.repo.Create(x, &auth)
}

func (b *biz) Login(c context.Context, user *entity.AuthEmailPassword) (*entity.TokenResponse, error) {
	return nil, nil
}

func (b *biz) IntrospectToken(ctx context.Context, accessToken string) (*jwt.RegisteredClaims, error) {
	claims, err := b.jwtProvider.ParseToken(ctx, accessToken)

	if err != nil {
		return nil, core.ErrUnauthorized.WithDebug(err.Error())
	}

	return claims, nil
}
