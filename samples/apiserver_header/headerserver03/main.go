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
	hydra.Conf.API("8072").Header(header.WithCrossDomain("http://192.168.5.94:8070", "http://localhost:8070"), header.WithAllowMethods("get", "post"), header.WithHeader("taosy-header", "test"))
	app.API("/hydratest/apiserver/header", funcHeader)
}

// apiserver-header中间件设置指定跨域配置demo

//1.1 使用 ./headerserver03 run
//1.2 请求不设置Origin头：http://localhost:8072/hydratest/apiserver/header 返回：只返回非allow的头信息
//1.3 请求设置Origin头,不在配置内：http://localhost:8072/hydratest/apiserver/header 返回：只返回非allow的头信息
//1.4 请求设置Origin头,在配置内：http://localhost:8072/hydratest/apiserver/header 返回：返回所有的非空头配置
func main() {
	app.Start()
}

var funcHeader = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver-header中间件设置指定跨域配置demo")
	headers := ctx.Response().GetHeaders()
	ctx.Log().Info("HeaderMap:", headers)
	ctx.Log().Info("GetSpecials:", ctx.Response().GetSpecials())
	return "success"
}
