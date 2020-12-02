package main

import (
	"fmt"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/api"

	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("apiserver_recovery"),
	hydra.WithClusterName("test"),
	hydra.WithRegistry("lm://."),
)

func init() {
	hydra.Conf.API(":50010", api.WithTimeout(10, 10))

	app.API("/api/localrpc", func(ctx context.IContext) (r interface{}) {
		panic(fmt.Errorf("主动抛出异常"))
	})
}

func main() {
	app.Start()
}
