package main

import (
	"fmt"
	"os"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithDebug(),
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("apiservertrace"),
	hydra.WithClusterName("taosytest"),
	hydra.WithRegistry("zk://192.168.0.101"),
)

func init() {
	hydra.Conf.API(":8070", api.WithTrace())
	app.API("/hydratest/apiserver/trace", funcTrace)
}

// 代码安装开启trace配置，跟踪cpu性能demo

//1.1 安装程序 ./servertrace01 conf install -cover
//1.2 使用 ./servertrace01 run -t cpu
//1.3 调用接口：http://localhost:8070/hydratest/apiserver/trace 判定配置是否正确
func main() {
	app.Start()
}

var funcTrace = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver_trace 测试代码指定trace-cpu配置")
	f, err := os.Stat("./cpu.pprof")
	if os.IsNotExist(err) {
		return fmt.Errorf("cpu.pprof 不存在，没有启动跟踪,%v", err)
	}
	f.Size()
	return "success"
}
