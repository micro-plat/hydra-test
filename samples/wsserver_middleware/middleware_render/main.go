package main

import (
	"github.com/micro-plat/hydra"

	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.WS),
	hydra.WithPlatName("hydra_test"),
	hydra.WithSystemName("ws_render"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("zk://192.168.0.101"),
)

//测试WS的render组件
//启动服务./middleware_render run
//访问 /xml  [返回200,content-type为application/xml,内容为组装的xml:<response><code>200</code><msg>success</msg></response>]
//访问 /json [返回200,content-type为application/json,内容为组装的json:{"msg":"success"}]
//访问 /plain [返回200,content-type为text/plain,内容为success]
func main() {
	app.WS("/xml", request)
	app.WS("/json", request)
	app.WS("/plain", request)
	app.Start()
}

func request(ctx hydra.IContext) interface{} {
	return map[string]interface{}{
		"msg": "success",
	}
}
