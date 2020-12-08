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
//启动服务./render_struct run
//	struct 不能转换为tengo.Object
func main() {
	app.API("/xml", request)
	app.API("/json", request)
	app.API("/plain", request)
	app.Start()
}

type result struct {
	Msg string
}

func request(ctx hydra.IContext) interface{} {
	return result{Msg: "success"}
}
