package main

import (
	"fmt"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/queue"
	"github.com/micro-plat/hydra/conf/vars/queue/queueredis"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/hydra/servers/http"
	"github.com/micro-plat/hydra/hydra/servers/mqc"
	"github.com/micro-plat/hydra/services"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(mqc.MQC, http.API),
	hydra.WithPlatName("hydra_test"),
	hydra.WithSystemName("mqc_registry_again"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

var mqcName = "mqcName"
var mqcService = "/mqc/service"
var printQueues = func() {
	for k, v := range services.MQC.GetQueues().Queues {
		fmt.Println(k, v.Queue, v.Service, v.Disable, v.Concurrency)
	}
}

func init() {
	hydra.Conf.API(":8070")
	hydra.Conf.Vars().Redis("5.79", "192.168.5.79:6379")
	hydra.Conf.Vars().Queue().Redis("xxx", "", queueredis.WithConfigName("5.79"))
	hydra.Conf.MQC("redis://xxx").Queue(queue.NewQueue(mqcName, mqcService))

	app.MQC("/mqc", func(ctx context.IContext) (r interface{}) { return }, mqcName)

	app.API("/mqc/add", func(ctx hydra.IContext) (r interface{}) {
		hydra.MQC.Add(mqcName, mqcService) //已存在
		printQueues()
		return
	})
}

//测试mqc服务注册重复注册
//启动服务  ./registry_assign run
//查看的服务启动前的消息队列数量[0]和启动后的消息队列信息[hydra_test:mqcName /mqc]
//反复请求 /mqc/add 查看消息队列数量和信息[hydra_test:mqcName /mqc]
func main() {
	fmt.Println("服务启动前的队列:", services.MQC.GetQueues().Queues)
	hydra.OnReady(func() {
		fmt.Println("服务启动后的队列")
		printQueues()
	})
	app.Start()
}
