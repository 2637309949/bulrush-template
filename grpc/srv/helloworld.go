package srv

import (
	pb "github.com/2637309949/bulrush-template/grpc/pb"
	"golang.org/x/net/context"
)

// SendMessage is used to implement pb.SendMessage.
type SendMessage struct{}

// SendEmail implements pb.SendMessage
func (s *SendMessage) SendEmail(ctx context.Context, in *pb.EmailRequest) (*pb.EmailReply, error) {
	return &pb.EmailReply{Status: pb.EmailReply_Success, Message: "send successfully"}, nil
}
