package authbiz

import (
	"errors"
	"github.com/caovanhoang63/hiholive/services/auth/module/auth/authmodel"
	"github.com/caovanhoang63/hiholive/shared/golang/core"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/pubsub"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type AuthBiz interface {
	Register(c context.Context, register *authmodel.AuthRegister) error
	Login(c context.Context, user *authmodel.AuthEmailPassword) (*authmodel.TokenResponse, error)
	ForgotPassword(c context.Context, email string) error
	ResetPasswordWithPin(c context.Context, email, pin, password string) error
	ResetPasswordWithRequester(c context.Context, requester core.Requester, password string) error
	CheckForgotPasswordPin(c context.Context, email, pin string) error
}

type AuthRepository interface {
	Create(ctx context.Context, register *authmodel.Auth) error
	FindByEmail(ctx context.Context, email string) (*authmodel.Auth, error)
	FindByUserId(ctx context.Context, id int) (*authmodel.Auth, error)
	ForgotPassword(c context.Context, email, pin string) error
	CheckForgotPasswordPin(c context.Context, email, pin string) error
	UpdatePassword(c context.Context, email, password string) error
}

type UserRepository interface {
	CreateUser(ctx context.Context, firstName, lastName, email string) (newId int, err error)
}

type Hasher interface {
	Hash(string) string
}

type authBiz struct {
	serviceContext srvctx.ServiceContext
	repo           AuthRepository
	userRepo       UserRepository
	jwtProvider    core.JWTProvider
	hasher         Hasher
	ps             pubsub.Pubsub
}

func (b *authBiz) CheckForgotPasswordPin(c context.Context, email, pin string) error {
	if err := b.repo.CheckForgotPasswordPin(c, email, pin); err != nil {
		if errors.Is(err, authmodel.ErrInvalidPin) {
			return core.ErrBadRequest.WithError(authmodel.ErrInvalidPin.Error())
		}
		return core.ErrInternalServerError.WithWrap(err)
	}
	return nil
}

func (b *authBiz) ForgotPassword(c context.Context, email string) error {
	if _, err := b.repo.FindByEmail(c, email); err != nil {
		return core.ErrNotFound
	}

	pin := core.GenSalt(6)
	if err := b.repo.ForgotPassword(c, email, pin); err != nil {
		if errors.Is(err, authmodel.ErrPinAlreadyExists) {
			return core.ErrBadRequest.WithError(authmodel.ErrPinAlreadyExists.Error())
		}
		return core.ErrInternalServerError.WithWrap(err)
	}
	_ = b.ps.Publish(c, core.TopicForgotPassword, pubsub.NewMessage(map[string]interface{}{
		"email": email,
		"pin":   pin,
	}))
	return nil
}

func (b *authBiz) ResetPasswordWithPin(c context.Context, email, pin, password string) error {
	if err := b.repo.CheckForgotPasswordPin(c, email, pin); err != nil {
		if errors.Is(err, authmodel.ErrInvalidPin) {
			return core.ErrBadRequest.WithError(authmodel.ErrInvalidPin.Error())
		}
		return core.ErrInternalServerError.WithWrap(err)
	}

	if err := b.repo.UpdatePassword(c, email, password); err != nil {
		return core.ErrInternalServerError.WithWrap(err)
	}
	return nil
}

func (b *authBiz) ResetPasswordWithRequester(c context.Context, requester core.Requester, password string) error {
	if requester == nil {
		return core.ErrUnauthorized
	}

	auth, err := b.repo.FindByUserId(c, requester.GetUserId())
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			return core.ErrBadRequest.WithError("Invalid username or password")
		}
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	if err = b.repo.UpdatePassword(c, auth.Email, password); err != nil {
		return core.ErrInternalServerError.WithWrap(err)
	}
	return nil
}

func NewAuthBiz(serviceContext srvctx.ServiceContext, repo AuthRepository, userRepo UserRepository, jwtProvider core.JWTProvider, hasher Hasher, ps pubsub.Pubsub) *authBiz {
	return &authBiz{
		serviceContext: serviceContext,
		repo:           repo,
		userRepo:       userRepo,
		jwtProvider:    jwtProvider,
		hasher:         hasher,
		ps:             ps,
	}
}

func (b *authBiz) Register(x context.Context, register *authmodel.AuthRegister) error {
	field, err := core.Validator.ValidateField(register)
	if err != nil {
		return core.ErrInvalidInput(field)
	}

	old, err := b.repo.FindByEmail(x, register.Email)
	if old != nil {
		return core.ErrBadRequest.WithError("Email already exists")
	}

	id, err := b.userRepo.CreateUser(x, register.FirstName, register.LastName, register.Email)
	if err != nil {
		return core.ErrInternalServerError.WithWrap(err)
	}

	salt := core.GenSalt(50)
	auth := authmodel.NewAuthWithEmailPassword(id, register.Email, salt, b.hasher.Hash(register.Password+salt))
	return b.repo.Create(x, &auth)
}

func (b *authBiz) Login(c context.Context, user *authmodel.AuthEmailPassword) (*authmodel.TokenResponse, error) {
	old, err := b.repo.FindByEmail(c, user.Email)
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			return nil, core.ErrBadRequest.WithError("Invalid username or password")
		}
		return nil, core.ErrInternalServerError.WithDebug(err.Error())
	}

	if b.hasher.Hash(user.Password+old.Salt) != old.Password {
		return nil, core.ErrBadRequest.WithError("Invalid username or password")
	}

	uid := core.NewUID(uint32(old.UserId), 1, 1)
	sub := uid.String()
	tid := uuid.New().String()

	tokenStr, expSecs, err := b.jwtProvider.IssueToken(c, tid, sub)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError("Invalid username or password").WithDebug(err.Error())
	}

	return &authmodel.TokenResponse{
		AccessToken: authmodel.Token{
			Token:     tokenStr,
			ExpiredIn: expSecs,
		},
	}, nil
}

func (b *authBiz) IntrospectToken(ctx context.Context, accessToken string) (*jwt.RegisteredClaims, error) {
	claims, err := b.jwtProvider.ParseToken(ctx, accessToken)

	if err != nil {
		return nil, core.ErrUnauthorized.WithDebug(err.Error())
	}

	return claims, nil
}
