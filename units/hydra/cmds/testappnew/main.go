package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithSystemName("apiserver"),
	hydra.WithClusterName("c"),
)

const newRedisAddr = "192.168.5.79:6379"

func main() {
	hydra.Conf.Vars().Redis("5.79", newRedisAddr)
	hydra.Conf.API("19003")
	app.Start()
}
