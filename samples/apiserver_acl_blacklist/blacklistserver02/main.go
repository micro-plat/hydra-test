package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/acl/blacklist"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("apiserverblacklist"),
	hydra.WithClusterName("taosytest"),
	hydra.WithRegistry("zk://192.168.0.101"),
)

func init() {
	hydra.Conf.API(":8071").BlackList(blacklist.WithEnable(), blacklist.WithIP("192.168.5.107"))
	app.API("/hydratest/apiserver/blacklist", funcBlackList)
}

//apiserver 黑名单中间件配置启用并且在限制名单内demo

//1.1  ./blacklistserver02 conf install -cover
//1.2 使用 ./blacklistserver02 run
//1.3 调用接口：http://localhost:8071/hydratest/apiserver/blacklist 通过机器192.168.5.107访问 可正常返回403/黑名单限制[%s]不允许访问
//1.4 调用接口：http://localhost:8071/hydratest/apiserver/blacklist 通过机器192.168.5.94访问 可正常返回200/success
func main() {
	app.Start()
}

var funcBlackList = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver_blacklist 黑名单中间件配置启用并且在限制名单内demo")
	return "success"
}
