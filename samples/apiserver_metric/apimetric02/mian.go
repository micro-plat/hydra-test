package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/metric"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("apiservermetric"),
	hydra.WithClusterName("taosytest"),
	hydra.WithRegistry("lm://./"),
)

func init() {
	hydra.Conf.API("8072").Metric("http://192.168.106.219:8086", "hydratest", "@every 5s", metric.WithUPName("", ""))
	app.API("/hydratest/apiservermetric/metric/succ", funcAPI1)
	app.API("/hydratest/apiservermetric/metric/fail", funcAPI2)
}

// apiserver_metric 配置被启用，根据配置统计信息保存demo

//1.1 运行程序 ./apimetric02 run
//1.2 随机的访问下面两个接口，查看influxdb中统计数据是否正确;定时上报时间是否生效;
//1.3 调用接口：http://localhost:8072/hydratest/apiservermetric/metric/succ
//1.4 调用接口：http://localhost:8072/hydratest/apiservermetric/metric/fail
func main() {
	app.Start()
}

var funcAPI1 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver_metric 配置被启用，根据配置统计信息保存demo")
	return "success"
}

var funcAPI2 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver_metric 配置被启用，根据配置统计信息保存demo")
	ctx.Response().Abort(678, "apiserver_metric 测试错误统计")
	return
}
