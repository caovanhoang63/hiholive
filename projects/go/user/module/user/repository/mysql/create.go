package mysql

import (
	"context"
	"github.com/caovanhoang63/hiholive/user/module/user/entity"
	"github.com/pkg/errors"
)

func (repo *mysqlRepo) CreateNewUser(ctx context.Context, data *entity.UserCreate) error {
	if err := repo.db.Create(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
