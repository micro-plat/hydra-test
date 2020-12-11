package main

import (
	"time"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/hydra/registry"
	"github.com/micro-plat/hydra/registry/registry/redis"
	"github.com/micro-plat/lib4go/logger"
)

var create = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("registry-redis-create")
	a, err := registry.GetRegistry(global.Def.RegistryAddr, logger.New("hydra"))
	if err != nil {
		return err
	}
	err = a.CreatePersistentNode("hydratest/node/persistent", "test")
	if err != nil {
		return err
	}
	err = a.CreateTempNode("hydratest/node/temp", "test")
	if err != nil {
		return err
	}
	rpath, err := a.CreateSeqNode("hydratest/node/seq", "test")
	if err != nil {
		return err
	}
	ctx.Log().Info("create.seq.path:", rpath)
	return ""
}

var update = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("registry-redis-update")
	a, err := registry.GetRegistry(global.Def.RegistryAddr, logger.New("hydra"))
	if err != nil {
		return err
	}
	err = a.Update("/hydratest/registry_redis/api/t/conf", `{"address":":18080","status":"start"}`)
	if err != nil {
		return err
	}
	return ""
}

var delete = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("registry-redis-delete")
	a, err := registry.GetRegistry(global.Def.RegistryAddr, logger.New("hydra"))
	if err != nil {
		return err
	}
	err = a.Delete("hydratest/node/persistent")
	if err != nil {
		return err
	}
	return ""
}

var exists = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("registry-redis-exists")
	a, err := registry.GetRegistry(global.Def.RegistryAddr, logger.New("hydra"))
	if err != nil {
		return err
	}
	b, err := a.Exists("/hydratest/registry_redis/api/t/conf")
	if err != nil {
		return err
	}
	b1, err := a.Exists("/hydratest/registry_redis/api/t/conf1")
	if err != nil {
		return err
	}
	b2, err := a.Exists("/hydratest/var")
	if err != nil {
		return err
	}
	return map[string]bool{
		"conf":  b,
		"conf1": b1,
		"var":   b2,
	}

}

var getvalue = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("registry-redis-getvalue")
	a, err := registry.GetRegistry(global.Def.RegistryAddr, logger.New("hydra"))
	if err != nil {
		return err
	}
	v, ver, err := a.GetValue("/hydratest/registry_redis/api/t/conf")
	if err != nil {
		return err
	}
	return map[string]interface{}{
		"version": ver,
		"value":   string(v),
	}
}

var getchildren = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("registry-redis-getchildren")
	a, err := registry.GetRegistry(global.Def.RegistryAddr, logger.New("hydra"))
	if err != nil {
		return err
	}
	v, ver, err := a.GetChildren("/hydratest/var")
	if err != nil {
		return err
	}
	return map[string]interface{}{
		"version":  ver,
		"children": v,
	}
}
var redisopts = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("redisOpts")

	redisObj, err := redis.NewRedisBy("", "", []string{"192.168.5.79:6379"}, 0, 100)

	if err != nil {
		ctx.Log().Error("redis.NewRedisBy:", err)
		return
	}

	opt := ctx.Request().GetString("opt")
	path := ctx.Request().GetString("path")

	ctx.Log().Debug("OPT:", opt)
	ctx.Log().Debug("path:", path)
	result := map[string]interface{}{}

	start := time.Now()

	switch opt {

	case "add":
		result["err"] = redisObj.CreatePersistentNode(path, "vvv")
	case "delete":
		result["err"] = redisObj.Delete(path)
	case "get":
		var bytes []byte
		bytes, result["version"], result["err"] = redisObj.GetValue(path)
		result["bytes"] = string(bytes)
	case "exists":
		result["exists"], result["err"] = redisObj.Exists(path)
	case "getchild":
		result["children"], result["version"], result["err"] = redisObj.GetChildren(path)
	case "hgetall":
		//	result["all"], result["err"] = redisObj.HGetAll(path).Result()
	}
	end := time.Now()
	result["range"] = end.Sub(start).Milliseconds()
	return result
}
