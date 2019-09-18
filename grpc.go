// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"net"

	tgrpc "github.com/2637309949/bulrush-template/grpc"

	ggrpc "google.golang.org/grpc"
)

const (
	port = ":50051"
)

func grpc() *ggrpc.Server {
	s := ggrpc.NewServer()
	tgrpc.GRPC(s)
	return s
}

func runGrpc() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	go func() {
		if err := grpc().Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
}
