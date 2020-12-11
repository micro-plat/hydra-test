package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/acl/limiter"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("registry"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("zk://192.168.0.101"),
)

func init() {
	hydra.Conf.API(":8080").Limit(
		limiter.WithRuleList(limiter.NewRule("/api/registry/multiple/query", 1, limiter.WithFallback(), limiter.WithReponse(200, "fallback"))),
		limiter.WithRuleList(limiter.NewRule("/void/api/registry/multiple/query", 1, limiter.WithFallback(), limiter.WithReponse(200, "fallback"))))
	//签名有返回值
	app.API("/api/registry/new", NewService, api.WithEncoding("UTF-8"))     //构建函数注册,并设置编码
	app.API("/api/registry/service", Service{})                             //对象注册
	app.API("/api/registry/func", ServiceFunc)                              //函数注册
	app.API("/api/registry/multiple", &Service2{}, api.WithEncoding("GBK")) //引用对象注册,并设置编码
	app.API("/api/registry/multiple", &Service3{})                          //同一服务注册多个对象
	app.API("/api/registry/*", &Service2{})                                 //同一注册对象注册不同服务
	app.API("/api", &Service4{})                                            //注册地址与handle前缀相同

	//签名没有返回值
	app.API("/void/api/registry/new", NewVoidService, api.WithEncoding("UTF-8"))     //构建函数注册,并设置编码
	app.API("/void/api/registry/service", VoidService{})                             //对象注册
	app.API("/void/api/registry/func", VoidServiceFunc)                              //函数注册
	app.API("/void/api/registry/multiple", &VoidService2{}, api.WithEncoding("GBK")) //引用对象注册,并设置编码
	app.API("/void/api/registry/multiple", &VoidService3{})                          //同一服务注册多个对象
	app.API("/void/api/registry/*", &VoidService2{})                                 //同一注册对象注册不同服务
	app.API("/void/api", &VoidService4{})                                            //注册地址与handle前缀相同
}

//测试服务注册，包括但不限于函数注册、对象注册、构建函数注册，restful服务，降级服务(不能直接被外部访问)，编码等及验证context中服务名称等是否正确
//启动服务
//访问服务
//  /api/registry/new [GET.POST] 查看编码(utf-8),服务名,日志打印
//  /api/registry/service [GET.POST] 查看服务名,日志打印
//  /api/registry/func [GET.POST] 查看服务名,日志打印
//  /api/registry/multiple [GET] 查看编码(gbk),服务名,日志打印(service2)
//  /api/registry/multiple [POST] 查看编码(gbk),服务名,日志打印(service3)
//  /api/registry/multiple/query [POST.GET] 1秒10个并发请求 9次降级,查看编码(gbk),服务名,日志打印,返回200,success
//  /api/registry/handle [POST.GET] 查看服务名,日志打印
//  /api/registry/query [POST.GET] 1秒10个并发请求,不降级,查看编码(utf-8),服务名,日志打印
//  /api/registry/queryfallback [POST.GET] 不可访问 返回404
//  /api [POST.GET] 不可访问 返回404
//  /api/api [POST.GET] 正常访问,返回200

//  /void/api/registry/new [GET.POST] 查看编码(utf-8),服务名,日志打印
//  /void/api/registry/service [GET.POST] 查看服务名,日志打印
//  /void/api/registry/func [GET.POST] 查看服务名,日志打印
//  /void/api/registry/multiple [GET] 查看编码(gbk),服务名,日志打印(voidservice2)
//  /void/api/registry/multiple [POST] 查看编码(gbk),服务名,日志打印(voidservice3)
//  /void/api/registry/multiple/query [POST.GET] 1秒10个并发请求 9次降级,查看编码(gbk),服务名,日志打印,返回200,success
//  /void/api/registry/handle [POST.GET] 查看服务名,日志打印
//  /void/api/registry/query [POST.GET] 1秒10个并发请求 不降级,查看编码(utf-8),服务名,日志打印
//  /void/api/registry/queryfallback [POST.GET] 不可访问 返回404
//  /void/api [POST.GET] 不可访问 返回404
//  /void/api/api [POST.GET] 正常访问,返回200

func main() {
	app.Start()
}
