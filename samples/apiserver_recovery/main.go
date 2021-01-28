package main

import (
	"fmt"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/api"

	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("apiserver_recovery"),
	hydra.WithClusterName("test"),
	hydra.WithRegistry("lm://."),
)

func init() {
	hydra.Conf.API("50010", api.WithTimeout(10, 10))

	app.API("/api/panic", func(ctx hydra.IContext) (r interface{}) {
		panic(fmt.Errorf("主动抛出异常"))
	})

	app.API("/api/normal", func(ctx hydra.IContext) (r interface{}) {
		return "normal.success"
	})
}

//测试recovery组件，检查响应内容是否正确
// 启动server : go run main.go run
//1. 请求 http://localhost:50010/api/panic 得到panic响应，错误码：510
//2. 请求 http://localhost:50010/api/normal 正常的相应 success,错误码：200
func main() {
	app.Start()
}
