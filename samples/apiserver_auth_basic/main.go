package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/auth/basic"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydra_test"),
	hydra.WithSystemName("auth_basic"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("zk://192.168.0.101"),
)

func init() {
	hydra.Conf.API(":8080").Basic(basic.WithExcludes("/api/exclude"), basic.WithEnable(), basic.WithUP("user", "pwd"))
	app.API("/api/", func(ctx context.IContext) (r interface{}) {
		ctx.Log().Info("api")
		return
	})
	app.API("/api/exclude", func(ctx context.IContext) (r interface{}) {
		ctx.Log().Info("api-exclude")
		return
	})
}

//启动服务
//访问 /api/exclude  [正常]
//访问 /api [返回错误码401]
//请求header设置["Authorization":"err_auth")] 访问 /api [返回错误码401]
//请求header设置["Authorization":"Basic " + base64.Encode("user:pwd")] 访问 /api [返回200]
func main() {
	app.Start()
}
