package main

import (
	"time"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/router"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("apiserverresponse"),
	hydra.WithClusterName("taosytest"),
	hydra.WithRegistry("lm://."),
)

func init() {
	hydra.Conf.API("8072")
	app.API("/hydratest/apiserver/response", funcResponse)
	app.API("/hydratest/apiserver/response1", funcResponse1)
	app.API("/hydratest/apiserver/response2", funcResponse2)
	app.API("/hydratest/apiserver/response3", funcResponse3, router.WithEncoding("gbk"))
	app.API("/hydratest/apiserver/response4", funcResponse4, router.WithEncoding("gb2312"))
	app.API("/hydratest/apiserver/response5", funcResponse5, router.WithEncoding("utf-8"))
	app.API("/hydratest/apiserver/response6", funcResponse6, router.WithEncoding("utf-8"))
	app.API("/hydratest/apiserver/response7", funcResponse7, router.WithEncoding("gbk"))
}

//apiserver_response-yaml 不同请求方式responce的数据编码demo

//1.2 使用 ./responceserver03 run

/*
1.1 请求路由无编码格式接口-Post-body-gbk-设置yaml头，返回struct         返回数据是gbk编码
1.2 请求路由无编码格式接口-Post-body-gb2312-设置yaml头，返回map     返回数据是gbk编码
1.3 请求路由无编码格式接口-Post-body-utf8-设置yaml头，数据yaml串        返回数据是gbk编码

2.1 请求路由有编码格式接口-Put-body-gbk-设置yaml头，数据json串         返回数据是gbk编码
2.2 请求路由有编码格式接口-Put-body-gb2312-设置yaml头，数据xml串      返回数据是gbk编码
2.3 请求路由有编码格式接口-Put-body-utf8-设置yaml头，数据为text        返回数据是gbk编码

3.1 头和路由都有编码-Get-body-gbk+utf8-不设置头，返回text文本         返回数据是utf8编码
3.2 头和路由都有编码-Get-body-utf8+gbk-不设置头，返回yaml文本          返回数据是gbk编码
*/

func main() {
	app.Start()
}

type ymal struct {
	Param1 string                 `yaml:"param1"`
	Param2 bool                   `yaml:"param2"`
	Param3 int32                  `yaml:"param3"`
	Param4 float32                `yaml:"param4"`
	Param5 []string               `yaml:"param5"`
	Param6 time.Time              `yaml:"param6" time_format:"2006/01/02 15:04:05"`
	Param7 map[string]interface{} `yaml:"param7"`
	Param8 []int                  `yaml:"param8"`
}

var defaultData = ymal{
	Param1: "34ddf#$*@大!@#$%^&*()_+~锅饭都是",
	Param2: true,
	Param3: 1024,
	Param4: 10.24,
	Param5: []string{"1", "2"},
	Param6: time.Now(),
	Param7: map[string]interface{}{"t1": 123, "t2": "sdfs@@###", "t3": 12.2},
	Param8: []int{1, 2},
}

var funcResponse = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver-response-yaml,路由不设置编码数据处理demo")
	ctx.Response().Header("Content-Type", "application/yaml")
	return &defaultData
}

var funcResponse1 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver-response-yaml,路由设置gbk编码数据处理demo")
	ctx.Response().Header("Content-Type", "application/yaml")
	input := map[string]interface{}{
		"param1": "34ddf#$*@大!@#$%^&*()_+~锅饭都是",
		"param2": true,
		"param3": 1024,
		"param4": 10.24,
		"param5": []string{"1", "2"},
		"param6": time.Now(),
		"param7": map[string]interface{}{"t1": 123, "t2": "sdfs@@###", "t3": 12.2},
		"param8": []int{1, 2},
	}
	return input
}

var funcResponse2 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver-response-yaml,路由设置gb2312编码数据处理demo")
	ctx.Response().Header("Content-Type", "application/yaml")
	return `param1: 34ddf#$*@大!@#$%^&*()_+~锅饭都是
	param2: true
	param3: 1024
	param4: 10.24
	param5:
	- "1"
	- "2"
	param6: 2020-12-09T10:59:37.979315945+08:00
	param7:
	  t1: 123
	  t2: sdfs@@###
	  t3: 12.2
	param8:
	- 1
	- 2`
}

//路由设置编码 gbk
var funcResponse3 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver-response-yaml,路由设置utf8编码数据处理demo")
	ctx.Response().Header("Content-Type", "application/yaml")
	return `{"param1":"34ddf#$*@大!@#$%^\u0026*()_+~锅饭都是","param2":true,"param3":1024,"param4":10.24,"param5":["1","2"],"param6":"2020-12-08T16:57:43.187670336+08:00","param7":{"t1":123,"t2":"sdfs@@###","t3":12.2},"param8":[1,2]}`
}

//路由设置编码 gb2312
var funcResponse4 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver-response-xml,路由设置utf8编码数据处理demo")
	ctx.Response().Header("Content-Type", "application/yaml")
	return `<xml><param1>34ddf#$*@大!@#$%^&amp;*()_+~锅饭都是</param1><param2>true</param2><param3>1024</param3><param4>10.24</param4><param5>1</param5><param5>2</param5><param6>2020-12-08T19:51:49.39880659+08:00</param6><param8>1</param8><param8>2</param8></xml>`
}

//路由设置编码 utf8
var funcResponse5 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver-response-yaml,路由设置utf8编码数据处理demo")
	ctx.Response().Header("Content-Type", "application/yaml")
	return "sdklfjdskjdfg三大类反击的上来看级三联书店开放564654654#&……%#kjhsdfkjh"
}

//路由设置编码 utf8
var funcResponse6 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver-response-yaml,路由设置utf8编码数据处理demo")
	return "sdklfjdskjdfg三大类反击的上来看级三联书店开放564654654#&……%#kjhsdfkjh"
}

//路由设置编码 gbk
var funcResponse7 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver-response-yaml,路由设置utf8编码数据处理demo")
	return `param1: 34ddf#$*@大!@#$%^&*()_+~锅饭都是
	param2: true
	param3: 1024
	param4: 10.24
	param5:
	- "1"
	- "2"
	param6: 2020-12-09T10:59:37.979315945+08:00
	param7:
	  t1: 123
	  t2: sdfs@@###
	  t3: 12.2
	param8:
	- 1
	- 2`
}
