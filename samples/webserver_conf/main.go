package main

import (
	"fmt"
	"time"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/static"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithDebug(),
	hydra.WithServerTypes(http.Web),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("webserverconf"),
	hydra.WithClusterName("test"),
	hydra.WithRegistry("lm://."),
)

func init() {
	vueconfig("conf")
	hydra.Conf.Web(":50003").Static(static.WithArchive("dist.zip"), static.WithRoot("./dist"))
}

func main() {
	app.Start()
}

func vueconfig(cur string) {
	hydra.Conf.Vars()["config"] = map[string]interface{}{
		"vue": map[string]interface{}{
			"api_addr":         fmt.Sprintf("//%s:50002", global.LocalIP()),
			"version":          time.Now().Format("20060102150405"),
			"currentComponent": cur,
		},
	}
	app.Web("/vue/config", func(ctx context.IContext) interface{} {
		data := map[string]interface{}{}
		ctx.APPConf().GetVarConf().GetObject("config", "vue", &data)
		return data
	})
}
