package main

import (
	"github.com/micro-plat/hydra"
	_ "github.com/micro-plat/hydra/components/caches/cache/redis"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var funcTrace func(ctx context.IContext) (r interface{}) = func(ctx context.IContext) (r interface{}) {

	return "success"
}

var app = hydra.NewApp(
	hydra.WithDebug(),
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("apiserverresponse"),
	hydra.WithClusterName("taosytest"),
	hydra.WithRegistry("zk://192.168.0.101"),
)

func init() {
	hydra.Conf.API(":8070", api.WithTrace())
	app.API("/hydratest/apiserver/response", funcTrace)
}

// apiserver-request-get请求demo

//1.1 安装程序 sudo ./servertrace01 conf install -cover
//1.2 使用 ./servertrace01 run -t cpu
//1.3 调用接口：http://192.168.5.94:8070/hydratest/apiserver/trace 判定配置是否正确
/*
1.路由不设置编码格式-Get-gbk-带中文和特殊符号正确编码数据
2.路由不设置编码格式-Get-gbk-带中文和特殊符号错误编码数据
3.路由不设置编码格式-Get-gb2312-带中文和特殊符号正确编码数据
4.路由不设置编码格式-Get-gb2312-带中文和特殊符号错误编码数据
3.路由不设置编码格式-Get-utf8-带中文和特殊符号正确编码数据
4.路由不设置编码格式-Get-utf8-带中文和特殊符号错误编码数据

1.路由设置正确编码格式-Get-gbk-带中文和特殊符号正确编码数据
2.路由设置正确编码格式-Get-gbk-带中文和特殊符号错误编码数据
3.路由设置正确编码格式-Get-gb2312-带中文和特殊符号正确编码数据
4.路由设置正确编码格式-Get-gb2312-带中文和特殊符号错误编码数据
3.路由设置正确编码格式-Get-utf8-带中文和特殊符号正确编码数据
4.路由设置正确编码格式-Get-utf8-带中文和特殊符号错误编码数据

1.路由设置错误编码格式-Get-gbk-带中文和特殊符号正确编码数据
2.路由设置错误编码格式-Get-gbk-带中文和特殊符号错误编码数据
3.路由设置错误编码格式-Get-gb2312-带中文和特殊符号正确编码数据
4.路由设置错误编码格式-Get-gb2312-带中文和特殊符号错误编码数据
3.路由设置错误编码格式-Get-utf8-带中文和特殊符号正确编码数据
4.路由设置错误编码格式-Get-utf8-带中文和特殊符号错误编码数据
*/

func main() {
	app.Start()
}
