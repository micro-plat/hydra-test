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
	hydra.WithRegistry("zk://192.168.0.101"),
)

func init() {
	hydra.Conf.Vars().RPC("rpc")
	hydra.Conf.API(":8070").Ras(ras.WithAuths(
		ras.New("/single/hydra/newversion/md5/auth@authserver.sas_debug",
			ras.WithRequest("/hydratest/apiserverras/test2"), //需要验证的路由
			// ras.WithRequired("test"),                        //设置必传字段
			// ras.WithSignAlias("newsign"),                    //设置新的签名字段名（默认为：sign）
			// ras.WithTimestampAlias("newtimestamp"),          //设置新的时间戳字段名（默认为：timestamp）
			// ras.WithUIDAlias("neweuid"), //设置新的签名配置id字段名（默认为：euid）
			// ras.WithDecryptName(""),                         //设置需要解密的字段（不设默认：全部）
			// ras.WithCheckTimestamp(true),   //设置是否需要时间戳校验（不设默认：true）
			// ras.WithParam("ext", "123456"), //设置扩展参数
			// ras.WithAuthDisable(), //该授权配置是否启用
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
	))
	app.API("/hydratest/apiserverras/test1", funcAPI1) //不加入签名验证,可直接返回成功
	app.API("/hydratest/apiserverras/test2", funcAPI2)
}

// apiserver_ras ras签名验证默认配置测试demo
//1.1 使用 ./rasserver02 conf install -cover
//1.1 使用 ./rasserver02 run

//1.2 调用不验签接口：http://localhost:8070/hydratest/apiserverras/test1  直接返回成功
//1.2 调用验签接口，sign不存在：http://localhost:8070/hydratest/apiserverras/test2?timestamp=131214152&euid=121213  403/远程认证失败:"{\"err\":\"sign值不能为空\"}"
//1.2 调用验签接口，timestamp不存在：http://localhost:8070/hydratest/apiserverras/test2?sign=131214152&euid=121213  返回参数错误
//1.2 调用验签接口，euid不存在：http://localhost:8070/hydratest/apiserverras/test  返回参数错误
//1.2 调用验签接口，延迟请求：http://localhost:8070/hydratest/apiserverras/test  请求过期
//1.2 调用验签接口，正常请求：http://localhost:8070/hydratest/apiserverras/test  返回成功
func main() {
	app.Start()
}

var funcAPI1 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver_ras ras签名验证默认配置测试demo")
	return "success"
}

var funcAPI2 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver_ras ras签名验证默认配置测试demo")
	return "success"
}

// var funcAPI2 = func(ctx hydra.IContext) (r interface{}) {
// 	ctx.Log().Info("apiserver_ras ras签名验证默认配置测试demo")
// 	input, err := ctx.Request().GetMap()
// 	if err != nil {
// 		ctx.Log().Errorf("参数获取失败:", err)
// 		ctx.Response().Abort(500, "参数获取失败")
// 		return
// 	}

// 	bt, _ := json.Marshal(input)
// 	content, status, err := components.Def.HTTP().GetRegularClient().Post("http://localhost:8070/hydratest/apiserverras/test1", string(bt))
// 	if err != nil {
// 		ctx.Log().Errorf("远程调用异常:", err)
// 		ctx.Response().Abort(500, "远程调用异常")
// 		return
// 	}

// 	ctx.Response().Abort(status, content)
// 	return
// }
