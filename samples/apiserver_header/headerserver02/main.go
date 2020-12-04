package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/header"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithDebug(),
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("apiserverheader"),
	hydra.WithClusterName("taosytest"),
	hydra.WithRegistry("zk://192.168.0.101"),
)

func init() {
	hydra.Conf.API(":8070").Header(header.WithCrossDomain())
	app.API("/hydratest/apiserver/header", funcHeader)
}

// apiserver-header中间件设置默认跨域配置demo

//1.1 安装程序 ./headerserver02 conf install -cover
//1.2 使用 ./headerserver02 run
//1.3 请求不设置Origin头：http://localhost:8070/hydratest/apiserver/header  返回：配置header中所有的非空header列表
//1.4 请求设置Origin头：http://localhost:8070/hydratest/apiserver/header  返回：配置header中所有的非空header列表
func main() {
	app.Start()
}

var funcHeader = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver-header中间件设置默认跨域配置demo")
	headers := ctx.Response().GetHeaders()
	ctx.Log().Info("HeaderMap:", headers)
	ctx.Log().Info("GetSpecials:", ctx.Response().GetSpecials())
	return "success"
}
