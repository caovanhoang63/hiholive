package streammysql

import (
	"github.com/caovanhoang63/hiholive/services/video/module/stream/streammodel"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type streamMysqlRepo struct {
	db *gorm.DB
}

func NewStreamMysqlRepo(db *gorm.DB) *streamMysqlRepo {
	return &streamMysqlRepo{
		db: db,
	}
}

func (s *streamMysqlRepo) Create(ctx context.Context, create *streammodel.StreamCreate) error {
	if err := s.db.Table(streammodel.Stream{}.TableName()).Create(&create).Error; err != nil {
		return err
	}
	return nil
}
