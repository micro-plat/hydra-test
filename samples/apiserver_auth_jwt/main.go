package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/auth/jwt"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/hydra/servers/http"
	xjwt "github.com/micro-plat/lib4go/security/jwt"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydra_test"),
	hydra.WithSystemName("auth_jwt"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("zk://192.168.0.101"),
)

func init() {
	secert := "123456"
	hbulder := hydra.Conf.API(":8080")
	//jwt.WithAuthURL("https://www.baidu.com") //设置跳转url
	//jwt.WithCookie() //设置jwt保存在cookie中
	hbulder.Jwt(jwt.WithEnable(), jwt.WithHeader(), jwt.WithSecret(secert), jwt.WithName("__jwt_"), jwt.WithMode(jwt.ModeHS512), jwt.WithExcludes("/api/exclude"))
	app.API("/api/", func(ctx context.IContext) (r interface{}) {
		ctx.Log().Info("api")
		ctx.Log().Info("user_auth:", ctx.User().Auth().Request())
		return
	})
	app.API("/api/exclude", func(ctx context.IContext) (r interface{}) {
		ctx.Log().Info("api-exclude", secert)
		jwt, err := xjwt.Encrypt(secert, jwt.ModeHS512, "jwt_", 120)
		if err != nil {
			return err
		}
		ctx.Response().Header("__jwt_", jwt)
		return
	})
}

//试jwt组件，包括且不限于：jwt参数是否正确工作、登录时设置用户信息，登录成功后获取用户信息、更新用户信息，及自动自动延时等
//启动服务
//访问 /api [__jwt_值为空,返回错误码302,跳转至设置好的URL]
//不设置跳转url, 访问api [__jwt_值为空,返回错误码401]
//不设置跳转url, 请求header设置__jwt_值错误的jwt验证串,访问 /api [返回错误码403]
//访问 /api/exclude  [正常],并且从响应的header中获取到jwt的正确的值
//访问 /api 在header中设置的__jwt_为正确的值 [返回200]  超过2分钟,再次访问 [Token is expired,返回403]

//访问 /api/exclude 获取一个新的jwt
//使用新的jwt 访问/api,查看用户认证信息,查看响应的header中jwt的过期时间是否延长

//通过秘钥123456,加密模式的HS512,生成新的jwt
//使用新的jwt 访问/api,查看用户认证信息是否更新
func main() {
	app.Start()
}
