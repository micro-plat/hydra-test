package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/micro-plat/hydra"
	componentrpc "github.com/micro-plat/hydra/components/rpcs/rpc"
	varconf "github.com/micro-plat/hydra/conf/vars/rpc"
	"github.com/micro-plat/hydra/hydra/servers/http"
	"github.com/micro-plat/hydra/hydra/servers/rpc"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API, rpc.RPC),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("rpcserver_request"),
	hydra.WithClusterName("test"),
	hydra.WithRegistry("lm://."),
)

func init() {
	hydra.Conf.RPC(":50025")
	hydra.Conf.API(":50026")

	hydra.Conf.Vars().RPC("rpc", varconf.WithRoundRobin())
	app.API("/api/request/:datatype", apiRequest)
	app.API("/api/bind/:resulttype/:datatype", apiBind)
	app.RPC("/rpc/proc", rpcProc)
	app.RPC("/rpc/bindmap", rpcBindMap)
	app.RPC("/rpc/bindstruct", rpcBindStruct)

}

// rpcserver_request 测试多个provider时默认ip轮寻负载均衡规则执行demo

//1.1 安装程序 ./rpcserver_request conf install -cover
//1.2 使用 ./rpcserver_request run
//1.1.1 http://localhost:50026/api/request/xml 调用xml 格式输出
//1.1.2 http://localhost:50026/api/request/json 调用json 格式输出
//1.1.3 http://localhost:50026/api/request/int 调用int 格式输出
//1.1.4 http://localhost:50026/api/request/int32 调用int32 格式输出
//1.1.5 http://localhost:50026/api/request/int64 调用int64 格式输出
//1.1.6 http://localhost:50026/api/request/float 调用float 格式输出
//1.1.7 http://localhost:50026/api/request/float32 调用float32 格式输出
//1.1.8 http://localhost:50026/api/request/float64 调用float64 格式输出
//1.1.9 http://localhost:50026/api/request/string 调用string 格式输出
//1.1.10 http://localhost:50026/api/request/time 调用time 格式输出
//1.1.11 http://localhost:50026/api/request/bool 调用bool 格式输出
//1.1.12 http://localhost:50026/api/request/structjson 调用structjson 格式输出
//1.1.13 http://localhost:50026/api/request/structxml 调用structxml 格式输出
//1.1.14 http://localhost:50026/api/request/structptrjson 调用structptrjson 格式输出
//1.1.15 http://localhost:50026/api/request/structptrxml 调用structptrxml 格式输出
//1.1.16 http://localhost:50026/api/request/mapintjson 调用mapintjson 格式输出
//1.1.17 http://localhost:50026/api/request/mapintxml 调用mapintxml格式输出
//1.1.18 http://localhost:50026/api/request/mapfloatjson 调用mapfloatjson格式输出
//1.1.19 http://localhost:50026/api/request/mapfloatxml 调用mapfloatxml格式输出
//1.1.20 http://localhost:50026/api/request/mapstringxmlarray 调用mapstringxmlarray格式输出
//1.1.21 http://localhost:50026/api/request/mapstringjsonarray 调用mapstringjsonarray格式输出
//1.1.22 http://localhost:50026/api/request/mapinterfacejson 调用mapinterfacejson格式输出
//1.1.23 http://localhost:50026/api/request/mapinterfacexml 调用mapinterfacexml格式输出
//1.1.24 http://localhost:50026/api/request/error 调用error 格式输出,状态码：400
//1.1.25 http://localhost:50026/api/request/errorcustom 调用errorcustom格式输出，状态码：400
//1.1.27 http://localhost:50026/api/request/body 调用xml 格式输出

func main() {
	app.Start()
}

type demoStrcutParent struct {
	Name     string             `json:"name" xml:"name" form:"name"`
	Age      int                `json:"age" xml:"age"  form:"age"`
	Children []*demoStrcutChild `json:"children" xml:"children"`
}

type demoStrcutChild struct {
	ID string `json:"id" xml:"id"`
}

var apiRequest = func(ctx hydra.IContext) (r interface{}) {
	dataType := ctx.Request().Path().Params().GetString("datatype")
	requestID := ctx.User().GetRequestID()
	ctx.Log().Infof("DataType:%s", dataType)

	opts := []componentrpc.RequestOption{}
	opts = append(opts, componentrpc.WithXRequestID(requestID))

	var input = map[string]interface{}{
		"datatype": dataType,
	}

	respones, err := hydra.C.RPC().GetRegularRPC().Request("/rpc/proc@hydratest", input, opts...)
	if err != nil {
		ctx.Log().Errorf("rpc 请求异常：%v", err)
		return
	}
	ctx.Log().Info("respones.Status:", respones.Status)
	return respones.Result
}

