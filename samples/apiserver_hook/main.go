package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("hook"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("zk://192.168.0.101"),
)

func init() {
	hydra.Conf.API(":8070", api.WithTimeout(10, 10))
	app.OnStarting(starting, http.API)               //添加服务启动函数
	app.OnClosing(closing, http.API)                 //添加服务关闭函数
	app.OnHandleExecuting(handleExecuting, http.API) //添加业务预处理钩子
	app.OnHandleExecuted(handleExecuted, http.API)   //添加业务后处理钩子
	app.API("/api", &APIServices{})                  //注册服务
}

// 使用zookeeper作为注册中心，验证系统勾子函数、handle勾子函数、对象关闭函数是否正确执行
// go build
// ./apiserver_hook run -r "zk://192.168.0.101" -c t 打印server-starting
// 访问/api 顺序打印 api-handleExectuting,api-Handling,api-Handle,api-Handled,api-handleExecuted
// 访问/api/query 顺序打印 api-handleExectuting,api-Handling,api-query-Handling,api-query-Handle,api-query-Handled,api-Handled,api-query-handleExecuted
// 关闭服务 顺序打印  server-closing service-Close
func main() {
	app.Start()
}
