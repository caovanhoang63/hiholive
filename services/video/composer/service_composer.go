package composer

import (
	"github.com/caovanhoang63/hiholive/services/video/module/channel/channelbiz"
	"github.com/caovanhoang63/hiholive/services/video/module/channel/repository/channelmysql"
	"github.com/caovanhoang63/hiholive/services/video/module/channel/transport/channelgin"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
	"github.com/gin-gonic/gin"
)

type ChannelService interface {
	CreateChannel() func(c *gin.Context)
}

func ComposeChannelAPIService(serviceCtx srvctx.ServiceContext) ChannelService {
	db := serviceCtx.MustGet(core.KeyCompMySQL).(core.GormComponent)

	channelRepo := channelmysql.NewChannelMysqlRepo(db.GetDB())
	biz := channelbiz.NewChannelBiz(channelRepo)
	userService := channelgin.NewChannelGinApi(biz, serviceCtx)
	return userService
}
