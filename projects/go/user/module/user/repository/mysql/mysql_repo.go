package mysql

import (
	"context"
	"gorm.io/gorm"
	"hiholive/projects/go/user/module/user/entity"
)

type mysqlRepo struct {
	db *gorm.DB
}

func (repo *mysqlRepo) DeleteUser(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (repo *mysqlRepo) FindUserById(ctx context.Context, id int) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *mysqlRepo) FindUserByIds(ctx context.Context, ids []int) ([]*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *mysqlRepo) UpdateUser(ctx context.Context, id int, data *entity.UserUpdate) error {
	//TODO implement me
	panic("implement me")
}

func NewMySQLRepository(db *gorm.DB) *mysqlRepo {
	return &mysqlRepo{db: db}
}
