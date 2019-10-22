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
  createAccount.Flags().StringVarP(&accountServiceURI, "accountserviceuri", "u", "", "account service uri (required)")
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

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		account := model.Account{
			AccountId: "1",
			Username: "dododuck",
			Email: "dododuck@example.com",
			PasswordHash: "test123",
			PasswordSalt: "test123",
			Description: "dodo duck lives here",
		}
		r, err := c.CreateAccount(ctx, &api.CreateAccountReq{Account: &account})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.Account)

		return nil
	},
}

func init() {
	RootCmd.AddCommand(createAccount)
}
