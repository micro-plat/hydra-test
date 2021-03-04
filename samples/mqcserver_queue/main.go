package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/hydra/conf/server/processor"
	"github.com/micro-plat/hydra/conf/server/queue"

	//"github.com/micro-plat/hydra/conf/vars/queue/queueredis"
	//"github.com/micro-plat/hydra/conf/vars/queue/lmq"
	"github.com/micro-plat/hydra/hydra/servers/http"
	"github.com/micro-plat/hydra/hydra/servers/mqc"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API, mqc.MQC),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("mqcserver"),
	hydra.WithClusterName("test"),
	//hydra.WithRegistry("zk://192.168.0.101"),
)

func init() {
	hydra.Conf.API("18072").Processor(processor.WithServicePrefix("api"))
	hydra.Conf.MQC("redis://queuename").Queue(queue.NewQueue("mqcserver:queue2", "/hydratest/mqcserver/queue2"))
	//hydra.Conf.Vars().Redis("redis", )
	hydra.Conf.Vars().Queue().Redis("queuename", "192.168.5.79:6379")
	//hydra.Conf.MQC(lmq.MQ)
	//hydra.Conf.Vars().Queue().LMQ("queuename")

	app.API("/hydratest/mqcserver/:queue", funcAPI)
	app.MQC("/hydratest/mqcserver/queue1", funcMQC1, "mqcserver:queue1")
	app.MQC("/hydratest/mqcserver/queue2", funcMQC2)
}

// mqcserver-queue 静态加载队列后，手动修改注册配置demo
//1.1 安装程序 ./mqcserver_queue conf install -v
//1.2 使用 ./mqcserver_queue run

//1.3 调用错误返回结果接口：http://localhost:8072/hydratest/mqcserver/queue1 都能执行指定的消息队列和mqc接受参数为：taosytest=queue1
//1.4 调用错误返回结果接口：http://localhost:8072/hydratest/mqcserver/queue2 都能执行指定的消息队列和mqc接受参数为：taosytest=queue2

//1.5 删除注册中心监听队列queue2  保存主节点，mqc服务自动重新启动
//1.6 调用错误返回结果接口：http://localhost:8072/hydratest/mqcserver/queue1 都能执行指定的消息队列和mqc接受参数为：taosytest=queue1
//1.7 调用错误返回结果接口：http://localhost:8072/hydratest/mqcserver/queue2 mqc-queue2队列不能执行，没有执行日志，在redis中查看消息存在

//1.8 添加注册中心监听队列queue2  保存主节点，mqc服务自动重新启动--mqc-queue2队列直接执行消费上面遗留的消息，接受参数：taosytest=queue2
//1.9 调用错误返回结果接口：http://localhost:8072/hydratest/mqcserver/queue1 都能执行指定的消息队列和mqc接受参数为：taosytest=queue1

//1.10 修改注册中心监听队列queue1-->queue3  保存主节点，mqc服务自动重启
//1.11 调用错误返回结果接口：http://localhost:8072/hydratest/mqcserver/queue1 mqc-queue1队列不能执行，没有执行日志，在redis中查看消息存在
//1.12 调用错误返回结果接口：http://localhost:8072/hydratest/mqcserver/queue2 都能执行指定的消息队列和mqc接受参数为：taosytest=queue2

//1.13 添加注册中心监听队列queue3-->queue1  保存主节点，mqc服务自动重新启动
//1.14 调用错误返回结果接口：http://localhost:8072/hydratest/mqcserver/queue1 都能执行指定的消息队列和mqc接受参数为：taosytest=queue1
//1.15 调用错误返回结果接口：http://localhost:8072/hydratest/mqcserver/queue2 都能执行指定的消息队列和mqc接受参数为：taosytest=queue2

//1.16 重复上面 1.3--1.15 数10次
func main() {
	app.Start()
}

var funcAPI = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("mqcserver-queue-api 静态加载队列后，手动修改注册配置demo")
	queue := ""
	value := ""
	p := ctx.Request().Path().Params()
	switch p.GetString("queue") {
	case "queue1":
		queue = "mqcserver:queue1"
		value = `{"taosytest":"queue1"}`
	case "queue2":
		queue = "mqcserver:queue2"
		value = `{"taosytest":"queue2"}`
	default:
		ctx.Log().Errorf("没有[%s]监听的队列", p.GetString("queue"))
		return
	}
	queueObj := components.Def.Queue().GetRegularQueue("queuename")
	if err := queueObj.Send(queue, value); err != nil {
		ctx.Log().Errorf("发送消息队列异常：%s", queue)
		return
	}
	return
}

var funcMQC1 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("mqcserver-queue-mqc1 静态加载队列后，手动修改注册配置demo")
	xmap := ctx.Request().GetMap()
	ctx.Log().Info("ctx.Request().GetMap()：", xmap)
	return
}

var funcMQC2 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("mqcserver-queue-mqc2 静态加载队列后，手动修改注册配置demo")
	xmap := ctx.Request().GetMap()
	ctx.Log().Info("ctx.Request().GetMap()：", xmap)
	return
}
