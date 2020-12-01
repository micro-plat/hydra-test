package main

import (
	"github.com/micro-plat/hydra/context"
)

type Service struct{}

//Handle 业务处理函数
func (s Service) Handle(ctx context.IContext) (r interface{}) {
	ctx.Log().Info("Service-Handle")
	return
}

type Service2 struct{}

//Handle 业务处理函数
func (s *Service2) Handle(ctx context.IContext) (r interface{}) {
	ctx.Log().Info("Service2-Handle")
	return
}

//QueryHandle 业务处理函数
func (s *Service2) QueryHandle(ctx context.IContext) (r interface{}) {
	ctx.Log().Info("Service2-query-Handle")
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
