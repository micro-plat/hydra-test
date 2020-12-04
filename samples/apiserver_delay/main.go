package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithDebug(),
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("apiserverdelay"),
	hydra.WithClusterName("taosytest"),
	hydra.WithRegistry("zk://192.168.0.101"),
)

func init() {
	hydra.Conf.API(":8070", api.WithTimeout(10, 10))
	app.API("/hydratest/apiserver/delay", funcHeader)
}

// apiserver-delay延迟中间件测试demo

//1.1 安装程序 ./apiserver_delay conf install -cover
//1.2 使用 ./apiserver_delay run
//1.3 不设置X-Add-Delay：http://localhost:8070/hydratest/apiserver/delay 观察日志X-Add-Delay的值和响应时间是否正常
//1.4 设置X-Add-Delay=3s：http://localhost:8070/hydratest/apiserver/delay 观察日志X-Add-Delay的值和响应时间是否正常>3s
//1.6 设置X-Add-Delay=11s：http://localhost:8070/hydratest/apiserver/delay 大于了链接超市时间，应该返回链接超时
func main() {
	app.Start()
}

var funcHeader = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver-delay延迟中间件测试demo")
	headers := ctx.Request().Headers()
	ctx.Log().Info("Header-delay:", headers.GetString("X-Add-Delay"))
	return "success"
}
