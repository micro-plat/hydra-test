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
	hydra.Conf.API("8072").BlackList(blacklist.WithEnable(), blacklist.WithIP("192.168.5.*", "192.168.*.107", "192.168.**"))
	app.API("/hydratest/apiserver/blacklist", funcBlackList)
}

//apiserver_blacklist 黑名单中间件配置启用*模糊匹配demo

//1.1 使用 ./blacklistserver03 run
//1.2 调用接口：http://localhost:8072/hydratest/apiserver/blacklist 通过机器192.168.5.107访问 可正常返回403/黑名单限制[%s]不允许访问
//1.3 调用接口：http://localhost:8072/hydratest/apiserver/blacklist 通过机器192.168.5.94访问 可正常返回403/黑名单限制[%s]不允许访问
func main() {
	app.Start()
}

var funcBlackList = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver_blacklist 黑名单中间件配置启用*模糊匹配demo")
	return "success"
}
