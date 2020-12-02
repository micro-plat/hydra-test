package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/acl/limiter"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydra_test"),
	hydra.WithSystemName("limiter_disbale"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

func init() {
	hydra.Conf.API(":8070").Limit(limiter.WithDisable(), limiter.WithRuleList(
		limiter.NewRule("/api", 0, limiter.WithMaxWait(1), limiter.WithFallback(), limiter.WithReponse(200, "success"))))
	app.API("/api", &Service{})
}

//  测试限流禁用
//  启动服务  ./limiter-enable run
//  访问 /api [GET,POST] 不降级,返回200,handle
//  访问 /api/query [GET.POST] 不降级,返回200,handle
//  访问 /api/order [GET.POST]  不降级,返回200,handle
func main() {
	app.Start()
}
