package svc

import (
	"st-key-manager/internal/config"
	"st-key-manager/internal/svc/dao"
)

type ServiceContext struct {
	Config config.Config
	Redis  *dao.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Redis:  dao.NewRedis(c.RedisUrl),
	}
}
