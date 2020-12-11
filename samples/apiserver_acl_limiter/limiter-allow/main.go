package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/acl/limiter"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("limiter_allow"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

func init() {
	hydra.Conf.API(":8070").Limit(limiter.WithEnable(), limiter.WithRuleList(
		limiter.NewRule("/api/query", 3, limiter.WithFallback(), limiter.WithReponse(200, "success")),
		limiter.NewRule("/api", 3, limiter.WithFallback(), limiter.WithReponse(200, "success"))))
	app.API("/api", &Service{})
}

//  测试限流启动，并发请求时超过允许数量的降级处理
//  启动服务  ./limiter_allow run
//  1秒10个并发请求访问 /api [GET] 返回200 4次降级[service-get-fallback],返回fallback, 6次不降级[Service-get-Handle]返回handle
//  1秒10个并发请求访问 /api [POST] 返回200 4次降级[service-fallback],返回fallback, 6次不降级[Service-Handle]返回handle
//  1秒10个并发请求访问 /api/query [GET.POST] 4次降级[service-query-fallback],返回fallback, 6次不降级[Service-query-Handle]返回handle
func main() {
	app.Start()
}
