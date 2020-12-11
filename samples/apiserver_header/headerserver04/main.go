package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/header"
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
	hydra.Conf.API(":8072").Header(header.WithCrossDomain(), header.WithHeader("taosy-header", "testx"))
	app.API("/hydratest/apiserver/header", funcHeader)
}

// apiserver-header中间件头信息覆盖demo

//1.1 使用 ./headerserver04 run
//1.2 请求头：http://localhost:8072/hydratest/apiserver/header 返回：返回的taosy-header=test
func main() {
	app.Start()
}

var funcHeader = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver-header中间件头信息覆盖demo")
	ctx.Response().Header("taosy-header", "test")
	headers := ctx.Response().GetHeaders()
	ctx.Log().Info("HeaderMap:", headers)
	ctx.Log().Info("GetSpecials:", ctx.Response().GetSpecials())
	return "success"
}
