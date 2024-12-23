package usersub

import (
	"github.com/caovanhoang63/hiholive/services/user/module/user/biz"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
)

type UserSub struct {
	biz        biz.UserBiz
	serviceCtx srvctx.ServiceContext
}

func NewUserSub(biz biz.UserBiz, serviceCtx srvctx.ServiceContext) *UserSub {
	return &UserSub{
		biz:        biz,
		serviceCtx: serviceCtx,
	}
}
