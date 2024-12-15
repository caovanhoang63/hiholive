package composer

import (
	"github.com/caovanhoang63/hiholive/services/video/module/category/ctgbiz"
	"github.com/caovanhoang63/hiholive/services/video/module/category/repository/ctgmysqlrepo"
	"github.com/caovanhoang63/hiholive/services/video/module/category/transport/ctggin"
	"github.com/caovanhoang63/hiholive/services/video/module/channel/channelbiz"
	"github.com/caovanhoang63/hiholive/services/video/module/channel/repository/channelmysql"
	"github.com/caovanhoang63/hiholive/services/video/module/channel/transport/channelgin"
	"github.com/caovanhoang63/hiholive/services/video/module/setting/repo/stmysqlrepo"
	"github.com/caovanhoang63/hiholive/services/video/module/setting/stbiz"
	"github.com/caovanhoang63/hiholive/services/video/module/setting/transport/stgin"
	"github.com/caovanhoang63/hiholive/services/video/module/stream/repository/streammysql"
	"github.com/caovanhoang63/hiholive/services/video/module/stream/streambiz"
	"github.com/caovanhoang63/hiholive/services/video/module/stream/transport/streamgin"
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

type StreamService interface {
	CreateStream() gin.HandlerFunc
}

func ComposeStreamAPIService(serviceCtx srvctx.ServiceContext) StreamService {
	db := serviceCtx.MustGet(core.KeyCompMySQL).(core.GormComponent)
	rd := serviceCtx.MustGet(core.KeyRedis).(core.RedisComponent)
	channelRepo := channelmysql.NewChannelMysqlRepo(db.GetDB())
	streamRepo := streammysql.NewStreamMysqlRepo(db.GetDB(), rd.GetClient())
	biz := streambiz.NewStreamBiz(streamRepo, channelRepo)
	streamService := streamgin.NewStreamApi(biz, serviceCtx)
	return streamService
}

type SystemSettingService interface {
	CreateSystemSetting() gin.HandlerFunc
	UpdateSystemSetting() gin.HandlerFunc
	FindSystemSetting() gin.HandlerFunc
	FindSystemSettingByName() gin.HandlerFunc
}

func ComposeSystemSettingApiService(serviceCtx srvctx.ServiceContext) SystemSettingService {
	db := serviceCtx.MustGet(core.KeyCompMySQL).(core.GormComponent)
	rd := serviceCtx.MustGet(core.KeyRedis).(core.RedisComponent)
	repo := stmysqlrepo.New(db.GetDB(), rd.GetClient())
	biz := stbiz.NewSystemSettingBiz(repo)
	service := stgin.NewGinApi(biz)
	return service
}

type CategoryService interface {
	CreateCategory() gin.HandlerFunc
	UpdateCategory() gin.HandlerFunc
	FindCategories() gin.HandlerFunc
	FindCategoryById() gin.HandlerFunc
	DeleteCategory() gin.HandlerFunc
}

func ComposeCategoryApiService(serviceCtx srvctx.ServiceContext) CategoryService {
	db := serviceCtx.MustGet(core.KeyCompMySQL).(core.GormComponent)
	ctgRepo := ctgmysqlrepo.NewMysqlRepo(db.GetDB())
	ctgBiz := ctgbiz.NewCategoryBiz(ctgRepo)
	service := ctggin.NewCategoryApi(ctgBiz)
	return service
}
