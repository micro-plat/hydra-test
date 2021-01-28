package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/components/dlock"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/hydra/servers/http"
	"github.com/micro-plat/hydra/registry"
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
	hydra.Conf.API("8070", api.WithTimeout(100, 100))
	app.API("/hydratest/apiserverDlock/get", funcAPI)
}

// apiserver_dlock 高并发下调用获取独占分布式锁测试demo
//1.1 安装程序 ./apiserverdlock02 conf install -cover
//1.2 使用 ./apiserverdlock02 run

//1.3 调用接口：http://localhost:8070/hydratest/apiserverDlock/get  观察日志是否有异常,1000并发的耗时情况
// 预期说明：因为count计数器不是线程安全的，所以如果try_count=succ_count 说明分布式锁在高并发情况下运行正常
func main() {
	app.Start()
}

var count = 0

var funcAPI = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver_dlock 高并发下调用获取独占分布式锁测试demo")
	count = 0
	wg := &sync.WaitGroup{}
	for i := 0; i < 1; i++ {
		start := time.Now()
		for j := 0; j < 10; j++ {
			wg.Add(1)
			go func(count *int) {
				defer wg.Done()
				regst, err := registry.CreateRegistry("zk://192.168.0.101", ctx.Log())
				if err != nil {
					ctx.Log().Errorf("获取注册中心异常，err:%v", err)
					return
				}
				dlockObj := dlock.NewLockByRegistry("tasoytest", regst)
				if err := dlockObj.Lock(); err == nil {
					defer dlockObj.Unlock()
					*count++
				}
			}(&count)
		}
		wg.Wait()
		end := time.Now()
		c := end.Sub(start)
		ctx.Log().Info("1000并发获取独占dlock耗时:", c.Seconds())
	}
	ctx.Log().Info("分布式锁获取情况汇总:", fmt.Sprintf(`{"succ_count":%d,"try_count":10000}`, count))
	return fmt.Sprintf(`{"succ_count":%d,"try_count":10000}`, count)
}
