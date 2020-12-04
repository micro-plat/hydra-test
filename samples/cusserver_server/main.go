package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro-plat/hydra"
)

var hydraApp = hydra.NewApp(
	hydra.WithServerTypes(CusServerName),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("cusserver_server"),
	hydra.WithClusterName("test"),
	hydra.WithRegistry("lm://."),
)

func init() {

	hydra.Conf.Custom(CusServerName, map[string]interface{}{
		"address": ":50018",
	}).Sub("router", &RouterList{
		List: []*Router{
			&Router{
				Service: "/customer/server/api",
			},
		},
	})

	Registry("/customer/server/api", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]interface{}{
			"data": "custom.server.api",
		})
	})
}

//启动服务:go run main.go run
//1. 请求 http://localhost:50018/customer/server/api 获取正常的响应[{"data":"custom.server.api"}] ，状态码：200
func main() {
	hydraApp.Start()
}
