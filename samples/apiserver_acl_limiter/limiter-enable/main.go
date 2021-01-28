package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/acl/limiter"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("limiter_enable"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

func init() {
	hydra.Conf.API("8070").Limit(limiter.WithEnable(), limiter.WithRuleList(
		limiter.NewRule("/api/order", 0, limiter.WithMaxWait(1), limiter.WithReponse(200, "success")), //未配置WithFallback()
		limiter.NewRule("/api/query", 0, limiter.WithMaxWait(1), limiter.WithFallback(), limiter.WithReponse(200, "success")),
		limiter.NewRule("/api", 0, limiter.WithMaxWait(1), limiter.WithFallback(), limiter.WithReponse(200, "success"))))
	app.API("/api", &Service{})
}

//  测试限流启动，降级、非降级时的处理流程
//  启动服务  ./limiter-enable run
//  访问 /api [GET] 降级[service-get-fallback],返回200,fallback-handle
//  访问 /api [POST] 降级[service-fallback],返回200,fallback-handle
//  访问 /api/query [GET.POST] 降级[service-query-fallback],返回200,fallback-handle
//  访问 /api/order [GET.POST] 不进行降级处理,返回200,succes
func main() {
	app.Start()
}
