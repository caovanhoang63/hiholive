package mysql

import (
	"context"
	"github.com/pkg/errors"
	"hiholive/projects/go/user/module/user/entity"
)

func (repo *mysqlRepo) CreateNewUser(ctx context.Context, data *entity.UserCreate) error {
	if err := repo.db.Create(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
