package main

import (
	"fmt"
	"time"

	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/hydra/registry"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/conf/server/header"
	"github.com/micro-plat/hydra/conf/vars/cache/cacheredis"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var hydraApp = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("apiserver_redis"),
	hydra.WithClusterName("test"),
	hydra.WithRegistry("lm://."),
)

const redisName = "5.79"
const redisAddr = "192.168.5.79:6379"
const cacheName = "cacheredis"

func init() {

	hydra.Conf.Vars().Redis(redisName, redisAddr)
	hydra.Conf.Vars().Cache().Redis(cacheName, "", cacheredis.WithConfigName(redisName))

	hydra.Conf.API(":50023", api.WithTimeout(10, 10)).Header(header.WithHeader("content-type", "application/json"))

	hydraApp.API("/api/redis/add", addHandler)
	hydraApp.API("/api/redis/set", setHandler)
	hydraApp.API("/api/redis/incr", incrementHandler)
	hydraApp.API("/api/redis/decr", decrementHandler)
	hydraApp.API("/api/redis/delete", deleteHandler)
	hydraApp.API("/api/redis/delay", delayHandler)
	hydraApp.API("/api/redis/gets", getsHandler)

	hydraApp.API("/api/redis/config", config)

}

// apiserver_db 数据库组件是否正确工作，修改配置是否自动生效（mysql）
// 1. 编译程序： go build
// 2. 启动程序：./apiserver_mysql run

// 3. 请求 http://localhost:50023/api/redis/add 添加一个值,再次添加报错[已存在]
// 4. 请求 http://localhost:50023/api/redis/set 设置值，存在即覆盖
// 5. 请求 http://localhost:50023/api/redis/incr 进行加处理,每次递增100
// 6. 请求 http://localhost:50023/api/redis/decr 进行减处理，每次递减50
// 7. 请求 http://localhost:50023/api/redis/delete 验证删除，添加值，判断是否存在，删除，再判定值是否存在
// 8. 请求 http://localhost:50023/api/redis/delay 验证delay 自动删除，设置值1,睡眠4秒，再检测值是否存在
// 8. 请求 http://localhost:50023/api/redis/gets 获取多个key值
// 8. 请求 http://localhost:50023/api/redis/config 修改cache的配置
func main() {
	hydraApp.Start()
}

const cachekey = "hydratest:apiserver:redis"

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
	cacheObj := hydra.C.Cache().GetRegularCache(cacheName)
	err := cacheObj.Set(cachekey, "10", -1)
	if err != nil {
		ctx.Log().Errorf("Cache.Set:%v", err)
		return err
	}
	val, err := cacheObj.Get(cachekey)
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

var config = func(ctx hydra.IContext) (r interface{}) {
	fmt.Println("0")
	regst, err := registry.GetRegistry(global.Def.RegistryAddr, global.Def.Log())
	if err != nil {
		return fmt.Errorf("NewRegistry:%v", err)
	}
	fmt.Println("1")
	dbpath := fmt.Sprintf("/hydratest/var/cache/%s", cacheName)
	err = regst.Update(dbpath, `{"proto":"redis","addrs":["192.168.0.101:6379"],"dial_timeout":10,"read_timeout":10,"write_timeout":10,"pool_size":10}`)
	if err != nil {
		return fmt.Errorf("UpdateCache:%v", err)
	}
	fmt.Println("2")
	path := "/hydratest/apiserver_redis/api/test/conf"
	err = regst.Update(path, `{"status":"start","address":":50023"}`)
	if err != nil {
		return fmt.Errorf("UpdateConf:%v", err)
	}
	fmt.Println("3")
	return "success"
}
