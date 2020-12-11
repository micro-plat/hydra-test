package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/acl/limiter"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.WS),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("ws_limit"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

func init() {
	hydra.Conf.WS(":8080").Limit(limiter.WithEnable(), limiter.WithRuleList(
		limiter.NewRule("/ws/query", 0, limiter.WithMaxWait(1), limiter.WithFallback(), limiter.WithReponse(200, "success")),
		limiter.NewRule("/ws", 0, limiter.WithMaxWait(1), limiter.WithFallback(), limiter.WithReponse(200, "success"))))
	app.WS("/ws", &Service{})

}

//  测试ws限流启动
//  启动服务./middleware_limit run
//  建立连接
//  发送数据{"service":"/ws"} 降级[service-get-fallback],返回200,fallback
//  发送数据{"service":"/ws/query"} 降级[Service-query-FallBack],返回200,fallback
//  发送数据{"service":"/ws/order"} 不进行降级处理,返回200,succes
func main() {
	app.Start()
}
