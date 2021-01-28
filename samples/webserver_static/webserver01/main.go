package main

import (
	"time"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.Web),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("webserverstatic"),
	hydra.WithClusterName("test"),
	hydra.WithRegistry("lm://."),
)

func init() {
	hydra.Conf.WEB("50005")

	hydra.Conf.Vars().Custom("config", "vue", map[string]interface{}{
		"api_addr":         "",
		"version":          time.Now().Format("20060102150405"),
		"currentComponent": "static",
	})
	app.Web("/vue/config", func(ctx hydra.IContext) interface{} {
		data := map[string]interface{}{}
		ctx.APPConf().GetVarConf().GetObject("config", "vue", &data)
		return data
	})
}

//默认是否已配置静态文件规则，及默认规则是否合理 (查找目录下文件夹)
//1. http://loaclhost:50005 正常响应页面，状态码：200
//2. http://loaclhost:50005/aaa.txt 正常响应页面，状态码：200
//3. http://loaclhost:50005/views/bbb.txt 状态码：404
//4. http://loaclhost:50005/views/aaa.exe 状态码：404
//5. http://loaclhost:50005/view/aaa.txt 状态码：404
//6. http://loaclhost:50005/web/aaa.txt 状态码：404
//7. http://loaclhost:50005/aaa.exe 状态码：404
//8. http://loaclhost:50005/aaa.so 状态码：404

func main() {
	app.Start()
}
