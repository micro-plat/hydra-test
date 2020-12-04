package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/components/dlock"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithDebug(),
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("apiserverDlock"),
	hydra.WithClusterName("taosytest"),
	hydra.WithRegistry("zk://192.168.0.101"),
)

func init() {
	hydra.Conf.API(":8070")
	app.API("/hydratest/apiserverDlock/get", funcAPI)
}

// apiserver_dlock 锁获取成功，服务器断电，程序崩溃和强退锁是否被释放测试demo
//1.1 安装程序 sudo ./apiserverdlock03 conf install -cover
//1.2 使用 ./apiserverdlock03 run

//1.3 调用接口获取分布式锁不释放：http://localhost:8070/hydratest/apiserverDlock/get  查看接口日志两次获取锁是否满足预期
//1.4 直接kill -9 杀死程序
//1.5 再重新启动程序，调用1.3接口   预期应该和第一次请求接口相同;  如果产生异常，说明分布式锁在该场景下不可用;
//1.6 在注册中心lm和fs中不存在分布式场景，所以暂时不用管
func main() {
	app.Start()
}

var funcAPI = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver_dlock 锁获取成功，服务器断电，程序崩溃和强退锁是否被释放测试demo")
	registry := ctx.APPConf().GetServerConf().GetRegistry()
	dlockObj := dlock.NewLockByRegistry("tasoytest", registry)
	if err := dlockObj.Lock(); err != nil {
		ctx.Log().Errorf("获取分布式锁异常:", err)
		// defer dlockObj.Unlock()
		return "success"
	}
	ctx.Log().Info("锁没获取成功")
	if err := dlockObj.TryLock(); err != nil {
		ctx.Log().Info("锁没有释放，不能获取成功，满足预期：", err)
	} else {
		defer dlockObj.Unlock()
	}
	return "success"
}
