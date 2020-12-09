package main

import (
	xmlM "encoding/xml"
	"time"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/router"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithDebug(),
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("apiserverresponse"),
	hydra.WithClusterName("taosytest"),
	hydra.WithRegistry("zk://192.168.0.101"),
)

func init() {
	hydra.Conf.API(":8070")
	app.API("/hydratest/apiserver/response", funcResponse)
	app.API("/hydratest/apiserver/response1", funcResponse1)
	app.API("/hydratest/apiserver/response2", funcResponse2)
	app.API("/hydratest/apiserver/response3", funcResponse3, router.WithEncoding("gbk"))
	app.API("/hydratest/apiserver/response4", funcResponse4, router.WithEncoding("gb2312"))
	app.API("/hydratest/apiserver/response5", funcResponse5, router.WithEncoding("utf-8"))
	app.API("/hydratest/apiserver/response6", funcResponse6, router.WithEncoding("utf-8"))
	app.API("/hydratest/apiserver/response7", funcResponse7, router.WithEncoding("gbk"))
}

//apiserver_response-xml 不同请求方式responce的数据编码demo
//1.2 使用 ./apiserver_response run

/*
1.1 请求路由无编码格式接口-Post-body-gbk-设置xml头，返回struct         返回数据是gbk编码
1.2 请求路由无编码格式接口-Post-body-gb2312-设置xml头，返回map     返回数据是gbk编码
1.3 请求路由无编码格式接口-Post-body-utf8-设置xml头，数据xml串        返回数据是gbk编码

2.1 请求路由有编码格式接口-Put-body-gbk-设置xml头，数据json串         返回数据是gbk编码
2.2 请求路由有编码格式接口-Put-body-gb2312-设置xml头，数据yaml串      返回数据是gbk编码
2.3 请求路由有编码格式接口-Put-body-utf8-设置xml头，数据为text        返回数据是gbk编码

3.1 头和路由都有编码-Get-body-gbk+utf8-不设置头，返回text文本         返回数据是utf8编码
3.2 头和路由都有编码-Get-body-utf8+gbk-不设置头，返回xml文本          返回数据是gbk编码
*/

func main() {
	app.Start()
}

type xml struct {
	Param1 string    `xml:"param1"`
	Param2 bool      `xml:"param2"`
	Param3 int32     `xml:"param3"`
	Param4 float32   `xml:"param4"`
	Param5 []string  `xml:"param5"`
	Param6 time.Time `xml:"param6" time_format:"2006/01/02 15:04:05"`
	Param8 []int     `xml:"param8"`
}

var defaultData = xml{
	Param1: "34ddf#$*@大!@#$%^&*()_+~锅饭都是",
	Param2: true,
	Param3: 1024,
	Param4: 10.24,
	Param5: []string{"1", "2"},
	Param6: time.Now(),
	Param8: []int{1, 2},
}

var funcResponse = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver-response-xml,路由不设置编码数据处理demo")
	ctx.Response().Header("Content-Type", "application/xml")
	return &defaultData
}

var funcResponse1 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver-response-xml,路由设置gbk编码数据处理demo")
	ctx.Response().Header("Content-Type", "application/xml")
	input := map[string]interface{}{
		"param1": "34ddf#$*@大!@#$%^&*()_+~锅饭都是",
		"param2": true,
		"param3": 1024,
		"param4": 10.24,
		"param5": []string{"1", "2"},
		"param6": time.Now(),
		"param8": []int{1, 2},
	}
	return input
}

var funcResponse2 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver-response-xml,路由设置gb2312编码数据处理demo")
	ctx.Response().Header("Content-Type", "application/xml")
	bt, _ := xmlM.Marshal(&defaultData)
	return string(bt)
}

//路由设置编码 gbk
var funcResponse3 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver-response-xml,路由设置utf8编码数据处理demo")
	ctx.Response().Header("Content-Type", "application/xml")
	return `{"param1":"34ddf#$*@大!@#$%^\u0026*()_+~锅饭都是","param2":true,"param3":1024,"param4":10.24,"param5":["1","2"],"param6":"2020-12-08T16:57:43.187670336+08:00","param8":[1,2]}`
}

//路由设置编码 gb2312
var funcResponse4 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver-response-xml,路由设置utf8编码数据处理demo")
	ctx.Response().Header("Content-Type", "application/xml")
	str := `param1: 34ddf#$*@大!@#$%^&*()_+~锅饭都是
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
	return str
}

//路由设置编码 utf8
var funcResponse5 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver-response-xml,路由设置utf8编码数据处理demo")
	ctx.Response().Header("Content-Type", "application/xml")
	return "sdklfjdskjdfg三大类反击的上来看级三联书店开放564654654#&……%#kjhsdfkjh"
}

//路由设置编码 utf8
var funcResponse6 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver-response-xml,路由设置utf8编码数据处理demo")
	return "sdklfjdskjdfg三大类反击的上来看级三联书店开放564654654#&……%#kjhsdfkjh"
}

//路由设置编码 gbk
var funcResponse7 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver-response-xml,路由设置utf8编码数据处理demo")
	bt, _ := xmlM.Marshal(&defaultData)
	return string(bt)
}
