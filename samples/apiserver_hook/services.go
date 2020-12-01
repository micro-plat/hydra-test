package main

import (
	"fmt"

	appconf "github.com/micro-plat/hydra/conf/app"
	"github.com/micro-plat/hydra/context"
)

var starting func(c appconf.IAPPConf) error = func(c appconf.IAPPConf) error {
	fmt.Printf("api-starting")
	return nil
}

var closing func(c appconf.IAPPConf) error = func(c appconf.IAPPConf) error {
	fmt.Printf("api-closing")
	return nil
}

var handleExecuting context.Handler = func(ctx context.IContext) interface{} {
	ctx.Log().Info("api-handleExectuting")
	return ""
}

var handleExecuted context.Handler = func(ctx context.IContext) interface{} {
	ctx.Log().Info("api-handleExecuted")
	return ""
}

//APIServices 测试服务
type APIServices struct{}

//Handling 业务预处理函数
func (s *APIServices) Handling(ctx context.IContext) (r interface{}) {
	ctx.Log().Info("api-Handling")
	return
}

//Handle 业务处理函数
func (s *APIServices) Handle(ctx context.IContext) (r interface{}) {
	ctx.Log().Info("api-Handle")
	return
}

//Handled 业务后处理函数
func (s *APIServices) Handled(ctx context.IContext) (r interface{}) {
	ctx.Log().Info("api-Handled")
	return
}
