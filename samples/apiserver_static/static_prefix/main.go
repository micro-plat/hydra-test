package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/static"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydra_test"),
	hydra.WithSystemName("static_prefix"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

func init() {
	hydra.Conf.API(":8070").Static(
		static.WithPrefix("/prefix"), static.WithRoot("../root"), static.WithEnable())

	app.API("/api", func(ctx hydra.IContext) (r interface{}) {
		ctx.Log().Info("api_static")
		return
	})

}

//测试 WithArchive配置是否正常
//启动服务 ./static_prefix run
//浏览器访问 http://localhost:8070/prefixexts.htm [返回200]
//浏览器访问 http://localhost:8070/prefixexts.html [返回200]
//浏览器访问 http://localhost:8070/prefixfirst.html [返回200]
//浏览器访问 http://localhost:8070/prefixa.html [404 找不到文件:../root/a.html ]
func main() {
	app.Start()
}
