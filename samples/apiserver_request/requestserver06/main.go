package main

import (
	"io/ioutil"
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
	hydra.Conf.API(":8070")
	app.API("/hydratest/apiserver/request/nil", funcRequest)
	app.API("/hydratest/apiserver/request/gbk", funcRequest1, router.WithEncoding("gbk"))
	app.API("/hydratest/apiserver/request/gb2312", funcRequest2, router.WithEncoding("gb2312"))
	app.API("/hydratest/apiserver/request/utf8", funcRequest3, router.WithEncoding("utf-8"))
}

// apiserver-request-文件数据提交demo
//1.2 使用 ./requestserver06 run
//1.4 请求的数据：
/*
通过postman构建如下请求：
1.1 请求路由无编码格式接口-Post-fromData-ut8-与头编码相同         返回正确的数据
1.2 请求路由无编码格式接口-Post-fromData-ut8-与头编码不同         无法获取数据
1.3 请求路由无编码格式接口-Post-fromData-gbk-与头编码相同         返回正确的数据
1.4 请求路由无编码格式接口-Post-fromData-gbk-与头编码不同         无法获取数据
1.5 请求路由无编码格式接口-Post-fromData-gb2312-与头编码相同       返回正确的数据
1.6 请求路由无编码格式接口-Post-fromData-gb2312-与头编码不同         无法获取数据

2.1 请求路由有编码格式接口-Post-fromData-ut8-与路由编码相同         返回正确的数据
2.2 请求路由有编码格式接口-Post-fromData-ut8-与路由编码不同         无法获取数据
2.3 请求路由有编码格式接口-Post-fromData-gbk-与路由编码相同         返回正确的数据
2.4 请求路由有编码格式接口-Post-fromData-gbk-与路由编码不同         无法获取数据
2.5 请求路由有编码格式接口-Post-fromData-gb2312-与路由编码相同       返回正确的数据
2.6 请求路由有编码格式接口-Post-fromData-gb2312-与路由编码不同         无法获取数据

3.1 头和路由都有编码-编码不同-Get-fromData-gbk-utf8-与头编码相同        无法获取数据
3.2 头和路由都有编码-编码不同-Get-fromData-gbk-utf8-与路由编码相同      返回正确的数据
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
	ctx.Log().Info("apiserver-request-yaml,路由不设置编码数据处理demo")
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
	_, raw, err := ctx.Request().GetFullRaw()
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetFullRaw()异常：%s", err)
	}
	ctx.Log().Info("----GetFullRaw data:", raw)
	jsonD, err := ctx.Request().GetJSON("param7")
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetJSON()异常：%s", err)
	}
	ctx.Log().Info("----GetJSON data:", string(jsonD))
	xmap, err := ctx.Request().GetMap()
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetMap()异常：%s", err)
	}
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

	fileIO, err := ctx.Request().GetFileBody("file1")
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetFileBody()异常：%s", err)
	}
	s, _ := ioutil.ReadAll(fileIO)
	ctx.Log().Info("----GetFileBody()：", string(s))
	fileName, err := ctx.Request().GetFileName("file1")
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetFileName()异常：%s", err)
	}
	ctx.Log().Info("----GetFileName()：", fileName)
	fileSize, err := ctx.Request().GetFileSize("file1")
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetFileSize()异常：%s", err)
	}
	ctx.Log().Info("----GetFileSize()：", fileSize)
	return "success"
}

//路由设置编码 gbk
var funcRequest1 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver-request-yaml,路由设置gbk编码数据处理demo")
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
	_, raw, err := ctx.Request().GetFullRaw()
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetFullRaw()异常：%s", err)
	}
	ctx.Log().Info("----GetFullRaw data:", raw)
	jsonD, err := ctx.Request().GetJSON("param7")
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetJSON()异常：%s", err)
	}
	ctx.Log().Info("----GetJSON data:", string(jsonD))
	xmap, err := ctx.Request().GetMap()
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetMap()异常：%s", err)
	}
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

	fileIO, err := ctx.Request().GetFileBody("file1")
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetFileBody()异常：%s", err)
	}
	s, _ := ioutil.ReadAll(fileIO)
	ctx.Log().Info("----GetFileBody()：", string(s))
	fileName, err := ctx.Request().GetFileName("file1")
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetFileName()异常：%s", err)
	}
	ctx.Log().Info("----GetFileName()：", fileName)
	fileSize, err := ctx.Request().GetFileSize("file1")
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetFileSize()异常：%s", err)
	}
	ctx.Log().Info("----GetFileSize()：", fileSize)
	return "success"
}

//路由设置编码 gb2312
var funcRequest2 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver-request-yaml,路由设置gb2312编码数据处理demo")
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
	_, raw, err := ctx.Request().GetFullRaw()
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetFullRaw()异常：%s", err)
	}
	ctx.Log().Info("----GetFullRaw data:", raw)
	jsonD, err := ctx.Request().GetJSON("param7")
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetJSON()异常：%s", err)
	}
	ctx.Log().Info("----GetJSON data:", string(jsonD))
	xmap, err := ctx.Request().GetMap()
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetMap()异常：%s", err)
	}
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
	fileIO, err := ctx.Request().GetFileBody("file1")
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetFileBody()异常：%s", err)
	}
	s, _ := ioutil.ReadAll(fileIO)
	ctx.Log().Info("----GetFileBody()：", string(s))
	fileName, err := ctx.Request().GetFileName("file1")
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetFileName()异常：%s", err)
	}
	ctx.Log().Info("----GetFileName()：", fileName)
	fileSize, err := ctx.Request().GetFileSize("file1")
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetFileSize()异常：%s", err)
	}
	ctx.Log().Info("----GetFileSize()：", fileSize)
	return "success"
}

//路由设置编码 utf8
var funcRequest3 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver-request-yaml,路由设置utf8编码数据处理demo")
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
	_, raw, err := ctx.Request().GetFullRaw()
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetFullRaw()异常：%s", err)
	}
	ctx.Log().Info("----GetFullRaw data:", raw)
	jsonD, err := ctx.Request().GetJSON("param7")
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetJSON()异常：%s", err)
	}
	ctx.Log().Info("----GetJSON data:", string(jsonD))
	xmap, err := ctx.Request().GetMap()
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetMap()异常：%s", err)
	}
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
	fileIO, err := ctx.Request().GetFileBody("file1")
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetFileBody()异常：%s", err)
	}
	s, _ := ioutil.ReadAll(fileIO)
	ctx.Log().Info("----GetFileBody()：", string(s))
	fileName, err := ctx.Request().GetFileName("file1")
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetFileName()异常：%s", err)
	}
	ctx.Log().Info("----GetFileName()：", fileName)
	fileSize, err := ctx.Request().GetFileSize("file1")
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetFileSize()异常：%s", err)
	}
	ctx.Log().Info("----GetFileSize()：", fileSize)
	return "success"
}
