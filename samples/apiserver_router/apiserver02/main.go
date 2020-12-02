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
	hydra.Conf.API(":50005")
	//同一个对象的Get/Post互斥
	app.API("/hydratest/apiserver/router", &SrvStruct{})
}

func main() {
	app.Start()
}
