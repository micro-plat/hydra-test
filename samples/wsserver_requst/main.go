package main

import (
	"fmt"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.WS),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("wsserver_request"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

func init() {
	hydra.Conf.WS(":8080")
	app.WS("/ws", func(ctx context.IContext) (r interface{}) {
		ctx.Log().Info("ws-Handle")
		headers := ctx.Request().Headers()
		fmt.Println("h:", headers)
		ctx.Log().Info("ws-Handle-headers:", headers)
		m, err := ctx.Request().GetMap()
		if err != nil {
			return err
		}
		ctx.Log().Info("ws-Handle-map:", m)
		ctx.Response().Header("Content-Type", "application/yaml;charset=utf-8")
		return m
	})
}

//ws服务访问，数据返回等，通过 http://coolaf.com/tool/chattest 测试
//启动服务./wsserver run
//建立与服务的连接
//给服务发送数据 {"service":"/ws","params":"params","data":1} 打印的header[Client-IP,Content-Type,X-Request-Id] 正确 打印的map正确 返回200
//设置不同的响应头部,返回200,返回不同格式的数据
func main() {
	app.Start()
}
