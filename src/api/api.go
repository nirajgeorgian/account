package api

import (
	"fmt"

	context "golang.org/x/net/context"

	app "github.com/nirajgeorgian/account/src/app"
	db "github.com/nirajgeorgian/account/src/db"
)

// AccountServer :- server base structure
type AccountServer struct {
	db *db.Database
}

// API :- base API strcture
type API struct {
	AccountServer AccountServer
	App           *app.App
	Config        *Config
}

// New :- new api instance
func New(a *app.App) (api *API, err error) {
	api = &API{App: a}
	api.Config, err = InitConfig()
	if err != nil {
		return nil, err
	}

	s := AccountServer{db: a.Database}
	api.AccountServer = s

	return api, nil
}

// CreateAccount :- CreateAccount rpc network call
func (s *AccountServer) CreateAccount(ctx context.Context, in *CreateAccountReq) (*CreateAccountRes, error) {
	fmt.Println("creating account")

	err := s.db.CreateAccount(ctx, in.Account)
	if err != nil {
		return nil, err
	}
	account := in.Account

	return &CreateAccountRes{Account: account}, nil
}
