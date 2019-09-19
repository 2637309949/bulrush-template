package srv

import (
	pb "github.com/2637309949/bulrush-template/grpc/pb"
	"golang.org/x/net/context"
)

// SendMessage is used to implement pb.SendMessage.
type SendMessage struct{}

// SendCode implements pb.SendCode
func (s *SendMessage) SendCode(ctx context.Context, in *pb.CodeRequest) (*pb.CodeReply, error) {
	return &pb.CodeReply{Status: pb.CodeReply_Success, Message: "send successfully"}, nil
}

// SendEmail implements pb.SendMessage
func (s *SendMessage) SendEmail(ctx context.Context, in *pb.EmailRequest) (*pb.EmailReply, error) {
	return &pb.EmailReply{Status: pb.EmailReply_Success, Message: "send successfully"}, nil
}
