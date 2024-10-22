package mysql

import (
	"github.com/caovanhoang63/hiholive/services/auth/module/auth/entity"
	"golang.org/x/net/context"
)

func (r *mysqlRepo) Create(ctx context.Context, register *entity.Auth) error {
	if err := r.db.WithContext(ctx).Create(&register).Error; err != nil {
		return err
	}
	return nil
}
