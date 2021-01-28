package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("apiservercode"),
	hydra.WithClusterName("taosytest"),
	hydra.WithRegistry("zk://192.168.0.101"),
)

func init() {
	hydra.Conf.API("8072")
	app.API("/hydratest/apiserver/cmd", funcCode)
	app.Web("/hydratest/apiserver/codeweb", funcCode1)
}

//apiserver 命令覆盖代码配置demo

//1.1 安装程序 ./servercode02 conf install -p hydratest1 -c taosytest1 -s apiservercode1 -S web -d -cover
//1.2 使用 ./servercode02 run -p hydratest1 -c taosytest1 -s apiservercode1 -S web -d
//1.3 调用接口：http://localhost:8089/hydratest/apiserver/codeweb 判定配置是否正确
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

var funcCode1 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver_code 测试程序命令指定配置")
	ctx.Log().Info("GetPlatName:", ctx.APPConf().GetServerConf().GetPlatName())
	ctx.Log().Info("GetClusterName:", ctx.APPConf().GetServerConf().GetClusterName())
	ctx.Log().Info("GetSysName:", ctx.APPConf().GetServerConf().GetSysName())
	ctx.Log().Info("GetServerType:", ctx.APPConf().GetServerConf().GetServerType())
	ctx.Log().Info("IsDebug:", global.IsDebug)
	return "success"
}
