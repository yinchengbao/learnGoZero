package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"login/internal/config"
	"login/internal/middleware"
	"login/model"
)

type ServiceContext struct {
	Config              config.Config
	UserModel           model.UserModel
	UserAgentMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	dbconn := sqlx.NewMysql(c.Db.DataSource)
	return &ServiceContext{
		Config:              c,
		UserModel:           model.NewUserModel(dbconn),
		UserAgentMiddleware: middleware.NewUserAgentMiddleware().Handle,
	}
}
