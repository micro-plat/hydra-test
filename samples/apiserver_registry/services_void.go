package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/context"
)

type VoidService struct{}

//Handle 业务处理函数
func (s VoidService) Handle(ctx hydra.IContext) {
	ctx.Log().Info("VoidService-Handle:")
	ctx.Log().Info("VoidService-Encoding:", ctx.Request().Path().GetEncoding())
	router, err := ctx.Request().Path().GetRouter()
	ctx.Log().Info("VoidService-Handle-router:", router, err)
	return
}

type VoidService2 struct{}

//Handle 业务处理函数
func (s *VoidService2) Handle(ctx hydra.IContext) {
	ctx.Log().Info("VoidService2-Handle")
	ctx.Log().Info("VoidService2-Encoding:", ctx.Request().Path().GetEncoding())
	router, err := ctx.Request().Path().GetRouter()
	ctx.Log().Info("VoidService2-Handle-router:", router, err)
	return
}

//QueryHandle 业务处理函数
func (s *VoidService2) QueryHandle(ctx hydra.IContext) {
	ctx.Log().Info("VoidService2-query-Handle")
	ctx.Log().Info("VoidService2-Encoding:", ctx.Request().Path().GetEncoding())
	router, err := ctx.Request().Path().GetRouter()
	ctx.Log().Info("VoidService2-query-router:", router, err)
	return
}

//Fallback 业务处理降级函数
func (s *VoidService2) Fallback(ctx hydra.IContext) {
	ctx.Log().Info("VoidService2-FallBack")
	router, err := ctx.Request().Path().GetRouter()
	ctx.Log().Info("VoidService2-router:", router, err)
	return
}

//QueryFallback 业务处理降级函数
func (s *VoidService2) QueryFallback(ctx hydra.IContext) {
	ctx.Log().Info("VoidService2-query-FallBack")
	router, err := ctx.Request().Path().GetRouter()
	ctx.Log().Info("VoidService2-fallback-router:", router, err)
	return
}

type VoidService3 struct{}

//PostHandle 业务处理函数
func (s *VoidService3) PostHandle(ctx hydra.IContext) {
	ctx.Log().Info("VoidService3-post-Handle")
	return
}

//NewVoidService 服务对象构建函数
var NewVoidService = func() (VoidService, error) {
	return VoidService{}, nil
}

//VoidServiceFunc
var VoidServiceFunc = func(ctx context.IContext) {
	ctx.Log().Info("void_func-Handle")
	return
}

type VoidService4 struct{}

//ApiHandle 业务处理函数
func (s *VoidService4) ApiHandle(ctx hydra.IContext) {
	ctx.Log().Info("VoidService4-post-Handle")
	return
}
