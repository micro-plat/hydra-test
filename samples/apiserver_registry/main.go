package main

import (
	"github.com/micro-plat/hydra"
	_ "github.com/micro-plat/hydra/components/caches/cache/redis"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydra_demo"),
	hydra.WithSystemName("registry"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("zk://192.168.0.101"),
)

func init() {
	hydra.Conf.API(":8080", api.WithTimeout(10, 10))
	app.API("/api/registry/new", NewService, api.WithEncoding("UTF-8"))     //构建函数注册,并设置编码
	app.API("/api/registry/multiple", &Service2{}, api.WithEncoding("GBK")) //引用对象注册,并设置编码
	app.API("/api/registry/service", Service{})                             //对象注册
	app.API("/api/registry/func", ServiceFunc)                              //函数注册
	app.API("/api/registry/*", &Service2{})                                 //同一注册对象注册不同服务
	app.API("/api/registry/multiple", &Service3{})                          //同一服务注册多个对象
}

func main() {
	app.Start()
}
