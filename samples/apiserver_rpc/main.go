package main

import (
	"fmt"

	"github.com/micro-plat/hydra"
	crpc "github.com/micro-plat/hydra/components/rpcs/rpc"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/global"

	"github.com/micro-plat/hydra/context"
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

	//1.请求 /api/localrpc
	//2.对比 requestID 和 /rpc/localrpc 是否一致

	app.API("/api/localrpc", func(ctx context.IContext) (r interface{}) {
		requestID := ctx.User().GetRequestID()
		ctx.Log().Info("/api/localrpc:RequestID:", requestID)
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
		res, err := request.RequestByCtx(ctx.Context(), "/rpc/localrpc@hydratest", data, crpc.WithXRequestID(requestID))
		if err != nil {
			ctx.Log().Error("RequestByCtx:", err)
			return
		}
		data["name"] = "Swap"
		res2, err := request.Swap("/rpc/localrpc@hydratest", ctx)
		if err != nil {
			ctx.Log().Error("Swap:", err)
			return
		}

		data["name"] = "Request"
		res3, err := request.Request("/rpc/localrpc@hydratest", data, crpc.WithXRequestID(requestID))
		if err != nil {
			ctx.Log().Error("Swap:", err)
			return
		}

		return map[string]interface{}{
			"api":  ctx.User().GetRequestID(),
			"rpc1": res.Result,
			"rpc2": res2.Result,
			"rpc3": res3.Result,
		}
	})

	localIP := global.LocalIP()

	app.API("/api/remoterpc", "rpc:///rpc/localrpc@hydratest")
	app.API("/api/remoterpcip", fmt.Sprintf("rpc://%s:50009/rpc/localrpc", localIP))

	app.RPC("/rpc/localrpc", func(ctx context.IContext) (r interface{}) {

		ctx.Log().Info("/rpc/localrpc:RequestID:", ctx.User().GetRequestID(), ctx.Request().GetString("name"))
		return ctx.User().GetRequestID()
	})

}

//启动服务
//  /api/localrpc [GET.POST]对比三种rpc 调用收集到的requestid
//  /api/remoterpc [GET.POST] 通过rpc服务地址的方式注册（api转rpc)
//  /api/remoterpcip [GET.POST] 通过ip+port方式注册（api转rpc)
func main() {
	app.Start()
}
