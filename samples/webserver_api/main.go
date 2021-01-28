package main

import (
	"time"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/static"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.Web),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("webserverapi"),
	hydra.WithClusterName("test"),
	hydra.WithRegistry("lm://."),
)

func init() {
	hydra.Conf.Web("50004").Static(static.WithArchive("dist.zip"), static.WithRoot("./"))
	hydra.Conf.Vars().Custom("config", "vue", map[string]interface{}{
		"api_addr":         "",
		"version":          time.Now().Format("20060102150405"),
		"currentComponent": "mixed",
	})

	app.Web("/vue/config", func(ctx hydra.IContext) interface{} {
		data := map[string]interface{}{}
		ctx.APPConf().GetVarConf().GetObject("config", "vue", &data)
		return data
	})

	app.Web("/self/api", func(ctx hydra.IContext) interface{} {
		return "self.api.success"
	})
}

//混合静态文件，api接口是否正常工作
//启动程序 go run main.go run
//1. http://localhost:50004 首页正常访问
//2. 点击页面中的 按钮，请求接口地址 正常返回
//3. http://localhost:50004/self/api 正常返回
func main() {
	app.Start()
}
