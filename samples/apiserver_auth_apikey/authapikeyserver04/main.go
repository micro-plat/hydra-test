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
	hydra.WithRegistry("zk://192.168.0.101"),
)

func init() {
	hydra.Conf.API(":8071").APIKEY("123456", apikey.WithSHA256Mode())
	app.API("/hydratest/apiserver/apikey", funcAPIKey)
}

//apiserver_apikey 中间件启用，sha256签名demo

//1.1  sudo ./authapikeyserver04 conf install -cover
//1.2 使用 ./authapikeyserver04 run
//1.5 签名串错误请求：http://localhost:8071/hydratest/apiserver/apikey?sign=ddffddffddff&timestamp=1925125121  返回 403/签名错误
//1.9 get请求编码：http://localhost:8071/hydratest/apiserver/apikey?param1=test&param2=%E4%B8%AD%E6%96%87%E6%95%B0%E6%8D%AE!%40%23%24%25%5E%26*()&timestamp=1925125121&sign=a80110bf04791371a3b18ce59030425a0e8218edddd82aedf4670dc09a851e8a 返回 200/success
func main() {
	app.Start()
}

var funcAPIKey = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver_apikey 中间件启用，sha256签名demo")
	return "success"
}
