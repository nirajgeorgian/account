package api

import (
	"context"

	"google.golang.org/grpc"

	"github.com/nirajgeorgian/account/src/model"
)

// Client :- account client structure
type Client struct {
	conn    *grpc.ClientConn
	service AccountServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := NewAccountServiceClient(conn)
	return &Client{conn, c}, nil
}

func (c *Client) CreateAccount(ctx context.Context, account model.Account) (*model.Account, error) {
	r, err := c.service.CreateAccount(
		ctx,
		&CreateAccountReq{Account: &account},
	)
	if err != nil {
		return nil, err
	}
	return r.Account, nil
}

func (c *Client) UpdateAccount(ctx context.Context, account model.Account) (*UpdateAccountRes, error) {
	r, err := c.service.UpdateAccount(
		ctx,
		&UpdateAccountReq{Account: &account},
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Client) ReadAccount(ctx context.Context, account_id string) (*ReadAccountRes, error) {
	r, err := c.service.ReadAccount(
		ctx,
		&ReadAccountReq{AccountId: account_id},
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Client) Auth(ctx context.Context, account model.Account) (*AuthRes, error) {
	r, err := c.service.Auth(
		ctx,
		&AuthReq{Account: &account},
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Client) Close() {
	c.conn.Close()
}
