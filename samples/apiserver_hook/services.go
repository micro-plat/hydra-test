package main

import (
	"fmt"

	"github.com/micro-plat/hydra"
	appconf "github.com/micro-plat/hydra/conf/app"
	"github.com/micro-plat/hydra/context"
)

var starting = func(c appconf.IAPPConf) error {
	fmt.Printf("server-starting")
	return nil
}

var closing = func(c appconf.IAPPConf) error {
	fmt.Printf("server-closing")
	return nil
}

var handleExecuting context.Handler = func(ctx hydra.IContext) interface{} {
	ctx.Log().Info("api-handleExectuting")
	return ""
}

var handleExecuted context.Handler = func(ctx hydra.IContext) interface{} {
	ctx.Log().Info("api-handleExecuted")
	return ""
}

//APIServices 测试服务
type APIServices struct{}

//Handling 业务预处理函数
func (s *APIServices) Handling(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("api-Handling")
	return
}

//Handle 业务处理函数
func (s *APIServices) Handle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("api-Handle")
	return
}

//Handled 业务后处理函数
func (s *APIServices) Handled(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("api-Handled")
	return
}

//QueryHandling 业务预处理函数
func (s *APIServices) QueryHandling(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("api-query-Handling")
	return
}

//QueryHandle 业务处理函数
func (s *APIServices) QueryHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("api-query-Handle")
	return
}

//QueryHandled 业务后处理函数
func (s *APIServices) QueryHandled(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("api-query-Handled")
	return
}

//Close 对象关闭函数
func (s *APIServices) Close() {
	fmt.Printf("service-Close")
	return
}
