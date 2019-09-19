package addition

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"
)

// Request defined Request global proxy
func Request(method string, url string, payload io.Reader) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}
	return client.Do(req)
}

// GRPC defined grpc call
func GRPC(address string, opts ...grpc.DialOption) func(func(*grpc.ClientConn, context.Context) interface{}) interface{} {
	opt := grpc.WithInsecure()
	if len(opts) > 0 {
		opt = opts[0]
	}
	conn, err := grpc.Dial(address, opt)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	return func(cb func(*grpc.ClientConn, context.Context) interface{}) interface{} {
		defer conn.Close()
		defer cancel()
		return cb(conn, ctx)
	}
}
