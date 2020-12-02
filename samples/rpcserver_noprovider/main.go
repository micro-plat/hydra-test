package main

import (
	"fmt"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/components/rpcs"
	crpc "github.com/micro-plat/hydra/components/rpcs/rpc"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/global"

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

type CallMethod func(rpcs.IRequest, hydra.IContext, string, string) (string, error)

var callmethods = map[string]CallMethod{}

func init() {

	callmethods["requestbyctx"] = requestbyctx
	callmethods["swap"] = swap
	callmethods["request"] = request

	localIP := global.LocalIP()
	addrs := map[string]string{}
	addrs["platname"] = "/rpc/not/exists@hydratest"
	addrs["ip"] = fmt.Sprintf("/rpc/not/exists@tcp://%s:50016", localIP)

	hydra.Conf.API(":50015", api.WithTimeout(10, 10))
	hydra.Conf.RPC(":50016")

	hydra.Conf.Vars().RPC("rpc")

	/*
		注册六种服务：

		   /api/requestbyctx/platname
		   /api/requestbyctx/ip

		   /api/swap/platname
		   /api/swap/ip

		   /api/request/platname
		   /api/request/ip

	*/

	for m := range callmethods {
		for name, addr := range addrs {
			fmt.Println("API:", fmt.Sprintf("/api/%s/%s", m, name))
			handler := Callback(addr, m)
			app.API(fmt.Sprintf("/api/%s/%s", m, name), handler)
		}
	}

}

func Callback(addr string, requestMethod string) func(hydra.IContext) interface{} {
	callmethod := callmethods[requestMethod]
	fmt.Println("callmethod:", callmethod)
	return func(ctx hydra.IContext) (r interface{}) {
		fmt.Println("r1")
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

func requestbyctx(rpcReq rpcs.IRequest, ctx hydra.IContext, addr, requestID string) (result string, err error) {

	data := map[string]interface{}{}
	data["name"] = "RequestByCtx"
	res, err := rpcReq.RequestByCtx(ctx.Context(), addr, data, crpc.WithXRequestID(requestID))
	if err != nil {
		return
	}
	result = res.Result
	return
}

func request(rpcReq rpcs.IRequest, ctx hydra.IContext, addr, requestID string) (result string, err error) {
	data := map[string]interface{}{}
	data["name"] = "Request"
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
