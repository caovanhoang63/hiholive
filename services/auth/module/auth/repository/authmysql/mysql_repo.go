package authmysql

import (
	"context"
	"errors"
	"fmt"
	"github.com/caovanhoang63/hiholive/services/auth/module/auth/authmodel"
	"github.com/caovanhoang63/hiholive/shared/golang/core"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

type authRepo struct {
	db *gorm.DB
	rd *redis.Client
}

func NewMySQLRepository(db *gorm.DB, rd *redis.Client) *authRepo {
	return &authRepo{db: db, rd: rd}
}

func (r *authRepo) CheckForgotPasswordPin(c context.Context, email, pin string) error {
	key := "forgot-password:" + email
	storedPin, err := r.rd.Get(c, key).Result()
	if errors.Is(err, redis.Nil) {
		return core.ErrRecordNotFound
	}
	if err != nil {
		return err
	}
	if storedPin != pin {
		return authmodel.ErrInvalidPin
	}
	return nil
}

func (r *authRepo) UpdatePassword(c context.Context, email, password string) error {
	key := "forgot-password:" + email

	err := r.db.Table(authmodel.Auth{}.TableName()).Where("email = ?", email).Update("password", password).Error
	if err != nil {
		return err
	}

	err = r.rd.Del(c, key).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *authRepo) ForgotPassword(c context.Context, email, pin string) error {
	key := "forgot-password:" + email

	ttl, err := r.rd.TTL(c, key).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return err
	}
	fmt.Println(ttl, err)

	if ttl > 0 {
		if ttl > time.Minute*2 {
			return authmodel.ErrPinAlreadyExists
		}
	}

	err = r.rd.Set(c, key, pin, 3*time.Minute).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *authRepo) FindByUserId(ctx context.Context, id int) (*authmodel.Auth, error) {
	data := &authmodel.Auth{}
	if err := r.db.WithContext(ctx).Table(data.TableName()).Where("user_id = ?", id).First(data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrRecordNotFound
		}
		return nil, err
	}
	return data, nil
}
