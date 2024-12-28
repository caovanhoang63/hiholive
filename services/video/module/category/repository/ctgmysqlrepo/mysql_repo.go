package ctgmysqlrepo

import (
	"errors"
	"github.com/caovanhoang63/hiholive/services/video/module/category/ctgmodel"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type mysqlRepo struct {
	db *gorm.DB
}

func NewMysqlRepo(db *gorm.DB) *mysqlRepo {
	return &mysqlRepo{db: db}
}

func (m *mysqlRepo) CreateCategory(ctx context.Context, create *ctgmodel.CategoryCreate) error {
	if err := m.db.Table(ctgmodel.Category{}.TableName()).Create(&create).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlRepo) UpdateCategory(ctx context.Context, id int, update *ctgmodel.CategoryUpdate) error {
	if err := m.db.Table(ctgmodel.Category{}.TableName()).Where("id = ?", id).Updates(&update).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlRepo) DeleteCategory(ctx context.Context, id int) error {
	if err := m.db.Table(ctgmodel.Category{}.TableName()).Where("id = ?", id).Update("status", 0).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlRepo) FindCategory(ctx context.Context, id int) (*ctgmodel.Category, error) {
	var data ctgmodel.Category
	if err := m.db.Where("id = ?", id).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrRecordNotFound
		}
		return nil, err
	}
	return &data, nil
}

func (m *mysqlRepo) IncreaseTotalContent(ctx context.Context, id int) error {
	if err := m.db.
		Table(ctgmodel.Category{}.TableName()).Where("id = ?", id).
		UpdateColumn("total_content", gorm.Expr("total_content + ? ", 1)).
		Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlRepo) FindCategories(ctx context.Context, filter *ctgmodel.CategoryFilter, paging *core.Paging) ([]ctgmodel.Category, error) {
	var result []ctgmodel.Category

	db := m.db.Table(ctgmodel.Category{}.TableName()).Where("status in (1)")
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
		last.Mask()
		paging.NextCursor = last.Uid.String()
	}

	return result, nil
}
