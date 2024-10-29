package rtmpc

import flvtag "github.com/yutopp/go-flv/tag"

type Pub struct {
	pb *Pubsub

	avcSeqHeader *flvtag.FlvTag
	lastKeyFrame *flvtag.FlvTag
}

// Publish TODO: Should check codec types and so on.
// In this example, checks only sequence headers and assume that AAC and AVC.
func (p *Pub) Publish(flv *flvtag.FlvTag) error {
	switch flv.Data.(type) {
	case *flvtag.AudioData, *flvtag.ScriptData:
		for _, sub := range p.pb.subs {
			_ = sub.onEvent(cloneView(flv))
		}

	case *flvtag.VideoData:
		d := flv.Data.(*flvtag.VideoData)
		if d.AVCPacketType == flvtag.AVCPacketTypeSequenceHeader {
			p.avcSeqHeader = flv
		}

		if d.FrameType == flvtag.FrameTypeKeyFrame {
			p.lastKeyFrame = flv
		}

		for _, sub := range p.pb.subs {
			if !sub.initialized {
				if p.avcSeqHeader != nil {
					_ = sub.onEvent(cloneView(p.avcSeqHeader))
				}
				if p.lastKeyFrame != nil {
					_ = sub.onEvent(cloneView(p.lastKeyFrame))
				}
				sub.initialized = true
				continue
			}

			_ = sub.onEvent(cloneView(flv))
		}

	default:
		panic("unexpected")
	}

	return nil
}

func (p *Pub) Close() error {
	return p.pb.Deregister()
}
