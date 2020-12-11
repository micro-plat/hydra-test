package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/hydra/conf/vars/queue/queueredis"
	"github.com/micro-plat/hydra/hydra/servers/http"
	"github.com/micro-plat/hydra/hydra/servers/rpc"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API, rpc.RPC),
	hydra.WithPlatName("hydra_test"),
	hydra.WithSystemName("session_rpc"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

func init() {
	hydra.Conf.API(":8070")
	hydra.Conf.Vars().Redis("5.79", "192.168.5.79:6379")
	hydra.Conf.Vars().Queue().Redis("xxx", "", queueredis.WithConfigName("5.79"))
	hydra.Conf.Vars().RPC("rpc")

	app.RPC("/rpc/proc", func(ctx hydra.IContext) (r interface{}) {
		ctx.Log().Info("rpc-session:", ctx.Log().GetSessionID())
		c := components.Def.Queue().GetRegularQueue("xxx")
		err := c.Send("mqc_session_t", ctx.Request().GetString("data"))
		if err != nil {
			return
		}
		return
	})

	app.API("/mqc/rpc", func(ctx hydra.IContext) (r interface{}) {
		respones, err := hydra.C.RPC().GetRegularRPC().Request("/rpc/proc@hydra_test", map[string]interface{}{
			"data": `{"key":"value"}`,
		})
		if err != nil {
			ctx.Log().Errorf("rpc 请求异常：%v", err)
			return
		}
		ctx.Log().Info("respones.Status:", respones.Status)
		return respones.Result
	})

}

//通过rpc,api内部调用，检查session_id是否正确的传到当前mqc服务
//启动mqc处理服务 ../session_server run
//启动服务./session_rpc run
//访问  /mqc/rpc  模拟rpc请求
//查看rpc请求的session_id与mqc处理服务收到session_id一致,且处理服务打印的数据正确 [返回200]
func main() {
	app.Start()
}
