package main

import (
	"time"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/vars/queue/queueredis"
	"github.com/micro-plat/hydra/conf/vars/redis"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/hydra/servers/mqc"
	"github.com/micro-plat/hydra/registry"
	"github.com/micro-plat/lib4go/logger"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(mqc.MQC),
	hydra.WithPlatName("hydra_test"),
	hydra.WithSystemName("mqc_registry_again"),
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
}

//消息队列服务器异常关闭后正常启动，服务是否自动恢复
//启动服务  ./mqcserver_exp run
//查询服务器在重新启动和禁用相互切换时是否正常
func main() {
	//go updateServerStatus()
	app.Start()
}

func updateServerStatus() {
	for k := 0; k < 10; k++ {
		time.Sleep(time.Second * 5)
		reg, _ := registry.NewRegistry("lm://./", logger.New("hydra"))
		path := "/hydra_test/mqc_registry_again/mqc/t/conf"
		reg.Update(path, `{"status":"s","addr":"redis://xxx"}`)
		time.Sleep(time.Second * 5)
		reg.Update(path, `{"status":"start","addr":"redis://xxx"}`)
	}

}
