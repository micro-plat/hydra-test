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
	hydra.WithRegistry("lm://."),
)

func init() {
	hydra.Conf.API(":8072").Metric("http://192.168.106.219:8086", "hydratest", "@every 5s", metric.WithDisable())
	app.API("/hydratest/apiservermetric/metric", funcAPI)
}

// apiserver_metric 配置被禁用，统计信息保存demo

//1.1 运行程序 ./apimetric01 run
//1.2 调用接口：http://localhost:8072/hydratest/apiservermetric/metric 查看influxdb中数据是否正确
func main() {
	app.Start()
}

var funcAPI = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver_metric 配置被禁用，统计信息保存demo")
	return "success"
}
