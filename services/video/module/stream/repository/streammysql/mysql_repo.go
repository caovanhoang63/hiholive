package streammysql

import (
	"encoding/json"
	"errors"
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

func (s *streamRepo) UpdateStream(ctx context.Context, id int, update *streammodel.StreamUpdate) error {
	if err := s.db.Model(&streammodel.Stream{}).Where("id = ?", id).Updates(update).Error; err != nil {
		return err
	}
	return nil
}

func NewStreamMysqlRepo(db *gorm.DB, rdClient *redis.Client) *streamRepo {
	return &streamRepo{
		db:       db,
		rdClient: rdClient,
	}
}

func (s *streamRepo) Create(ctx context.Context, create *streammodel.StreamCreate) error {
	tx := s.db.Begin()

	create.UnMask()
	if err := tx.Table(streammodel.Stream{}.TableName()).Create(&create).Error; err != nil {
		tx.Rollback()
		return err
	}

	create.Mask(core.DbTypeStream)

	streamStateData, _ := json.Marshal(core.StreamState{
		Uid:   create.Uid.String(),
		State: "pending",
	})

	createData, _ := json.Marshal(create)

	pipeline := s.rdClient.Pipeline()
	pipeline.SetEx(ctx, fmt.Sprintf("streamKey:%s", create.StreamKey), streamStateData, 4*time.Hour)
	pipeline.SetNX(ctx, fmt.Sprintf("stream:%d", create.Id), createData, 4*time.Hour)

	if _, err := pipeline.Exec(ctx); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (s *streamRepo) FindStreamByID(ctx context.Context, id int) (*streammodel.Stream, error) {
	var stream streammodel.Stream
	r, err := s.rdClient.Get(ctx, fmt.Sprintf("streamKey:%d", id)).Result()
	if err == nil {
		err = json.Unmarshal([]byte(r), &stream)
		if err == nil {
			return &stream, nil
		}
	}
	if err = s.db.Preload("Category").Preload("Channel").Where("id = ?", id).First(&stream).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &stream, nil
}

func (s *streamRepo) FindStreams(ctx context.Context, filter *streammodel.StreamFilter, paging *core.Paging) ([]streammodel.StreamList, error) {
	var result []streammodel.StreamList

	db := s.db.Table(streammodel.Stream{}.TableName()).Where("status in (1)")

	if filter != nil {
		if filter.CategoryId != 0 {
			db = db.Where("category_id = ?", filter.CategoryId)
		}
		if filter.ChannelId != 0 {
			db = db.Where("channel_id = ?", filter.ChannelId)
		}
		if filter.State != "" {
			db = db.Where("state = ?", filter.State)
		}

		if filter.Title != "" {
			db = db.Where("title LIKE ?", "%"+filter.Title+"%")
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	db.Preload("Category").Preload("Channel").Find(&result)

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
