package streammysql

import (
	"fmt"
	"github.com/caovanhoang63/hiholive/services/video/module/stream/streammodel"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"time"
)

type streamRepo struct {
	db       *gorm.DB
	rdClient *redis.Client
}

func NewStreamMysqlRepo(db *gorm.DB, rdClient *redis.Client) *streamRepo {
	return &streamRepo{
		db:       db,
		rdClient: rdClient,
	}
}

func (s *streamRepo) Create(ctx context.Context, create *streammodel.StreamCreate) error {
	tx := s.db.Begin()

	if err := tx.Table(streammodel.Stream{}.TableName()).Create(&create).Error; err != nil {
		tx.Rollback()
		return err
	}

	create.Mask(core.DbTypeStream)

	if err := s.rdClient.SetEx(ctx, fmt.Sprintf("stream:%s", create.StreamKey), create.Uid, 3600*time.Second).Err(); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
