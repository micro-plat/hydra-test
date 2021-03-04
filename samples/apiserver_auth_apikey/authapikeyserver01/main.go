package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/auth/apikey"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("apiserverapikey"),
	hydra.WithClusterName("taosytest"),
	hydra.WithRegistry("lm://."),
)

func init() {
	hydra.Conf.API("8072").APIKEY("12345678", apikey.WithDisable(), apikey.WithMD5Mode())
	app.API("/hydratest/apiserver/apikey", funcAPIKey)
}

//apiserver_apikey 中间件被禁用demo

//1.1 使用 ./authapikeyserver01 run
//1.2 不签名请求：http://localhost:8072/hydratest/apiserver/apikey  返回 200/success
//1.3 随意签名请求：http://localhost:8072/hydratest/apiserver/apikey?sign=34fvfefg45sdf  返回 200/success
func main() {
	app.Start()
}

var funcAPIKey = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver_apikey 中间件被禁用demo")
	return "success"
}
