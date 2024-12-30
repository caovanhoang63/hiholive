package channelmysql

import (
	"errors"
	"github.com/caovanhoang63/hiholive/services/video/module/channel/channelmodel"
	"github.com/caovanhoang63/hiholive/shared/golang/core"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type channelMysqlRepo struct {
	db *gorm.DB
}

func (c *channelMysqlRepo) UpdateChannelName(ctx context.Context, userId int, userName, displayName string) error {
	//TODO implement me
	panic("implement me")
}

func (c *channelMysqlRepo) FindChannelByUserName(ctx context.Context, userName string) (*channelmodel.Channel, error) {
	var data channelmodel.Channel
	if err := c.db.Where("user_name = ?", userName).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrRecordNotFound
		}
		return nil, err
	}
	return &data, nil
}

func NewChannelMysqlRepo(db *gorm.DB) *channelMysqlRepo {
	return &channelMysqlRepo{db: db}
}

func (c *channelMysqlRepo) Create(ctx context.Context, create *channelmodel.ChannelCreate) error {
	if err := c.db.Table(channelmodel.Channel{}.TableName()).Create(create).Error; err != nil {
		return err
	}
	return nil
}

func (c *channelMysqlRepo) FindUserChannel(ctx context.Context, userId int) (*channelmodel.Channel, error) {
	var data channelmodel.Channel
	if err := c.db.Where("user_id = ?", userId).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrRecordNotFound
		}
		return nil, err
	}
	return &data, nil
}
func (c *channelMysqlRepo) FindChannelById(ctx context.Context, channelId int) (*channelmodel.Channel, error) {
	var data channelmodel.Channel
	if err := c.db.Where("id = ?", channelId).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrRecordNotFound
		}
		return nil, err
	}
	return &data, nil
}

func (c *channelMysqlRepo) FindChannels(ctx context.Context, filter *channelmodel.ChannelFilter, paging *core.Paging) ([]channelmodel.Channel, error) {
	var result []channelmodel.Channel

	db := c.db.Table(channelmodel.Channel{}.TableName()).Where("status in (1)")

	if filter != nil {
		if filter.IsLive != nil {
			// TODO: Implement this
		}
		if filter.UserName != "" {
			db = db.Where("user_name LIKE ?", "%"+filter.UserName+"%")
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
