package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/static"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydra_test"),
	hydra.WithSystemName("static_rewrite"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

func init() {
	hydra.Conf.API(":8070").Static(
		static.WithRewriters("/", "/indextest.htm", "/defaulttest.html"), static.WithRoot("../root"), static.WithFirstPage("first.html"), static.WithEnable())

	app.API("/api", func(ctx hydra.IContext) (r interface{}) {
		ctx.Log().Info("api_static")
		return
	})

}

//测试static rewriters,firstPage配置是否正常
//设置静态根目录 ./root ,重写页面为first.html
//启动服务  ./static_rewrite run
//浏览器访问 http://localhost:8070/ 进行重写 [返回200]
//浏览器访问 http://localhost:8070/indextest.htm 进行重写 [返回200]
//浏览器访问 http://localhost:8070/defaulttest.html 进行重写 [返回200]
func main() {
	app.Start()
}
