package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/acl/limiter"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydra_test"),
	hydra.WithSystemName("limiter_enable"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

func init() {
	hydra.Conf.API(":8070").Limit(limiter.WithEnable(), limiter.WithRuleList(
		limiter.NewRule("/api/query", 0, limiter.WithMaxWait(1), limiter.WithFallback(), limiter.WithReponse(200, "success")),
		limiter.NewRule("/api/order", 0, limiter.WithMaxWait(1), limiter.WithReponse(200, "success")), //未配置WithFallback()
		limiter.NewRule("/api", 0, limiter.WithMaxWait(1), limiter.WithFallback(), limiter.WithReponse(200, "success"))))
	app.API("/api", &Service{})
	app.API("/api2", &Service{})
}

//  测试限流启动，降级、非降级时的处理流程
//  启动服务  ./limiter-enable run
//  访问 /api [GET] 降级[service-get-fallback],返回200,fallback
//  访问 /api [POST] 降级[service-fallback],返回200,fallback
//  访问 /api/query [GET.POST] 降级[service-get-fallback],返回200,fallback
//  访问 /api/order [GET.POST] 不进行降级处理,返回200,success
//  访问 /api2* [GET.POST] 不降级,返回200,handle
func main() {
	app.Start()
}
