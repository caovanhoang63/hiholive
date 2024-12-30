package videocomposer

import (
	"github.com/caovanhoang63/hiholive/services/video/module/category/ctgbiz"
	"github.com/caovanhoang63/hiholive/services/video/module/category/repository/ctgmysqlrepo"
	"github.com/caovanhoang63/hiholive/services/video/module/category/transport/ctgsub"
	"github.com/caovanhoang63/hiholive/services/video/module/channel/channelbiz"
	"github.com/caovanhoang63/hiholive/services/video/module/channel/repository/channelmysql"
	"github.com/caovanhoang63/hiholive/services/video/module/channel/transport/channelsub"
	"github.com/caovanhoang63/hiholive/services/video/module/stream/repository/streammysql"
	"github.com/caovanhoang63/hiholive/services/video/module/stream/streambiz"
	"github.com/caovanhoang63/hiholive/services/video/module/stream/transport/strmsub"
	"github.com/caovanhoang63/hiholive/shared/golang/core"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/pubsub"
)

func ComposeStreamSubscriber(serviceCtx srvctx.ServiceContext) *strmsub.StreamSub {
	db := serviceCtx.MustGet(core.KeyCompMySQL).(core.GormComponent)
	rd := serviceCtx.MustGet(core.KeyRedis).(core.RedisComponent)
	ps := serviceCtx.MustGet(core.KeyCompRabbitMQ).(pubsub.Pubsub)
	channelRepo := channelmysql.NewChannelMysqlRepo(db.GetDB())
	streamRepo := streammysql.NewStreamMysqlRepo(db.GetDB(), rd.GetClient())
	biz := streambiz.NewStreamBiz(streamRepo, channelRepo, ps)
	streamSub := strmsub.NewStreamSub(biz, serviceCtx)
	return streamSub
}

func ComposeChannelSubcriber(serviceCtx srvctx.ServiceContext) *channelsub.ChannelSub {
	db := serviceCtx.MustGet(core.KeyCompMySQL).(core.GormComponent)
	ps := serviceCtx.MustGet(core.KeyCompRabbitMQ).(pubsub.Pubsub)
	channelRepo := channelmysql.NewChannelMysqlRepo(db.GetDB())
	biz := channelbiz.NewChannelBiz(channelRepo, nil, ps)
	channelService := channelsub.NewChannelSub(biz, serviceCtx)
	return channelService
}

func ComposeCategorySubscriber(serviceCtx srvctx.ServiceContext) *ctgsub.CategorySub {
	db := serviceCtx.MustGet(core.KeyCompMySQL).(core.GormComponent)
	repo := ctgmysqlrepo.NewMysqlRepo(db.GetDB())
	biz := ctgbiz.NewCategoryBiz(repo)
	service := ctgsub.NewCategorySub(biz, serviceCtx)
	return service
}
