package biz

import (
	"context"
	"github.com/caovanhoang63/hiholive/services/user/module/user/entity"
	"github.com/caovanhoang63/hiholive/shared/go/core"
)

type UserBiz interface {
	CreateNewUser(ctx context.Context, data *entity.UserCreate) error
	DeleteUser(ctx context.Context, id int) error
	FindUserById(ctx context.Context, id int) (*entity.User, error)
	FindUserByIds(ctx context.Context, ids []int) ([]*entity.User, error)
	UpdateUser(ctx context.Context, id int, data *entity.UserUpdate) error
}

type UserRepository interface {
	CreateNewUser(ctx context.Context, data *entity.UserCreate) error
	DeleteUser(ctx context.Context, id int) error
	FindUserById(ctx context.Context, id int) (*entity.User, error)
	FindUserByIds(ctx context.Context, ids []int) ([]*entity.User, error)
	UpdateUser(ctx context.Context, id int, data *entity.UserUpdate) error
}

type biz struct {
	repository UserRepository
}

func (b *biz) CreateNewUser(ctx context.Context, data *entity.UserCreate) error {

	if field, err := core.Validator.ValidateField(data); err != nil {
		return core.ErrInvalidInput(field)
	}
	data.Gender = entity.Other

	if err := b.repository.CreateNewUser(ctx, data); err != nil {
		return core.ErrInternalServerError.WithError(err.Error())
	}

	return nil
}

func (b *biz) DeleteUser(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (b *biz) FindUserById(ctx context.Context, id int) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (b *biz) FindUserByIds(ctx context.Context, ids []int) ([]*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (b *biz) UpdateUser(ctx context.Context, id int, data *entity.UserUpdate) error {
	//TODO implement me
	panic("implement me")
}

func NewBiz(repository UserRepository) *biz {
	return &biz{repository: repository}
}
