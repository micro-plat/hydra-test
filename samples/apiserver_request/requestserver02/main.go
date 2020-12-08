package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var funcTrace = func(ctx hydra.IContext) (r interface{}) {
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

	//1. 路由编码 （1.不舍值，2,正确，3错误
	//2. 格式： form,body,fromdata, form+url,body+url,fromdata+url
	//3.contenttype: kv, json,yml,text,xml,
	//4. encoding : gbk,gb2312,utf8
	//5. 中文特殊字符

	// for contenttype {
	// 	for encoding{
	// 		for  xxx {
	// 			idx ++
	// app.API("/hydratest/apiserver/response"+idx, callback(contetype,encoding,xxx))
}

// func  callback(contentType,encoding,xxx string) func(context.IContext) interface{}{

// 	return func(ctx context.IContext) interface{}{
// 	}
// }

// apiserver-request-post请求demo

//1.1 安装程序 ./servertrace01 conf install -cover
//1.2 使用 ./servertrace01 run
//1.3 调用接口：http://localhost:8070/hydratest/apiserver/trace 判定配置是否正确

/*
1.路由不设置编码格式-Post-from-正确contentType-gbk-带中文和特殊符号正确编码数据
1.路由不设置编码格式-Post-from-正确contentType-gbk-带中文和特殊符号错误编码数据
2.路由不设置编码格式-Post-from-正确contentType-gb2312-带中文和特殊符号正确编码数据
2.路由不设置编码格式-Post-from-正确contentType-gb2312-带中文和特殊符号错误编码数据
3.路由不设置编码格式-Post-from-正确contentType-utf8-带中文和特殊符号正确编码数据
3.路由不设置编码格式-Post-from-正确contentType-utf8-带中文和特殊符号错误编码数据
3.路由不设置编码格式-Post-from-错误contentType-utf8-带中文和特殊符号正确编码数据
1.路由不设置编码格式-Post-from-错误contentType-gbk-带中文和特殊符号正确编码数据

1.路由不设置编码格式-Post-from+url-错误contentType-gbk-带中文和特殊符号正确编码数据

1.路由不设置编码格式-Post-body-json-正确contentType-gbk-带中文和特殊符号正确编码数据
1.路由不设置编码格式-Post-body-json-正确contentType-gbk-带中文和特殊符号错误编码数据
2.路由不设置编码格式-Post-body-json-正确contentType-gb2312-带中文和特殊符号正确编码数据
2.路由不设置编码格式-Post-body-json-正确contentType-gb2312-带中文和特殊符号错误编码数据
3.路由不设置编码格式-Post-body-json-正确contentType-utf8-带中文和特殊符号正确编码数据
3.路由不设置编码格式-Post-body-json-正确contentType-utf8-带中文和特殊符号错误编码数据
1.路由不设置编码格式-Post-body-json-错误contentType-gbk-带中文和特殊符号正确编码数据
2.路由不设置编码格式-Post-body-json-错误contentType-gb2312-带中文和特殊符号正确编码数据
3.路由不设置编码格式-Post-body-json-错误contentType-utf8-带中文和特殊符号正确编码数据
1.路由不设置编码格式-Post-body-xml-正确contentType-gbk-带中文和特殊符号正确编码数据
1.路由不设置编码格式-Post-body-xml-正确contentType-gbk-带中文和特殊符号错误编码数据
2.路由不设置编码格式-Post-body-xml-正确contentType-gb2312-带中文和特殊符号正确编码数据
2.路由不设置编码格式-Post-body-xml-正确contentType-gb2312-带中文和特殊符号错误编码数据
3.路由不设置编码格式-Post-body-xml-正确contentType-utf8-带中文和特殊符号正确编码数据
3.路由不设置编码格式-Post-body-xml-正确contentType-utf8-带中文和特殊符号错误编码数据
1.路由不设置编码格式-Post-body-xml-错误contentType-gbk-带中文和特殊符号正确编码数据
2.路由不设置编码格式-Post-body-xml-错误contentType-gb2312-带中文和特殊符号正确编码数据
3.路由不设置编码格式-Post-body-xml-错误contentType-utf8-带中文和特殊符号正确编码数据
1.路由不设置编码格式-Post-body-text-正确contentType-gbk-带中文和特殊符号正确编码数据
1.路由不设置编码格式-Post-body-text-正确contentType-gbk-带中文和特殊符号错误编码数据
2.路由不设置编码格式-Post-body-text-正确contentType-gb2312-带中文和特殊符号正确编码数据
2.路由不设置编码格式-Post-body-text-正确contentType-gb2312-带中文和特殊符号错误编码数据
3.路由不设置编码格式-Post-body-text-正确contentType-utf8-带中文和特殊符号正确编码数据
3.路由不设置编码格式-Post-body-text-正确contentType-utf8-带中文和特殊符号错误编码数据
1.路由不设置编码格式-Post-body-text-错误contentType-gbk-带中文和特殊符号正确编码数据
2.路由不设置编码格式-Post-body-text-错误contentType-gb2312-带中文和特殊符号正确编码数据
3.路由不设置编码格式-Post-body-text-错误contentType-utf8-带中文和特殊符号正确编码数据
1.路由不设置编码格式-Post-body-yaml-正确contentType-gbk-带中文和特殊符号正确编码数据
1.路由不设置编码格式-Post-body-yaml-正确contentType-gbk-带中文和特殊符号错误编码数据
2.路由不设置编码格式-Post-body-yaml-正确contentType-gb2312-带中文和特殊符号正确编码数据
2.路由不设置编码格式-Post-body-yaml-正确contentType-gb2312-带中文和特殊符号错误编码数据
3.路由不设置编码格式-Post-body-yaml-正确contentType-utf8-带中文和特殊符号正确编码数据
3.路由不设置编码格式-Post-body-yaml-正确contentType-utf8-带中文和特殊符号错误编码数据
1.路由不设置编码格式-Post-body-yaml-错误contentType-gbk-带中文和特殊符号正确编码数据
2.路由不设置编码格式-Post-body-yaml-错误contentType-gb2312-带中文和特殊符号正确编码数据
3.路由不设置编码格式-Post-body-yaml-错误contentType-utf8-带中文和特殊符号正确编码数据
1.路由不设置编码格式-Post-fromdata-json-正确contentType-gbk-带中文和特殊符号正确编码数据
1.路由不设置编码格式-Post-fromdata-json-正确contentType-gbk-带中文和特殊符号错误编码数据
2.路由不设置编码格式-Post-fromdata-json-正确contentType-gb2312-带中文和特殊符号正确编码数据
2.路由不设置编码格式-Post-fromdata-json-正确contentType-gb2312-带中文和特殊符号错误编码数据
3.路由不设置编码格式-Post-fromdata-json-正确contentType-utf8-带中文和特殊符号正确编码数据
3.路由不设置编码格式-Post-fromdata-json-正确contentType-utf8-带中文和特殊符号错误编码数据
1.路由不设置编码格式-Post-fromdata-json-错误contentType-gbk-带中文和特殊符号正确编码数据
2.路由不设置编码格式-Post-fromdata-json-错误contentType-gb2312-带中文和特殊符号正确编码数据
3.路由不设置编码格式-Post-fromdata-json-错误contentType-utf8-带中文和特殊符号正确编码数据
1.路由不设置编码格式-Post-fromdata-xml-正确contentType-gbk-带中文和特殊符号正确编码数据
1.路由不设置编码格式-Post-fromdata-xml-正确contentType-gbk-带中文和特殊符号错误编码数据
2.路由不设置编码格式-Post-fromdata-xml-正确contentType-gb2312-带中文和特殊符号正确编码数据
2.路由不设置编码格式-Post-fromdata-xml-正确contentType-gb2312-带中文和特殊符号错误编码数据
3.路由不设置编码格式-Post-fromdata-xml-正确contentType-utf8-带中文和特殊符号正确编码数据
3.路由不设置编码格式-Post-fromdata-xml-正确contentType-utf8-带中文和特殊符号错误编码数据
1.路由不设置编码格式-Post-fromdata-xml-错误contentType-gbk-带中文和特殊符号正确编码数据
2.路由不设置编码格式-Post-fromdata-xml-错误contentType-gb2312-带中文和特殊符号正确编码数据
3.路由不设置编码格式-Post-fromdata-xml-错误contentType-utf8-带中文和特殊符号正确编码数据
1.路由不设置编码格式-Post-fromdata-text-正确contentType-gbk-带中文和特殊符号正确编码数据
1.路由不设置编码格式-Post-fromdata-text-正确contentType-gbk-带中文和特殊符号错误编码数据
2.路由不设置编码格式-Post-fromdata-text-正确contentType-gb2312-带中文和特殊符号正确编码数据
2.路由不设置编码格式-Post-fromdata-text-正确contentType-gb2312-带中文和特殊符号错误编码数据
3.路由不设置编码格式-Post-fromdata-text-正确contentType-utf8-带中文和特殊符号正确编码数据
3.路由不设置编码格式-Post-fromdata-text-正确contentType-utf8-带中文和特殊符号错误编码数据
1.路由不设置编码格式-Post-fromdata-text-错误contentType-gbk-带中文和特殊符号正确编码数据
2.路由不设置编码格式-Post-fromdata-text-错误contentType-gb2312-带中文和特殊符号正确编码数据
3.路由不设置编码格式-Post-fromdata-text-错误contentType-utf8-带中文和特殊符号正确编码数据
1.路由不设置编码格式-Post-fromdata-yaml-正确contentType-gbk-带中文和特殊符号正确编码数据
1.路由不设置编码格式-Post-fromdata-yaml-正确contentType-gbk-带中文和特殊符号错误编码数据
2.路由不设置编码格式-Post-fromdata-yaml-正确contentType-gb2312-带中文和特殊符号正确编码数据
2.路由不设置编码格式-Post-fromdata-yaml-正确contentType-gb2312-带中文和特殊符号错误编码数据
3.路由不设置编码格式-Post-fromdata-yaml-正确contentType-utf8-带中文和特殊符号正确编码数据
3.路由不设置编码格式-Post-fromdata-yaml-正确contentType-utf8-带中文和特殊符号错误编码数据
1.路由不设置编码格式-Post-fromdata-yaml-错误contentType-gbk-带中文和特殊符号正确编码数据
2.路由不设置编码格式-Post-fromdata-yaml-错误contentType-gb2312-带中文和特殊符号正确编码数据
3.路由不设置编码格式-Post-fromdata-yaml-错误contentType-utf8-带中文和特殊符号正确编码数据

1.路由设置正确编码格式-Post-from-正确contentType-gbk-带中文和特殊符号正确编码数据
1.路由设置正确编码格式-Post-from-正确contentType-gbk-带中文和特殊符号错误编码数据
2.路由设置正确编码格式-Post-from-正确contentType-gb2312-带中文和特殊符号正确编码数据
2.路由设置正确编码格式-Post-from-正确contentType-gb2312-带中文和特殊符号错误编码数据
3.路由设置正确编码格式-Post-from-正确contentType-utf8-带中文和特殊符号正确编码数据
3.路由设置正确编码格式-Post-from-正确contentType-utf8-带中文和特殊符号错误编码数据
3.路由设置正确编码格式-Post-from-错误contentType-utf8-带中文和特殊符号正确编码数据
1.路由设置正确编码格式-Post-from-错误contentType-gbk-带中文和特殊符号正确编码数据
1.路由设置正确编码格式-Post-body-json-正确contentType-gbk-带中文和特殊符号正确编码数据
1.路由设置正确编码格式-Post-body-json-正确contentType-gbk-带中文和特殊符号错误编码数据
2.路由设置正确编码格式-Post-body-json-正确contentType-gb2312-带中文和特殊符号正确编码数据
2.路由设置正确编码格式-Post-body-json-正确contentType-gb2312-带中文和特殊符号错误编码数据
3.路由设置正确编码格式-Post-body-json-正确contentType-utf8-带中文和特殊符号正确编码数据
3.路由设置正确编码格式-Post-body-json-正确contentType-utf8-带中文和特殊符号错误编码数据
1.路由设置正确编码格式-Post-body-json-错误contentType-gbk-带中文和特殊符号正确编码数据
2.路由设置正确编码格式-Post-body-json-错误contentType-gb2312-带中文和特殊符号正确编码数据
3.路由设置正确编码格式-Post-body-json-错误contentType-utf8-带中文和特殊符号正确编码数据
1.路由设置正确编码格式-Post-body-xml-正确contentType-gbk-带中文和特殊符号正确编码数据
1.路由设置正确编码格式-Post-body-xml-正确contentType-gbk-带中文和特殊符号错误编码数据
2.路由设置正确编码格式-Post-body-xml-正确contentType-gb2312-带中文和特殊符号正确编码数据
2.路由设置正确编码格式-Post-body-xml-正确contentType-gb2312-带中文和特殊符号错误编码数据
3.路由设置正确编码格式-Post-body-xml-正确contentType-utf8-带中文和特殊符号正确编码数据
3.路由设置正确编码格式-Post-body-xml-正确contentType-utf8-带中文和特殊符号错误编码数据
1.路由设置正确编码格式-Post-body-xml-错误contentType-gbk-带中文和特殊符号正确编码数据
2.路由设置正确编码格式-Post-body-xml-错误contentType-gb2312-带中文和特殊符号正确编码数据
3.路由设置正确编码格式-Post-body-xml-错误contentType-utf8-带中文和特殊符号正确编码数据
1.路由设置正确编码格式-Post-body-text-正确contentType-gbk-带中文和特殊符号正确编码数据
1.路由设置正确编码格式-Post-body-text-正确contentType-gbk-带中文和特殊符号错误编码数据
2.路由设置正确编码格式-Post-body-text-正确contentType-gb2312-带中文和特殊符号正确编码数据
2.路由设置正确编码格式-Post-body-text-正确contentType-gb2312-带中文和特殊符号错误编码数据
3.路由设置正确编码格式-Post-body-text-正确contentType-utf8-带中文和特殊符号正确编码数据
3.路由设置正确编码格式-Post-body-text-正确contentType-utf8-带中文和特殊符号错误编码数据
1.路由设置正确编码格式-Post-body-text-错误contentType-gbk-带中文和特殊符号正确编码数据
2.路由设置正确编码格式-Post-body-text-错误contentType-gb2312-带中文和特殊符号正确编码数据
3.路由设置正确编码格式-Post-body-text-错误contentType-utf8-带中文和特殊符号正确编码数据
1.路由设置正确编码格式-Post-body-yaml-正确contentType-gbk-带中文和特殊符号正确编码数据
1.路由设置正确编码格式-Post-body-yaml-正确contentType-gbk-带中文和特殊符号错误编码数据
2.路由设置正确编码格式-Post-body-yaml-正确contentType-gb2312-带中文和特殊符号正确编码数据
2.路由设置正确编码格式-Post-body-yaml-正确contentType-gb2312-带中文和特殊符号错误编码数据
3.路由设置正确编码格式-Post-body-yaml-正确contentType-utf8-带中文和特殊符号正确编码数据
3.路由设置正确编码格式-Post-body-yaml-正确contentType-utf8-带中文和特殊符号错误编码数据
1.路由设置正确编码格式-Post-body-yaml-错误contentType-gbk-带中文和特殊符号正确编码数据
2.路由设置正确编码格式-Post-body-yaml-错误contentType-gb2312-带中文和特殊符号正确编码数据
3.路由设置正确编码格式-Post-body-yaml-错误contentType-utf8-带中文和特殊符号正确编码数据
1.路由设置正确编码格式-Post-fromdata-json-正确contentType-gbk-带中文和特殊符号正确编码数据
1.路由设置正确编码格式-Post-fromdata-json-正确contentType-gbk-带中文和特殊符号错误编码数据
2.路由设置正确编码格式-Post-fromdata-json-正确contentType-gb2312-带中文和特殊符号正确编码数据
2.路由设置正确编码格式-Post-fromdata-json-正确contentType-gb2312-带中文和特殊符号错误编码数据
3.路由设置正确编码格式-Post-fromdata-json-正确contentType-utf8-带中文和特殊符号正确编码数据
3.路由设置正确编码格式-Post-fromdata-json-正确contentType-utf8-带中文和特殊符号错误编码数据
1.路由设置正确编码格式-Post-fromdata-json-错误contentType-gbk-带中文和特殊符号正确编码数据
2.路由设置正确编码格式-Post-fromdata-json-错误contentType-gb2312-带中文和特殊符号正确编码数据
3.路由设置正确编码格式-Post-fromdata-json-错误contentType-utf8-带中文和特殊符号正确编码数据
1.路由设置正确编码格式-Post-fromdata-xml-正确contentType-gbk-带中文和特殊符号正确编码数据
1.路由设置正确编码格式-Post-fromdata-xml-正确contentType-gbk-带中文和特殊符号错误编码数据
2.路由设置正确编码格式-Post-fromdata-xml-正确contentType-gb2312-带中文和特殊符号正确编码数据
2.路由设置正确编码格式-Post-fromdata-xml-正确contentType-gb2312-带中文和特殊符号错误编码数据
3.路由设置正确编码格式-Post-fromdata-xml-正确contentType-utf8-带中文和特殊符号正确编码数据
3.路由设置正确编码格式-Post-fromdata-xml-正确contentType-utf8-带中文和特殊符号错误编码数据
1.路由设置正确编码格式-Post-fromdata-xml-错误contentType-gbk-带中文和特殊符号正确编码数据
2.路由设置正确编码格式-Post-fromdata-xml-错误contentType-gb2312-带中文和特殊符号正确编码数据
3.路由设置正确编码格式-Post-fromdata-xml-错误contentType-utf8-带中文和特殊符号正确编码数据
1.路由设置正确编码格式-Post-fromdata-text-正确contentType-gbk-带中文和特殊符号正确编码数据
1.路由设置正确编码格式-Post-fromdata-text-正确contentType-gbk-带中文和特殊符号错误编码数据
2.路由设置正确编码格式-Post-fromdata-text-正确contentType-gb2312-带中文和特殊符号正确编码数据
2.路由设置正确编码格式-Post-fromdata-text-正确contentType-gb2312-带中文和特殊符号错误编码数据
3.路由设置正确编码格式-Post-fromdata-text-正确contentType-utf8-带中文和特殊符号正确编码数据
3.路由设置正确编码格式-Post-fromdata-text-正确contentType-utf8-带中文和特殊符号错误编码数据
1.路由设置正确编码格式-Post-fromdata-text-错误contentType-gbk-带中文和特殊符号正确编码数据
2.路由设置正确编码格式-Post-fromdata-text-错误contentType-gb2312-带中文和特殊符号正确编码数据
3.路由设置正确编码格式-Post-fromdata-text-错误contentType-utf8-带中文和特殊符号正确编码数据
1.路由设置正确编码格式-Post-fromdata-yaml-正确contentType-gbk-带中文和特殊符号正确编码数据
1.路由设置正确编码格式-Post-fromdata-yaml-正确contentType-gbk-带中文和特殊符号错误编码数据
2.路由设置正确编码格式-Post-fromdata-yaml-正确contentType-gb2312-带中文和特殊符号正确编码数据
2.路由设置正确编码格式-Post-fromdata-yaml-正确contentType-gb2312-带中文和特殊符号错误编码数据
3.路由设置正确编码格式-Post-fromdata-yaml-正确contentType-utf8-带中文和特殊符号正确编码数据
3.路由设置正确编码格式-Post-fromdata-yaml-正确contentType-utf8-带中文和特殊符号错误编码数据
1.路由设置正确编码格式-Post-fromdata-yaml-错误contentType-gbk-带中文和特殊符号正确编码数据
2.路由设置正确编码格式-Post-fromdata-yaml-错误contentType-gb2312-带中文和特殊符号正确编码数据
3.路由设置正确编码格式-Post-fromdata-yaml-错误contentType-utf8-带中文和特殊符号正确编码数据

1.路由设置错误编码格式-Post-from-正确contentType-gbk-带中文和特殊符号正确编码数据
1.路由设置错误编码格式-Post-from-正确contentType-gbk-带中文和特殊符号错误编码数据
2.路由设置错误编码格式-Post-from-正确contentType-gb2312-带中文和特殊符号正确编码数据
2.路由设置错误编码格式-Post-from-正确contentType-gb2312-带中文和特殊符号错误编码数据
3.路由设置错误编码格式-Post-from-正确contentType-utf8-带中文和特殊符号正确编码数据
3.路由设置错误编码格式-Post-from-正确contentType-utf8-带中文和特殊符号错误编码数据
3.路由设置错误编码格式-Post-from-错误contentType-utf8-带中文和特殊符号正确编码数据
1.路由设置错误编码格式-Post-from-错误contentType-gbk-带中文和特殊符号正确编码数据
1.路由设置错误编码格式-Post-body-json-正确contentType-gbk-带中文和特殊符号正确编码数据
1.路由设置错误编码格式-Post-body-json-正确contentType-gbk-带中文和特殊符号错误编码数据
2.路由设置错误编码格式-Post-body-json-正确contentType-gb2312-带中文和特殊符号正确编码数据
2.路由设置错误编码格式-Post-body-json-正确contentType-gb2312-带中文和特殊符号错误编码数据
3.路由设置错误编码格式-Post-body-json-正确contentType-utf8-带中文和特殊符号正确编码数据
3.路由设置错误编码格式-Post-body-json-正确contentType-utf8-带中文和特殊符号错误编码数据
1.路由设置错误编码格式-Post-body-json-错误contentType-gbk-带中文和特殊符号正确编码数据
2.路由设置错误编码格式-Post-body-json-错误contentType-gb2312-带中文和特殊符号正确编码数据
3.路由设置错误编码格式-Post-body-json-错误contentType-utf8-带中文和特殊符号正确编码数据
1.路由设置错误编码格式-Post-body-xml-正确contentType-gbk-带中文和特殊符号正确编码数据
1.路由设置错误编码格式-Post-body-xml-正确contentType-gbk-带中文和特殊符号错误编码数据
2.路由设置错误编码格式-Post-body-xml-正确contentType-gb2312-带中文和特殊符号正确编码数据
2.路由设置错误编码格式-Post-body-xml-正确contentType-gb2312-带中文和特殊符号错误编码数据
3.路由设置错误编码格式-Post-body-xml-正确contentType-utf8-带中文和特殊符号正确编码数据
3.路由设置错误编码格式-Post-body-xml-正确contentType-utf8-带中文和特殊符号错误编码数据
1.路由设置错误编码格式-Post-body-xml-错误contentType-gbk-带中文和特殊符号正确编码数据
2.路由设置错误编码格式-Post-body-xml-错误contentType-gb2312-带中文和特殊符号正确编码数据
3.路由设置错误编码格式-Post-body-xml-错误contentType-utf8-带中文和特殊符号正确编码数据
1.路由设置错误编码格式-Post-body-text-正确contentType-gbk-带中文和特殊符号正确编码数据
1.路由设置错误编码格式-Post-body-text-正确contentType-gbk-带中文和特殊符号错误编码数据
2.路由设置错误编码格式-Post-body-text-正确contentType-gb2312-带中文和特殊符号正确编码数据
2.路由设置错误编码格式-Post-body-text-正确contentType-gb2312-带中文和特殊符号错误编码数据
3.路由设置错误编码格式-Post-body-text-正确contentType-utf8-带中文和特殊符号正确编码数据
3.路由设置错误编码格式-Post-body-text-正确contentType-utf8-带中文和特殊符号错误编码数据
1.路由设置错误编码格式-Post-body-text-错误contentType-gbk-带中文和特殊符号正确编码数据
2.路由设置错误编码格式-Post-body-text-错误contentType-gb2312-带中文和特殊符号正确编码数据
3.路由设置错误编码格式-Post-body-text-错误contentType-utf8-带中文和特殊符号正确编码数据
1.路由设置错误编码格式-Post-body-yaml-正确contentType-gbk-带中文和特殊符号正确编码数据
1.路由设置错误编码格式-Post-body-yaml-正确contentType-gbk-带中文和特殊符号错误编码数据
2.路由设置错误编码格式-Post-body-yaml-正确contentType-gb2312-带中文和特殊符号正确编码数据
2.路由设置错误编码格式-Post-body-yaml-正确contentType-gb2312-带中文和特殊符号错误编码数据
3.路由设置错误编码格式-Post-body-yaml-正确contentType-utf8-带中文和特殊符号正确编码数据
3.路由设置错误编码格式-Post-body-yaml-正确contentType-utf8-带中文和特殊符号错误编码数据
1.路由设置错误编码格式-Post-body-yaml-错误contentType-gbk-带中文和特殊符号正确编码数据
2.路由设置错误编码格式-Post-body-yaml-错误contentType-gb2312-带中文和特殊符号正确编码数据
3.路由设置错误编码格式-Post-body-yaml-错误contentType-utf8-带中文和特殊符号正确编码数据
1.路由设置错误编码格式-Post-fromdata-json-正确contentType-gbk-带中文和特殊符号正确编码数据
1.路由设置错误编码格式-Post-fromdata-json-正确contentType-gbk-带中文和特殊符号错误编码数据
2.路由设置错误编码格式-Post-fromdata-json-正确contentType-gb2312-带中文和特殊符号正确编码数据
2.路由设置错误编码格式-Post-fromdata-json-正确contentType-gb2312-带中文和特殊符号错误编码数据
3.路由设置错误编码格式-Post-fromdata-json-正确contentType-utf8-带中文和特殊符号正确编码数据
3.路由设置错误编码格式-Post-fromdata-json-正确contentType-utf8-带中文和特殊符号错误编码数据
1.路由设置错误编码格式-Post-fromdata-json-错误contentType-gbk-带中文和特殊符号正确编码数据
2.路由设置错误编码格式-Post-fromdata-json-错误contentType-gb2312-带中文和特殊符号正确编码数据
3.路由设置错误编码格式-Post-fromdata-json-错误contentType-utf8-带中文和特殊符号正确编码数据
1.路由设置错误编码格式-Post-fromdata-xml-正确contentType-gbk-带中文和特殊符号正确编码数据
1.路由设置错误编码格式-Post-fromdata-xml-正确contentType-gbk-带中文和特殊符号错误编码数据
2.路由设置错误编码格式-Post-fromdata-xml-正确contentType-gb2312-带中文和特殊符号正确编码数据
2.路由设置错误编码格式-Post-fromdata-xml-正确contentType-gb2312-带中文和特殊符号错误编码数据
3.路由设置错误编码格式-Post-fromdata-xml-正确contentType-utf8-带中文和特殊符号正确编码数据
3.路由设置错误编码格式-Post-fromdata-xml-正确contentType-utf8-带中文和特殊符号错误编码数据
1.路由设置错误编码格式-Post-fromdata-xml-错误contentType-gbk-带中文和特殊符号正确编码数据
2.路由设置错误编码格式-Post-fromdata-xml-错误contentType-gb2312-带中文和特殊符号正确编码数据
3.路由设置错误编码格式-Post-fromdata-xml-错误contentType-utf8-带中文和特殊符号正确编码数据
1.路由设置错误编码格式-Post-fromdata-text-正确contentType-gbk-带中文和特殊符号正确编码数据
1.路由设置错误编码格式-Post-fromdata-text-正确contentType-gbk-带中文和特殊符号错误编码数据
2.路由设置错误编码格式-Post-fromdata-text-正确contentType-gb2312-带中文和特殊符号正确编码数据
2.路由设置错误编码格式-Post-fromdata-text-正确contentType-gb2312-带中文和特殊符号错误编码数据
3.路由设置错误编码格式-Post-fromdata-text-正确contentType-utf8-带中文和特殊符号正确编码数据
3.路由设置错误编码格式-Post-fromdata-text-正确contentType-utf8-带中文和特殊符号错误编码数据
1.路由设置错误编码格式-Post-fromdata-text-错误contentType-gbk-带中文和特殊符号正确编码数据
2.路由设置错误编码格式-Post-fromdata-text-错误contentType-gb2312-带中文和特殊符号正确编码数据
3.路由设置错误编码格式-Post-fromdata-text-错误contentType-utf8-带中文和特殊符号正确编码数据
1.路由设置错误编码格式-Post-fromdata-yaml-正确contentType-gbk-带中文和特殊符号正确编码数据
1.路由设置错误编码格式-Post-fromdata-yaml-正确contentType-gbk-带中文和特殊符号错误编码数据
2.路由设置错误编码格式-Post-fromdata-yaml-正确contentType-gb2312-带中文和特殊符号正确编码数据
2.路由设置错误编码格式-Post-fromdata-yaml-正确contentType-gb2312-带中文和特殊符号错误编码数据
3.路由设置错误编码格式-Post-fromdata-yaml-正确contentType-utf8-带中文和特殊符号正确编码数据
3.路由设置错误编码格式-Post-fromdata-yaml-正确contentType-utf8-带中文和特殊符号错误编码数据
1.路由设置错误编码格式-Post-fromdata-yaml-错误contentType-gbk-带中文和特殊符号正确编码数据
2.路由设置错误编码格式-Post-fromdata-yaml-错误contentType-gb2312-带中文和特殊符号正确编码数据
3.路由设置错误编码格式-Post-fromdata-yaml-错误contentType-utf8-带中文和特殊符号正确编码数据
*/

func main() {
	app.Start()
}
