package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("apiservercli"),
	hydra.WithClusterName("taosytest"),
	hydra.WithRegistry("zk://192.168.0.101"),
)

func init() {
	hydra.Conf.API(":8071")
	app.API("/hydratest/apiserver/cli", funcCli)
	app.Web("/hydratest/apiserver/cliweb", funcCli)
}

//apiserver 通过命令指定服务配置demo

//1.1 通过命令重新指定服务配置 sudo ./apiserver_cli conf install -p hydratest1 -c taosytest1 -s apiservercli1 -S web -cover
//1.2 使用 ./apiserver_cli run -p hydratest1 -c taosytest1 -s apiservercli1 -S web
//1.3 调用接口：http://localhost:8089/hydratest/apiserver/cliweb 判定配置是否正确
func main() {
	app.Start()
}

var funcCli = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver_code 测试程序代码安装指定配置")
	ctx.Log().Info("GetPlatName:", ctx.APPConf().GetServerConf().GetPlatName())
	ctx.Log().Info("GetClusterName:", ctx.APPConf().GetServerConf().GetClusterName())
	ctx.Log().Info("GetSysName:", ctx.APPConf().GetServerConf().GetSysName())
	ctx.Log().Info("GetServerType:", ctx.APPConf().GetServerConf().GetServerType())
	ctx.Log().Info("IsDebug:", global.IsDebug)
	return "success"
}
