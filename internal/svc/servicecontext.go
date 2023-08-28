package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"login/internal/config"
	"login/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	dbconn := sqlx.NewMysql(c.Db.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(dbconn),
	}
}
