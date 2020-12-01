package main

import (
	"fmt"

	"github.com/micro-plat/hydra"
	_ "github.com/micro-plat/hydra/components/caches/cache/redis"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var funcAPI1 func(ctx context.IContext) (r interface{}) = func(ctx context.IContext) (r interface{}) {
	ctx.Log().Info("apiserver_code 测试程序代码安装指定配置")
	if ctx.APPConf().GetServerConf().GetPlatName() != "hydratest_debug" {
		return fmt.Errorf("PlatName 数据错误,%s", ctx.APPConf().GetServerConf().GetPlatName())
	}
	if ctx.APPConf().GetServerConf().GetClusterName() != "taosytest" {
		return fmt.Errorf("GetClusterName 数据错误,%s", ctx.APPConf().GetServerConf().GetClusterName())
	}
	if ctx.APPConf().GetServerConf().GetSysName() != "apiservercode" {
		return fmt.Errorf("GetSysName 数据错误,%s", ctx.APPConf().GetServerConf().GetSysName())
	}
	if ctx.APPConf().GetServerConf().GetServerType() != "api" {
		return fmt.Errorf("GetServerType 数据错误,%s", ctx.APPConf().GetServerConf().GetServerType())
	}
	b, err := ctx.APPConf().GetServerConf().GetRegistry().Exists("/hydratest_debug/apiservercode/api/taosytest/conf")
	if err != nil {
		return fmt.Errorf("GetRegistry 数据错误,%v", err)
	}
	if !b {
		return fmt.Errorf("GetRegistry 注册中心服务主节点不存在,%v", b)
	}
	if !global.IsDebug {
		return fmt.Errorf("IsDebug 数据错误,%v", global.IsDebug)
	}
	return "success"
}

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
	app.API("/hydratest/apiserver/code", funcAPI1)
}

// 代码指定服务配置demo

//1.1 安装程序 sudo ./apiserver_code conf install -cover
//1.2 使用 ./apiserver_code run
//1.3 调用接口：http://192.168.5.94:8070/hydratest/apiserver/code 判定配置是否正确
func main() {
	app.Start()
}
