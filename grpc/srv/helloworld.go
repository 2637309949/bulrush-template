package srv

import (
	pb "github.com/2637309949/bulrush-template/grpc/pb"
	"golang.org/x/net/context"
)

// Greeter is used to implement pb.GreeterServer.
type Greeter struct{}

// SayHello implements helloworld.GreeterServer
func (s *Greeter) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
