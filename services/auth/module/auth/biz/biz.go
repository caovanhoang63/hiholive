package biz

import (
	"github.com/caovanhoang63/hiholive/services/auth/module/auth/entity"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
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

type biz struct {
	serviceContext srvctx.ServiceContext
	repo           AuthRepository
}

func NewAuthBiz(serviceContext srvctx.ServiceContext, repo AuthRepository) *biz {
	return &biz{serviceContext: serviceContext, repo: repo}
}

func (b *biz) Register(x context.Context, register *entity.AuthRegister) error {
	field, err := core.Validator.ValidateField(register)

	old, err := b.repo.FindByEmail(x, register.Email)
	if old != nil {
		return errors.Errorf("Email existed!")
	}

	if err != nil {
		return errors.Errorf("INVALID_%s", field)
	}

	salt := core.GenSalt(50)
	auth := entity.NewAuthWithEmailPassword(0, register.Email, salt, register.Password)
	return b.repo.Create(x, &auth)
}

func (b *biz) Login(c context.Context, user *entity.AuthEmailPassword) (*entity.TokenResponse, error) {
	return nil, nil
}
