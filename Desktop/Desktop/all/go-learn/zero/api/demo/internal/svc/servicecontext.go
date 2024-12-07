package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"zero-study/api/demo/internal/config"
	"zero-study/api/demo/internal/middleware"
	"zero-study/api/model"
)

type ServiceContext struct {
	Config         config.Config
	TestMiddleware rest.Middleware
	UserModel      model.UserModel
	UserDataModel  model.UserDataModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		UserModel:      model.NewUserModel(sqlx.NewMysql(c.DB.Datasource), c.Cache),
		UserDataModel:  model.NewUserDataModel(sqlx.NewMysql(c.DB.Datasource), c.Cache),
		TestMiddleware: middleware.NewTestMiddleware().Handle,
	}
}
