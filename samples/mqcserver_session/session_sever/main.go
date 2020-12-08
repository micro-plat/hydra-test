package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/queue"
	"github.com/micro-plat/hydra/conf/vars/queue/queueredis"
	"github.com/micro-plat/hydra/conf/vars/redis"
	"github.com/micro-plat/hydra/hydra/servers/mqc"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(mqc.MQC),
	hydra.WithPlatName("hydra_test"),
	hydra.WithSystemName("session_server"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

func init() {
	hydra.Conf.Vars().Redis("5.79", redis.New([]string{"192.168.5.79:6379"}))
	hydra.Conf.Vars().Queue().Redis("xxx", queueredis.New(queueredis.WithConfigName("5.79")))
	hydra.Conf.MQC("redis://xxx").Queue(queue.NewQueue("mqc_session_t", "/mqc"))
	app.MQC("/mqc", func(ctx hydra.IContext) (r interface{}) {
		ctx.Log().Info("mqc-header:", ctx.Request().Headers())
		m, err := ctx.Request().GetMap()
		if err != nil {
			return err
		}
		ctx.Log().Info("mqc-map:", m)
		return
	})
}

func main() {
	app.Start()

}
