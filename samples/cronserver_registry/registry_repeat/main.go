package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/hydra/servers/cron"
	"github.com/micro-plat/hydra/hydra/servers/http"
	"github.com/micro-plat/hydra/services"
	"github.com/micro-plat/lib4go/logger"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(cron.CRON, http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("cron_registry_repeat"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

var cronName = "@every 20s"
var cronService = "/cron"
var printTasks = func(log logger.ILogger) {
	for k, v := range services.CRON.GetTasks().Tasks {
		log.Debug(k, v.Cron, v.Service, v.Disable)
	}
}

func init() {
	hydra.Conf.API(":8070")
	hydra.Conf.CRON()

	app.CRON("/cron", func(ctx hydra.IContext) (r interface{}) { return })

	app.API("/cron/add", func(ctx hydra.IContext) (r interface{}) {
		hydra.CRON.Add(cronName, cronService)
		printTasks(ctx.Log())
		return
	})

	app.API("/cron/remove", func(ctx hydra.IContext) (r interface{}) {
		hydra.CRON.Remove(cronName, cronService)
		printTasks(ctx.Log())
		return
	})

}

//测试动态注册与取消（同一任务多次注册与取消)
//启动服务  ./registry_repeat run
//查看的服务启动前的任务数量[0]和启动后的任务信息
//反复请求 /cron/add和/cron/remove 查看任务数量和信息
//查看/cron 执行频率[每20s一次],重复动态注册对执行没有影响
func main() {
	hydra.CRON.Add(cronName, cronService)
	hydra.CRON.Remove(cronName, cronService)
	app.Start()
}
