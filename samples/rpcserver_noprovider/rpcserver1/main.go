package main

import (
	"fmt"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/components/rpcs"
	crpc "github.com/micro-plat/hydra/components/rpcs/rpc"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/hydra/registry"

	"github.com/micro-plat/hydra/hydra/servers/http"
	"github.com/micro-plat/hydra/hydra/servers/rpc"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API, rpc.RPC),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("rpcserver_noprovider"),
	hydra.WithClusterName("test"),
	hydra.WithRegistry("lm://."),
)

type RpcCallMethod func(rpcs.IRequest, hydra.IContext, string, string) (string, error)

var callmethods = map[string]RpcCallMethod{}

func init() {

	hydra.Conf.API(":50015", api.WithTimeout(10, 10))
	hydra.Conf.RPC(":50016")
	hydra.Conf.Vars().RPC("rpc")

	localIP := global.LocalIP()

	callmethods["requestbyctx"] = requestbyctx
	callmethods["swap"] = swap
	callmethods["request"] = request

	platAddr := "/rpc/not/exists@hydratest"
	ipAddr := fmt.Sprintf("/rpc/not/exists@tcp://%s:50016", localIP)

	/*
		注册六种服务：
		/api/requestbyctx/platname
		/api/requestbyctx/ip

		/api/swap/platname
		/api/swap/ip

		/api/request/platname
		/api/request/ip
	*/

	app.API("/api/requestbyctx/platname", handleCallback(platAddr, "requestbyctx"))
	app.API("/api/requestbyctx/ip", handleCallback(ipAddr, "requestbyctx"))
	app.API("/api/swap/platname", handleCallback(platAddr, "swap"))
	app.API("/api/swap/ip", handleCallback(ipAddr, "swap"))
	app.API("/api/request/platname", handleCallback(platAddr, "request"))
	app.API("/api/request/ip", handleCallback(ipAddr, "request"))

	app.API("/api/rpc/start", rpcStart)
	app.API("/api/rpc/stop", rpcStop)

	app.RPC("/rpc/local", rpcLocal)
	//app.RPC("/rpc/not/exists", rpcLocal)

}

func handleCallback(addr string, requestMethod string) func(hydra.IContext) interface{} {
	callmethod := callmethods[requestMethod]
	return func(ctx hydra.IContext) (r interface{}) {
		request, err := hydra.C.RPC().GetRPC()
		if err != nil {
			ctx.Log().Error("GetRPC:", err)
			return
		}
		requestID := ctx.User().GetRequestID()
		result, err := callmethod(request, ctx, addr, requestID)
		if err != nil {
			ctx.Log().Errorf("%s:%v", requestMethod, err)
			return err
		}
		return result
	}
}

//测试有端口启用，但无服务的情况
//启动服务:go run main.go run

//1.请求 http://localhost:50015/api/requestbyctx/platname 报：没有找到有效的服务路径, /rpc/not/exists
//2.请求 http://localhost:50015/api/requestbyctx/ip 报：没有找到有效的服务路径, /rpc/not/exists

//3.请求 http://localhost:50015/api/swap/platname 报：没有找到有效的服务路径, /rpc/not/exists
//4.请求 http://localhost:50015/api/swap/ip 报：没有找到有效的服务路径, /rpc/not/exists

//5.请求 http://localhost:50015/api/request/platname 报：没有找到有效的服务路径, /rpc/not/exists
//6.请求 http://localhost:50015/api/request/ip 报：没有找到有效的服务路径, /rpc/not/exists

func main() {
	app.Start()
}

func requestbyctx(rpcReq rpcs.IRequest, ctx hydra.IContext, addr, requestID string) (result string, err error) {

	data := map[string]interface{}{}
	res, err := rpcReq.RequestByCtx(ctx.Context(), addr, data, crpc.WithXRequestID(requestID))
	if err != nil {
		return
	}
	result = res.Result
	return
}

func request(rpcReq rpcs.IRequest, ctx hydra.IContext, addr, requestID string) (result string, err error) {
	data := map[string]interface{}{}
	res, err := rpcReq.Request(addr, data, crpc.WithXRequestID(requestID))
	if err != nil {
		return
	}
	result = res.Result
	return
}

func swap(rpcReq rpcs.IRequest, ctx hydra.IContext, addr, requestID string) (result string, err error) {
	res, err := rpcReq.Swap(addr, ctx)
	if err != nil {
		return
	}
	result = res.Result
	return
}

func rpcLocal(ctx hydra.IContext) interface{} {
	return "rpclocal.success"
}

func rpcStart(ctx hydra.IContext) interface{} {
	regst, err := registry.GetRegistry(global.Def.RegistryAddr, global.Def.Log())
	if err != nil {
		return err
	}
	err = regst.Update("/hydratest/rpcserver_noprovider/rpc/test/conf", `{"address":":50016","status":"start"}`)
	if err != nil {
		ctx.Log().Error("Update:", err)
		return err
	}
	return "success"
}

func rpcStop(ctx hydra.IContext) interface{} {
	regst, err := registry.GetRegistry(global.Def.RegistryAddr, global.Def.Log())
	if err != nil {
		return err
	}
	err = regst.Update("/hydratest/rpcserver_noprovider/rpc/test/conf", `{"address":":50016","status":"stop"}`)
	if err != nil {
		return err
	}
	return "success"
}
