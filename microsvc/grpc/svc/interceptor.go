package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func unaryInterceptor(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("===[in ]", info.FullMethod)

	m, err := handler(ctx, req)

	log.Println("===[out]", m)

	return m, err
}

type wrappedStream struct {
	grpc.ServerStream
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	log.Println("[wrap recv]")
	return w.ServerStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	log.Println("[wrap send]")
	return w.ServerStream.SendMsg(m)
}

func newWrappedStream(s grpc.ServerStream) grpc.ServerStream {
	return &wrappedStream{s}
}

func streamInterceptor(srv interface{}, ss grpc.ServerStream,
	info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println("[start stream]", info.FullMethod)
	err := handler(srv, newWrappedStream(ss))
	log.Println("[close stream]", info.FullMethod)
	return err
}
