package strmgrpc

import (
	"github.com/caovanhoang63/hiholive/services/video/module/stream/streambiz"
	"github.com/caovanhoang63/hiholive/shared/golang/proto/pb"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx"
	"golang.org/x/net/context"
)

type streamGRPC struct {
	biz    streambiz.StreamBiz
	srvctx srvctx.ServiceContext
}

func NewStreamGRPC(biz streambiz.StreamBiz, srvctx srvctx.ServiceContext) *streamGRPC {
	return &streamGRPC{
		biz:    biz,
		srvctx: srvctx,
	}
}
func (s *streamGRPC) FindStreamById(ctx context.Context, req *pb.FindStreamReq) (*pb.FindStreamResp, error) {
	r, err := s.biz.FindStreamById(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.FindStreamResp{
		Title:     r.Title,
		State:     r.State,
		Status:    int32(r.Status),
		ChannelId: int32(r.ChannelId),
	}, nil
}
