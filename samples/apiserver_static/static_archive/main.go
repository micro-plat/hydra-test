package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/static"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydra_test"),
	hydra.WithSystemName("static_archive"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

func init() {
	hydra.Conf.API(":8070").Static(
		static.WithArchive("../root/archive"), static.WithRoot("../root"), static.WithEnable())

	app.API("/api", func(ctx hydra.IContext) (r interface{}) {
		ctx.Log().Info("api_static")
		return
	})

}

//测试 WithArchive配置是否正常

//启动服务 ./static_archive run
//浏览器访问 http://localhost:8070/archive.txt 程序对../root/archive.zip进行解压至当前临时目录 [返回200]

//将WithArchive的配置改为../root/archive.tar,重启服务
//浏览器访问 http://localhost:8070/archive.txt 程序对../root/archive.tar进行解压至当前临时目录 [返回200]

//将WithArchive的配置改为../root/archive.tar.gz,重启服务
//浏览器访问 http://localhost:8070/archive.txt 程序对../root/archive.tar.gz进行解压至当前临时目录 [返回200]
func main() {
	app.Start()
}
