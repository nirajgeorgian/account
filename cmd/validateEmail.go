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
  validateEmail.Flags().StringVarP(&accountServiceURI, "accountserviceuri", "u", "", "account service uri (required)")
  validateEmail.MarkFlagRequired("accountserviceuri")
  viper.BindPFlag("accountserviceuri", createAccount.Flags().Lookup("accountserviceuri"))
}

var validateEmail = &cobra.Command{
  Use: "validateEmail",
  Short: "validate an account with email with gRPC server on:3000",
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

		r, err := c.ValidateEmail(ctx, &api.ValidateEmailReq{Email: "dododuck@example.com"})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %t", r.Success)

		return nil
	},
}

func init() {
	RootCmd.AddCommand(validateEmail)
}
