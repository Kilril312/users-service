package grpc

import (
	"log"
	"net"

	userpb "github.com/Kilril312/project-protos/proto/user"
	"github.com/Kilril312/users-service/internal/user"
	"google.golang.org/grpc"
)

func RunGRPC(svc *user.Service) error {
	netList, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	srv := grpc.NewServer()
	handler := NewHandler(svc)

	userpb.RegisterUserServiceServer(srv, handler)

	log.Printf("Starting gRPC server on %s", netList.Addr())

	return srv.Serve(netList)
}
