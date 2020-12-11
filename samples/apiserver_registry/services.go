package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/context"
)

type Service struct{}

//Handle 业务处理函数
func (s Service) Handle(ctx hydra.IContext) (r interface{}) {
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
func (s *Service2) Handle(ctx hydra.IContext) (r interface{}) {
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
func (s *Service2) QueryHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("Service2-query-Handle")
	ctx.Log().Info("Service-Encoding:", ctx.Request().Path().GetEncoding())
	router, err := ctx.Request().Path().GetRouter()
	if err != nil {
		return err
	}
	ctx.Log().Info("Service2-query-router:", router)
	return
}

//Fallback 业务处理降级函数
func (s *Service2) Fallback(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("Service2-FallBack")
	router, err := ctx.Request().Path().GetRouter()
	if err != nil {
		return err
	}
	ctx.Log().Info("Service2-router:", router)
	return
}

//QueryFallback 业务处理降级函数
func (s *Service2) QueryFallback(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("Service2-query-FallBack")
	router, err := ctx.Request().Path().GetRouter()
	if err != nil {
		return err
	}
	ctx.Log().Info("Service2-fallback-router:", router)
	return
}

type Service3 struct{}

//PostHandle 业务处理函数
func (s *Service3) PostHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("Service3-post-Handle")
	return
}

//NewService 服务对象构建函数
var NewService = func() (Service, error) {
	return Service{}, nil
}

//ServiceFunc
var ServiceFunc = func(ctx context.IContext) (r interface{}) {
	ctx.Log().Info("func-Handle")
	return
}

type Service4 struct{}

//ApiHandle 业务处理函数
func (s *Service4) ApiHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("Service4-post-Handle")
	return
}
