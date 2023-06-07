package main

import (
	"flag"
	"fmt"

	"github.com/ninekiss/go-zero-demo/quickstartmicro/template/internal/config"
	templateServer "github.com/ninekiss/go-zero-demo/quickstartmicro/template/internal/server/template"
	"github.com/ninekiss/go-zero-demo/quickstartmicro/template/internal/svc"
	"github.com/ninekiss/go-zero-demo/quickstartmicro/template/template"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/template.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		template.RegisterTemplateServer(grpcServer, templateServer.NewTemplateServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
