package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/auth/ras"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.WS),
	hydra.WithPlatName("hydra_test"),
	hydra.WithSystemName("ws_ras"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("zk://192.168.0.101"),
)

func init() {
	hydra.Conf.Vars().RPC("rpc")
	hydra.Conf.WS(":8080").Ras(ras.WithAuths(
		ras.New("/single/hydra/newversion/md5/auth@authserver.sas_debug",
			ras.WithRequest("/test2"), //需要验证的路由
			ras.WithConnect(
				ras.WithConnectChar("kv", "&"), //设置加密串拼接模式和链接字符 如：kv  &
				ras.WithConnectSortByData(),    //只排序数据字段，不排序secrect   排序类型 三选一
				ras.WithSecretConnect(
					ras.WithSecretHeadMode("&"), //设置secrect与数据串之间的拼接方式,并将secret串拼接到数据串的头部  密钥拼接类型三选一
				),
			),
		),
	))

	//不加入签名验证,可直接返回成功
	app.WS("/test1", func(ctx hydra.IContext) (r interface{}) {
		ctx.Log().Info("apiserver_ras ras签名验证默认配置测试demo")
		return "success"
	})

	app.WS("/test2", func(ctx hydra.IContext) (r interface{}) {
		ctx.Log().Info("apiserver_ras ras签名验证默认配置测试demo")
		return "success"
	})
}

//测试 wsserver的ras工作
//启动服务./middleware_ras run
//建立连接
//发送数据 {"service":"/test1"} [返回200, 不进行验证]
//发送数据 {"service":"/test2"} [返回403,euid值不能为空, 中间件生效]
func main() {
	app.Start()
}
