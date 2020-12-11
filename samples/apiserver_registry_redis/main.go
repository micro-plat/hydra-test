package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/vars/queue/queueredis"
	"github.com/micro-plat/hydra/conf/vars/redis"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/hydra/servers/cron"
	"github.com/micro-plat/hydra/hydra/servers/http"
	"github.com/micro-plat/hydra/hydra/servers/mqc"
	"github.com/micro-plat/hydra/hydra/servers/rpc"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API, mqc.MQC, cron.CRON, rpc.RPC),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("registry_redis"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("redis://192.168.5.79:6379"),
)

func init() {
	hydra.Conf.API(":8080")
	hydra.Conf.Vars().Redis("5.79", redis.New([]string{"192.168.5.79:6379"}))
	hydra.Conf.Vars().Queue().Redis("xxx", queueredis.New(queueredis.WithConfigName("5.79")))
	hydra.Conf.MQC("redis://xxx")

	app.CRON("/cron", func(ctx context.IContext) (r interface{}) { return })
	app.MQC("/mqc", create)
	app.RPC("/rpc", create)
	app.API("/reg/create", create)
	app.API("/reg/update", update)
	app.API("/reg/delete", delete)
	app.API("/reg/exists", exists)
	app.API("/reg/getvalue", getvalue)
	app.API("/reg/getchildren", getchildren)
}

//编译并进行服务配置安装 ./apiserver_registry_redis conf install
//查看redis中key为hydra_test:registry_redis:api:t:conf的值,以及conf下router和中间件的配置的值
//启动服务./apiserver_registry_redis run
//查看redis中key为hydra_test:registry_redis:api:t:servers的值,以及conf下router和中间件的配置的值
//查看redis中key为hydra_test:services:api:providers下的值

//多次调用/reg/create   [创建永久节点,临时节点,顺序节点的创建的值,顺序节点的顺序,过期时间正确]
//调用/reg/delete      [创建的的永久节点被正确删除]
//调用/reg/update      [更新api的conf节点,conf节点值被更新,服务器重启]
//调用/reg/exists      [返回节点的判断正确]
//调用/reg/getvalue    [返回conf节点的值正确]
//调用/reg/getchildren [返回conf节点的下的子节点正确]

//关闭服务
//进行配置覆盖 ./apiserver_registry_redis conf install -cover [配置还原]

func main() {
	app.Start()
}
