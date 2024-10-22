package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"hiholive/projects/go/hls_mux/component/ffmpegc"
	"hiholive/projects/go/user/composer"
	"hiholive/shared/go/shared"
	"hiholive/shared/go/srvctx"
	"hiholive/shared/go/srvctx/components/ginc"
	"hiholive/shared/go/srvctx/components/gormc"
	"hiholive/shared/go/srvctx/components/jwtc"
	"os"
)

func newServiceCtx() srvctx.ServiceContext {
	return srvctx.NewServiceContext(
		srvctx.WithName("Demo Microservices"),
		srvctx.WithComponent(ginc.NewGin(shared.KeyCompGIN)),
		srvctx.WithComponent(gormc.NewGormDB(shared.KeyCompMySQL, "")),
		srvctx.WithComponent(jwtc.NewJWT(shared.KeyCompJWT)),
		srvctx.WithComponent(NewConfig()),
	)
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start service",
	Run: func(cmd *cobra.Command, args []string) {
		serviceCtx := newServiceCtx()

		ffmpeg := ffmpegc.NewFfmpeg(serviceCtx).WithConfig(nil)
		ffmpeg.NewStream("test")
	},
}

func SetupRoutes(router *gin.RouterGroup, serviceCtx srvctx.ServiceContext) {
	userService := composer.ComposeUserAPIService(serviceCtx)

	tasks := router.Group("/user")
	{
		tasks.GET("", userService.GetUserProfile())
	}
}

func Execute() {
	rootCmd.AddCommand(outEnvCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
