package main

import (
	"github.com/micro-plat/hydra"
	crpc "github.com/micro-plat/hydra/components/rpcs/rpc"
	"github.com/micro-plat/hydra/conf/server/api"

	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("rpcserver_noprovider"),
	hydra.WithClusterName("test"),
	hydra.WithRegistry("lm://."),
)

func init() {
	hydra.Conf.API(":50015", api.WithTimeout(10, 10))

	hydra.Conf.Vars().RPC("rpc")

	rpcAddr := "/rpc/not/exists@hydratest"

	//1.请求 /api/localrpc
	//2.对比 requestID 和 /rpc/localrpc 是否一致

	app.API("/api/requestbyctx", func(ctx context.IContext) (r interface{}) {
		requestID := ctx.User().GetRequestID()
		ctx.Log().Info("/api/requestbyctx:RequestID:", requestID)
		data, err := ctx.Request().GetMap()
		if err != nil {
			ctx.Log().Error("GetMap:", err)
			return
		}
		request, err := hydra.C.RPC().GetRPC()
		if err != nil {
			ctx.Log().Error("GetRPC:", err)
			return
		}
		data["name"] = "RequestByCtx"
		res, err := request.RequestByCtx(ctx.Context(), rpcAddr, data, crpc.WithXRequestID(requestID))
		if err != nil {
			ctx.Log().Error("RequestByCtx:", err)
			return
		}

		return map[string]interface{}{
			"api":  ctx.User().GetRequestID(),
			"rpc1": res.Result,
		}
	})

	app.API("/api/swap", func(ctx context.IContext) (r interface{}) {
		requestID := ctx.User().GetRequestID()
		ctx.Log().Info("/api/swap:RequestID:", requestID)
		data, err := ctx.Request().GetMap()
		if err != nil {
			ctx.Log().Error("GetMap:", err)
			return
		}
		request, err := hydra.C.RPC().GetRPC()
		if err != nil {
			ctx.Log().Error("GetRPC:", err)
			return
		}

		data["name"] = "Swap"
		res2, err := request.Swap(rpcAddr, ctx)
		if err != nil {
			ctx.Log().Error("Swap:", err)
			return
		}

		return map[string]interface{}{
			"api":  ctx.User().GetRequestID(),
			"rpc2": res2.Result,
		}
	})

	app.API("/api/request", func(ctx context.IContext) (r interface{}) {
		requestID := ctx.User().GetRequestID()
		ctx.Log().Info("/api/request:RequestID:", requestID)
		data, err := ctx.Request().GetMap()
		if err != nil {
			ctx.Log().Error("GetMap:", err)
			return
		}
		request, err := hydra.C.RPC().GetRPC()
		if err != nil {
			ctx.Log().Error("GetRPC:", err)
			return
		}

		data["name"] = "Request"
		res3, err := request.Request(rpcAddr, data, crpc.WithXRequestID(requestID))
		if err != nil {
			ctx.Log().Error("Request:", err)
			return
		}

		return map[string]interface{}{
			"api":  ctx.User().GetRequestID(),
			"rpc3": res3.Result,
		}
	})
}

//启动服务:go run main.go run
//1.请求 http://localhost:50015/api/requestbyctx 报：没有找到有效的服务路径, /rpc/not/exists
//2.请求 http://localhost:50015/api/swap 报：没有找到有效的服务路径, /rpc/not/exists
//3.请求 http://localhost:50015/api/request 报：没有找到有效的服务路径, /rpc/not/exists
func main() {
	app.Start()
}
