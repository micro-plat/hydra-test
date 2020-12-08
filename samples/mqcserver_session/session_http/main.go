package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/hydra/conf/vars/queue/queueredis"
	"github.com/micro-plat/hydra/conf/vars/redis"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydra_test"),
	hydra.WithSystemName("session_http"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

func init() {
	hydra.Conf.API(":8070")
	hydra.Conf.Vars().Redis("5.79", redis.New([]string{"192.168.5.79:6379"}))
	hydra.Conf.Vars().Queue().Redis("xxx", queueredis.New(queueredis.WithConfigName("5.79")))

	app.API("/mqc/http", func(ctx hydra.IContext) (r interface{}) {
		c := components.Def.Queue().GetRegularQueue("xxx")
		ctx.Log().Info("request.sessiom_id:", ctx.Log().GetSessionID())
		err := c.Send("mqc_session_t", `{"key":"value"}`)
		if err != nil {
			return
		}
		return
	})

}

//通过rpc,api内部调用，检查session_id是否正确的传到当前mqc服务
//启动mqc处理服务 ../session_server run
//启动服务./session_http run
//访问  /mqc/http 查看请求的session_id与mqc处理服务收到session_id一致,且处理服务打印的数据正确 [返回200]
func main() {
	app.Start()
}
