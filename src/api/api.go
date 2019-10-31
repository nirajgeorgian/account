// Copyright 2019 nirajgeorgian. All rights reserved.
// source: src/api/api.go
// package: api

package api

import (
	"fmt"
	"net"
	"time"
	"sync"

	log "github.com/Sirupsen/logrus"

	"go.opencensus.io/trace"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/exporter/zipkin"

	openzipkin "github.com/openzipkin/zipkin-go"
	zipkinHTTP "github.com/openzipkin/zipkin-go/reporter/http"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"github.com/spf13/viper"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc/reflection"

	app "github.com/nirajgeorgian/account/src/app"
	db "github.com/nirajgeorgian/account/src/db"
)

// AccountServer gives a base structure for account server.
type AccountServer struct {
	mu sync.Mutex
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

	ctx, span := trace.StartSpan(ctx, "account.grpc.api.Auth")
	defer span.End()

	account, err := s.db.Auth(ctx, in.Account)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, err
	}

	token, err := s.Encode(account)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, err
	}

	span.Annotate([]trace.Attribute{
		trace.StringAttribute("fetch", "Auth"),
	}, "account /src/api")

	return &AuthRes{Token: token, Valid: true}, nil
}

// ValidateUsername validates username and returns weather the username is available or not.
func (s *AccountServer) ValidateUsername(ctx context.Context, in *ValidateUsernameReq) (*ValidateUsernameRes, error) {
	fmt.Println("validating username")

	ctx, span := trace.StartSpan(ctx, "account.grpc.api.ValidateUsername")
	defer span.End()

	success, err := s.db.ValidateUsername(ctx, in.Username)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, err
	}

	span.Annotate([]trace.Attribute{
		trace.StringAttribute("fetch", "ValidateUsername"),
	}, "account /src/api")

	return &ValidateUsernameRes{Success: success}, nil
}

// ValidateEmail validates email and returns weather the email is available or not.
func (s *AccountServer) ValidateEmail(ctx context.Context, in *ValidateEmailReq) (*ValidateEmailRes, error) {
	fmt.Println("validating email")

	ctx, span := trace.StartSpan(ctx, "account.grpc.api.ValidateEmail")
	defer span.End()

	success, err := s.db.ValidateEmail(ctx, in.Email)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, err
	}

	span.Annotate([]trace.Attribute{
		trace.StringAttribute("fetch", "ValidateEmail"),
	}, "account /src/api")

	return &ValidateEmailRes{Success: success}, nil
}

// CreateAccount creates a rpc network call to create an user profile
func (s *AccountServer) CreateAccount(ctx context.Context, in *CreateAccountReq) (*CreateAccountRes, error) {
	log.Println("server: CreateAccount")

	ctx, span := trace.StartSpan(ctx, "account.grpc.api.CreateAccount")
	defer span.End()

	accountPresent, err := s.db.ValidateUsername(ctx, in.Account.Username)
	if err  != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, err
	}
	if accountPresent {
		err := status.Error(codes.NotFound, "account already exist")
		return nil, err
	}

	accountPresent, err = s.db.ValidateEmail(ctx, in.Account.Email)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, err
	}
	if accountPresent {
		err := status.Error(codes.NotFound, "account already exist")
		return nil, err
	}

	account, err := s.db.CreateAccount(ctx, in.Account)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, err
	}

	span.Annotate([]trace.Attribute{
		trace.StringAttribute("fetch", "CreateAccount"),
	}, "account /src/api")

	return &CreateAccountRes{Account: account}, nil
}

func (s *AccountServer) ConfirmAccount(ctx context.Context, in *ConfirmAccountReq) (*ConfirmAccountRes, error) {
	log.Println("server: ConfirmAccount")

	ctx, span := trace.StartSpan(ctx, "account.grpc.api.CreateAccount")
	defer span.End()

	accountPresent, err := s.db.ValidateUsername(ctx, in.Username)
	if err  != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, err
	}
	if accountPresent {
		err := status.Error(codes.NotFound, "account already exist")
		return nil, err
	}

	_, err = s.db.ValidateEmail(ctx, in.Email)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, err
	}

	confirmed, err := s.db.ConfirmAccount(ctx, in.Email, in.Token)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, err
	}

	if confirmed {
		return &ConfirmAccountRes{
			Success: true,
		}, nil
	}
	return &ConfirmAccountRes{
		Success: false,
	}, nil
}

// UpdateAccount creates a rpc network call to update an user profile
func (s *AccountServer) UpdateAccount(ctx context.Context, in *UpdateAccountReq) (*UpdateAccountRes, error) {
	fmt.Println("updating profile")

	ctx, span := trace.StartSpan(ctx, "account.grpc.account.UpdateAccount")
	defer span.End()

	account, err := s.db.UpdateAccount(ctx, in.Account)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, err
	}

	span.Annotate([]trace.Attribute{
		trace.StringAttribute("fetch", "UpdateAccount"),
	}, "account /src/api")

	return &UpdateAccountRes{Success: true, Account: account}, nil
}

// ReadAccount creates a rpc network call to read user account information
func (s *AccountServer) ReadAccount(ctx context.Context, in *ReadAccountReq) (*ReadAccountRes, error) {
	fmt.Println("reading single profile")

	ctx, span := trace.StartSpan(ctx, "account.grpc.account.ReadAccount")
	defer span.End()

	account, err := s.db.ReadAccount(ctx, in.AccountId)
	if err != nil {
		span.SetStatus(trace.Status{
			Code:    trace.StatusCodeUnknown,
			Message: err.Error(),
		})
		return nil, err
	}

	span.Annotate([]trace.Attribute{
		trace.StringAttribute("fetch", "ReadAccount"),
	}, "account /src/api")

	return &ReadAccountRes{Account: account}, nil
}

func (s *AccountServer) AccountCheck(ctx context.Context, in *AccountHealthCheckReq) (*AccountHealthCheckRes, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if in.Service == "" {
		// check the server overall health status.
		return &AccountHealthCheckRes{
			Status: AccountHealthCheckRes_SERVING,
		}, nil
	}

	statusMap := make(map[string]AccountHealthCheckRes_ServingStatus)
	if status, ok := statusMap[in.Service]; ok {
		return &AccountHealthCheckRes{
			Status: status,
		}, nil
	}

	return nil, status.Error(codes.NotFound, "unknown service")
}

// ListenGRPC serves the account API on GRPC server
func ListenGRPC(api *API, port int) error {
	if err := view.Register(ocgrpc.DefaultServerViews...); err != nil {
		log.Fatalf("Failed to register ocgrpc server views: %v", err)
	}

	localURL := viper.GetString("localendpoint")
	zipkinURL := viper.GetString("zipkinendpoint")

	localEndpoint, err := openzipkin.NewEndpoint("account-svc", localURL)
	if err != nil {
		log.Fatalf("Failed to create the local zipkinEndpoint: %v", err)
	}

	reporter := zipkinHTTP.NewReporter(zipkinURL + "/api/v2/spans")
	ze := zipkin.NewExporter(reporter, localEndpoint)
	trace.RegisterExporter(ze)

	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

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
		grpc.StatsHandler(&ocgrpc.ServerHandler{}),
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
