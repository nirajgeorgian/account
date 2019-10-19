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
  updateAccount.Flags().StringVarP(&accountServiceURI, "accountserviceuri", "u", "", "account service uri (required)")
  updateAccount.MarkFlagRequired("accountserviceuri")
  viper.BindPFlag("accountserviceuri", createAccount.Flags().Lookup("accountserviceuri"))
}

var updateAccount = &cobra.Command{
  Use: "updateAccount",
  Short: "update an account with gRPC server on:3000",
	RunE: func(cmd *cobra.Command, args []string) error {
		address     := viper.GetString("accountserviceuri")

		// Set up a connection to the server.
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := api.NewAccountServiceClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
		defer cancel()

		account := model.Account{
			AccountId: "1",
			Username: "updateusername",
			Email: "test",
			Description: "update 1 desc",
		}
		r, err := c.UpdateAccount(ctx, &api.UpdateAccountReq{Account: &account})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %t %s", r.Success, r.Account)

		return nil
	},
}

func init() {
	RootCmd.AddCommand(updateAccount)
}
