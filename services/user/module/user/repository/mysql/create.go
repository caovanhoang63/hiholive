package mysql

import (
	"context"
	"github.com/caovanhoang63/hiholive/services/user/module/user/usermodel"
	"github.com/pkg/errors"
)

func (repo *mysqlRepo) CreateNewUser(ctx context.Context, data *usermodel.UserCreate) error {
	if err := repo.db.Create(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
