package authgrpcrepo

import (
	"fmt"
	pb2 "github.com/caovanhoang63/hiholive/shared/go/proto/pb"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type rpcClient struct {
	client pb2.UserServiceClient
}

func NewClient(client pb2.UserServiceClient) *rpcClient {
	return &rpcClient{client: client}
}

func (c *rpcClient) CreateUser(ctx context.Context, firstName, lastName, email string) (newId int, err error) {
	resp, err := c.client.CreateUser(ctx, &pb2.CreateUserReq{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	})

	fmt.Println("resp: ", resp)
	fmt.Println("err: ", err)

	if err != nil {
		return 0, errors.WithStack(err)
	}

	return int(resp.Id), nil
}
