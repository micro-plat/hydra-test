package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/acl/whitelist"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("whitelist-enable"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

func init() {
	hydra.Conf.API("8070").WhiteList(whitelist.WithEnable(), whitelist.WithIPList(whitelist.NewIPList([]string{"/api"}, "192.168.5.115")))
	app.API("/api", func(ctx hydra.IContext) (r interface{}) { return "success" })
}

//中间件白名单启用状态配置测试
//启动服务 ./whitelist-enable run
//通过机器192.168.5.115 访问/api [返回200]
//通过机器192.168.4.171 访问/api [不允许访问服务[/api],返回403]
func main() {
	app.Start()
}
