package cmd

import (
	"context"
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	api "github.com/nirajgeorgian/account/src/api"
	model "github.com/nirajgeorgian/account/src/model"
)

func init() {
  createAuth.Flags().StringVarP(&AccountServiceURI, "accountserviceuri", "u", "localhost:3001", "account service uri (required)")
  createAuth.MarkFlagRequired("accountserviceuri")
  viper.BindPFlag("accountserviceuri", createAccount.Flags().Lookup("accountserviceuri"))
}

var createAuth = &cobra.Command{
  Use: "auth",
  Short: "authenticate an account with gRPC server on:3001",
	RunE: func(cmd *cobra.Command, args []string) error {
		address     := viper.GetString("accountserviceuri")

		// Set up a connection to the server.
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := api.NewAccountServiceClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		account := model.Account{
			Username: "test",
			Email: "test",
			PasswordHash: "passwordtest123",
			PasswordSalt: "test123",
			Description: "dodo duck lives here",
		}
		r, err := c.Auth(ctx, &api.AuthReq{Account: &account})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.Token)

		return nil
	},
}

func init() {
	RootCmd.AddCommand(createAuth)
}
