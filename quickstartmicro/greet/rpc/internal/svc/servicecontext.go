package svc

import "github.com/ninekiss/go-zero-demo/quickstartmicro/greet/rpc/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
