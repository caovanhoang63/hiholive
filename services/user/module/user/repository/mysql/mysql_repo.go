package mysql

import (
	"context"
	"errors"
	"github.com/caovanhoang63/hiholive/services/user/module/user/usermodel"
	"github.com/caovanhoang63/hiholive/shared/golang/core"
	"gorm.io/gorm"
)

type mysqlRepo struct {
	db *gorm.DB
}

func (repo *mysqlRepo) FindUserByUserName(ctx context.Context, userName string) (*usermodel.User, error) {
	var user usermodel.User
	if err := repo.db.WithContext(ctx).Where("user_name = ?", userName).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrRecordNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (repo *mysqlRepo) UpdateUserRole(ctx context.Context, id int, role string) error {
	if err := repo.db.Table(usermodel.User{}.TableName()).Where("id = ?", id).Update("system_role", role).Error; err != nil {
		return err
	}
	return nil
}

func (repo *mysqlRepo) DeleteUser(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (repo *mysqlRepo) FindUserByIds(ctx context.Context, ids []int) ([]usermodel.User, error) {
	var user []usermodel.User
	if err := repo.db.Where("id IN (?)", ids).Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *mysqlRepo) UpdateUser(ctx context.Context, id int, data *usermodel.UserUpdate) error {
	if err := repo.db.WithContext(ctx).Table(usermodel.User{}.TableName()).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func (repo *mysqlRepo) UpdateUserName(ctx context.Context, id int, name *usermodel.UserNameAndDisplayName) error {
	if err := repo.db.WithContext(ctx).Table(usermodel.User{}.TableName()).Where("id = ?", id).Updates(name).Error; err != nil {
		return err
	}
	return nil
}

func NewMySQLRepository(db *gorm.DB) *mysqlRepo {
	return &mysqlRepo{db: db}
}
