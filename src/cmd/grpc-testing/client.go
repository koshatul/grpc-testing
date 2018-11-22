package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

var cmdClient = &cobra.Command{
	Use:   "client [word]",
	Short: "Run Client",
	Run:   clientCommand,
	Args:  cobra.MaximumNArgs(1),
}

func init() {
	rootCmd.AddCommand(cmdClient)
}

func clientCommand(cmd *cobra.Command, args []string) {
	// Set up a connection to the server.
	creds, _ := credentials.NewClientTLSFromFile(viper.GetString("grpc.client-cert-file"), "")
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", viper.GetString("grpc.host"), viper.GetInt("grpc.port")), grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := "worldish"
	if len(args) >= 1 {
		name = strings.Join(args, " ")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
