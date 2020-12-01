package main

import (
	"github.com/micro-plat/hydra"
	_ "github.com/micro-plat/hydra/components/caches/cache/redis"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydra_demo"),
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
	app.API("/api/test/hook", &APIServices{})        //注册服务
}

func main() {
	app.Start()
}
