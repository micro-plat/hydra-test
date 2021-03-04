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
	hydra.Conf.API("8072").APIKEY("12345678", apikey.WithSHA1Mode())
	app.API("/hydratest/apiserver/apikey", funcAPIKey)
}

//apiserver_apikey 中间件启用，sha1签名demo

//1.1 使用 ./authapikeyserver03 run
//1.2 签名串错误请求：http://localhost:8072/hydratest/apiserver/apikey?sign=ddffddffddff&timestamp=1925125121  返回 403/签名错误
//1.3 get请求编码：http://localhost:8072/hydratest/apiserver/apikey?param1=test&param2=%E4%B8%AD%E6%96%87%E6%95%B0%E6%8D%AE!%40%23%24%25%5E%26*()&timestamp=1925125121&sign=0d22d560d8db28e3ec84763b41c84ca6dd5540b4 返回 200/success
func main() {
	app.Start()
}

var funcAPIKey = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver_apikey 中间件启用，sha1签名demo")
	return "success"
}
