package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/hydra/hydra/servers/http"
	"github.com/micro-plat/hydra/registry"
)

var hydraApp = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("cusserve_registry"),
	hydra.WithClusterName("test"),
	hydra.WithRegistry("custom://."),
)

func init() {
	//注册 自定义的配置中心
	registry.Register(registrycustom, &customFactory{
		opts: &registry.Options{},
	})
	hydra.Conf.API(":50019")

	hydraApp.API("/custom/registry/api", func(ctx hydra.IContext) interface{} {
		regist, err := registry.NewRegistry(global.Def.GetRegistryAddr(), ctx.Log())
		if err != nil {
			ctx.Log().Error("registry.NewRegistry:", err)
		}
		data, version, err := regist.GetValue("/hydratest/cusserve_registry/api/test/conf")
		if err != nil {
			ctx.Log().Error("regist.GetValue:", err)
		}

		return map[string]interface{}{
			"conf_data": string(data),
			"version":   version,
		}
	})

}

//启动服务
//1. 编译程序
//2. 执行 ./cusserve_registry conf install 安装配置信息
//3. 启动服务 ./cusserve_registry run
//4. 请求 http://localhost:50019/custom/registry/api 获取正常的响应[conf_data,version] ，状态码：200
func main() {
	hydraApp.Start()
}
