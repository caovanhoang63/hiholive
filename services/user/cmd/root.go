package cmd

import (
	"fmt"
	"github.com/caovanhoang63/hiholive/services/user/common"
	"github.com/caovanhoang63/hiholive/services/user/composer"
	"github.com/caovanhoang63/hiholive/services/user/proto/pb"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx/components/ginc"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx/components/ginc/middlewares"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx/components/gormc"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx/components/jwtc"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"net"
	"time"

	"net/http"
	"os"
)

func newServiceCtx() srvctx.ServiceContext {
	return srvctx.NewServiceContext(
		srvctx.WithName("Demo Microservices"),
		srvctx.WithComponent(ginc.NewGin(core.KeyCompGIN)),
		srvctx.WithComponent(gormc.NewGormDB(core.KeyCompMySQL, "")),
		srvctx.WithComponent(jwtc.NewJWT(core.KeyCompJWT)),
		srvctx.WithComponent(NewConfig()),
	)
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start service",
	Run: func(cmd *cobra.Command, args []string) {
		serviceCtx := newServiceCtx()

		logger := srvctx.GlobalLogger().GetLogger("service")

		// Make some delay for DB ready (migration)
		// remove it if you already had your own DB
		time.Sleep(time.Second * 5)

		if err := serviceCtx.Load(); err != nil {
			logger.Fatal(err)
		}

		ginComp := serviceCtx.MustGet(core.KeyCompGIN).(core.GINComponent)

		router := ginComp.GetRouter()
		router.Use(gin.Recovery(), middlewares.Logger(serviceCtx), middlewares.Recovery(serviceCtx))

		router.Use(middlewares.Cors())
		router.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "pong"})
		})

		go StartGRPCServices(serviceCtx)

		v1 := router.Group("/v1")

		SetupRoutes(v1, serviceCtx)

		if err := router.Run(fmt.Sprintf(":%d", ginComp.GetPort())); err != nil {
			logger.Fatal(err)
		}
	},
}

func SetupRoutes(router *gin.RouterGroup, serviceCtx srvctx.ServiceContext) {
	userService := composer.ComposeUserAPIService(serviceCtx)

	tasks := router.Group("user")
	{
		tasks.GET(":id", userService.GetUserProfile())
		tasks.GET("", userService.ListUser())
	}
}
func StartGRPCServices(serviceCtx srvctx.ServiceContext) {
	configComp := serviceCtx.MustGet(core.KeyCompConf).(common.Config)
	logger := serviceCtx.Logger("grpc")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", configComp.GetGRPCPort()))

	if err != nil {
		log.Fatal(err)
	}

	logger.Infof("GRPC Server is listening on %d ...\n", configComp.GetGRPCPort())

	s := grpc.NewServer()

	pb.RegisterUserServiceServer(s, composer.ComposeUserGRPCService(serviceCtx))

	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}

func Execute() {
	rootCmd.AddCommand(outEnvCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
