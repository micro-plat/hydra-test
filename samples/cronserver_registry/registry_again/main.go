package main

import (
	"fmt"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/task"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/hydra/servers/cron"
	"github.com/micro-plat/hydra/hydra/servers/http"
	"github.com/micro-plat/hydra/services"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(cron.CRON, http.API),
	hydra.WithPlatName("hydra_test"),
	hydra.WithSystemName("_registry_again"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

var cronName = "@every 20s"
var cronService = "/cron"
var printTasks = func() {
	for k, v := range services.CRON.GetTasks().Tasks {
		fmt.Println(k, v.Cron, v.Service, v.Disable)
	}
}

func init() {
	hydra.Conf.API(":8070")
	hydra.Conf.CRON().Task(task.NewTask(cronName, cronService))

	app.CRON("/cron", func(ctx context.IContext) (r interface{}) { return }, cronName)

	app.API("/cron/add", func(ctx hydra.IContext) (r interface{}) {
		hydra.CRON.Add(cronName, cronService) //已存在
		printTasks()
		return
	})
}

//测试cron服务注册重复注册
//启动服务  ./registry_again run
//查看的服务启动前的任务数量[0个]和服务启动后的任务[1个]
//反复请求 /cron/add 查看任务数量[1个]和信息
func main() {
	fmt.Println("服务启动前的任务:", services.CRON.GetTasks().Tasks)
	hydra.OnReady(func() {
		fmt.Println("服务启动后的任务")
		printTasks()
	})
	app.Start()
}
