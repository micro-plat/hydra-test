package main

import (
	"time"

	"github.com/micro-plat/hydra"
	_ "github.com/micro-plat/hydra/components/caches/cache/gocache"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/conf/server/header"
	"github.com/micro-plat/hydra/conf/vars/cache/gocache"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var hydraApp = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("apiserver_gocache"),
	hydra.WithClusterName("test"),
	hydra.WithRegistry("lm://."),
)

const cacheName = "cachegocache"

func init() {

	hydra.Conf.Vars().Cache().GoCache(cacheName, gocache.WithCleanupInterval(time.Minute*5))

	hydra.Conf.API(":50023", api.WithTimeout(10, 10)).Header(header.WithHeader("content-type", "application/json"))

	hydraApp.API("/api/gocache/add", addHandler)
	hydraApp.API("/api/gocache/set", setHandler)
	hydraApp.API("/api/gocache/incr", incrementHandler)
	hydraApp.API("/api/gocache/decr", decrementHandler)
	hydraApp.API("/api/gocache/incrint", incrementIntHandler)
	hydraApp.API("/api/gocache/decrint", decrementIntHandler)
	hydraApp.API("/api/gocache/delete", deleteHandler)
	hydraApp.API("/api/gocache/delay", delayHandler)
	hydraApp.API("/api/gocache/gets", getsHandler)

}

// apiserver_db 数据库组件是否正确工作，修改配置是否自动生效（mysql）
// 1. 编译程序： go build
// 2. 启动程序：./apiserver_mysql run

// 3. 请求 http://localhost:50023/api/gocache/add 添加一个值,再次添加报错[已存在]
// 4. 请求 http://localhost:50023/api/gocache/set 设置值，存在即覆盖
// 5. 请求 http://localhost:50023/api/gocache/incr 进行加处理,每次递增100
// 6. 请求 http://localhost:50023/api/gocache/decr 进行减处理，每次递减50
// 7. 请求 http://localhost:50023/api/gocache/delete 验证删除，添加值，判断是否存在，删除，再判定值是否存在
// 8. 请求 http://localhost:50023/api/gocache/delay 验证delay 自动删除，设置值1,睡眠4秒，再检测值是否存在
// 8. 请求 http://localhost:50023/api/gocache/gets 获取多个key值
func main() {
	hydraApp.Start()
}

const cachekey = "hydratest:apiserver:gocache"

var addHandler = func(ctx hydra.IContext) (r interface{}) {
	cacheObj := hydra.C.Cache().GetRegularCache(cacheName)
	err := cacheObj.Add(cachekey, "1", -1)
	if err != nil {
		ctx.Log().Errorf("Cache.Add:%v", err)
		return err
	}
	val, err := cacheObj.Get(cachekey)
	if err != nil {
		ctx.Log().Errorf("Cache.Get:%v", err)
		return err
	}
	return map[string]interface{}{
		"add.val": val,
	}
}

var setHandler = func(ctx hydra.IContext) (r interface{}) {
	key := ctx.Request().GetString("key", cachekey)
	val := ctx.Request().GetString("val", "10")
	cacheObj := hydra.C.Cache().GetRegularCache(cacheName)
	err := cacheObj.Set(key, val, -1)
	if err != nil { 
		ctx.Log().Errorf("Cache.Set:%v", err)
		return err
	}
	val, err = cacheObj.Get(key)
	if err != nil {
		ctx.Log().Errorf("Cache.Get:%v", err)
		return err
	}
	return map[string]interface{}{
		"set.val": val,
	}
}

var incrementHandler = func(ctx hydra.IContext) (r interface{}) {
	cacheObj := hydra.C.Cache().GetRegularCache(cacheName)
	newval, err := cacheObj.Increment(cachekey, 100)
	if err != nil {
		ctx.Log().Errorf("Cache.Increment:%v", err)
		return err
	}
	return map[string]interface{}{
		"incr.val": newval,
	}
}

var decrementHandler = func(ctx hydra.IContext) (r interface{}) {
	cacheObj := hydra.C.Cache().GetRegularCache(cacheName)
	newval, err := cacheObj.Decrement(cachekey, 50)
	if err != nil {
		ctx.Log().Errorf("Cache.Decrement:%v", err)
		return err
	}
	return map[string]interface{}{
		"decr.val": newval,
	}
}

var incrementIntHandler = func(ctx hydra.IContext) (r interface{}) {
	newcachekey := "hydra:apiserver:gocache-int"
	cacheObj := hydra.C.Cache().GetRegularCache(cacheName)

	newval, err := cacheObj.Increment(newcachekey, 100)
	if err != nil {
		ctx.Log().Errorf("Cache.Increment:%v", err)
		return err
	}
	return map[string]interface{}{
		"incr.val": newval,
	}
}

var decrementIntHandler = func(ctx hydra.IContext) (r interface{}) {
	newcachekey := "hydra:apiserver:gocache-int"
	cacheObj := hydra.C.Cache().GetRegularCache(cacheName)
	newval, err := cacheObj.Decrement(newcachekey, 50)
	if err != nil {
		ctx.Log().Errorf("Cache.Decrement:%v", err)
		return err
	}
	return map[string]interface{}{
		"decr.val": newval,
	}
}

var existsHandler = func(ctx hydra.IContext) (r interface{}) {
	cacheObj := hydra.C.Cache().GetRegularCache(cacheName)
	newval := cacheObj.Exists(cachekey)
	return map[string]interface{}{
		"exists.val": newval,
	}
}

var deleteHandler = func(ctx hydra.IContext) (r interface{}) {
	cacheObj := hydra.C.Cache().GetRegularCache(cacheName)

	deleteKey := cachekey + "delete"
	cacheObj.Set(deleteKey, "1", -1)
	isExistBefore := cacheObj.Exists(deleteKey)
	err := cacheObj.Delete(deleteKey)
	if err != nil {
		ctx.Log().Errorf("Cache.Delete:%v", err)
		return err
	}
	isExistAfter := cacheObj.Exists(deleteKey)
	return map[string]interface{}{
		"before": isExistBefore,
		"after":  isExistAfter,
	}
}

var delayHandler = func(ctx hydra.IContext) (r interface{}) {
	cacheObj := hydra.C.Cache().GetRegularCache(cacheName)

	delayKey := cachekey + "delay"
	cacheObj.Set(delayKey, "1", -1)
	isExistBefore := cacheObj.Exists(delayKey)
	err := cacheObj.Delay(delayKey, 3)
	if err != nil {
		ctx.Log().Errorf("Cache.Delay:%v", err)
		return err
	}
	time.Sleep(4 * time.Second)
	isExistAfter := cacheObj.Exists(delayKey)

	return map[string]interface{}{
		"before": isExistBefore,
		"after":  isExistAfter,
	}
}

var getsHandler = func(ctx hydra.IContext) (r interface{}) {
	cacheObj := hydra.C.Cache().GetRegularCache(cacheName)

	key1 := cachekey + "1"
	key2 := cachekey + "2"
	key3 := cachekey + "3"

	cacheObj.Set(key1, "1", -1)
	cacheObj.Set(key2, "2", -1)
	cacheObj.Set(key3, "3", -1)

	vals, err := cacheObj.Gets(key1, key2, key3)
	if err != nil {
		ctx.Log().Errorf("Cache.Gets:%v", err)
		return err
	}

	return vals
}
