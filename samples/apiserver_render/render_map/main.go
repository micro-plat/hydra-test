package main

import (
	"github.com/micro-plat/hydra"

	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithPlatName("hydra-t"),
	hydra.WithSystemName("apiserver"),
	hydra.WithClusterName("render"),
	hydra.WithServerTypes(http.API),
)

//测试render组件，将普通的字符串render为xml, 特殊的json,或plain等格式
//启动服务./render_map run
//访问 /xml  [返回200,content-type为application/xml,内容为组装的xml:<response><code>200</code><msg>success</msg></response>]
//访问 /json [返回200,content-type为application/json,内容为组装的json:{"msg":"success"}]
//访问 /plain [返回200,content-type为text/plain,内容为success]
func main() {
	app.API("/xml", request)
	app.API("/json", request)
	app.API("/plain", request)
	app.Start()
}

func request(ctx hydra.IContext) interface{} {
	return map[string]interface{}{
		"msg": "success",
	}
}
