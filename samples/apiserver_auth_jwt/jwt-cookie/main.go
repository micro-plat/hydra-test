package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/auth/jwt"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydra_test"),
	hydra.WithSystemName("auth_jwt"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

func init() {
	hydra.Conf.API(":8080").Jwt(
		jwt.WithEnable(), jwt.WithCookie(), jwt.WithSecret("123456"),
		jwt.WithName("__jwt_"), jwt.WithMode(jwt.ModeHS512), jwt.WithExcludes("/api/getjwt"))

	app.API("/api/", func(ctx hydra.IContext) (r interface{}) {
		ctx.Log().Info("api")
		ctx.Log().Info("user_auth:", ctx.User().Auth().Request())
		return
	})

	app.API("/api/getjwt", func(ctx hydra.IContext) (r interface{}) {
		ctx.Log().Info("api-getjwt")
		ctx.User().Auth().Response("jwt_data")
		return
	})
}

//测试jwt参数保存在cookie中是否正确工作、设置用户信息，登录成功后获取用户信息、更新用户信息，及自动自动延时等
//启动服务
//cookie为空  访问 /api  [返回错误码401]
//cookie设置__jwt_值为错误验证串,访问 /api  [返回错误码403]
//访问 /api/getjwt  [正常],并且从响应的cookie中获取到jwt的正确的值
//cookie设置__jwt_值为正确验证串 访问 /api [返回200]  超过2分钟,再次访问 [Token is expired,返回403]

//访问 /api/getjwt 获取一个新的jwt
//使用新的jwt 访问/api,查看用户认证信息,查看响应的cookie中jwt的过期时间是否延长

//通过秘钥123456,加密模式的HS512,在线生成新的jwt
//使用新的jwt 访问/api,查看用户认证信息是否更新
func main() {
	app.Start()
}
