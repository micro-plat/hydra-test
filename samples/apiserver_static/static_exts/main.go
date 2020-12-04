package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/static"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydra_test"),
	hydra.WithSystemName("static_exts"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

func init() {
	hydra.Conf.API(":8070").Static(
		static.WithExts(".htm"), static.AppendExts(".js"), static.WithRoot("../root"), static.WithEnable())

	app.API("/api", func(ctx hydra.IContext) (r interface{}) {
		ctx.Log().Info("api_static")
		return
	})

}

//测试 WithExts配置是否正常
//启动服务 ./static_exts run
//浏览器访问 http://localhost:8070/exts.htm [返回200]
//浏览器访问 http://localhost:8070/exts.htm [返回200]
//浏览器访问 http://localhost:8070/exts.html [浏览器404 page not found]
//将WithExts(".htm")改为static.WithExts("*")
//重新编译启动服务
//浏览器再次访问 http://localhost:8070/exts.htm [返回200]
//浏览器再次访问 http://localhost:8070/exts.htm [返回200]
//浏览器再次访问 http://localhost:8070/exts.html  [返回200]
func main() {
	app.Start()
}
