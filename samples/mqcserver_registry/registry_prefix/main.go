package main

import (
	"fmt"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/vars/queue/queueredis"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/hydra/hydra/servers/mqc"
	"github.com/micro-plat/hydra/services"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(mqc.MQC),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("mqc_registry_prefix"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

func init() {
	global.MQConf.PlatNameAsPrefix(false) //设置不拼接平台名为前缀
	hydra.Conf.Vars().Redis("5.79", "192.168.5.79:6379")
	hydra.Conf.Vars().Queue().Redis("xxx", "", queueredis.WithConfigName("5.79"))
	hydra.Conf.MQC("redis://xxx")
	app.MQC("/mqc", func(ctx context.IContext) (r interface{}) {
		ctx.Log().Info("mqc")
		return
	}, "mqcName")
	app.MQC("/mqc2", func(ctx context.IContext) (r interface{}) {
		ctx.Log().Info("mqc")
		return
	}, "mqcName2")
}

//测试mqc服务注册时队列名不拼接平台名为前缀
//启动服务  ./registry_assign run
//查看的服务启动前的消息队列数量[mqcName,mqcName2]和启动后的消息队列信息
func main() {
	fmt.Println("服务启动前的队列")
	for k, v := range services.MQC.GetQueues().Queues {
		fmt.Println(k, v.Queue, v.Service, v.Disable, v.Concurrency)
	}
	hydra.OnReady(func() {
		fmt.Println("服务启动后的队列")
		for k, v := range services.MQC.GetQueues().Queues {
			fmt.Println(k, v.Queue, v.Service, v.Disable, v.Concurrency)
		}
	})
	app.Start()
}
