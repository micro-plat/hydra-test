package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/auth/apikey"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.WS),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("ws_apikey"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

func init() {
	hydra.Conf.WS("8180").APIKEY("12345678", apikey.WithSHA1Mode())
	app.WS("/ws", func(ctx hydra.IContext) (r interface{}) {
		ctx.Log().Info("ws中间件启用")
		return "success"
	})
}

//测试ws的apikey中间件启用
//启动服务
//建立连接
//发送数据{"service":"/ws","sign":"ddffddffddff","timestamp":"1925125121"} 返回 403/签名错误
//发送数据{"service":"/ws","a":"1","b":"1","timestamp":"1925125121","sign":"6e0daed014cde48a3b19a9afc2a9089e63f4e06b"} 返回 200/success
func main() {
	app.Start()
}
