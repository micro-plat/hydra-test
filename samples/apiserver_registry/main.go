package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/acl/limiter"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydra_test"),
	hydra.WithSystemName("registry"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("zk://192.168.0.101"),
)

func init() {
	hydra.Conf.API(":8080").Limit(limiter.WithRuleList(limiter.NewRule("/api/registry/multiple/query", 0, limiter.WithMaxWait(3), limiter.WithFallback(), limiter.WithReponse(200, "success"))))
	app.API("/api/registry/new", NewService, api.WithEncoding("UTF-8"))     //构建函数注册,并设置编码
	app.API("/api/registry/service", Service{})                             //对象注册
	app.API("/api/registry/func", ServiceFunc)                              //函数注册
	app.API("/api/registry/multiple", &Service2{}, api.WithEncoding("GBK")) //引用对象注册,并设置编码
	app.API("/api/registry/multiple", &Service3{})                          //同一服务注册多个对象
	app.API("/api/registry/*", &Service2{})                                 //同一注册对象注册不同服务
}

//启动服务
//  /api/registry/new [GET.POST] 查看编码(utf-8),服务名,日志打印
//  /api/registry/service [GET.POST] 查看服务名,日志打印
//  /api/registry/func [GET.POST] 查看服务名,日志打印
//  /api/registry/multiple [GET] 查看编码(gbk),服务名,日志打印(service2)
//  /api/registry/multiple [POST] 查看编码(gbk),服务名,日志打印(service3)
//  /api/registry/multiple/query [POST.GET] 降级,查看编码(gbk),服务名,日志打印
//  /api/registry/handle [POST.GET] 查看服务名,日志打印
//  /api/registry/query [POST.GET] 不降级,查看编码(utf-8),服务名,日志打印
//  /api/registry/queryfallback [POST.GET] 不可访问
func main() {
	app.Start()
}
