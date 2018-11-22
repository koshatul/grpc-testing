package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/reflection"
)

var cmdServer = &cobra.Command{
	Use:   "server",
	Short: "Start Server",
	Run:   serverCommand,
	Args:  cobra.NoArgs,
}

func init() {
	rootCmd.AddCommand(cmdServer)
}

func serverCommand(cmd *cobra.Command, args []string) {
	creds, _ := credentials.NewServerTLSFromFile(viper.GetString("grpc.server-cert-file"), viper.GetString("grpc.server-key-file"))
	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", viper.GetString("grpc.host"), viper.GetInt("grpc.port")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %s", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
