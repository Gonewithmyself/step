package main

import (
	"context"
	"step/microsvc/grpc/proto"
)

type hello struct {
}

func (h *hello) Say(ctx context.Context, q *proto.Query) (*proto.Result, error) {
	return &proto.Result{Mean: "Hello"}, nil
}
