/*
 * @Description:
 * @Autor: taoshouyin
 * @Date: 2021-09-26 09:54:02
 * @LastEditors: taoshouyin
 * @LastEditTime: 2021-09-27 10:32:11
 */
package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/cron"
	"github.com/micro-plat/hydra/conf/server/mqc"
	"github.com/micro-plat/hydra/conf/vars/queue/queueredis"
	"github.com/micro-plat/hydra/conf/vars/redis"
	cronx "github.com/micro-plat/hydra/hydra/servers/cron"
	"github.com/micro-plat/hydra/hydra/servers/http"
	mqcx "github.com/micro-plat/hydra/hydra/servers/mqc"
	"github.com/micro-plat/hydra/hydra/servers/rpc"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API, rpc.RPC, cronx.CRON, mqcx.MQC),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("apiserver_dbr"),
	hydra.WithClusterName("test"),
	hydra.WithRegistry("mysql://hbsv2x_dev:123456dev@hbsv2x_dev#192.168.0.36"),
)

func init() {

	redisAddr := "192.168.5.79:6379"
	hydra.Conf.API("28080")
	hydra.Conf.RPC("28081")
	hydra.Conf.Vars().Redis("5.79", redisAddr, redis.WithPoolSize(100))
	hydra.Conf.Vars().Queue().Redis("queue", "", queueredis.WithConfigName("5.79"))
	hydra.Conf.MQC(mqc.WithRedis("queue"), mqc.WithSharding(1))
	hydra.Conf.CRON(cron.WithSharding(1))

	hydra.S.CRON("/cron", func(ctx hydra.IContext) (r interface{}) {
		ctx.Log().Info("-------------db测试----------------")
		return
	}, "@every 10s")
	hydra.S.MQC("/mqc", mqcRecve, "apiserverdbr:redis:queue1")
	hydra.S.RPC("/rpc", rpcServer)

	hydra.S.API("/reg/reqrpc", reqRPC)
	hydra.S.API("/reg/create", create)
	hydra.S.API("/reg/update", update)
	hydra.S.API("/reg/delete", delete)
	hydra.S.API("/reg/exists", exists)
	hydra.S.API("/reg/getvalue", getvalue)
	hydra.S.API("/reg/getchildren", getchildren)

	hydra.S.API("/redis/opts", dbropts)

}

//编译并进行服务配置安装 ./apiserver_registry_redis conf install
//查看redis中key为hydratest:registry_redis:api:t:conf的值,以及conf下router和中间件的配置的值
//启动服务./apiserver_registry_redis run
//查看redis中key为hydratest:registry_redis:api:t:servers的值,以及conf下router和中间件的配置的值
//查看redis中key为hydratest:services:api:providers下的值

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
