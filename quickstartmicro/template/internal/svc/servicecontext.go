package svc

import "github.com/ninekiss/go-zero-demo/quickstartmicro/template/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
