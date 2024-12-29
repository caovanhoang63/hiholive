package strmsub

import (
	"fmt"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/pubsub"
	"github.com/caovanhoang63/hiholive/shared/golang/subengine"
	"golang.org/x/net/context"
	"strconv"
)

type UpdateStreamViewCountData struct {
	Id        float64 `json:"id"`
	View      float64 `json:"view"`
	TimeStamp string  `json:"timeStamp"`
}

func (s *StreamSub) UpdateStreamViewCount() subengine.ConsumerJob {
	return subengine.ConsumerJob{
		Title: "Update stream view count",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			data, ok := message.Data.(map[string]interface{})
			fmt.Println(data, ok)
			if !ok {
				return fmt.Errorf("invalid message data format, expected map[string]interface{}")
			}

			idStr, ok := data["id"].(string)
			if !ok {
				return fmt.Errorf("invalid or missing 'id' in message data")
			}
			id, _ := strconv.Atoi(idStr)

			viewStr, ok := data["view"].(float64)
			if !ok {
				return fmt.Errorf("invalid or missing 'view' in message data")
			}
			view := int(viewStr)

			if err := s.biz.UpdateStreamView(ctx, id, view); err != nil {
				return fmt.Errorf("failed to update stream view count: %w", err)
			}

			return nil
		},
	}
}
