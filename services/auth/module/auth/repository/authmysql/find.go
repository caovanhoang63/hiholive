package authmysql

import (
	"errors"
	"github.com/caovanhoang63/hiholive/services/auth/module/auth/authmodel"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

func (r *mysqlRepo) FindByEmail(ctx context.Context, email string) (*authmodel.Auth, error) {
	data := &authmodel.Auth{}
	if err := r.db.WithContext(ctx).Table(data.TableName()).Where("email = ?", email).First(data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrRecordNotFound
		}
		return nil, err
	}
	return data, nil
}
