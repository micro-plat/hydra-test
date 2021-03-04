package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/auth/jwt"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("auth_jwt"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

func init() {
	secert := "12345678"
	hbulder := hydra.Conf.API("8080")
	hbulder.Jwt(jwt.WithEnable(), jwt.WithHeader(), jwt.WithAuthURL("https://www.baidu.com"), jwt.WithSecret(secert), jwt.WithName("__jwt_"), jwt.WithMode(jwt.ModeHS512), jwt.WithExcludes("/api/exclude"))

	app.API("/api", func(ctx hydra.IContext) (r interface{}) {
		ctx.Log().Info("api")
		ctx.Log().Info("user_auth:", ctx.User().Auth().Request())
		return
	})
}

//测试jwt配置 跳转url
//启动服务 ./jwt-authurl run
//浏览器访问 /api [__jwt_值为空,返回错误码302,跳转至设置好的URL]
func main() {
	app.Start()
}
