package main

import (
	"time"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/conf/server/metric"
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
	hydra.Conf.API(":8080", api.WithTimeout(10, 10)).Metric("192.168.0.101:8090", "1", "cron", metric.WithDisable(), metric.WithUPName("upnem", "1223456"))
	app.API("/api", func(ctx context.IContext) (r interface{}) {
		ctx.Log().Info("api-Handle")
		return "success"
	})
	app.API("/api2", func(ctx context.IContext) (r interface{}) {
		ctx.Log().Info("api3-Handle")
		time.Sleep(time.Second * 2)
		return "success"
	})
}

//使用zookeeper作为注册中心，通过代码指定服务器的关键配置，手动修改后自动更新
//go build
//安装配置 ./apiserver_conf conf install -r "zk://192.168.0.101" -c t
//启动服务 ./apiserver_conf run -r "zk://192.168.0.101" -c t [成功] 可访问/api
//更新zk节点 /hydra_test/run/api/t/conf 的address值为8070 [服务器重新.配置更新完成]
//更新zk节点 /hydra_test/run/api/t/conf 的status值为stop [关闭api服务,不重启,配置更新完成]
//更新zk节点 /hydra_test/run/api/t/conf 的status值为start或者空值 [重启api服务成功,配置更新完成]
//更新zk节点 /hydra_test/run/api/t/conf 的wTimeout值为1  重启服务,无法访问/api2
//更新zk节点 /hydra_test/run/api/t/conf 的wTimeout值为0  重启服务,正常访问/api2
//更新zk节点 /hydra_test/run/api/t/conf 的wTimeout值为3  重启服务,正常访问/api2
//更新zk节点 /hydra_test/run/api/t/conf 的rTimeout值为1  重启服务,上传的大文件(600M),i/o timeout 返回400
//更新zk节点 /hydra_test/run/api/t/conf 的rTimeout值为0  重启服务,上传的大文件(600M), 正常访问/api3
//更新zk节点 /hydra_test/run/api/t/conf 的rhTimeout值为1 暂时没有实现header读取超时的情况
//更新zk节点 /hydra_test/run/api/t/conf 的rhTimeout值为0 重启服务,正常访问/api
//更新zk节点 /hydra_test/run/api/t/conf 的dn值为1
//更新zk节点 /hydra_test/run/api/t/conf 的dn值为0

//conf下默认的节点监控
//更新zk节点 /hydra_test/run/api/t/conf/router,将action/GET删除 [保存conf节点,重启配置成功,配置更新完成] GET无法访问/api,POST可访问/api
//更新zk节点 /hydra_test/run/api/t/conf/metric,将disable改为true [保存conf节点,重启配置成功,配置更新完成]

func main() {
	app.Start()
}
