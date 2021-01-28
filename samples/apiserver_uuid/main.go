package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/components/uuid"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/hydra/servers/http"
	"github.com/micro-plat/lib4go/concurrent/cmap"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("apiserverUUID"),
	hydra.WithClusterName("taosytest"),
	hydra.WithRegistry("lm://."),
)

func init() {
	hydra.Conf.API("8072", api.WithTimeout(100, 100))
	app.API("/hydratest/apiserverUUID/get", funcAPI)
}

// apiserver_uuid uuid同集群下并发获取demo
//1.1 使用 ./apiserver_uuid run

//1.2 调用接口：http://localhost:8072/hydratest/apiserverUUID/get  观察日志是否有异常,1000并发的耗时情况
func main() {
	app.Start()
}

var uuidMap = cmap.New(1)

var funcAPI = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver_uuid 静态加载队列后，手动修改注册配置demo")
	uuidMap.Clear()
	clusterID := ctx.APPConf().GetServerConf().GetClusterName()
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		start := time.Now()
		for j := 0; j < 1000; j++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				compare(clusterID)
			}()
			wg.Add(1)
			go func() {
				defer wg.Done()
				compare("123456")
			}()
		}
		wg.Wait()
		end := time.Now()
		c := end.Sub(start)
		ctx.Log().Info("两个集群同时1000并发获取uuid耗时:", c.Seconds())
		time.Sleep(time.Second)
	}
	return
}

func compare(clusterID string) {
	uid := uuid.GetSUUID(clusterID).Get().ToString()
	if uuidMap.Has(uid) {
		panic(fmt.Sprintf("uuid重复:%s:%D", uid, uuidMap.Count()))
	}
	uuidMap.Set(uid, "1")
}
