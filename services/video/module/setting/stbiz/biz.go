package stbiz

import (
	"errors"
	"github.com/caovanhoang63/hiholive/services/video/module/setting/stmodel"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"golang.org/x/net/context"
)

type SystemSettingBiz interface {
	Create(ctx context.Context, requester core.Requester, create *stmodel.SettingCreate) error
	Update(ctx context.Context, requester core.Requester, name string, update *stmodel.SettingUpdate) error
	FindByName(ctx context.Context, name string) (*stmodel.Setting, error)
	FindByCondition(ctx context.Context, filter *stmodel.Filter, paging *core.Paging) ([]stmodel.Setting, error)
}

type SystemSettingRepo interface {
	Create(ctx context.Context, create *stmodel.SettingCreate) error
	Update(ctx context.Context, name string, update *stmodel.SettingUpdate) error
	FindByName(ctx context.Context, name string) (*stmodel.Setting, error)
	FindById(ctx context.Context, id int) (*stmodel.Setting, error)
	FindByCondition(ctx context.Context, filter *stmodel.Filter, paging *core.Paging) ([]stmodel.Setting, error)
}

type systemSettingBizImpl struct {
	repo SystemSettingRepo
}

func NewSystemSettingBiz(repo SystemSettingRepo) SystemSettingBiz {
	return &systemSettingBizImpl{repo: repo}
}
func (s *systemSettingBizImpl) Create(ctx context.Context, requester core.Requester, create *stmodel.SettingCreate) error {
	if requester.GetRole() != "admin" {
		return core.ErrForbidden
	}

	old, err := s.repo.FindByName(ctx, create.Name)
	if err != nil && !errors.Is(err, core.ErrRecordNotFound) {
		return core.ErrInternalServerError.WithWrap(err)
	}
	if old != nil {
		return core.ErrConflict.WithError("Setting already exists")
	}

	if err = s.repo.Create(ctx, create); err != nil {
		return core.ErrInternalServerError.WithWrap(err)
	}

	return nil

}

func (s *systemSettingBizImpl) Update(ctx context.Context, requester core.Requester, name string, update *stmodel.SettingUpdate) error {
	if requester.GetRole() != "admin" {
		return core.ErrForbidden
	}

	_, err := s.repo.FindByName(ctx, name)
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			return core.ErrNotFound
		}
		return core.ErrInternalServerError.WithTrace(err)
	}

	if err = s.repo.Update(ctx, name, update); err != nil {
		return core.ErrInternalServerError.WithTrace(err)
	}

	return nil
}

func (s *systemSettingBizImpl) FindByName(ctx context.Context, name string) (*stmodel.Setting, error) {
	data, err := s.repo.FindByName(ctx, name)
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			return nil, core.ErrNotFound
		}
		return nil, core.ErrInternalServerError.WithTrace(err)
	}
	return data, nil
}

func (s *systemSettingBizImpl) FindByCondition(ctx context.Context, filter *stmodel.Filter, paging *core.Paging) ([]stmodel.Setting, error) {
	data, err := s.repo.FindByCondition(ctx, filter, paging)
	if err != nil {
		return nil, core.ErrInternalServerError.WithTrace(err)
	}

	return data, nil
}
