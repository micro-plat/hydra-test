package main

import (
	"fmt"
	"sync/atomic"

	"github.com/micro-plat/hydra"
	appconf "github.com/micro-plat/hydra/conf/app"
	"github.com/micro-plat/hydra/context"
)

var starting = func(c appconf.IAPPConf) error {
	atomic.AddInt64(&index, 1)
	fmt.Println("server-starting", index)
	return nil
}

var closing = func(c appconf.IAPPConf) error {
	atomic.AddInt64(&index, 1)
	fmt.Println("server-closing", index)
	return nil
}

var handleExecuting context.Handler = func(ctx hydra.IContext) interface{} {
	atomic.AddInt64(&index, 1)
	ctx.Log().Info("api-handleExectuting", index)
	return ""
}

var handleExecuted context.Handler = func(ctx hydra.IContext) interface{} {
	atomic.AddInt64(&index, 1)
	ctx.Log().Info("api-handleExecuted", index)
	return ""
}

//APIServices 测试服务
type APIServices struct{}

//Handling 业务预处理函数
func (s *APIServices) Handling(ctx hydra.IContext) (r interface{}) {
	atomic.AddInt64(&index, 1)
	ctx.Log().Info("api-Handling", index)
	return
}

//Handle 业务处理函数
func (s *APIServices) Handle(ctx hydra.IContext) (r interface{}) {
	atomic.AddInt64(&index, 1)
	ctx.Log().Info("api-Handle", index)
	return
}

//Handled 业务后处理函数
func (s *APIServices) Handled(ctx hydra.IContext) (r interface{}) {
	atomic.AddInt64(&index, 1)
	ctx.Log().Info("api-Handled", index)
	return
}

//QueryHandling 业务预处理函数
func (s *APIServices) QueryHandling(ctx hydra.IContext) (r interface{}) {
	atomic.AddInt64(&index, 1)
	ctx.Log().Info("api-query-Handling", index)
	return
}

//QueryHandle 业务处理函数
func (s *APIServices) QueryHandle(ctx hydra.IContext) (r interface{}) {
	atomic.AddInt64(&index, 1)
	ctx.Log().Info("api-query-Handle", index)
	return
}

//QueryHandled 业务后处理函数
func (s *APIServices) QueryHandled(ctx hydra.IContext) (r interface{}) {
	atomic.AddInt64(&index, 1)
	ctx.Log().Info("api-query-Handled", index)
	return
}

//Close 对象关闭函数
func (s *APIServices) Close() {
	atomic.AddInt64(&index, 1)
	fmt.Println("service-Close", index)
	return
}
