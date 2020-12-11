package main

import (
	"github.com/micro-plat/hydra"

	"github.com/micro-plat/hydra/hydra/servers/http"
)

//服务器各种返回结果
func main() {
	app := hydra.NewApp(
		hydra.WithPlatName("hydratest"),
		hydra.WithSystemName("apiserver"),
		hydra.WithClusterName("proxy-b"),
		hydra.WithRegistry("fs://./"),
		hydra.WithServerTypes(http.API),
	)

	app.API("/request", request)
	app.Start()
}
func request(ctx hydra.IContext) interface{} {
	return "success"
}
