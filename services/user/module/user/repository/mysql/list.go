package mysql

import (
	"github.com/caovanhoang63/hiholive/services/user/module/user/usermodel"
	"github.com/caovanhoang63/hiholive/shared/golang/core"
	"golang.org/x/net/context"
)

func (repo *mysqlRepo) FindUsersWithCondition(ctx context.Context, filter *usermodel.UserFilter, paging *core.Paging) ([]usermodel.User, error) {
	var result []usermodel.User

	db := repo.db.Table(usermodel.User{}.TableName()).Where("status in (1)")

	if filter != nil {
		if filter.UserName != "" {
			db = db.Where("user_name like ?", "%"+filter.UserName+"%")
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
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
