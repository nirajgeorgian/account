package api

import (
	"fmt"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	app "github.com/nirajgeorgian/account/src/app"
	db "github.com/nirajgeorgian/account/src/db"
)

// AccountServer :- server base structure
type AccountServer struct {
	db *db.Database
}

// Client :- account client structure
type Client struct {
	conn    *grpc.ClientConn
	service AccountServiceClient
}

// API :- base API strcture
type API struct {
	App           *app.App
	Config        *Config
	AccountServer AccountServer
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

func ListenGRPC(api *API, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	RegisterAccountServiceServer(grpcServer, &api.AccountServer)
	reflection.Register(grpcServer)

	fmt.Printf("starting to listen on tcp: %q\n", lis.Addr().String())

	return grpcServer.Serve(lis)
}
