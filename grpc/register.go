package grpc

import (
	pb "github.com/2637309949/bulrush-template/grpc/pb"
	srv "github.com/2637309949/bulrush-template/grpc/srv"
	"google.golang.org/grpc"
)

// Init for register srv
func Init(s *grpc.Server) {
	pb.RegisterMessageServer(s, &srv.Message{})
}
