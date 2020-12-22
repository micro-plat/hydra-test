package main

import (
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
	hydra.WithRegistry("lm://./"),
)

func init() {
	hydra.Conf.API(":8072")
	app.API("/hydratest/apiserver/request/nil", funcRequest)
	app.API("/hydratest/apiserver/request/gbk", funcRequest1, router.WithEncoding("gbk"))
	app.API("/hydratest/apiserver/request/gb2312", funcRequest2, router.WithEncoding("gb2312"))
	app.API("/hydratest/apiserver/request/utf8", funcRequest3, router.WithEncoding("utf-8"))
}

// apiserver-request-xml数据处理demo

//1.1 使用 ./requestserver02 run
//1.2 请求的数据：<xml><param1>34ddf#$*@大!@#$%^&amp;*()_+~锅饭都是</param1><param2>true</param2><param3>1024</param3><param4>10.24</param4><param5>1</param5><param5>2</param5><param6>2020-12-08T19:51:49.39880659+08:00</param6><param7><t1>123</t1><t2>sdfs@@###</t2><t3>12.2</t3></param7><param8>1</param8><param8>2</param8></xml>
/*
通过postman构建如下请求：
1.1 请求路由无编码格式接口-Post-body-gbk-与头编码相同         返回正确的数据
1.2 请求路由无编码格式接口-Post-body-gbk-与头编码不同         返回错误的数据
1.3 请求路由无编码格式接口-Post-body-gb2312-与头编码相同      返回正确的数据
1.4 请求路由无编码格式接口-Post-body-gb2312-与头编码不同      返回错误的数据
1.5 请求路由无编码格式接口-Post-body-utf8-与头编码相同        返回正确的数据
1.6 请求路由无编码格式接口-Post-body-utf8-与头编码不同        返回错误的数据
1.7 请求路由无编码格式接口-Get-body-gbk-与头编码相同         返回正确的数据
1.8 请求路由无编码格式接口-Get-body-gbk-与头编码不同         返回错误的数据
1.9 请求路由无编码格式接口-Get-body-gb2312-与头编码相同      返回正确的数据
1.10 请求路由无编码格式接口-Get-body-gb2312-与头编码不同      返回错误的数据
1.11 请求路由无编码格式接口-Get-body-utf8-与头编码相同        返回正确的数据
1.12 请求路由无编码格式接口-Get-body-utf8-与头编码不同        返回错误的数据

2.1 请求路由有编码格式接口-Post-body-gbk-与路由编码相同         返回正确的数据
2.2 请求路由有编码格式接口-Post-body-gbk-与路由编码不同         返回错误的数据
2.3 请求路由有编码格式接口-Post-body-gb2312-与路由编码相同      返回正确的数据
2.4 请求路由有编码格式接口-Post-body-gb2312-与路由编码不同      返回错误的数据
2.5 请求路由有编码格式接口-Post-body-utf8-与路由编码相同        返回正确的数据
2.6 请求路由有编码格式接口-Post-body-utf8-与路由编码不同        返回错误的数据
2.7 请求路由有编码格式接口-Get-body-gbk-与路由编码相同         返回正确的数据
2.8 请求路由有编码格式接口-Get-body-gbk-与路由编码不同         返回错误的数据
2.9 请求路由有编码格式接口-Get-body-gb2312-与路由编码相同      返回正确的数据
2.10 请求路由有编码格式接口-Get-body-gb2312-与路由编码不同      返回错误的数据
2.11 请求路由有编码格式接口-Get-body-utf8-与路由编码相同        返回正确的数据
2.12 请求路由有编码格式接口-Get-body-utf8-与路由编码不同        返回错误的数据

3.1 头和路由都有编码-编码相同-Post-body-gbk-编码相同        返回正确的数据
3.2 头和路由都有编码-编码相同-Post-body-gbk-编码不同        返回错误的数据
3.3 头和路由都有编码-编码相同-Post-body-gb2312-编码相同     返回正确的数据
3.4 头和路由都有编码-编码相同-Post-body-gb2312-编码不同     返回错误的数据
3.5 头和路由都有编码-编码相同-Post-body-utf8-编码相同       返回正确的数据
3.6 头和路由都有编码-编码相同-Post-body-utf8-编码不同       返回错误的数据
3.7 头和路由都有编码-编码相同-Get-body-gbk-编码相同        返回正确的数据
3.8 头和路由都有编码-编码相同-Get-body-gbk-编码不同        返回错误的数据
3.9 头和路由都有编码-编码相同-Get-body-gb2312-编码相同     返回正确的数据
3.10 头和路由都有编码-编码相同-Get-body-gb2312-编码不同     返回错误的数据
3.11 头和路由都有编码-编码相同-Get-body-utf8-编码相同       返回正确的数据
3.12 头和路由都有编码-编码相同-Get-body-utf8-编码不同       返回错误的数据

4.1 头和路由都有编码-编码不同-Post-body-gbk-utf8-与头编码相同        无法获取数据
4.2 头和路由都有编码-编码不同-Post-body-gbk-utf8-与路由编码相同      返回正确的数据
4.3 头和路由都有编码-编码不同-Get-body-gbk-utf8-与头编码相同        无法获取数据
4.4 头和路由都有编码-编码不同-Get-body-gbk-utf8-与路由编码相同      返回正确的数据
*/

