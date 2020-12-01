package main

import (
	"github.com/micro-plat/hydra"
	_ "github.com/micro-plat/hydra/components/caches/cache/redis"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/hydra/servers/http"
	"github.com/micro-plat/lib4go/types"
)

var app = hydra.NewApp(
	hydra.WithDebug(),
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("apiservertrace"),
	hydra.WithClusterName("taosytest"),
	hydra.WithRegistry("zk://192.168.0.101"),
)

func init() {
	hydra.Conf.API(":8070", api.WithTrace())
	app.API("/hydratest/apiserver/:trace", funcTrace)
}

// 检查各种输入和输出的数据demo

//1.1 安装程序 sudo ./servertrace06 conf install -cover
//1.2 使用默认端口监听 ./servertrace06 run
//1.3 调用接口：http://192.168.5.94:8070/hydratest/apiserver/:1?param1=test&param2=中文数据$%#^##@  get请求带中文key-value数据
//1.3 调用接口：http://192.168.5.94:8070/hydratest/apiserver/:2   post-body-json：{"param1":"test","param2":"中文数据"}
//1.3 调用接口：http://192.168.5.94:8070/hydratest/apiserver/:3   post-body-json：{"param1":"test","param2":"中文数据"}
//1.3 调用接口：http://192.168.5.94:8070/hydratest/apiserver/:4   post-body-xml：<sites><site><name>菜鸟教程</name><url>www.runoob.com</url></site></sites>   不编码传输
//1.3 调用接口：http://192.168.5.94:8070/hydratest/apiserver/:5   post-body-xml：<sites><site><name>菜鸟教程</name><url>www.runoob.com</url></site></sites>   编码后传输
//1.3 调用接口：http://192.168.5.94:8070/hydratest/apiserver/:6   post-body-text：param1testparam2中文数据34%%   不编码传输
//1.3 调用接口：http://192.168.5.94:8070/hydratest/apiserver/:7   post-body-text：text：param1testparam2中文数据34%%   编码后传输baidu
//1.3 调用接口：http://192.168.5.94:8070/hydratest/apiserver/:8   post-fromdata-json：{"param1":"test","param2":"中文数据"}   编码后传输
//1.3 调用接口：http://192.168.5.94:8070/hydratest/apiserver/:9   post-fromdata-json：{"param1":"test","param2":"中文数据"}   不编码传输
//1.3 调用接口：http://192.168.5.94:8070/hydratest/apiserver/:10   post-fromdata-xml：<sites><site><name>菜鸟教程</name><url>www.runoob.com</url></site></sites>  编码后传输
//1.3 调用接口：http://192.168.5.94:8070/hydratest/apiserver/:11   post-fromdata-xml：<sites><site><name>菜鸟教程</name><url>www.runoob.com</url></site></sites>  不编码传输
//1.3 调用接口：http://192.168.5.94:8070/hydratest/apiserver/:12   post-fromdata-text：param1testparam2中文数据34%%   编码后传输
//1.3 调用接口：http://192.168.5.94:8070/hydratest/apiserver/:13   post-fromdata-text：param1testparam2中文数据34%%   不编码传输
//1.3 调用接口：http://192.168.5.94:8070/hydratest/apiserver/:14?ppostram1=test&param2=中文数据  post-bodyjson-kv混合传输： {"param13":"test","param4":"中文数据"}
//1.3 调用接口：http://192.168.5.94:8070/hydratest/apiserver/:15?param1=test&param2=中文数据  post-bodyxml：-kv混合传输：<sites><site><name>菜鸟教程</name><url>www.runoob.com</url></site></sites>
//1.3 调用接口：http://192.168.5.94:8070/hydratest/apiserver/:16?param1=test&param2=中文数据  post-bodytext-kv混合传输： param1testparam2中文数据34%%
//1.3 调用接口：http://192.168.5.94:8070/hydratest/apiserver/:17?param1=test&param2=中文数据  post-body-kv混合传输： {"param13":"test","param4":"中文数据"}

func main() {
	app.Start()
}

var funcTrace func(ctx context.IContext) (r interface{}) = func(ctx context.IContext) (r interface{}) {
	ctx.Log().Info("apiserver_trace 检查各种输入和输出的数据demo")
	mapT := ctx.Request().Path().Params()
	trace := mapT.GetString("trace")
	xmap, err := ctx.Request().GetMap()
	if err != nil {
		ctx.Log().Errorf("errerrerrerr:%v", err)
	}
	switch trace {
	case "1":
		d := types.XMap{"param1": "test", "param2": "中文数据"}
		if len(d) != len(xmap) {
			ctx.Log().Errorf("get请求时 数据错误1:%v", xmap)
			return "fail"
		}
		for k, v := range xmap {
			if _, ok := d[k]; !ok || (ok && d[k] != v) {
				ctx.Log().Errorf("get请求时 数据错误2 %v", xmap)
				return "fail"
			}
			delete(d, k)
		}
		if len(d) > 0 {
			ctx.Log().Errorf("get请求时 数据错误3:%v", xmap)
			return "fail"
		}
	}

	return "success"
}
