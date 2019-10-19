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

	account, err := s.db.CreateAccount(ctx, in.Account)
	if err != nil {
		return nil, err
	}

	return &CreateAccountRes{Account: account}, nil
}

func (s *AccountServer) UpdateAccount(ctx context.Context, in *UpdateAccountReq) (*UpdateAccountRes, error) {
	fmt.Println("updating profile")

	account, err := s.db.UpdateAccount(ctx, in.Account)
	if err != nil {
		return nil, err
	}

	return &UpdateAccountRes{Success: true, Account: account}, nil
}

func (s *AccountServer) ReadAccount(ctx context.Context, in *ReadAccountReq) (*ReadAccountRes, error) {
	fmt.Println("reading single profile")

	account, err := s.db.ReadAccount(ctx, in.AccountId)
	if err != nil {
		return nil, err
	}

	return &ReadAccountRes{Account: account}, nil
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
