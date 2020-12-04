package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/vars/queue/queueredis"
	"github.com/micro-plat/hydra/conf/vars/redis"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/hydra/servers/http"
	"github.com/micro-plat/hydra/hydra/servers/mqc"
	"github.com/micro-plat/hydra/registry"
	"github.com/micro-plat/lib4go/logger"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(mqc.MQC, http.API),
	hydra.WithPlatName("hydra_test"),
	hydra.WithSystemName("mqcserver_exp"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

var confPath = "/hydra_test/mqcserver_exp/mqc/t/conf"
var reg, _ = registry.NewRegistry("lm://./", logger.New("hydra"))

func init() {
	hydra.Conf.API(":8070")
	hydra.Conf.Vars().Redis("5.79", redis.New([]string{"192.168.5.79:6379"}))
	hydra.Conf.Vars().Queue().Redis("xxx", queueredis.New(queueredis.WithConfigName("5.79")))
	hydra.Conf.MQC("redis://xxx")

	app.MQC("/mqc", func(ctx context.IContext) (r interface{}) { return }, "mqcName")

	app.API("/mqc/start", func(ctx context.IContext) (r interface{}) {
		return reg.Update(confPath, `{"status":"start","addr":"redis://xxx"}`)
	})

	app.API("/mqc/stop", func(ctx context.IContext) (r interface{}) {
		return reg.Update(confPath, `{"status":"stop","addr":"redis://xxx"}`)
	})

}

//消息队列服务器异常关闭后正常启动，服务是否自动恢复
//启动服务  ./mqcserver_exp run
//反复调用  /mqc/start和mqc/stop 查看服务器在重新启动和禁用相互切换时是否正常
func main() {
	app.Start()
}
