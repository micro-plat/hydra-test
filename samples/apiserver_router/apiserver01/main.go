package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithDebug(),
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("apiserverrouter"),
	hydra.WithClusterName("test"),
	hydra.WithRegistry("lm://."),
)

func init() {
	hydra.Conf.API(":50004")
	app.API("/hydratest/apiserver/router", &GetStruct{})
	app.API("/hydratest/apiserver/router", &PostStruct{})
	app.API("/hydratest/apiserver/router", &PutStruct{})
	app.API("/hydratest/apiserver/router", &DeleteStruct{})

	//测试路由重复注册，或POST,GET,PUT,DELETE等已被占用时的服务注册,在增加注册会报错误
	app.API("/hydratest/apiserver/router", &AddStruct{})
}

//同一个地址注册Action用完,启动会报错
func main() {
	app.Start()
}
