package authmysql

import (
	"github.com/caovanhoang63/hiholive/services/auth/module/auth/authmodel"
	"golang.org/x/net/context"
)

func (r *mysqlRepo) FindByEmail(ctx context.Context, email string) (*authmodel.Auth, error) {
	data := &authmodel.Auth{}
	if err := r.db.WithContext(ctx).Table(data.TableName()).Where("email = ?", email).First(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
