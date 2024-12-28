package ctgsub

import (
	"errors"
	"github.com/caovanhoang63/hiholive/services/video/module/category/ctgbiz"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx/components/pubsub"
	"github.com/caovanhoang63/hiholive/shared/go/subengine"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type CategorySub struct {
	biz        ctgbiz.CategoryBiz
	serviceCtx srvctx.ServiceContext
}

func NewCategorySub(biz ctgbiz.CategoryBiz, serviceCtx srvctx.ServiceContext) *CategorySub {
	return &CategorySub{
		biz:        biz,
		serviceCtx: serviceCtx,
	}
}

func (s *CategorySub) IncreaseTotalContent() subengine.ConsumerJob {
	return subengine.ConsumerJob{
		Title: "Increase Category Total Content When a stream create ",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			data, ok := message.Data.(map[string]any)
			if !ok {
				return errors.New("invalid data format")
			}

			uid, ok := data["categoryId"]
			if !ok {
				return errors.New("invalid data")
			}

			if id, err := core.FromBase58(uid.(string)); err == nil {
				err = s.biz.IncreaseTotalContent(ctx, int(id.GetLocalID()))
				if err != nil {
					log.Println(err)
					return err
				}
				return nil
			}
			return nil
		},
	}
}
