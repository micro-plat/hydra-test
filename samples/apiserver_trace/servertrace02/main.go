package main

import (
	"fmt"
	"os"

	"github.com/micro-plat/hydra"
	_ "github.com/micro-plat/hydra/components/caches/cache/redis"
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

// 代码安装开启trace配置，跟踪mem性能demo

//1.1 安装程序 sudo ./servertrace02 conf install -cover
//1.2 使用 ./servertrace02 run -t mem
//1.3 调用接口：http://localhost:8070/hydratest/apiserver/trace 判定配置是否正确
func main() {
	app.Start()
}

var funcTrace = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver_trace 测试代码指定trace-mem配置")
	f, err := os.Stat("./mem.pprof")
	if os.IsNotExist(err) {
		return fmt.Errorf("mem.pprof 不存在，没有启动跟踪,%v", err)
	}
	f.Size()
	return "success"
}
