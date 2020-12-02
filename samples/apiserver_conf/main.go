package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydra_test"),
	hydra.WithSystemName("conf"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("zk://192.168.0.101"),
)

func init() {
	hydra.Conf.API(":8080", api.WithTimeout(10, 10))
	app.API("/api", func(ctx context.IContext) (r interface{}) {
		ctx.Log().Info("api-Handle")
		return "success"
	})

}

//使用zookeeper作为注册中心，通过代码指定服务器的关键配置，手动修改后自动更新
//go build
//安装配置 ./apiserver_conf conf install -r "zk://192.168.0.101" -c t
//启动服务 ./apiserver_conf run -r "zk://192.168.0.101" -c t [成功] 可访问/api
//更新zk节点 /hydra_test/run/api/t/conf 的address值为8070 [服务器重新.配置更新完成]
//更新zk节点 /hydra_test/run/api/t/conf 的status值为stop [关闭api服务,不重启,配置更新完成]
//更新zk节点 /hydra_test/run/api/t/conf 的status值为start [重启api服务成功,配置更新完成]
//更新zk节点 /hydra_test/run/api/t/conf/router,将action/GET删除 [缓存配置清除,重启配置成功,配置更新完成] GET无法访问/api,POST可访问/api
func main() {
	app.Start()
}
