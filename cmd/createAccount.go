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

var AccountServiceURI string

func init() {
  createAccount.Flags().StringVarP(&AccountServiceURI, "accountserviceuri", "u", "", "account service uri (required)")
  createAccount.MarkFlagRequired("accountserviceuri")
  viper.BindPFlag("accountserviceuri", createAccount.Flags().Lookup("accountserviceuri"))
}

var createAccount = &cobra.Command{
  Use: "createAccount",
  Short: "create an account with gRPC server on:3000",
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
			PasswordHash: "test123",
			PasswordSalt: "test123",
			Description: "dodo duck lives here",
		}
		r, err := c.CreateAccount(ctx, &api.CreateAccountReq{Account: &account})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.Account.AccountId)

		return nil
	},
}

func init() {
	RootCmd.AddCommand(createAccount)
}
