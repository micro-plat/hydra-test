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

// 代码安装开启trace配置，跟踪web性能demo

//1.1 安装程序 ./servertrace05 conf install -cover
//1.2 使用默认端口监听 ./servertrace05 run -t web
//1.3 调用接口：http://localhost:8070/hydratest/apiserver/trace 判定配置是否正确
//1.4 调用性能跟踪web：http://0.0.0.0:19999/debug/pprof/ 判定配置是否开启

//2.1 使用命令指定traceweb端口号 ./servertrace05 run -t web -tp 19998
//2.2 调用性能跟踪web：http://0.0.0.0:19998/debug/pprof/ 判定配置是否开启
func main() {
	app.Start()
}

var funcTrace = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver_trace 测试代码指定trace-web配置")
	f, err := os.Stat("./trace.out")
	if os.IsNotExist(err) {
		return fmt.Errorf("trace.out 不存在，没有启动跟踪,%v", err)
	}
	f.Size()
	return "success"
}
