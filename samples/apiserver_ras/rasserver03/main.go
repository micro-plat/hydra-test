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
	hydra.Conf.API(":8070").Ras(ras.WithAuths(
		ras.New("/single/hydra/newversion/md5/auth@authserver.sas_debug",
			ras.WithRequest("/hydratest/apiserverras/test"), //需要验证的路由
			ras.WithRequired("test1"),                       //设置必传字段
			ras.WithSignAlias("newsign"),                    //设置新的签名字段名（默认为：sign）
			ras.WithTimestampAlias("newtimestamp"),          //设置新的时间戳字段名（默认为：timestamp）
			ras.WithUIDAlias("neweuid"),                     //设置新的签名配置id字段名（默认为：euid）
			ras.WithDecryptName("test1"),                    //设置需要解密的字段（不设默认：全部）
			// ras.WithCheckTimestamp(true),   //设置是否需要时间戳校验（不设默认：true）
			ras.WithParam("ext", "123456"), //设置扩展参数
			ras.WithAuthDisable(),          //该授权配置是否启用
			ras.WithConnect(
				ras.WithConnectChar("kv", "*"), //设置加密串拼接模式和链接字符 如：kv  &
				// ras.WithConnectSortByData(),    //只排序数据字段，不排序secrect   排序类型三选一
				ras.WithConnectSortAll(), //排序所有字段，包括数据，secrect 排序类型三选一
				// ras.WithConnectSortStatic("feld1", "feld2"), //使用指定的字段进行排序  排序类型三选一
				ras.WithSecretConnect(
					ras.WithSecretName("sect", "kv"), //设置secrect的键名称  (不设默认：只拼接密钥串)
					// ras.WithSecretHeadMode("&"),      //设置secrect与数据串之间的拼接方式,并将secret串拼接到数据串的头部  密钥拼接类型三选一
					ras.WithSecretTailMode("="), //设置secrect与数据串之间的拼接方式，并将secret串拼接到数据串的尾部 密钥拼接类型三选一
					// ras.WithSecretHeadAndTailMode("*"), //设置secrect与数据串之间的拼接方式，并将secret串拼接到数据串的头部和尾部 密钥拼接类型三选一
				),
			),
		),
		ras.New("/single/hydra/newversion/md5/auth@authserver.sas_debug",
			ras.WithRequest("/hydratest/apiserverras/test1"), //需要验证的路由
			ras.WithRequired("test1"),                        //设置必传字段
			ras.WithSignAlias("newsign"),                     //设置新的签名字段名（默认为：sign）
			ras.WithTimestampAlias("newtimestamp"),           //设置新的时间戳字段名（默认为：timestamp）
			ras.WithUIDAlias("neweuid"),                      //设置新的签名配置id字段名（默认为：euid）
			ras.WithDecryptName("test1"),                     //设置需要解密的字段（不设默认：全部）
			// ras.WithCheckTimestamp(true),   //设置是否需要时间戳校验（不设默认：true）
			ras.WithParam("ext", "123456"), //设置扩展参数
			// ras.WithAuthDisable(),          //该授权配置是否启用
			ras.WithConnect(
				ras.WithConnectChar("kv", "*"), //设置加密串拼接模式和链接字符 如：kv  &
				// ras.WithConnectSortByData(),    //只排序数据字段，不排序secrect   排序类型三选一
				ras.WithConnectSortAll(), //排序所有字段，包括数据，secrect 排序类型三选一
				// ras.WithConnectSortStatic("feld1", "feld2"), //使用指定的字段进行排序  排序类型三选一
				ras.WithSecretConnect(
					ras.WithSecretName("sect", "kv"), //设置secrect的键名称  (不设默认：只拼接密钥串)
					// ras.WithSecretHeadMode("&"),      //设置secrect与数据串之间的拼接方式,并将secret串拼接到数据串的头部  密钥拼接类型三选一
					ras.WithSecretTailMode("="), //设置secrect与数据串之间的拼接方式，并将secret串拼接到数据串的尾部 密钥拼接类型三选一
					// ras.WithSecretHeadAndTailMode("*"), //设置secrect与数据串之间的拼接方式，并将secret串拼接到数据串的头部和尾部 密钥拼接类型三选一
				),
			),
		),
	))
	app.API("/hydratest/apiserverras/test", funcAPI) //不加入签名验证，通过该接口发起间接调用
	app.API("/hydratest/apiserverras/test1", funcAPI)

}

// apiserver_ras ras签名验证自定义配置测试demo1
//1.1 使用 ./rasserver03 conf install -cover
//1.1 使用 ./rasserver03 run

//1.2 调用签名配置被禁用接口：http://localhost:8070/hydratest/apiserverras/test  直接返回成功
//1.2 调用验签接口，newsign不存在：http://localhost:8070/hydratest/apiserverras/test1?test1=1&neweuid=test1  没有检测新的sign字段
//1.2 调用验签接口，newtimestamp不存在：http://localhost:8070/hydratest/apiserverras/test1?test1=1&neweuid=test1&sign=123456  返回参数错误
//1.2 调用验签接口，neweuid不存在：http://localhost:8070/hydratest/apiserverras/test1?test1=1&sign=123456  远程认证失败:"{\"err\":\"neweuid值不能为空\"}",(400)
//1.2 调用验签接口，必传字段不存在：http://localhost:8070/hydratest/apiserverras/test1?neweuid=test1  远程认证失败:"{\"err\":\"test1值不能为空\"}",(400)
//1.2 调用验签接口，按照配置进行错误签名请求：http://localhost:8070/hydratest/apiserverras/test1  返回签名失败
//1.2 调用验签接口，正常签名延迟请求：http://localhost:8070/hydratest/apiserverras/test1  请求过期
//1.2 调用验签接口，按照配置进行正常签名请求：http://localhost:8070/hydratest/apiserverras/test1  返回成功
func main() {
	app.Start()
}

var funcAPI = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver_ras ras签名验证默认配置测试demo")
	return "success"
}
