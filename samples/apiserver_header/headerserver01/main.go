package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("apiserverheader"),
	hydra.WithClusterName("taosytest"),
	hydra.WithRegistry("lm://."),
)

func init() {
	hydra.Conf.API("8072").Header()
	app.API("/hydratest/apiserver/header", funcHeader)
}

// apiserver-header中间件设置空配置demo

//1.1 使用 ./headerserver01 run
//1.2 调用接口：http://localhost:8072/hydratest/apiserver/header 判定配置是否正确
func main() {
	app.Start()
}

var funcHeader = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver-header中间件设置空配置demo")
	headers := ctx.Response().GetHeaders()
	ctx.Log().Info("HeaderMap:", headers)
	ctx.Log().Info("GetSpecials:", ctx.Response().GetSpecials())
	return "success"
}
