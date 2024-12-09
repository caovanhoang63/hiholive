package biz

import (
	"context"
	"errors"
	"github.com/caovanhoang63/hiholive/services/user/module/user/usermodel"
	"github.com/caovanhoang63/hiholive/shared/go/core"
)

type UserBiz interface {
	CreateNewUser(ctx context.Context, data *usermodel.UserCreate) error
	DeleteUser(ctx context.Context, requester core.Requester, id int) error
	FindUserById(ctx context.Context, id int) (*usermodel.User, error)
	FindUserByIds(ctx context.Context, ids []int) ([]*usermodel.User, error)
	FindUsersWithCondition(ctx context.Context, filter *usermodel.UserFilter, paging *core.Paging) ([]usermodel.User, error)
	UpdateUser(ctx context.Context, requester core.Requester, id int, data *usermodel.UserUpdate) error
}

type UserRepo interface {
	CreateNewUser(ctx context.Context, data *usermodel.UserCreate) error
	DeleteUser(ctx context.Context, id int) error
	FindUserById(ctx context.Context, id int) (*usermodel.User, error)
	FindUserByIds(ctx context.Context, ids []int) ([]*usermodel.User, error)
	UpdateUser(ctx context.Context, id int, data *usermodel.UserUpdate) error
	FindUsersWithCondition(ctx context.Context, filter *usermodel.UserFilter, paging *core.Paging) ([]usermodel.User, error)
}

type userBiz struct {
	repo UserRepo
}

func (b *userBiz) GetUserRole(ctx context.Context, userId int) (string, error) {
	user, err := b.repo.FindUserById(ctx, userId)
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			return "", core.ErrNotFound
		} else {
			return "", core.ErrInternalServerError.WithWrap(err)
		}
	}
	return string(user.Role), nil
}

func (b *userBiz) FindUsersWithCondition(ctx context.Context, filter *usermodel.UserFilter, paging *core.Paging) ([]usermodel.User, error) {

	if field, err := core.Validator.ValidateField(filter); err != nil {
		return nil, core.ErrInvalidInput(field)
	}

	users, err := b.repo.FindUsersWithCondition(ctx, filter, paging)
	if err != nil {
		return nil, core.ErrInternalServerError.WithWrap(err)
	}

	return users, nil
}

func (b *userBiz) CreateNewUser(ctx context.Context, data *usermodel.UserCreate) error {

	if field, err := core.Validator.ValidateField(data); err != nil {
		return core.ErrInvalidInput(field)
	}
	data.Gender = usermodel.Other

	if err := b.repo.CreateNewUser(ctx, data); err != nil {
		return core.ErrInternalServerError.WithError(err.Error())
	}

	return nil
}

func (b *userBiz) DeleteUser(ctx context.Context, requester core.Requester, id int) error {
	//TODO implement me
	panic("implement me")
}

func (b *userBiz) FindUserById(ctx context.Context, id int) (*usermodel.User, error) {
	user, err := b.repo.FindUserById(ctx, id)
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			return nil, core.ErrNotFound
		} else {
			return nil, core.ErrInternalServerError.WithWrap(err)
		}
	}

	if user == nil || user.Status != 1 {
		return nil, core.ErrNotFound
	}

	return user, nil
}

func (b *userBiz) FindUserByIds(ctx context.Context, ids []int) ([]*usermodel.User, error) {
	//TODO implement me
	panic("implement me")
}

func (b *userBiz) UpdateUser(ctx context.Context, requester core.Requester, id int, data *usermodel.UserUpdate) error {
	//TODO implement me
	panic("implement me")
}

func NewBiz(repository UserRepo) *userBiz {
	return &userBiz{repo: repository}
}
