package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/task"
	"github.com/micro-plat/hydra/hydra/servers/cron"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(cron.CRON, http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("task_repeat"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

func init() {
	hydra.Conf.API(":8070")
	hydra.Conf.CRON().Task(task.NewTask("@every 20s", "/cron"), task.NewTask("@every 20s", "/cron"))

	app.CRON("/cron", func(ctx hydra.IContext) (r interface{}) { return })
	app.CRON("/cron2", func(ctx hydra.IContext) (r interface{}) { return }, "@every 20s")

	app.API("/getcron", getCron)
}

//测试cron_task静态注册,重复注册
//启动服务  ./cronserver_task run
//访问 /getcron  打印任务配置[3个]  [{@every 20s /cron},{@every 20s /cron},{@every 20s /cron2}]
//查看 /cron和cron2 执行频率[每20s执行一次],静态注册的重复对执行没有影响
func main() {
	app.Start()
}
