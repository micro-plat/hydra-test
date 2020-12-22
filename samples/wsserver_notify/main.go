package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/vars/queue/queueredis"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/hydra/servers/http"
	"github.com/micro-plat/hydra/hydra/servers/http/ws"
	"github.com/micro-plat/hydra/hydra/servers/mqc"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.WS, http.API, mqc.MQC),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("wsserver_notify"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

func init() {
	hydra.Conf.API(":8070")
	hydra.Conf.WS(":8080")
	hydra.Conf.Vars().Redis("5.79", "192.168.5.79:6379")
	hydra.Conf.Vars().Queue().Redis("queue", "", queueredis.WithConfigName("5.79"))
	hydra.Conf.MQC("redis://queue")
	app.API("/notify", func(ctx context.IContext) (r interface{}) {
		ctx.Log().Info("notify-Handle")
		reqID := ctx.Request().GetString("req_id")
		ctx.Log().Info("notify-Handle-reqID:", reqID)
		err := ws.DataExchange.Notify(reqID, map[string]interface{}{
			"msg": "success",
		})
		if err != nil {
			return err
		}
		return "success"
	})
	app.WS("/id/get", func(ctx context.IContext) (r interface{}) {
		ctx.Log().Info("ws-Handle")
		reqID := ctx.User().GetRequestID()
		ctx.Log().Info("ws-Handle-reqID:", reqID)
		return reqID
	})
}

//通过ws.DataExchange.Notify发送通知消息 [需要启动mqc服务]
//启动服务./wsserver_notify run
//建立与服务的连接
//客户端给服务发送数据 {"service":"/id/get"} 返回200 获取到requestID
//通过http 访问/notify 传入参数req_id=requestID [notify调用成功,消息队列发送成功,消息队列消费成功,客户端收到通知消息,消息正确,状态200]
func main() {
	app.Start()
}
