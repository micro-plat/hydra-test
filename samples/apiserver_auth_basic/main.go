package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/auth/basic"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("auth_basic"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

func init() {
	hydra.Conf.API(":8080").Basic(basic.WithExcludes("/api/exclude"), basic.WithEnable(), basic.WithUP("user", "pwd"))
	app.API("/api", func(ctx hydra.IContext) (r interface{}) {
		ctx.Log().Info("api")
		return
	})
	app.API("/api/exclude", func(ctx hydra.IContext) (r interface{}) {
		ctx.Log().Info("api-exclude")
		return
	})
}

//测试basic组件，验证用户名密码正确、错误时服务响应内容
//启动服务
//PostMan 访问 /api/exclude  [正常]
//PostMan 访问 /api [返回错误码401]
//PostMan 请求header设置["Authorization":"err_auth")] 访问 /api [返回错误码401]
//PostMan 请求header设置["Authorization":"Basic " + base64.Encode("user:pwd")] 访问 /api [返回200]

//使用谷歌浏览器(版本 86.0.4240.198) 访问 /api/exclude  [正常,返回200]
//使用谷歌浏览器(版本 86.0.4240.198) 访问 /api          [返回错误码401,弹出登录框]
//在登录输入框中输入err_user和serr_pwd 访问/api  [返回错误码401,弹出登录框]
//在登录输入框中输入user和pwd 访问/api  [正常,返回200]
func main() {
	app.Start()
}
