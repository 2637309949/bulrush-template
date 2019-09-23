// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"context"
	"log"
	"net/http"

	"github.com/2637309949/bulrush"
	"github.com/2637309949/bulrush-template/addition"
	"github.com/2637309949/bulrush-template/conf"
	"github.com/2637309949/bulrush-template/grpc/pb"
	"github.com/gin-gonic/gin"
	"github.com/kataras/go-events"
	"google.golang.org/grpc"
)

func helloGrpc(router *gin.RouterGroup, event events.EventEmmiter) {
	router.GET("/test/grpc", func(c *gin.Context) {
		srv := &struct {
			Srv1 string
		}{}
		if err := conf.Conf.Unmarshal("grpc", srv); err != nil {
			panic(err)
		}
		ret := addition.GRPC(srv.Srv1)(func(conn *grpc.ClientConn, ctx context.Context) interface{} {
			cli := pb.NewMessageClient(conn)
			r, err := cli.SendEmail(ctx, &pb.EmailRequest{To: []string{"2637309949@qq.com"}})
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}
			return r
		})
		c.JSON(http.StatusOK, ret)
	})
}

// RegisterGRPC defined test routes
func RegisterGRPC(ri *bulrush.ReverseInject) {
	ri.Register(helloGrpc)
}
