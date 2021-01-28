package main

import (
	"fmt"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/hydra/hydra/servers/http"
	"github.com/micro-plat/hydra/hydra/servers/rpc"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API, rpc.RPC),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("rpcserIP"),
	hydra.WithClusterName("taosytest"),
	hydra.WithRegistry("lm://."),
)

func init() {
	hydra.Conf.RPC("8073")
	hydra.Conf.API("8072")
	hydra.Conf.Vars().RPC("rpc")
	app.API("/hydratest/rpcserver/apiip/fail", funcAPI)
	app.RPC("/hydratest/rpcserver/rpcip/fail", funcRPC)
	app.API("/hydratest/rpcserver/apiip/succ", funcAPI1)
	app.RPC("/hydratest/rpcserver/rpcip/succ", funcRPC1)
}

// rpcserver-ip 通过ip地址直接访问rpc服务demo

//1.1 使用 ./rpcserver_ip run
//1.2 调用错误返回结果接口：http://localhost:8072/hydratest/rpcserver/apiip/fail 观察日志中rpc如参是否正确 返回值： 666/rpc服务返回异常
//1.3 调用正确返回结果接口：http://localhost:8072/hydratest/rpcserver/apiip/succ 观察日志中rpc如参是否正确 返回值： 200/rpcsuccess
func main() {
	app.Start()
}

var funcAPI = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("rpcserver-ip-api rpc错返回访问demo")
	url := fmt.Sprintf("/hydratest/rpcserver/rpcip/fail@%s:8073", global.LocalIP())
	input := map[string]interface{}{
		"taosytest": "123456",
	}
	ctx.Log().Info("url:", url)
	respones, err := components.Def.RPC().GetRegularRPC().Request(url, input)
	if err != nil {
		ctx.Log().Errorf("rpc 请求异常：%v", err)
		return
	}
	ctx.Log().Info("respones.IsSuccess():", respones.IsSuccess())
	ctx.Log().Info("respones.GetStatus():", respones.GetStatus())
	ctx.Log().Info("respones.GetResult():", respones.GetResult())
	ctx.Response().Abort(respones.GetStatus(), respones.GetResult())
	return
}

var funcRPC = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("rpcserver-ip-rpc rpc错返回访问demo")
	xMap := ctx.Request().GetMap()
	ctx.Log().Info("ctx.Request().GetMap()：", xMap)
	ctx.Response().Abort(666, "rpc服务返回异常")
	return
}

var funcAPI1 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("rpcserver-ip-api rpc正确返回访问demo")
	url := fmt.Sprintf("/hydratest/rpcserver/rpcip/succ@%s:8073", global.LocalIP())
	input := map[string]interface{}{
		"taosytest": "654321",
	}
	ctx.Log().Info("url:", url)
	respones, err := components.Def.RPC().GetRegularRPC().Request(url, input)
	if err != nil {
		ctx.Log().Errorf("rpc 请求异常：%v", err)
		return
	}
	ctx.Log().Info("respones.IsSuccess():", respones.IsSuccess())
	ctx.Log().Info("respones.GetStatus():", respones.GetStatus())
	ctx.Log().Info("respones.GetResult():", respones.GetResult())
	return respones.GetResult()
}

var funcRPC1 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("rpcserver-ip-rpc rpc正确返回访问demo")
	xMap := ctx.Request().GetMap()
	ctx.Log().Info("ctx.Request().GetMap()：", xMap)
	return "rpcsuccess"
}
