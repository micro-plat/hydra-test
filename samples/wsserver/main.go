package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.WS),
	hydra.WithPlatName("hydra_test"),
	hydra.WithSystemName("wsserver"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("zk://192.168.0.101"),
)

func init() {
	hydra.Conf.WS(":8080", api.WithTimeout(10, 10), api.WithEnable(), api.WithTrace()).Jwt().Basic().Limit()
	app.WS("/ws", func(ctx context.IContext) (r interface{}) {
		ctx.Log().Info("ws-Handle")
		return "success"
	})
	app.WS("/ws/timeout", func(ctx context.IContext) (r interface{}) {
		ctx.Log().Info("ws-timeout-Handle")
		return "success"
	})

}

//使用zookeeper作为注册中心，当配置未创建，已创建，已修改，已启动，已关闭等情况的服务器状态，及修改服务超时时间后工作是否正常

//go build
//启动服务 /wsserver run [服务监听配置]
//关闭服务
//安装配置 ./wsserver conf install
//启动服务 ./wsserver run [成功]

//更新zk节点 /hydra_test/wsserver/ws/t/conf 的address值为8070 [服务器重启.配置更新完成]
//更新zk节点 /hydra_test/wsserver/ws/t/conf 的status值为空,start 服务器进行重启成功
//更新zk节点 /hydra_test/wsserver/ws/t/conf 的status值为stop 服务器关闭api服务,不进行重启
//更新zk节点 /hydra_test/wsserver/ws/t/conf 的trace值为false 服务器重启

//更新zk节点 /hydra_test/wsserver/ws/t/auth/jwt   的disbale值为true [服务器重启.配置更新完成]
//更新zk节点 /hydra_test/wsserver/ws/t/auth/basic 的disbale值为true [服务器重启.配置更新完成]
//更新zk节点 /hydra_test/wsserver/ws/t/acl/limit  的disbale值为true [服务器重启.配置更新完成]
func main() {
	app.Start()
}
