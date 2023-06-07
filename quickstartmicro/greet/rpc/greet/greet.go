// Code generated by goctl. DO NOT EDIT.
// Source: greet.proto

package greet

import (
	"context"

	"github.com/ninekiss/go-zero-demo/quickstartmicro/greet/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Placeholder = pb.Placeholder

	Greet interface {
		Ping(ctx context.Context, in *Placeholder, opts ...grpc.CallOption) (*Placeholder, error)
	}

	defaultGreet struct {
		cli zrpc.Client
	}
)

func NewGreet(cli zrpc.Client) Greet {
	return &defaultGreet{
		cli: cli,
	}
}

func (m *defaultGreet) Ping(ctx context.Context, in *Placeholder, opts ...grpc.CallOption) (*Placeholder, error) {
	client := pb.NewGreetClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}
