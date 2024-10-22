package cmd

import (
	"fmt"
	"github.com/caovanhoang63/hiholive/hls_mux/component/ffmpegc"
	"github.com/caovanhoang63/hiholive/shared/shared"
	"github.com/caovanhoang63/hiholive/shared/srvctx"
	"github.com/caovanhoang63/hiholive/shared/srvctx/components/ginc"
	"github.com/caovanhoang63/hiholive/shared/srvctx/components/gormc"
	"github.com/caovanhoang63/hiholive/shared/srvctx/components/jwtc"
	"github.com/spf13/cobra"

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

func Execute() {
	rootCmd.AddCommand(outEnvCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
