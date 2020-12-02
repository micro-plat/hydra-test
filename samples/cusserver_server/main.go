package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/services"
)

var hydraApp = hydra.NewApp(
	hydra.WithServerTypes(CusServerName),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("cusserver_server"),
	hydra.WithClusterName("test"),
	hydra.WithRegistry("lm://."),
)

var CusRouters *services.ORouter

func init() {

	CusRouters = services.NewORouter()

	hydra.S.RegisterServer(CusServerName, func(g *services.Unit, ext ...interface{}) error {
		return CusRouters.Add(g.Path, g.Service, g.Actions, ext...)
	})

	hydra.Conf.Custom(CusServerName, map[string]interface{}{
		"address": ":50018",
	})

	hydraApp.Custom(CusServerName, "/customer/server/api", func(ctx hydra.IContext) interface{} {
		return "custom.server.api"
	})

}

//启动服务:go run main.go run
//1. 请求 http://localhost:50018/customer/server/api 获取正常的响应[custom.server.api] ，状态码：200
func main() {
	hydraApp.Start()
}
