package main

import (
	"fmt"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/hydra/hydra/servers/http"
	"github.com/micro-plat/hydra/hydra/servers/rpc"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API, rpc.RPC),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("rpcserbalance"),
	hydra.WithClusterName("test"),
	hydra.WithRegistry("zk://192.168.0.101"),
)

func init() {
	hydra.Conf.RPC("8073")
	hydra.Conf.API("8072")
	hydra.Conf.Vars().RPC("rpc")
	app.API("/hydratest/rpcserbalance/apiip", funcAPI)
	app.RPC("/hydratest/rpcserbalance/rpcip", funcRPC)
}

// rpcserver_balance 测试多个provider时默认本地有限负载均衡规则执行demo

//1.1 安装程序 ./rpcserverbalance01 conf install -v
//1.2 使用 ./rpcserverbalance01 run
//1.3 拷贝一份执行程序到其他pc主机上
//1.4 调用接口执行循环访问rpc：http://localhost:8072/hydratest/rpcserbalance/apiip 观察两台服务器的执行日志，只访问本地rpc的服务器
func main() {
	app.Start()
}

var funcAPI = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("rpcserver_balance 测试多个provider时默认本地有限负载均衡规则执行demo")
	input := map[string]interface{}{
		"taosytest": "123456",
	}
	for i := 0; i < 12; i++ {
		url := fmt.Sprintf("/hydratest/rpcserbalance/rpcip@hydratest")
		respones, err := components.Def.RPC().GetRegularRPC().Request(url, input)
		if err != nil {
			ctx.Log().Errorf("rpc 请求异常：%v", err)
			return
		}
		ctx.Log().Info("respones.Status:", respones.GetStatus())
	}
	return "success"
}

var funcRPC = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("rpcserver_balance 测试多个provider时默认本地有限负载均衡规则执行demo")
	xMap := ctx.Request().GetMap()
	ctx.Log().Info("ctx.Request().GetMap()：", xMap)
	return "success"
}
