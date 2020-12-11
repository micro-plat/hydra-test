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
	hydra.WithRegistry("lm://."),
)

func init() {
	hydra.Conf.API(":8072").BlackList(blacklist.WithDisable(), blacklist.WithIP("192.168.5.107"))
	app.API("/hydratest/apiserver/blacklist", funcBlackList)
}

//apiserver_blacklist 黑名单中间件配置被禁用demo

//1.1 使用 ./blacklistserver01 run
//1.2 调用接口：http://localhost:8072/hydratest/apiserver/blacklist 通过机器192.168.5.107访问 可正常返回200/success
func main() {
	app.Start()
}

var funcBlackList = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver_blacklist 黑名单中间件配置被禁用demo")
	return "success"
}
