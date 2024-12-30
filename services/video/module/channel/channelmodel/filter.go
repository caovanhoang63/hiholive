package channelmodel

type ChannelFilter struct {
	IsLive   *bool  `json:"isLive" form:"isLive"`
	UserName string `json:"userName" form:"userName"`
}
