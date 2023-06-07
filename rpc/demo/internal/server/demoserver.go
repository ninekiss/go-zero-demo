// Code generated by goctl. DO NOT EDIT.
// Source: demo.proto

package server

import (
	"context"

	"github.com/ninekiss/go-zero-demo/rpc/demo/demo"
	"github.com/ninekiss/go-zero-demo/rpc/demo/internal/logic"
	"github.com/ninekiss/go-zero-demo/rpc/demo/internal/svc"
)

type DemoServer struct {
	svcCtx *svc.ServiceContext
	demo.UnimplementedDemoServer
}

func NewDemoServer(svcCtx *svc.ServiceContext) *DemoServer {
	return &DemoServer{
		svcCtx: svcCtx,
	}
}

func (s *DemoServer) Ping(ctx context.Context, in *demo.Request) (*demo.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
