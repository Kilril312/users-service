package grpc

import (
	"log"
	"net"

	"github.com/Kilril312/users-service/internal/user"
	"google.golang.org/grpc"
)

func RunGRPC(svc *user.Service) error {
	netList, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	srv := grpc.NewServer()

	userpb.RegisterUserServiceServer(srv, NewHandler(svc))

	log.Printf("Starting gRPC server on %s", netList.Addr())

	return srv.Serve(netList)
}
