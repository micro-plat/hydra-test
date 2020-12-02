package main

import (
	"github.com/micro-plat/hydra"
	_ "github.com/micro-plat/hydra/components/caches/cache/redis"
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
	hydra.Conf.API(":8070").Header()
	app.API("/hydratest/apiserver/header", funcHeader)
}

// apiserver-header中间件设置空配置demo

//1.1 安装程序 sudo ./headerserver01 conf install -cover
//1.2 使用 ./headerserver01 run
//1.3 调用接口：http://localhost:8070/hydratest/apiserver/header 判定配置是否正确
func main() {
	app.Start()
}

var funcHeader = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver-header中间件设置空配置demo")
	headers := ctx.Response().GetHeaders()
	ctx.Log().Info("HeaderMap:", headers)
	return "success"
}
