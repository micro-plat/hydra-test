package main

import (
	"github.com/micro-plat/hydra"
	_ "github.com/micro-plat/hydra/components/caches/cache/redis"
	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithDebug(),
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("apiservercode"),
	hydra.WithClusterName("taosytest"),
	hydra.WithRegistry("zk://192.168.0.101"),
)

func init() {
	hydra.Conf.API(":8070")
	app.API("/hydratest/apiserver/cmd", funcCode)
}

// apiserver命令重新指定注册中心类型demo

//1.1 使用 ./servercode03 run -r lm://. -p hydratest1 -c taosytest1 -s apiservercode1
//1.2 调用接口：http://localhost:8070/hydratest/apiserver/cmd 判定配置是否正确
func main() {
	app.Start()
}

var funcCode = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver_code 测试程序命令指定配置")
	ctx.Log().Info("GetPlatName:", ctx.APPConf().GetServerConf().GetPlatName())
	ctx.Log().Info("GetClusterName:", ctx.APPConf().GetServerConf().GetClusterName())
	ctx.Log().Info("GetSysName:", ctx.APPConf().GetServerConf().GetSysName())
	ctx.Log().Info("GetServerType:", ctx.APPConf().GetServerConf().GetServerType())
	ctx.Log().Info("IsDebug:", global.IsDebug)
	return "success"
}
