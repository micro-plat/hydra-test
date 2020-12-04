package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/static"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydra_test"),
	hydra.WithSystemName("static_exclude"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

func init() {
	hydra.Conf.API(":8070").Static(
		static.WithExclude(".zip", ".css", "/exclude/"), static.WithRoot("../root"), static.WithEnable())

	app.API("/api", func(ctx hydra.IContext) (r interface{}) {
		ctx.Log().Info("api_static")
		return
	})

}

//测试 WithArchive配置是否正常
//启动服务 ./static_exclude run
//浏览器访问 http://localhost:8070/archive.zip [日志打印请求路径和状态,浏览器404 page not found]
//浏览器访问 http://localhost:8070/exclude.css [日志打印请求路径和状态,浏览器404 page not found]
//浏览器访问 http://localhost:8070/exclude/home.html[日志打印请求路径和状态,浏览器404 page not found]
//浏览器访问 http://localhost:8070/first.html [正常 返回200]
func main() {
	app.Start()
}
