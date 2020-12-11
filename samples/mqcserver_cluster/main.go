package main

import (
	"github.com/micro-plat/hydra"
	mqcconf "github.com/micro-plat/hydra/conf/server/mqc"
	"github.com/micro-plat/hydra/conf/vars/queue/queueredis"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/hydra/servers/http"
	"github.com/micro-plat/hydra/hydra/servers/mqc"
	"github.com/micro-plat/hydra/registry"
	"github.com/micro-plat/lib4go/logger"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(mqc.MQC, http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("mqc_cluster"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("zk://192.168.0.101"),
)
var confPath = "/hydra_test/mqc_cluster/mqc/t/conf"
var reg, _ = registry.GetRegistry("zk://192.168.0.101", logger.New("hydra"))

func init() {
	hydra.Conf.API(":8070")
	hydra.Conf.Vars().Redis("5.79", "192.168.5.79:6379")
	hydra.Conf.Vars().Queue().Redis("xxx", "", queueredis.WithConfigName("5.79"))
	hydra.Conf.MQC("redis://xxx", mqcconf.WithMasterSlave()) //设置为主从模式

	app.MQC("/mqc", func(ctx context.IContext) (r interface{}) { return }, "mqcName")

	app.API("/mqc/master", func(ctx context.IContext) (r interface{}) {
		return reg.Update(confPath, `{"status":"start","sharding":1,"addr":"redis://xxx"}`)
	})

	app.API("/mqc/sharding", func(ctx context.IContext) (r interface{}) {
		return reg.Update(confPath, `{"status":"start","sharding":2,"addr":"redis://xxx"}`)
	})

	app.API("/mqc/p2p", func(ctx context.IContext) (r interface{}) {
		return reg.Update(confPath, `{"status":"start","sharding":0,"addr":"redis://xxx"}`)
	})
}

//消息队列服务器异常关闭后正常启动，服务是否自动恢复
//启动服务 ./cluster_master run 启动3个mqc服务(仅一个服务开启api服务)
//反复调用/mqc/master,mqc/sharding,mqc/p2p 查看服务器在模式相互切换时是否正常
func main() {
	app.Start()
}
