package authmysql

import (
	"github.com/caovanhoang63/hiholive/services/auth/module/auth/authmodel"
	"golang.org/x/net/context"
)

func (r *authRepo) Create(ctx context.Context, register *authmodel.Auth) error {
	if err := r.db.WithContext(ctx).Create(&register).Error; err != nil {
		return err
	}
	return nil
}
