package main

import (
	"github.com/micro-plat/hydra/context"
)

type Service struct{}

//Handle 业务处理函数
func (s Service) Handle(ctx context.IContext) (r interface{}) {
	ctx.Log().Info("Service-Handle:")
	ctx.Log().Info("Service-Encoding:", ctx.Request().Path().GetEncoding())
	router, err := ctx.Request().Path().GetRouter()
	if err != nil {
		return err
	}
	ctx.Log().Info("Service-Handle-router:", router)
	return
}

type Service2 struct{}

//Handle 业务处理函数
func (s *Service2) Handle(ctx context.IContext) (r interface{}) {
	ctx.Log().Info("Service2-Handle")
	ctx.Log().Info("Service-Encoding:", ctx.Request().Path().GetEncoding())
	router, err := ctx.Request().Path().GetRouter()
	if err != nil {
		return err
	}
	ctx.Log().Info("Service2-Handle-router:", router)
	return
}

//QueryHandle 业务处理函数
func (s *Service2) QueryHandle(ctx context.IContext) (r interface{}) {
	ctx.Log().Info("Service2-query-Handle")
	ctx.Log().Info("Service-Encoding:", ctx.Request().Path().GetEncoding())
	router, err := ctx.Request().Path().GetRouter()
	if err != nil {
		return err
	}
	ctx.Log().Info("Service2-query-router:", router)
	return
}

func (s *Service2) QueryFallback(ctx context.IContext) (r interface{}) {
	ctx.Log().Info("Service2-query-FallBack")
	router, err := ctx.Request().Path().GetRouter()
	if err != nil {
		return err
	}
	ctx.Log().Info("Service2-fallback-router:", router)
	return
}

type Service3 struct{}

//QueryHandle 业务处理函数
func (s *Service3) PostHandle(ctx context.IContext) (r interface{}) {
	ctx.Log().Info("Service3-post-Handle")
	return
}

//NewService 服务对象构建函数
var NewService func() (Service, error) = func() (Service, error) {
	return Service{}, nil
}

//ServiceFunc
var ServiceFunc func(context.IContext) interface{} = func(ctx context.IContext) (r interface{}) {
	ctx.Log().Info("func-Handle")
	return
}