var apiBind = func(ctx hydra.IContext) (r interface{}) {
	resultType := ctx.Request().Path().Params().GetString("resulttype")
	dataType := ctx.Request().Path().Params().GetString("datatype")

	rpcURL := fmt.Sprintf("/rpc/bind%s@hydratest", resultType)
	var input interface{}
	var contentType string
	opts := []componentrpc.RequestOption{}
	switch dataType {
	case "form":
		input = `name=1&age=2`
		contentType = "application/x-www-form-urlencoded"
	case "xml":
		input = `<demoStrcutParent><name>demo.xml1</name><age>38</age><children><id>xml007</id></children></demoStrcutParent>`
		contentType = "application/xml"
	case "json":
		input = `{"name":"demo.json1","age":38,"children":[{"id":"json007"}]}`
		contentType = "application/json"
	}
	opts = append(opts, componentrpc.WithContentType(contentType))
	//fmt.Println("content-type:", contentType)
	respones, err := hydra.C.RPC().GetRegularRPC().Request(rpcURL, input, opts...)
	if err != nil {
		ctx.Log().Errorf("rpc 请求异常：%v", err)
		return
	}
	ctx.Log().Info("respones.Status:", respones.Status)
	return respones.Result
}

var rpcProc = func(ctx hydra.IContext) (r interface{}) {
	dataType := ctx.Request().GetString("datatype")
	var input interface{}
	switch dataType {
	case "xml":
		input = "<xml><a>1</a></xml>"
		ctx.Response().ContentType("appliction/xml")

	case "json":
		ctx.Response().ContentType("appliction/json")
		return `{"name":"demo","age":38}`
	case "int":
		input = 1
	case "int32":
		input = int32(32)
	case "int64":
		input = int64(64)
	case "float":
		input = 1.0
	case "float32":
		input = float32(32.1)
	case "float64":
		input = float64(64.2)
	case "string":
		input = "success"
	case "time":
		input = time.Now()
	case "bool":
		input = true
	case "structjson":
		ctx.Response().ContentType("appliction/json")
		input = demoStrcutParent{Name: "demo.json1", Age: 38, Children: []*demoStrcutChild{&demoStrcutChild{ID: "json007"}}}
	case "structxml":
		ctx.Response().ContentType("appliction/xml")
		input = demoStrcutParent{Name: "demo.xml1", Age: 38, Children: []*demoStrcutChild{&demoStrcutChild{ID: "xml007"}}}
	case "structptrjson":
		ctx.Response().ContentType("appliction/json")
		input = &demoStrcutParent{Name: "demo.json2", Age: 38, Children: []*demoStrcutChild{&demoStrcutChild{ID: "json007"}}}
	case "structptrxml":
		ctx.Response().ContentType("appliction/xml")
		input = &demoStrcutParent{Name: "demo.xml2", Age: 38, Children: []*demoStrcutChild{&demoStrcutChild{ID: "json007"}}}
	case "mapstringjson":
		ctx.Response().ContentType("appliction/json")
		input = map[string]string{
			"order": "123456",
		}
	case "mapstringxml": //不支持map 转xml
		ctx.Response().ContentType("appliction/xml")
		input = map[string]string{
			"order": "78910",
		}
	case "mapintjson":
		ctx.Response().ContentType("appliction/json")
		input = map[string]int{
			"order": 12345678910123456,
		}
	case "mapintxml": //不支持map 转xml
		ctx.Response().ContentType("appliction/xml")
		input = map[string]int{
			"order": 12345678910123456,
		}
	case "mapfloatjson":
		ctx.Response().ContentType("appliction/json")
		input = map[string]float64{
			"order": 12345678910123456.1,
		}
	case "mapfloatxml": //不支持map 转xml
		ctx.Response().ContentType("appliction/xml")
		input = map[string]float64{
			"order": 12345678910123456.2,
		}
	case "mapstringxmlarray": //不支持Array和Slice 类型
		ctx.Response().ContentType("appliction/xml")
		input = []map[string]string{
			map[string]string{"order": "78911"},
			map[string]string{"order": "78912"},
		}
	case "mapstringjsonarray":
		ctx.Response().ContentType("appliction/json")
		input = []map[string]string{
			map[string]string{"order": "78911"},
			map[string]string{"order": "78912"},
		}
	case "mapinterfacejson":
		ctx.Response().ContentType("appliction/json")

		input = map[string]interface{}{
			"product": map[string]string{
				"price": "100",
			},
		}
	case "mapinterfacexml":
		ctx.Response().ContentType("appliction/xml")
		input = map[string]interface{}{
			"product": map[string]string{
				"price": "200",
			},
		}
	case "error":
		input = errors.New("原生错误消息")
	case "errorcustom":
		input = errs.NewError(201, "无需处理")
	case "body":
		ctx.Log().Info(ctx.Request().GetBody())
		r, err := ctx.Request().GetBody()
		if err != nil {
			return err
		}
		input = r
	default:
		input = "default value"
	}

	return input
}

var rpcBindMap = func(ctx hydra.IContext) (r interface{}) {
	result := types.XMap{}
	err := ctx.Request().Bind(&result)
	if err != nil {
		ctx.Log().Errorf("Request.Bind:%v", err)
		return err
	}
	return result
}

var rpcBindStruct = func(ctx hydra.IContext) (r interface{}) {

	result := &demoStrcutParent{}
	m, err := ctx.Request().GetMap()
	ctx.Log().Info("GetMap:", m, err)
	err = ctx.Request().Bind(result)
	if err != nil {
		ctx.Log().Errorf("Request.Bind:%v", err)
		return err
	}
	return result
}
