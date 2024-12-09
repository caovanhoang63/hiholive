package channelmysql

import (
	"errors"
	"github.com/caovanhoang63/hiholive/services/video/module/channel/channelmodel"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type channelMysqlRepo struct {
	db *gorm.DB
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
