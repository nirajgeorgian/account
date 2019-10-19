package cmd

import (
	"context"
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	api "github.com/nirajgeorgian/account/src/api"
)

func init() {
  readAccount.Flags().StringVarP(&accountServiceURI, "accountserviceuri", "u", "", "account service uri (required)")
  readAccount.MarkFlagRequired("accountserviceuri")
  viper.BindPFlag("accountserviceuri", createAccount.Flags().Lookup("accountserviceuri"))
}

var readAccount = &cobra.Command{
  Use: "readAccount",
  Short: "readAccount an account with gRPC server on:3000",
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

		r, err := c.ReadAccount(ctx, &api.ReadAccountReq{AccountId: "8f48be25-7b21-471c-a8dc-562adec0835e"})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.Account)

		return nil
	},
}

func init() {
	RootCmd.AddCommand(readAccount)
}
