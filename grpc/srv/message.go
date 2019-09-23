package srv

import (
	pb "github.com/2637309949/bulrush-template/grpc/pb"
	"golang.org/x/net/context"
)

// Message is used to implement pb.Messages.
type Message struct{}

// SendCode implements pb.Message
func (s *Message) SendCode(ctx context.Context, in *pb.CodeRequest) (*pb.CodeReply, error) {
	return &pb.CodeReply{Status: pb.CodeReply_Success, Message: "send successfully"}, nil
}

// SendEmail implements pb.Message
func (s *Message) SendEmail(ctx context.Context, in *pb.EmailRequest) (*pb.EmailReply, error) {
	return &pb.EmailReply{Status: pb.EmailReply_Success, Message: "send successfully"}, nil
}
