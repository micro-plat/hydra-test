package main

import (
	"time"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydra_test"),
	hydra.WithSystemName("run"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("zk://192.168.0.101"),
)

func init() {
	hydra.Conf.API(":8080", api.WithTimeout(10, 10))
	app.API("/api", func(ctx context.IContext) (r interface{}) {
		ctx.Log().Info("api-Handle")
		time.Sleep(time.Second * 2)
		return "success"
	})

}

//go build
//启动服务 ./apiserver_run run -r "zk://192.168.0.101" -c t [服务监听配置]
//关闭服务
//安装配置 ./apiserver_run conf install -r "zk://192.168.0.101" -c t
//启动服务 ./apiserver_run run -r "zk://192.168.0.101" -c t [成功]
//更新zk节点 /hydra_test/run/api/t/conf 的status值为空,start 服务器进行重启成功
//更新zk节点 /hydra_test/run/api/t/conf 的status值为stop 服务器关闭api服务,不进行重启
//更新zk节点 /hydra_test/run/api/t/conf 的rTimeout,wTimeout值为1, 无法访问/api
func main() {
	app.Start()
}
