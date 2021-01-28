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
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("apiserver_run"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("zk://192.168.0.101"),
)

func init() {
	hydra.Conf.API("8080", api.WithTimeout(10, 10))

	app.API("/api", func(ctx context.IContext) (r interface{}) {
		ctx.Log().Info("api-Handle")
		time.Sleep(time.Second * 2)
		return "success"
	})

	app.API("/file", func(ctx context.IContext) (r interface{}) {
		ctx.Log().Info("file-Handle")
		_, err := ctx.Request().GetFileBody("file")
		ctx.Log().Info("GetFileBody:", err)
		if err != nil {
			return err
		}
		return "success"
	})

}

//使用zookeeper作为注册中心，当配置未创建，已创建，已修改，已启动，已关闭等情况的服务器状态，及修改服务超时时间后工作是否正常

//go build
//启动服务 ./apiserver_run run [服务监听配置]
//关闭服务
//安装配置 ./apiserver_run conf install [安装配置成功]
//启动服务 ./apiserver_run run [成功]
//更新zk节点 /hydratest/run/api/t/conf 的address值为8070 [服务器重新.配置更新完成]
//更新zk节点 /hydratest/run/api/t/conf 的status值为空,start 服务器进行重启成功
//更新zk节点 /hydratest/run/api/t/conf 的status值为stop 服务器关闭api服务,不进行重启
//更新zk节点 /hydratest/run/api/t/conf 的wTimeout值为1, 无法访问/api
//更新zk节点 /hydratest/run/api/t/conf 的wTimeout值为0, 正常访问/api
//更新zk节点 /hydratest/run/api/t/conf 的wTimeout值为3, 正常访问/api
//更新zk节点 /hydratest/run/api/t/conf 的rTimeout值为1, 上传的大文件(1.3G) 访问/file [i/o timeout 返回400]
//更新zk节点 /hydratest/run/api/t/conf 的rTimeout值为0, 上传的大文件(1.3G) 访问/file [正常]
func main() {
	app.Start()
}
