package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/components/dlock"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/hydra/servers/http"
	"github.com/micro-plat/hydra/registry"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("apiserverDlock"),
	hydra.WithClusterName("taosytest"),
	hydra.WithRegistry("lm://."),
)

func init() {
	hydra.Conf.API(":8072", api.WithTimeout(100, 100))
	app.API("/hydratest/apiserverDlock/get", funcAPI)
}

// apiserver_dlock 高并发下调用偿试获取分布式锁测试demo
//1.1 使用 ./apiserverdlock01 run

//1.2 调用接口：http://localhost:8072/hydratest/apiserverDlock/get  观察日志是否有异常,1000并发的耗时情况
// 预期说明：因为count计数器不是线程安全的,atomicCount计数器是线程安全的，所以如果atomicCount=count 说明分布式锁在高并发情况下运行正常
func main() {
	app.Start()
}

var atomicCount = int64(0)
var count = 0

var funcAPI = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver_dlock 高并发下调用偿试获取分布式锁测试demo")
	count = 0
	atomicCount = 0
	wg := &sync.WaitGroup{}
	for i := 0; i < 1; i++ {
		start := time.Now()
		for j := 0; j < 10; j++ {
			wg.Add(1)
			go func(count *int, atomicCount *int64) {
				defer wg.Done()
				regst, err := registry.CreateRegistry("zk://192.168.0.101", ctx.Log())
				if err != nil {
					ctx.Log().Errorf("获取注册中心异常，err:%v", err)
					return
				}
				dlockObj := dlock.NewLockByRegistry("tasoytest", regst)
				if err := dlockObj.TryLock(); err == nil {
					defer dlockObj.Unlock()
					*count++
					atomic.AddInt64(atomicCount, 1)
				}
			}(&count, &atomicCount)
		}
		wg.Wait()
		end := time.Now()
		c := end.Sub(start)
		ctx.Log().Info("1000并发获取trydlock耗时:", c.Seconds())
	}
	ctx.Log().Info("分布式锁获取情况汇总:", fmt.Sprintf(`{"atomicCount":%d,"count":%d}`, count, atomicCount))
	return fmt.Sprintf(`{"atomicCount":%d,"count":%d}`, count, atomicCount)
}
