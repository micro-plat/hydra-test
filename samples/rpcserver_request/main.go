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
//1.4 调用接口执行循环访问rpc：http://localhost:8070/hydratest/rpcserbalance/apiip 观察两台服务器的执行日志，轮流访问两台服务器
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
