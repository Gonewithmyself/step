package main

import (
	"net"
	"step/microsvc/grpc/proto"

	"google.golang.org/grpc"
)

func main() {
	addr := "0.0.0.0:9999"
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	svr := grpc.NewServer(
		grpc.UnaryInterceptor(unaryInterceptor),
		grpc.StreamInterceptor(streamInterceptor),
	)

	proto.RegisterMonsterServer(svr, &monster{
		chatroom: chatroom{
			conns: map[int32]proto.Monster_ChatServer{},
		},
	})

	proto.RegisterHelloServer(svr, &hello{})

	svr.Serve(ln)
}
