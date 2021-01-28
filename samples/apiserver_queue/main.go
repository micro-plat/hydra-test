package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/hydra/conf/vars/cache/cacheredis"
	"github.com/micro-plat/hydra/conf/vars/queue/mqtt"
	"github.com/micro-plat/hydra/conf/vars/queue/queueredis"
	"github.com/micro-plat/hydra/conf/vars/redis"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/hydra/servers/http"
	"github.com/micro-plat/hydra/hydra/servers/mqc"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(mqc.MQC, http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("apiserver_queue"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("zk://192.168.0.101"),
)

func init() {
	hydra.Conf.API("8070")

	hydra.Conf.Vars().Redis("5.79", "192.168.5.79:6379", redis.WithPoolSize(10))
	hydra.Conf.Vars().Queue().MQTT("mqtt", "192.168.0.219:8883", mqtt.WithDialTimeout(500), mqtt.WithUP("mqtt", "abc123$"), mqtt.WithCert("./ca.pem"))
	hydra.Conf.Vars().Queue().Redis("xxx", "", queueredis.WithConfigName("5.79"))
	hydra.Conf.Vars().Cache().Redis("xxx", "", cacheredis.WithConfigName("5.79"))
	hydra.Conf.MQC("redis://xxx")

	app.MQC("/mqc", func(ctx context.IContext) (r interface{}) {
		m := ctx.Request().GetMap()
		ctx.Log().Info("mqc_message:", m)
		return
	}, "mqc_apiserver")

	app.API("/add/redisqueue", func(ctx hydra.IContext) (r interface{}) {
		c := components.Def.Queue().GetRegularQueue("xxx")
		c.Send("mqc_apiserver", `{"key":"value"}`)
		return
	})
	app.API("/add/mqttqueue", func(ctx hydra.IContext) (r interface{}) {
		c := components.Def.Queue().GetRegularQueue("mqtt")
		c.Send("mqc_apiserver", `{"key":"value"}`)
		return
	})

	app.API("/test/redisqueue", func(ctx hydra.IContext) (r interface{}) {
		c := components.Def.Queue().GetRegularQueue("mqtt")
		err := c.Send("mqc_apiserver_test", `{"key":"value"}`)
		if err != nil {
			return err
		}

		return nil
	})

}

//queue组件是否正确工作,修改配置是否自动生效(redis,mqtt)
//启动服务  ./apiserver_queue run
//1. 访问  /add/redisqueue 添加消息队列, /mqc主动的消费消息队列,打印Message: map[key:value]
//2. 访问  /test/redisqueue 查看组件各个功能是否正常
//3. 修改zk节点 /hydratest/apiserver_queue/mqc-api/t/conf 的addr值为"mqtt://mqtt"
//4. 重启服务,连接mqtt
//5. 访问  /add/mqttqueue 添加消息队列, /mqc主动的消费消息队列,打印Message: map[key:value]
func main() {
	app.Start()
}
