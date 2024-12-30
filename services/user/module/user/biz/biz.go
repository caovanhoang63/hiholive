package biz

import (
	"context"
	"errors"
	"fmt"
	"github.com/caovanhoang63/hiholive/services/user/module/user/usermodel"
	"github.com/caovanhoang63/hiholive/shared/golang/core"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/pubsub"
	"strings"
)

type UserBiz interface {
	CreateNewUser(ctx context.Context, data *usermodel.UserCreate) error
	DeleteUser(ctx context.Context, requester core.Requester, id int) error
	FindUserById(ctx context.Context, id int) (*usermodel.User, error)
	FindUserByIds(ctx context.Context, ids []int) ([]usermodel.User, error)
	FindUsersWithCondition(ctx context.Context, filter *usermodel.UserFilter, paging *core.Paging) ([]usermodel.User, error)
	UpdateUser(ctx context.Context, requester core.Requester, id int, data *usermodel.UserUpdate) error
	UpdateToRoleStreamer(ctx context.Context, id int) error
	UpdateUserName(ctx context.Context, requester core.Requester, id int, name *usermodel.UserNameAndDisplayName) error
}

type UserRepo interface {
	CreateNewUser(ctx context.Context, data *usermodel.UserCreate) error
	FindUserByUserName(ctx context.Context, userName string) (*usermodel.User, error)
	DeleteUser(ctx context.Context, id int) error
	FindUserById(ctx context.Context, id int) (*usermodel.User, error)
	FindUserByIds(ctx context.Context, ids []int) ([]usermodel.User, error)
	UpdateUser(ctx context.Context, id int, data *usermodel.UserUpdate) error
	FindUsersWithCondition(ctx context.Context, filter *usermodel.UserFilter, paging *core.Paging) ([]usermodel.User, error)
	UpdateUserRole(ctx context.Context, id int, role string) error
	UpdateUserName(ctx context.Context, id int, name *usermodel.UserNameAndDisplayName) error
}

type userBiz struct {
	repo UserRepo
	ps   pubsub.Pubsub
}

func NewBiz(repository UserRepo, ps pubsub.Pubsub) *userBiz {
	return &userBiz{
		repo: repository,
		ps:   ps,
	}
}

func (b *userBiz) UpdateToRoleStreamer(ctx context.Context, id int) error {
	user, err := b.repo.FindUserById(ctx, id)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, core.ErrRecordNotFound) {
			return core.ErrNotFound
		} else {
			return core.ErrInternalServerError.WithWrap(err)
		}
	}

	if user.SystemRole == usermodel.RoleStreamer {
		return nil
	}

	if user.SystemRole != usermodel.RoleUser {
		return core.ErrForbidden.WithError("Admin & Moderator cannot become a streamer")
	}

	if err = b.repo.UpdateUserRole(ctx, id, usermodel.RoleStreamer); err != nil {
		fmt.Println(err)
		return core.ErrInternalServerError.WithWrap(err)
	}
	return nil
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
	return user.SystemRole, nil
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
	data.SystemRole = usermodel.RoleUser

	// Default userName and DisplayName
	displayName := strings.Split(data.Email, "@")[0]
	displayName = displayName + core.GenSalt(5)
	data.DisplayName = displayName
	data.UserName = displayName

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

func (b *userBiz) FindUserByIds(ctx context.Context, ids []int) ([]usermodel.User, error) {
	user, err := b.repo.FindUserByIds(ctx, ids)
	if err != nil {
		return nil, core.ErrInternalServerError.WithWrap(err)
	}
	return user, nil
}

func (b *userBiz) UpdateUser(ctx context.Context, requester core.Requester, id int, data *usermodel.UserUpdate) error {
	if requester != nil {
		if id != requester.GetUserId() {
			return core.ErrForbidden
		}
	}
	if field, err := core.Validator.ValidateField(data); err != nil {
		return core.ErrInvalidInput(field)
	}
	oldUser, err := b.repo.FindUserById(ctx, id)
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			return core.ErrNotFound
		}
		return core.ErrInternalServerError.WithWrap(err)
	}
	if oldUser.Status != 1 {
		return core.ErrNotFound
	}

	if err = b.repo.UpdateUser(ctx, id, data); err != nil {
		return core.ErrInternalServerError.WithWrap(err)
	}
	return nil
}

func (b *userBiz) UpdateUserName(ctx context.Context, requester core.Requester, id int, name *usermodel.UserNameAndDisplayName) error {
	if requester != nil {
		if id != requester.GetUserId() {
			return core.ErrForbidden
		}
	}
	if field, err := core.Validator.ValidateField(name); err != nil {
		return core.ErrInvalidInput(field)
	}

	if strings.ToLower(name.UserName) != name.UserName {
		return core.ErrInvalidInput("userName")
	}

	if strings.ToLower(name.UserName) != strings.ToLower(name.DisplayName) {
		return core.ErrInvalidInput("displayName")
	}

	oldUser, err := b.repo.FindUserById(ctx, id)
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			return core.ErrNotFound
		}
		return core.ErrInternalServerError.WithWrap(err)
	}
	if oldUser.Status != 1 {
		return core.ErrNotFound
	}

	oldName, err := b.repo.FindUserByUserName(ctx, name.UserName)
	if err != nil && !errors.Is(err, core.ErrRecordNotFound) {
		return core.ErrInternalServerError.WithWrap(err)
	}
	if oldName != nil && err == nil && oldName.Id != oldUser.Id {
		return core.ErrConflict
	}

	if err = b.repo.UpdateUserName(ctx, id, name); err != nil {
		return core.ErrInternalServerError.WithWrap(err)
	}

	_ = b.ps.Publish(ctx, core.TopicUpdateUserName, pubsub.NewMessage(map[string]interface{}{
		"id":          id,
		"userName":    name.UserName,
		"displayName": name.DisplayName,
	}))
	return nil
}
