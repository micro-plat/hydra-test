package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/auth/ras"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithDebug(),
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("apiserverras"),
	hydra.WithClusterName("taosytest"),
	hydra.WithRegistry("lm://."),
)

func init() {
	hydra.Conf.API(":8070").Ras(ras.WithDisable(),
		ras.WithAuths(
			ras.New("/single/hydra/newversion/md5/auth@authserver.sas_debug",
				ras.WithRequest("/hydratest/apiserverras/test"), //需要验证的路由
				ras.WithRequired("test"),                        //设置必传字段
				ras.WithSignAlias("newsign"),                    //设置新的签名字段名（默认为：sign）
				ras.WithTimestampAlias("newtimestamp"),          //设置新的时间戳字段名（默认为：timestamp）
				ras.WithUIDAlias("neweuid"),                     //设置新的签名配置id字段名（默认为：euid）
				// ras.WithDecryptName(""),                         //设置需要解密的字段（不设默认：全部）
				ras.WithCheckTimestamp(true),   //设置是否需要时间戳校验（不设默认：true）
				ras.WithParam("ext", "123456"), //设置扩展参数
				ras.WithAuthDisable(),          //该授权配置是否启用
				ras.WithConnect(
					ras.WithConnectChar("kv", "&"), //设置加密串拼接模式和链接字符 如：kv  &
					ras.WithConnectSortByData(),    //只排序数据字段，不排序secrect   排序类型 三选一
					// ras.WithConnectSortAll(),                    //排序所有字段，包括数据，secrect 排序类型 三选一
					// ras.WithConnectSortStatic("feld1", "feld2"), //使用指定的字段进行排序  排序类型 三选一
					ras.WithSecretConnect(
						// ras.WithSecretName("name", "kv"), //设置secrect的键名称  (不设默认：只拼接密钥串)
						ras.WithSecretHeadMode("&"), //设置secrect与数据串之间的拼接方式,并将secret串拼接到数据串的头部  密钥拼接类型三选一
						// ras.WithSecretTailMode("="),        //设置secrect与数据串之间的拼接方式，并将secret串拼接到数据串的尾部 密钥拼接类型三选一
						// ras.WithSecretHeadAndTailMode("*"), //设置secrect与数据串之间的拼接方式，并将secret串拼接到数据串的头部和尾部 密钥拼接类型三选一
					),
				),
			),
		),
	)
	app.API("/hydratest/apiserverras/test", funcAPI)

}

// apiserver_ras ras签名验证禁用测试demo
//1.1 使用 ./cronserver_cycle run

//1.2 调用接口：http://localhost:8070/hydratest/apiserverras/test  禁用情况下，无任何签名参数也可以正常返回
func main() {
	app.Start()
}

var funcAPI = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver_ras ras签名验证禁用测试demo")
	return "success"
}