func main() {
	app.Start()
}

type xml struct {
	Param1 string                 `json:"param1"`
	Param2 bool                   `json:"param2"`
	Param3 int32                  `json:"param3"`
	Param4 float32                `json:"param4"`
	Param5 []string               `json:"param5"`
	Param6 time.Time              `json:"param6" time_format:"2006/01/02 15:04:05"`
	Param7 map[string]interface{} `json:"param7"`
	Param8 []int                  `json:"param8"`
}

var defaultData = xml{
	Param1: "34ddf#$*@大!@#$%^&*()_+~锅饭都是",
	Param2: true,
	Param3: 1024,
	Param4: 10.24,
	Param5: []string{"1", "2"},
	Param6: time.Now(),
	Param7: map[string]interface{}{"t1": 123, "t2": "sdfs@@###", "t3": 12.2},
	Param8: []int{1, 2},
}

//路由不设置编码
var funcRequest = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver-request-json,路由不设置编码数据处理demo")
	if err := ctx.Request().Check(); err != nil {
		ctx.Log().Errorf("ctx.Request().Check()异常：%s", err)
	}
	param := &xml{}
	if err := ctx.Request().Bind(param); err != nil {
		ctx.Log().Errorf("ctx.Request().Bind()异常：%s", err)
	}
	ctx.Log().Info("----bind data:", param)
	body, err := ctx.Request().GetBody()
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetBody()异常：%s", err)
	}
	ctx.Log().Info("----body data:", string(body))
	bodyRaw, raw, err := ctx.Request().GetFullRaw()
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetFullRaw()异常：%s", err)
	}
	ctx.Log().Info("----GetFullRaw data:", string(bodyRaw), raw)

	tm, err := ctx.Request().GetDatetime("param6")
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetDatetime()异常：%s", err)
	}
	xmap := ctx.Request().GetMap()

	ctx.Log().Info("----GetMap()：", xmap)
	jsonD, err := ctx.Request().GetJSON("param7")
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetJSON()异常：%s", err)
	}
	ctx.Log().Info("----GetJSON data:", string(jsonD))
	ctx.Log().Info("----GetDatetime()：", tm)
	ctx.Log().Info("----Keys data:", ctx.Request().Keys())
	ctx.Log().Info("----Len data:", ctx.Request().Len())
	ctx.Log().Info("----GetArray data:", ctx.Request().GetArray("param5"))
	ctx.Log().Info("----GetInt32 data:", ctx.Request().GetInt32("param3"))
	ctx.Log().Info("----GetFloat32 data:", ctx.Request().GetFloat32("param4"))
	ctx.Log().Info("----GetBool data:", ctx.Request().GetBool("param2"))
	return "success"
}

