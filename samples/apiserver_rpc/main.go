package main

import (
	"fmt"

	"github.com/micro-plat/hydra"
	crpc "github.com/micro-plat/hydra/components/rpcs/rpc"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/global"

	"github.com/micro-plat/hydra/hydra/servers/http"
	"github.com/micro-plat/hydra/hydra/servers/rpc"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API, rpc.RPC),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("apiserver_rpc"),
	hydra.WithClusterName("test"),
	hydra.WithRegistry("lm://."),
)

func init() {
	hydra.Conf.API(":50008", api.WithTimeout(10, 10))
	hydra.Conf.RPC(":50009")

	hydra.Conf.Vars().RPC("rpc")
	localIP := global.LocalIP()

	//1.请求 /api/localrpc
	//2.对比 requestID 和 /rpc/localrpc 是否一致

	app.API("/api/requestbyctx", rpcRequestByCtx)
	app.API("/api/swap", rpcSwap)
	app.API("/api/request", rpcRequest)

	app.API("/api/remoterpc", "rpc:///rpc/localrpc@hydratest")
	rpcSrv:=fmt.Sprintf("rpc:///rpc/localrpc@%s:50009", localIP)
	
 	app.API("/api/remoterpcip", rpcSrv) 
	app.RPC("/rpc/localrpc", func(ctx hydra.IContext) (r interface{}) {
		ctx.Log().Info("/rpc/localrpc:RequestID:", ctx.User().GetRequestID(), ctx.Request().GetString("name"))
		return ctx.User().GetRequestID()
	})

}

//启动服务
// 1. 请求http://localhost:50008/api/requestbyctx 对比响应数据中rpc和api一致
// 2. 请求http://localhost:50008/api/swap 对比响应数据中rpc和api一致
// 3. 请求http://localhost:50008/api/request 对比响应数据中rpc和api一致

// 4. 请求http://localhost:50008/api/remoterpc 通过rpc服务地址的方式注册（api转rpc)
// 5. 请求http://localhost:50008/api/remoterpcip 通过ip+port方式注册（api转rpc)
func main() {
	app.Start()
}

var rpcRequestByCtx = func(ctx hydra.IContext) (r interface{}) {
	requestID := ctx.User().GetRequestID()
	data := map[string]interface{}{
		"name": "RequestByCtx",
	}
	request, err := hydra.C.RPC().GetRPC()
	if err != nil {
		ctx.Log().Error("GetRPC:", err)
		return
	}
	res, err := request.RequestByCtx(ctx.Context(), "/rpc/localrpc@hydratest", data, crpc.WithXRequestID(requestID))
	if err != nil {
		ctx.Log().Error("RequestByCtx:", err)
		return
	}

	return map[string]interface{}{
		"api": requestID,
		"rpc": res.Result,
	}
}

var rpcSwap = func(ctx hydra.IContext) (r interface{}) {

	request, err := hydra.C.RPC().GetRPC()
	if err != nil {
		ctx.Log().Error("GetRPC:", err)
		return
	}
	res2, err := request.Swap("/rpc/localrpc@hydratest", ctx)
	if err != nil {
		ctx.Log().Error("Swap:", err)
		return
	}

	return map[string]interface{}{
		"api": ctx.User().GetRequestID(),
		"rpc": res2.Result,
	}
}

var rpcRequest = func(ctx hydra.IContext) (r interface{}) {
	requestID := ctx.User().GetRequestID()
	data := map[string]interface{}{
		"name": "Request",
	}
	request, err := hydra.C.RPC().GetRPC()
	if err != nil {
		ctx.Log().Error("GetRPC:", err)
		return
	}

	res3, err := request.Request("/rpc/localrpc@hydratest", data, crpc.WithXRequestID(requestID))
	if err != nil {
		ctx.Log().Error("Swap:", err)
		return
	}

	return map[string]interface{}{
		"api": ctx.User().GetRequestID(),
		"rpc": res3.Result,
	}
}
