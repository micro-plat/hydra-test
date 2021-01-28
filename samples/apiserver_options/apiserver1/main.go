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
	hydra.WithServerTypes(http.API, http.Web),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("apiserver_option"),
	hydra.WithClusterName("test"),
	hydra.WithRegistry("lm://."),
)

func init() {
	hydra.Conf.API("50011")
	hydra.Conf.Web("50012").Static(static.WithArchive("dist.zip"), static.WithRoot("./"))
	hydra.Conf.Vars().Custom("config", "vue", map[string]interface{}{
		"api_addr":         fmt.Sprintf("//%s:50011", global.LocalIP()),
		"version":          time.Now().Format("20060102150405"),
		"currentComponent": "options",
	})

	app.Web("/vue/config", func(ctx hydra.IContext) interface{} {
		data := map[string]interface{}{}
		ctx.APPConf().GetVarConf().GetObject("config", "vue", &data)
		return data
	})

	app.API("/options", func(ctx context.IContext) (r interface{}) {
		return "success"
	})
}

//测试option请求是否正确工作（未设置夸域头信息）
//启动server: go run main.go run
//1. 访问： http://localhost:50012 正常获取到页面
//2. 点击页面[Options]按钮，发起后端请求，报夸域错误 (has been blocked by CORS policy: No 'Access-Control-Allow-Origin' header is present on the requested resource.)
func main() {
	app.Start()
}
