package main

import (
	"github.com/micro-plat/hydra"
	_ "github.com/micro-plat/hydra/components/caches/cache/redis"
	"github.com/micro-plat/hydra/conf/server/acl/blacklist"
	"github.com/micro-plat/hydra/context"
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
	hydra.Conf.API(":8071").BlackList(blacklist.WithEnable(), blacklist.WithIP("192.*.5.111", "192.168.5.107"))
	app.API("/hydratest/apiserver/blacklist", funcBlackList)
}

//apiserver 黑名单中间件配置启用混合匹配demo

//1.1  sudo ./blacklistserver04 conf install -cover
//1.2 使用 ./blacklistserver04 run
//1.3 调用接口：http://192.168.5.94:8071/hydratest/apiserver/blacklist 通过机器192.168.5.107访问 可正常返回403/黑名单限制[%s]不允许访问
//1.4 调用接口：http://192.168.5.94:8071/hydratest/apiserver/blacklist 通过机器192.168.5.94访问 可正常返回200/success
func main() {
	app.Start()
}

var funcBlackList func(ctx context.IContext) (r interface{}) = func(ctx context.IContext) (r interface{}) {
	ctx.Log().Info("apiserver_blacklist 黑名单中间件配置启用混合匹配demo")
	return "success"
}
