package main

import (
	"flag"
	"fmt"

	"github.com/ninekiss/go-zero-demo/quickstartmicro/greet/api/internal/config"
	"github.com/ninekiss/go-zero-demo/quickstartmicro/greet/api/internal/handler"
	"github.com/ninekiss/go-zero-demo/quickstartmicro/greet/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/greet.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
