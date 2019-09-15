package api

import (
	"context"

	"google.golang.org/grpc"

	"github.com/nirajgeorgian/account/src/model"
	"github.com/nirajgeorgian/account/src/api"
)

// Client :- account client structure
type Client struct {
	conn    *grpc.ClientConn
	service api.AccountServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := api.NewAccountServiceClient(conn)
	return &Client{conn, c}, nil
}

func (c *Client) CreateAccount(ctx context.Context, account model.Account) (*model.Account, error) {
	r, err := c.service.CreateAccount(
		ctx,
		&api.CreateAccountReq{Account: &account},
	)
	if err != nil {
		return nil, err
	}
	return r.Account, nil
}

func (c *Client) Close() {
	c.conn.Close()
}