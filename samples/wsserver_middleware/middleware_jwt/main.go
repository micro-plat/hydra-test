package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/auth/jwt"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.WS),
	hydra.WithPlatName("hydra_test"),
	hydra.WithSystemName("ws_jwt"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

func init() {
	hydra.Conf.WS(":8080").Jwt(
		jwt.WithEnable(), jwt.WithHeader(), jwt.WithSecret("123456"),
		jwt.WithName("__jwt_"), jwt.WithMode(jwt.ModeHS512), jwt.WithExcludes("/ws/getjwt"))

	app.WS("/ws", func(ctx hydra.IContext) (r interface{}) {
		ctx.Log().Info("api")
		ctx.Log().Info("user_auth:", ctx.User().Auth().Request())
		return
	})

	app.WS("/ws/getjwt", func(ctx hydra.IContext) (r interface{}) {
		ctx.Log().Info("api-getjwt")
		ctx.User().Auth().Response("jwt_data")
		return
	})
}

//测试wsserver jwt中间件工作
//启动服务 ./middleware_jwt run
//发送数据 {"service":"/ws"} [__jwt_为空,返回401].
//发送数据 {"service":"/ws/getjwt"} [正常],并且从响应的header中获取到jwt的正确的值
//header设置__jwt_值为正确验证串 访问 /api [返回200]  超过2分钟,再次访问 [Token is expired,返回403]

//访问 /api/getjwt 获取一个新的jwt
//使用新的jwt 访问/api,查看用户认证信息,查看响应的header中jwt的过期时间是否延长

//通过秘钥123456,加密模式的HS512,在线生成新的jwt
//使用新的jwt 访问/api,查看用户认证信息是否更新
func main() {
	app.Start()
}
