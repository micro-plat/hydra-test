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
	hydra.Conf.API(":8071").APIKEY("123456", apikey.WithMD5Mode(), apikey.WithExcludes("/hydratest/apiserver/*", "/hydratest/apiserver1/apikey"))
	app.API("/hydratest/apiserver/apikey", funcAPIKey)
	app.API("/hydratest/apiserver1/apikey", funcAPIKey)
	app.API("/hydratest/apiserver1/apikey1", funcAPIKey)
}

//apiserver_apikey 中间件启用，md5路径被排除demo

//1.1  sudo ./authapikeyserver02 conf install -cover
//1.2 使用 ./authapikeyserver02 run
//1.3 模糊匹配被排除路径请求：http://localhost:8071/hydratest/apiserver/apikey  返回 200/success
//1.4 精确匹配被排除路径请求：http://localhost:8071/hydratest/apiserver1/apikey  返回 200/success
//1.5 签名串错误请求：http://localhost:8071/hydratest/apiserver1/apikey1?sign=ddffddffddff&timestamp=1925125121  返回 403/签名错误
//1.6 缺少必要参数timestamp请求：http://localhost:8071/hydratest/apiserver1/apikey1?sign=ddffddffddff  返回 406/timestamp值不能为空
//1.7 缺少必要参数sign请求：http://localhost:8071/hydratest/apiserver1/apikey1?timestamp=1925125121  返回 401/sign值不能为空
//1.8 get请求不编码：http://localhost:8071/hydratest/apiserver1/apikey1?param1=test&param2=中文数据!@$%^&*()&timestamp=1925125121&sign=5905ec70b8f1ab903ff224afe8282d6e 返回  403/签名错误
//1.9 get请求编码：http://localhost:8071/hydratest/apiserver1/apikey1?param1=test&param2=%E4%B8%AD%E6%96%87%E6%95%B0%E6%8D%AE!%40%23%24%25%5E%26*()&timestamp=1925125121&sign=f6751e6ee103776a4550bd3445a2f258 返回 200/success
//1.10 post-from 不编码：http://localhost:8071/hydratest/apiserver1/apikey1  param1=test&param2=中文数据!@#$%^&*()&timestamp=1925125121&sign=f6751e6ee103776a4550bd3445a2f258 返回  403/签名错误
//1.11 post-from 编码：http://localhost:8071/hydratest/apiserver1/apikey1    param1=test&param2=%E4%B8%AD%E6%96%87%E6%95%B0%E6%8D%AE!%40%23%24%25%5E%26*()&timestamp=1925125121&sign=f6751e6ee103776a4550bd3445a2f258 返回 200/success
//1.12 post-fromdata-json：http://localhost:8071/hydratest/apiserver1/apikey1  {"param1":"test","param2":"中文数据!@#$%^&*()","timestamp":"1925125121","sign":"f6751e6ee103776a4550bd3445a2f258"} 返回 200/success
//1.13 post-body-json：http://localhost:8071/hydratest/apiserver1/apikey1      {"param1":"test","param2":"中文数据!@#$%^&*()","timestamp":"1925125121","sign":"f6751e6ee103776a4550bd3445a2f258"} 返回 200/success
func main() {
	app.Start()
}

var funcAPIKey = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver_apikey 中间件启用，md5路径被排除demo")
	return "success"
}
