package mysql

import (
	"github.com/caovanhoang63/hiholive/services/auth/module/auth/entity"
	"golang.org/x/net/context"
)

func (r *mysqlRepo) FindByEmail(ctx context.Context, email string) (*entity.Auth, error) {
	data := &entity.Auth{}
	if err := r.db.WithContext(ctx).Table(data.TableName()).Where("email = ?", email).First(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
