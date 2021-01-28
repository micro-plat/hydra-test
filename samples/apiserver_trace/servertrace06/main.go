package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("apiservertrace"),
	hydra.WithClusterName("taosytest"),
	hydra.WithRegistry("lm://."),
)

func init() {
	hydra.Conf.API("8070", api.WithTrace())
	app.API("/hydratest/apiserver/trace", funcTrace)
}

//apiserver_trace 检查各种输入和输出的数据demo

//1.1 使用默认端口监听 ./servertrace06 run

//以下请求使用postman调用
//1.2 调用接口：http://localhost:8070/hydratest/apiserver/trace?param1=test&param2=中文数据$%#^##@  get请求带中文key-value数据
//1.3 调用接口：http://localhost:8070/hydratest/apiserver/trace   post-body-json：{"param1":"test","param2":"中文数据$%#^##@"}
//1.4 调用接口：http://localhost:8070/hydratest/apiserver/trace   post-body-json：{"param1":"test","param2":"中文数据$%#^##@"}
//1.5 调用接口：http://localhost:8070/hydratest/apiserver/trace   post-body-xml：<xml><param1>test</param1><param2>中文数据$%#^##@</param2></xml>
//1.6 调用接口：http://localhost:8070/hydratest/apiserver/trace   post-body-text：param1testparam2中文数据$%#^##@   不编码传输
//1.7 调用接口：http://localhost:8070/hydratest/apiserver/trace   post-fromdata-json：{"param1":"test","param2":"中文数据$%#^##@"}   编码后传输
//1.8 调用接口：http://localhost:8070/hydratest/apiserver/trace   post-fromdata-xml：<xml><param1>test</param1><param2>中文数据$%#^##@</param2></xml>
//1.9 调用接口：http://localhost:8070/hydratest/apiserver/trace   post-fromdata-text：param1testparam2中文数据$%#^##@   编码后传输
//1.10 调用接口：http://localhost:8070/hydratest/apiserver/trace?param1=test&param2=中文数据$%#^##@  post-bodyjson-kv混合传输： {"param13":"test","param4":"中文数据$%#^##@"}
//1.11 调用接口：http://localhost:8070/hydratest/apiserver/trace?param1=test&param2=中文数据$%#^##@  post-bodyjson-kv混合传输： {"param1":"test","param4":"中文数据$%#^##@"}
//1.12 调用接口：http://localhost:8070/hydratest/apiserver/trace?param1=test&param2=中文数据$%#^##@  post-bodyxml：-kv混合传输：<xml><param1>test</param3><param2>中文数据$%#^##@</param4></xml>
//1.13 调用接口：http://localhost:8070/hydratest/apiserver/trace?param1=test&param2=中文数据$%#^##@  post-bodyxml：-kv混合传输：<xml><param1>test</param1><param2>中文数据$%#^##@</param2></xml>
//1.14 调用接口：http://localhost:8070/hydratest/apiserver/trace?param1=test&param2=中文数据$%#^##@  post-bodytext-kv混合传输： param1testparam2中文数据34%%
//1.15 调用接口：http://localhost:8070/hydratest/apiserver/trace?param1=test&param2=中文数据$%#^##@  post-body-kv混合传输： {"param13":"test","param4":"中文数据"}
//1.16 调用接口：http://localhost:8070/hydratest/apiserver/trace?param1=test&param2=中文数据$%#^##@  post-body-kv混合传输： {"param1":"test","param4":"中文数据"}

func main() {
	app.Start()
}

var funcTrace = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver_trace 检查各种输入和输出的数据demo")
	xmap := ctx.Request().GetMap()

	ctx.Log().Info("trance监控的GetMap：", xmap)
	return "success"
}
