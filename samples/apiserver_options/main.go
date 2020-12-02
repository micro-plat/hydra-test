package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/static"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

func funcAPI1(ctx context.IContext) (r interface{}) {
	return "success"
}

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API, http.Web),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("apiserver_option"),
	hydra.WithClusterName("test"),
	hydra.WithRegistry("zk://192.168.0.101"),
)

func init() {
	hydra.Conf.API(":58071")
	hydra.Conf.Web(":58072").Static(static.WithArchive("dist.zip"), static.WithRoot("./dist"))
	app.API("/options", funcAPI1)
}

//测试option请求是否正确工作
func main() {
	app.Start()
}
