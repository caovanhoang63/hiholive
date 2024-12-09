package mysql

import (
	"context"
	"github.com/caovanhoang63/hiholive/services/user/module/user/usermodel"
	"gorm.io/gorm"
)

type mysqlRepo struct {
	db *gorm.DB
}

func (repo *mysqlRepo) DeleteUser(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (repo *mysqlRepo) FindUserByIds(ctx context.Context, ids []int) ([]*usermodel.User, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *mysqlRepo) UpdateUser(ctx context.Context, id int, data *usermodel.UserUpdate) error {
	//TODO implement me
	panic("implement me")
}

func NewMySQLRepository(db *gorm.DB) *mysqlRepo {
	return &mysqlRepo{db: db}
}
