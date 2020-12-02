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
	hydra.WithDebug(),
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("apiserverUUID"),
	hydra.WithClusterName("taosytest"),
	hydra.WithRegistry("zk://192.168.0.101"),
)

func init() {
	hydra.Conf.API(":8070", api.WithTimeout(100, 100))
	app.API("/hydratest/apiserverUUID/get", funcAPI)
}

// apiserver_uuid uuid同集群下并发获取demo
//1.1 安装程序 sudo ./apiserver_uuid conf install -cover
//1.2 使用 ./apiserver_uuid run

//1.3 调用接口：http://localhost:8070/hydratest/apiserverUUID/get  观察日志是否有异常,1000并发的耗时情况
func main() {
	app.Start()
}

var uuidMap cmap.ConcurrentMap

var funcAPI = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("apiserver_uuid 静态加载队列后，手动修改注册配置demo")
	defer func() {
		e := recover()
		if e != nil {
			ctx.Log().Errorf("uuid测试失败,%v", e)
		}
	}()
	if uuidMap == nil {
		uuidMap = cmap.New(1)
	}
	uuidMap.Clear()
	clusterID := ctx.APPConf().GetServerConf().GetClusterName()
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		a := time.Now()
		for j := 0; j < 1000; j++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				uid := uuid.Get(clusterID).ToString()
				if err := compare(uid); err != nil {
					panic(err)
				}
			}()
			// wg.Add(1)
			// go func() {
			// 	defer wg.Done()
			// 	uid := uuid.Get("123456").ToString()
			// 	if err := compare(uid); err != nil {
			// 		panic(err)
			// 	}
			// }()
		}

		wg.Wait()
		b := time.Now()
		c := b.Sub(a)
		ctx.Log().Info("两个集群同时1000并发获取uuid耗时:", c.Seconds())
		time.Sleep(time.Second)
	}
	uuidMap.Clear()
	return
}

func compare(uid string) error {
	if uuidMap.Has(uid) {
		return fmt.Errorf("uuid重复")
	}
	uuidMap.Set(uid, "1")
	return nil
}
