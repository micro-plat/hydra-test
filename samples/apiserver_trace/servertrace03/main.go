package main

import (
	"fmt"
	"os"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("apiservertrace"),
	hydra.WithClusterName("taosytest"),
	hydra.WithRegistry("lm://."),
)

func init() {
	hydra.Conf.API("8072", api.WithTrace())
	app.API("/hydratest/apiserver/trace", funcTrace)
}

//apiserver_trace 代码安装开启trace配置，跟踪block性能demo

//1.1 使用 ./servertrace03 run -t block
//1.2 调用接口：http://localhost:8072/hydratest/apiserver/trace 判定配置是否正确
func main() {
	app.Start()
}

var funcTrace = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver_trace 测试代码指定trace-block配置")
	_, err := os.Stat("./block.pprof")
	if os.IsNotExist(err) {
		return fmt.Errorf("block.pprof 不存在，没有启动跟踪,%v", err)
	}
	return "success"
}
