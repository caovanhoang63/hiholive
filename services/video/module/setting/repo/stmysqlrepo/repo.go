package stmysqlrepo

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/caovanhoang63/hiholive/services/video/module/setting/stmodel"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type repo struct {
	db       *gorm.DB
	rdClient *redis.Client
}

func New(db *gorm.DB, rdClient *redis.Client) *repo {
	return &repo{
		db:       db,
		rdClient: rdClient,
	}
}

func (r *repo) Create(ctx context.Context, create *stmodel.SettingCreate) error {
	tx := r.db.Begin()
	if err := tx.Table(stmodel.Setting{}.TableName()).Create(create).Error; err != nil {
		tx.Rollback()
		return err
	}
	b, _ := json.Marshal(create.Value)
	if err := r.rdClient.Set(ctx, fmt.Sprintf("system_setting:%s", create.Name), b, 0).Err(); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (r *repo) Update(ctx context.Context, name string, update *stmodel.SettingUpdate) error {
	tx := r.db.Begin()
	if err := tx.Table(stmodel.Setting{}.TableName()).Where("id = ?", name).Updates(&update).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := r.rdClient.SetNX(ctx, fmt.Sprintf("system_setting:%s", name), update.Value, 0).Err(); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r *repo) FindByName(ctx context.Context, name string) (*stmodel.Setting, error) {
	var data stmodel.Setting

	if err := r.db.Table(stmodel.Setting{}.TableName()).Where("name = ?", name).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrRecordNotFound
		}
		return nil, err
	}

	return &data, nil
}

func (r *repo) FindById(ctx context.Context, id int) (*stmodel.Setting, error) {
	var data stmodel.Setting

	if err := r.db.Table(stmodel.Setting{}.TableName()).Where("id = ?", id).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrRecordNotFound
		}
		return nil, err
	}

	return &data, nil
}

func (r *repo) FindByCondition(ctx context.Context, filter *stmodel.Filter, paging *core.Paging) ([]stmodel.Setting, error) {
	var result []stmodel.Setting

	db := r.db.Table(stmodel.Setting{}.TableName()).Where("status in (1)")
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if filter != nil {
		if filter.GtUpdatedAt != nil {
			db = db.Where("updated_at >?", filter.GtUpdatedAt)
		}
		if filter.GtCreatedAt != nil {
			db = db.Where("created_at >?", filter.GtCreatedAt)
		}
		if filter.LtUpdatedAt != nil {
			db = db.Where("updated_at <?", filter.LtUpdatedAt)
		}
		if filter.LtCreatedAt != nil {
			db = db.Where("created_at <?", filter.LtCreatedAt)
		}
		if filter.Name != "" {
			db = db.Where("name LIKE ?", filter.Name+"%")
		}
	}

	// paging
	if v := paging.FakeCursor; v != "" {
		uid, err := core.FromBase58(v)
		if err != nil {
			return nil, err
		}
		db = db.Where("id  < ? ", uid.GetLocalID())
	} else {
		db = db.Offset(paging.GetOffSet())
	}

	if err := db.Limit(paging.Limit).Order("id desc").Find(&result).Error; err != nil {
		return nil, err
	}

	if len(result) > 0 {
		last := result[len(result)-1]
		last.Mask(core.DbTypeSystemSetting)
		paging.NextCursor = last.Uid.String()
	}

	return result, nil
}
