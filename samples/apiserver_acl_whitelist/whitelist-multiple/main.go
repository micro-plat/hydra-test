package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/acl/whitelist"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydra_test"),
	hydra.WithSystemName("whitelist-multiple"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

func init() {
	hydra.Conf.API(":8070").WhiteList(whitelist.WithEnable(), whitelist.WithIPList(whitelist.NewIPList([]string{"/api1", "api2"}, "192.168.5.115", "192.168.4.171")))
	app.API("/api1", func(ctx hydra.IContext) (r interface{}) { return "success" })
	app.API("/api2", func(ctx hydra.IContext) (r interface{}) { return "success" })
}

//中间件白名单禁用状态配置测试
//启动服务 ./whitelist-multiple run
//通过机器192.168.5.115 访问/api1 [返回200]
//通过机器192.168.5.115 访问/api2 [返回200]
//通过机器192.168.4.171 访问/api1 [返回200]
//通过机器192.168.4.171 访问/api2 [返回200]
//通过机器192.168.5.106 访问/api1 [返回403]
//通过机器192.168.5.106 访问/api2 [返回403]
func main() {
	app.Start()
}
