package main

import (
	httpx "net/http"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/components"
	httpopt "github.com/micro-plat/hydra/conf/vars/http"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydra_test"),
	hydra.WithSystemName("apiserver_http"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

func init() {
	hydra.Conf.API(":8070")
	hydra.Conf.Vars().HTTP("http", httpopt.WithConnTimeout(30), httpopt.WithRequestTimeout(10))

	app.API("/client/api", func(ctx hydra.IContext) (r interface{}) {
		c := components.Def.HTTP().GetRegularClient("http")
		ctx.Log().Info("request.sessiom_id:", ctx.Log().GetSessionID())
		content, status, err := c.Request(
			"POST", "http://127.0.0.1:8099/api", `{"key":"value"}`, "UTF-8",
			httpx.Header{"Content-type": []string{"application/json"}, "X-Request-Id": []string{ctx.Log().GetSessionID()}},
			&httpx.Cookie{Name: "key", Value: "value"})
		ctx.Log().Info("request.content:", string(content))
		ctx.Log().Info("request.status:", status)
		return err
	})

	app.API("/client/timeout", func(ctx hydra.IContext) (r interface{}) {
		c := components.Def.HTTP().GetRegularClient("http")
		ctx.Log().Info("request.sessiom_id:", ctx.Log().GetSessionID())
		content, status, err := c.Request("POST", "http://127.0.0.1:8099/timeout", "", "", httpx.Header{}, &httpx.Cookie{})
		ctx.Log().Info("request.content:", string(content))
		ctx.Log().Info("request.status:", status)
		return err
	})
}

//queue组件是否正确工作,修改配置是否自动生效(redis,mqtt)
//启动服务  ./http_server run
//启动客户端服务  ./apiserver_http run
//访问  /client/api 查看服务器打印的session_id,body,method,header,encoding等信息
//访问  /client/timeout 超过请求时间.返回400 i/o timeout

func main() {
	app.Start()
}
