// Copyright 2019 nirajgeorgian. All rights reserved.
// source: src/api/api.go


package api

import (
	"fmt"
	"net"
	"time"

	log "github.com/Sirupsen/logrus"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc/reflection"

	app "github.com/nirajgeorgian/account/src/app"
	db "github.com/nirajgeorgian/account/src/db"
)

// AccountServer gives a base structure for account server.
type AccountServer struct {
	db *db.Database
}

// API gives a base structure for Account API
type API struct {
	App           *app.App
	Config        *Config
	AccountServer AccountServer
}

// New new api instance
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

// Auth resolves with username and password and returns JSON web token
func (s *AccountServer) Auth(ctx context.Context, in *AuthReq) (*AuthRes, error) {
	log.Println("server: Auth")

	account, err := s.db.Auth(ctx, in.Account)
	if err != nil {
		return nil, err
	}

	token, err := s.Encode(account)
	if err != nil {
		return nil, err
	}

	// create token
	return &AuthRes{Token: token, Valid: true}, nil
}

// ValidateUsername validates username and returns weather the username is available or not.
func (s *AccountServer) ValidateUsername(ctx context.Context, in *ValidateUsernameReq) (*ValidateUsernameRes, error) {
	fmt.Println("validating username")

	success, err := s.db.ValidateUsername(ctx, in.Username)
	if err != nil {
		return nil, err
	}

	return &ValidateUsernameRes{Success: success}, nil
}

// ValidateEmail validates email and returns weather the email is available or not.
func (s *AccountServer) ValidateEmail(ctx context.Context, in *ValidateEmailReq) (*ValidateEmailRes, error) {
	fmt.Println("validating email")

	success, err := s.db.ValidateEmail(ctx, in.Email)
	if err != nil {
		return nil, err
	}

	return &ValidateEmailRes{Success: success}, nil
}

// CreateAccount creates a rpc network call to create an user profile
func (s *AccountServer) CreateAccount(ctx context.Context, in *CreateAccountReq) (*CreateAccountRes, error) {
	log.Println("server: CreateAccount")

	account, err := s.db.CreateAccount(ctx, in.Account)
	if err != nil {
		return nil, err
	}

	return &CreateAccountRes{Account: account}, nil
}

// UpdateAccount creates a rpc network call to update an user profile
func (s *AccountServer) UpdateAccount(ctx context.Context, in *UpdateAccountReq) (*UpdateAccountRes, error) {
	fmt.Println("updating profile")

	account, err := s.db.UpdateAccount(ctx, in.Account)
	if err != nil {
		return nil, err
	}

	return &UpdateAccountRes{Success: true, Account: account}, nil
}

// ReadAccount creates a rpc network call to read user account information
func (s *AccountServer) ReadAccount(ctx context.Context, in *ReadAccountReq) (*ReadAccountRes, error) {
	fmt.Println("reading single profile")

	account, err := s.db.ReadAccount(ctx, in.AccountId)
	if err != nil {
		return nil, err
	}

	return &ReadAccountRes{Account: account}, nil
}

// ListenGRPC serves the account API on GRPC server
func ListenGRPC(api *API, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	recoveryFunc := func(p interface{}) (err error) {
		log.Errorf("recovered from panic: %v", p)
		return fmt.Errorf("recovered from panic: %v", p)
	}
	recoveryOpts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(recoveryFunc),
	}
	grpcServer := grpc.NewServer(
		grpc.ConnectionTimeout(time.Minute*30),
		grpc.MaxRecvMsgSize(1024*1024*128),
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(recoveryOpts...),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_recovery.StreamServerInterceptor(recoveryOpts...),
		),
	)
	RegisterAccountServiceServer(grpcServer, &api.AccountServer)
	reflection.Register(grpcServer)

	log.Printf("starting HTTP/2 gRPC API server: %q\n", lis.Addr().String())

	return grpcServer.Serve(lis)
}

func (s *AccountServer) Auth(ctx context.Context, in *AuthReq) (*AuthRes, error) {
	log.Println("server: Auth")

	account, err := s.db.Auth(ctx, in.Account)
	if err != nil {
		return nil, err
	}

	token, err := s.Encode(account)
	if err != nil {
		return nil, err
	}

	// create token
	return &AuthRes{Token: token, Valid: true}, nil
}
