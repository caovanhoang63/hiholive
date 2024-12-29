package cmd

import (
	"github.com/caovanhoang63/hiholive/services/hls_mux/component/ffmpegc"
	"github.com/caovanhoang63/hiholive/services/hls_mux/module/transport/hlssub"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx/components/pubsub"
	"github.com/caovanhoang63/hiholive/shared/go/subengine"
)

func StartSubscriberServices(ffmpeg *ffmpegc.Ffmpeg, serviceCtx srvctx.ServiceContext) {
	pb := serviceCtx.MustGet(core.KeyCompRabbitMQ).(pubsub.Pubsub)
	service := hlssub.NewHLSSub(ffmpeg)
	engine := subengine.NewEngine(serviceCtx, pb)
	engine.Subscribe(core.TopicStreamEnded, service.OnStopStream())
	_ = engine.Start()
}
