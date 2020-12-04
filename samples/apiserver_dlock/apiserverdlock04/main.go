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
	app.API("/hydratest/apiserverDlock/getLock", funcAPI)
	app.API("/hydratest/apiserverDlock/getTryLock", funcAPI1)
}

// apiserver_dlock 获取锁时，注册中心掉线测试demo
//1.1 安装程序 ./apiserverdlock04 conf install -cover
//1.2 使用 ./apiserverdlock04 run

//1.3 调用接口获取分布式锁：http://localhost:8070/hydratest/apiserverDlock/getLock  能够获取独占锁成功
//1.4 调用接口获取分布式锁：http://localhost:8070/hydratest/apiserverDlock/getTryLock  能够尝试获取锁成功
//1.5 断开本机网络
//1.6 调用接口获取分布式锁：http://localhost:8070/hydratest/apiserverDlock/getLock  获取锁失败，:zk: could not connect to the server
//1.7 调用接口获取分布式锁：http://localhost:8070/hydratest/apiserverDlock/getTryLock   获取锁失败，:zk: could not connect to the server
//1.8 恢复网络，注册中心自动链接成功
//1.9 调用接口获取分布式锁：http://localhost:8070/hydratest/apiserverDlock/getLock  能够获取独占锁成功
//1.10 调用接口获取分布式锁：http://localhost:8070/hydratest/apiserverDlock/getTryLock  能够尝试获取锁成功
//1.11 在注册中心lm和fs中不存在分布式场景，所以暂时不用管
func main() {
	app.Start()
}

var funcAPI = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver_dlock 获取锁时，注册中心掉线测试demo")
	registry := ctx.APPConf().GetServerConf().GetRegistry()
	dlockObj := dlock.NewLockByRegistry("tasoytest", registry)
	if err := dlockObj.Lock(); err != nil {
		ctx.Log().Errorf("获取分布式锁异常:", err)
		return
	}
	defer dlockObj.Unlock()
	return "success"
}

var funcAPI1 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver_dlock 获取锁时，注册中心掉线测试demo")
	registry := ctx.APPConf().GetServerConf().GetRegistry()
	dlockObj := dlock.NewLockByRegistry("tasoytest", registry)
	if err := dlockObj.TryLock(); err != nil {
		ctx.Log().Errorf("获取分布式锁异常:", err)
		return
	}
	defer dlockObj.Unlock()
	return "success"
}
