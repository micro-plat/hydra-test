package main

import (
	"time"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithDebug(),
	hydra.WithServerTypes(http.Web),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("webserverstatic"),
	hydra.WithClusterName("test"),
	hydra.WithRegistry("lm://."),
)

func init() {
	vueconfig("static")
	hydra.Conf.Web(":50005")
}

//默认是否已配置静态文件规则，及默认规则是否合理
//1. http://loaclhost:50005 正常相应页面
//2. http://loaclhost:50005/aaa.txt 正常相应页面
//3. http://loaclhost:50005/views/bbb.txt 状态们404
//4. http://loaclhost:50005/views/aaa.exe 状态们404
//5. http://loaclhost:50005/view/aaa.txt 状态们404
//6. http://loaclhost:50005/web/aaa.txt 状态们404
//7. http://loaclhost:50005/aaa.exe 状态们404
//8. http://loaclhost:50005/aaa.so 状态们404

func main() {
	app.Start()
}

func vueconfig(cur string) {
	hydra.Conf.Vars()["config"] = map[string]interface{}{
		"vue": map[string]interface{}{
			"api_addr":         "",
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
