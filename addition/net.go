package addition

import (
	"context"
	"io"
	"net/http"
	"time"

	"google.golang.org/grpc"
	"gopkg.in/gomail.v2"
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
		Logger.Error("did not connect: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	return func(cb func(*grpc.ClientConn, context.Context) interface{}) interface{} {
		defer conn.Close()
		defer cancel()
		return cb(conn, ctx)
	}
}

// DialEmail defined send email
func DialEmail(ret func(m *gomail.Message) *gomail.Message) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "alex@example.com")
	// m.SetHeader("To", "bob@example.com", "cora@example.com")
	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	// m.SetHeader("Subject", "Hello!")
	// m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	// m.Attach("/home/Alex/lolcat.jpg")
	d := gomail.NewDialer("smtp.example.com", 587, "user", "123456")
	if err := d.DialAndSend(ret(m)); err != nil {
		Logger.Error("did not connect: %v", err)
		return err
	}
	return nil
}
