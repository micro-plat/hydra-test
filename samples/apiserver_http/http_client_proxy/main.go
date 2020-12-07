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
	hydra.WithSystemName("http_client_proxy"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

func init() {
	hydra.Conf.API(":8070")
	hydra.Conf.Vars().HTTP("http", httpopt.WithProxy("http://127.0.0.1:6550")) //设置代理

	app.API("/client/api", func(ctx hydra.IContext) (r interface{}) {
		c := components.Def.HTTP().GetRegularClient("http")
		ctx.Log().Info("request.sessiom_id:", ctx.Log().GetSessionID())
		content, status, err := c.Request("GET", "https://www.google.com/", "", "UTF-8", httpx.Header{}, &httpx.Cookie{})
		ctx.Log().Info("request.content:", string(content))
		ctx.Log().Info("request.status:", status)
		return err
	})
}

//测试http组件代理
//启动服务  ./http_client_proxy run
//访问  /client/api     [访问正常,返回200]  查看代理成功

func main() {
	app.Start()
}