//路由设置编码 gbk
var funcRequest1 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver-request-json,路由设置gbk编码数据处理demo")
	if err := ctx.Request().Check(); err != nil {
		ctx.Log().Errorf("ctx.Request().Check()异常：%s", err)
	}
	param := &xml{}
	if err := ctx.Request().Bind(param); err != nil {
		ctx.Log().Errorf("ctx.Request().Bind()异常：%s", err)
	}
	ctx.Log().Info("----bind data:", param)
	body, err := ctx.Request().GetBody()
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetBody()异常：%s", err)
	}
	ctx.Log().Info("----body data:", string(body))
	bodyRaw, raw, err := ctx.Request().GetFullRaw()
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetFullRaw()异常：%s", err)
	}
	ctx.Log().Info("----GetFullRaw data:", string(bodyRaw), raw)
	xmap := ctx.Request().GetMap()

	ctx.Log().Info("----GetMap()：", xmap)
	tm, err := ctx.Request().GetDatetime("param6")
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetDatetime()异常：%s", err)
	}
	ctx.Log().Info("----GetDatetime()：", tm)
	ctx.Log().Info("----Keys data:", ctx.Request().Keys())
	ctx.Log().Info("----Len data:", ctx.Request().Len())
	ctx.Log().Info("----GetArray data:", ctx.Request().GetArray("param5"))
	ctx.Log().Info("----GetInt32 data:", ctx.Request().GetInt32("param3"))
	ctx.Log().Info("----GetFloat32 data:", ctx.Request().GetFloat32("param4"))
	ctx.Log().Info("----GetBool data:", ctx.Request().GetBool("param2"))
	return "success"
}

//路由设置编码 gb2312
var funcRequest2 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver-request-json,路由设置gb2312编码数据处理demo")
	if err := ctx.Request().Check(); err != nil {
		ctx.Log().Errorf("ctx.Request().Check()异常：%s", err)
	}
	param := &xml{}
	if err := ctx.Request().Bind(param); err != nil {
		ctx.Log().Errorf("ctx.Request().Bind()异常：%s", err)
	}
	ctx.Log().Info("----bind data:", param)
	body, err := ctx.Request().GetBody()
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetBody()异常：%s", err)
	}
	ctx.Log().Info("----body data:", string(body))
	bodyRaw, raw, err := ctx.Request().GetFullRaw()
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetFullRaw()异常：%s", err)
	}
	ctx.Log().Info("----GetFullRaw data:", string(bodyRaw), raw)
	xmap := ctx.Request().GetMap()

	ctx.Log().Info("----GetMap()：", xmap)
	tm, err := ctx.Request().GetDatetime("param6")
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetDatetime()异常：%s", err)
	}
	ctx.Log().Info("----GetDatetime()：", tm)
	ctx.Log().Info("----Keys data:", ctx.Request().Keys())
	ctx.Log().Info("----Len data:", ctx.Request().Len())
	ctx.Log().Info("----GetArray data:", ctx.Request().GetArray("param5"))
	ctx.Log().Info("----GetInt32 data:", ctx.Request().GetInt32("param3"))
	ctx.Log().Info("----GetFloat32 data:", ctx.Request().GetFloat32("param4"))
	ctx.Log().Info("----GetBool data:", ctx.Request().GetBool("param2"))
	return "success"
}

//路由设置编码 utf8
var funcRequest3 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver-request-json,路由设置utf8编码数据处理demo")
	if err := ctx.Request().Check(); err != nil {
		ctx.Log().Errorf("ctx.Request().Check()异常：%s", err)
	}
	param := &xml{}
	if err := ctx.Request().Bind(param); err != nil {
		ctx.Log().Errorf("ctx.Request().Bind()异常：%s", err)
	}
	ctx.Log().Info("----bind data:", param)
	body, err := ctx.Request().GetBody()
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetBody()异常：%s", err)
	}
	ctx.Log().Info("----body data:", string(body))
	bodyRaw, raw, err := ctx.Request().GetFullRaw()
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetFullRaw()异常：%s", err)
	}
	ctx.Log().Info("----GetFullRaw data:", string(bodyRaw), raw)
	xmap := ctx.Request().GetMap()

	ctx.Log().Info("----GetMap()：", xmap)
	tm, err := ctx.Request().GetDatetime("param6")
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetDatetime()异常：%s", err)
	}
	ctx.Log().Info("----GetDatetime()：", tm)
	ctx.Log().Info("----Keys data:", ctx.Request().Keys())
	ctx.Log().Info("----Len data:", ctx.Request().Len())
	ctx.Log().Info("----GetArray data:", ctx.Request().GetArray("param5"))
	ctx.Log().Info("----GetInt32 data:", ctx.Request().GetInt32("param3"))
	ctx.Log().Info("----GetFloat32 data:", ctx.Request().GetFloat32("param4"))
	ctx.Log().Info("----GetBool data:", ctx.Request().GetBool("param2"))
	return "success"
}
