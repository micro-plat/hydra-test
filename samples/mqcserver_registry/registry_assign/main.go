package main

import (
	"fmt"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/vars/queue/queueredis"
	"github.com/micro-plat/hydra/conf/vars/redis"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/hydra/servers/mqc"
	"github.com/micro-plat/hydra/services"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(mqc.MQC),
	hydra.WithPlatName("hydra_test"),
	hydra.WithSystemName("mqc_registry_assign"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

func init() {
	hydra.Conf.Vars().Redis("5.79", redis.New([]string{"192.168.5.79:6379"}))
	hydra.Conf.Vars().Queue().Redis("xxx", queueredis.New(queueredis.WithConfigName("5.79")))
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

//测试mqc服务注册时指定队列名
//启动服务  ./registry_assign run
//查看的服务启动前的消息队列数量[0]和启动后的消息队列信息
func main() {
	fmt.Println("服务启动前的队列:", services.MQC.GetQueues().Queues)
	hydra.OnReady(func() {
		fmt.Println("服务启动后的队列")
		for k, v := range services.MQC.GetQueues().Queues {
			fmt.Println(k, v.Queue, v.Service, v.Disable, v.Concurrency)
		}
	})
	app.Start()
}
