package videocomposer

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
	"github.com/caovanhoang63/hiholive/services/video/module/stream/transport/strmgin"
	"github.com/caovanhoang63/hiholive/services/video/module/stream/transport/strmgrpc"
	"github.com/caovanhoang63/hiholive/services/video/module/upload/transport/uploadgin"
	"github.com/caovanhoang63/hiholive/services/video/module/upload/uploadbiz"
	"github.com/caovanhoang63/hiholive/shared/golang/core"
	"github.com/caovanhoang63/hiholive/shared/golang/proto/pb"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/pubsub"
	"github.com/caovanhoang63/hiholive/shared/golang/uploadprovider"

	"github.com/gin-gonic/gin"
)

type ChannelService interface {
	CreateChannel() func(c *gin.Context)
	FindUserChannel() func(c *gin.Context)
	FindChannelById() func(c *gin.Context)
	FindChannels() func(c *gin.Context)
}

func ComposeChannelAPIService(serviceCtx srvctx.ServiceContext, repo channelbiz.UserRepo) ChannelService {
	db := serviceCtx.MustGet(core.KeyCompMySQL).(core.GormComponent)
	ps := serviceCtx.MustGet(core.KeyCompRabbitMQ).(pubsub.Pubsub)
	channelRepo := channelmysql.NewChannelMysqlRepo(db.GetDB())
	biz := channelbiz.NewChannelBiz(channelRepo, repo, ps)
	userService := channelgin.NewChannelGinApi(biz, serviceCtx)
	return userService
}

type StreamService interface {
	CreateStream() gin.HandlerFunc
	GetStreamById() gin.HandlerFunc
	FindStreams() gin.HandlerFunc
}

func ComposeStreamAPIService(serviceCtx srvctx.ServiceContext) StreamService {
	db := serviceCtx.MustGet(core.KeyCompMySQL).(core.GormComponent)
	rd := serviceCtx.MustGet(core.KeyRedis).(core.RedisComponent)
	ps := serviceCtx.MustGet(core.KeyCompRabbitMQ).(pubsub.Pubsub)
	channelRepo := channelmysql.NewChannelMysqlRepo(db.GetDB())
	streamRepo := streammysql.NewStreamMysqlRepo(db.GetDB(), rd.GetClient())
	biz := streambiz.NewStreamBiz(streamRepo, channelRepo, ps)
	streamService := strmgin.NewStreamApi(biz, serviceCtx)
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

func ComposeStreamGRPCService(serviceCtx srvctx.ServiceContext) pb.StreamServiceServer {
	db := serviceCtx.MustGet(core.KeyCompMySQL).(core.GormComponent)
	rd := serviceCtx.MustGet(core.KeyRedis).(core.RedisComponent)
	ps := serviceCtx.MustGet(core.KeyCompRabbitMQ).(pubsub.Pubsub)
	channelRepo := channelmysql.NewChannelMysqlRepo(db.GetDB())
	streamRepo := streammysql.NewStreamMysqlRepo(db.GetDB(), rd.GetClient())
	biz := streambiz.NewStreamBiz(streamRepo, channelRepo, ps)
	service := strmgrpc.NewStreamGRPC(biz, serviceCtx)
	return service
}

type UploadService interface {
	UploadImage() gin.HandlerFunc
}

func ComposeUploadAPIService(serviceCtx srvctx.ServiceContext) UploadService {
	provider := serviceCtx.MustGet(core.KeyS3).(uploadprovider.UploadProvider)
	biz := uploadbiz.NewUploadBiz(provider)
	service := uploadgin.NewUploadGin(biz, serviceCtx)
	return service
}
