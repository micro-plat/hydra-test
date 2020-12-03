package main

import (
	"github.com/micro-plat/hydra"
)

type Service struct{}

//Handle 业务处理函数
func (s *Service) Handle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("Service-Handle")
	return "handle"
}

//QueryHandle 业务处理函数
func (s *Service) QueryHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("Service-query-Handle")
	return "handle"
}

//GetHandle 业务处理函数
func (s *Service) GetHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("Service-get-Handle")
	return "handle"
}

//Fallback 业务处理降级函数
func (s *Service) Fallback(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("Service-FallBack")
	return "fallback"
}

//QueryFallback 业务处理降级函数
func (s *Service) QueryFallback(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("Service-query-FallBack")
	return "fallback"
}

//GetFallback 业务处理降级函数
func (s *Service) GetFallback(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("Service-get-FallBack")
	return "fallback"
}
