package mysql

import (
	"errors"
	"github.com/caovanhoang63/hiholive/services/user/module/user/usermodel"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

func (repo *mysqlRepo) FindUserById(ctx context.Context, id int) (*usermodel.User, error) {
	var user usermodel.User

	if err := repo.db.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrRecordNotFound
		}
		return nil, err
	}
	return &user, nil
}
