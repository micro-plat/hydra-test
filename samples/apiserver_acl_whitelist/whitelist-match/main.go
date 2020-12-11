package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/acl/whitelist"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("whitelist-match"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

func init() {
	hydra.Conf.API(":8070").WhiteList(whitelist.WithEnable(), whitelist.WithIPList(whitelist.NewIPList([]string{"/api1/*/last", "/api2/**"}, "192.168.*.115", "192.168.4.*")))
	app.API("/api1/sub", func(ctx hydra.IContext) (r interface{}) { return "success" })
	app.API("/api1/sub/last", func(ctx hydra.IContext) (r interface{}) { return "success" })
	app.API("/api2/sub", func(ctx hydra.IContext) (r interface{}) { return "success" })
	app.API("/api2/sub/last", func(ctx hydra.IContext) (r interface{}) { return "success" })
}

//中间件白名单禁用状态配置测试
//启动服务 ./whitelist-match run
//通过机器192.168.5.115 访问/api1/sub,/api1/sub/last,/api2/sub,/api2/sub/last [返回200]
//通过机器192.168.4.171 访问/api1/sub,/api1/sub/last,/api2/sub,/api2/sub/last [返回200]
//通过机器192.168.5.106 访问/api1/sub [返回200]
//通过机器192.168.5.106 访问/api1/sub/last,api2/sub,api2/sub/last [返回403]
func main() {
	app.Start()
}
