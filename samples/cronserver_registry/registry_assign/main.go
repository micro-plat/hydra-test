package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/hydra/servers/cron"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(cron.CRON),
	hydra.WithPlatName("hydra_test"),
	hydra.WithSystemName("cron_registry_assign"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

func init() {
	hydra.Conf.API(":8070")
	hydra.Conf.CRON()

	app.CRON("/cron", func(ctx hydra.IContext) (r interface{}) { return }, "@every 20s")
	app.CRON("/cron2", func(ctx hydra.IContext) (r interface{}) { return }, "@every 10s")
}

//服务注册时指定执行周期
//启动服务  ./registry_assign run
//查看 /cron[10s一次] 和 /cron2[20s一次] 的执行周期
func main() {
	app.Start()
}
