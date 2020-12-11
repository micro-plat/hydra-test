package main

import (
	"fmt"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/task"
	"github.com/micro-plat/hydra/hydra/servers/cron"
	"github.com/micro-plat/hydra/hydra/servers/http"
	"github.com/micro-plat/hydra/services"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(cron.CRON, http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("task_conf"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

var printTasks = func() {
	for k, v := range services.CRON.GetTasks().Tasks {
		fmt.Println(k, v.Cron, v.Service, v.Disable)
	}
}

func init() {
	hydra.Conf.API(":8070")
	hydra.Conf.CRON().Task(task.NewTask("@every 5s", "/cron"), task.NewTask("@every 10s", "/cron2"))

	app.CRON("/cron", func(ctx hydra.IContext) (r interface{}) { return })
	app.CRON("/cron2", func(ctx hydra.IContext) (r interface{}) { return })

	app.API("/cron/add", addCron1)
	app.API("/cron/delete", deleteCron1)
	app.API("/cron/update", updateCron2)
}

//通过task静态注册，并在线增加、删除、修改注册配置
//启动服务 ./task_conf run
//1 访问 /cron/add    还原配置,服务重启     查看执行频率[/cron,/cron2每10s一次]
//2 访问 /cron/delete 删除/cron,服务重启   查看执行频率[/cron不执行,/cron2每10s一次]
//3 访问 /cron/add    添加/cron,服务重启   查看执行频率[/cron,/cron2每10s一次]
//4 访问 /cron/update 修改/cron2,服务重启   查看执行频率[/cron每10s一次,/cron2每1分钟一次]
//顺序执行1-4 重复10次
func main() {
	app.Start()
}
