package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/registry"
	"github.com/micro-plat/lib4go/logger"
)

var create = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("registry-redis-create")
	a, err := registry.GetRegistry("redis://192.168.5.79:6379", logger.New("hydra"))
	if err != nil {
		return err
	}
	err = a.CreatePersistentNode("hydra_test/node/persistent", "test")
	if err != nil {
		return err
	}
	err = a.CreateTempNode("hydra_test/node/temp", "test")
	if err != nil {
		return err
	}
	rpath, err := a.CreateSeqNode("hydra_test/node/seq", "test")
	if err != nil {
		return err
	}
	ctx.Log().Info("create.seq.path:", rpath)
	return ""
}

var update = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("registry-redis-update")
	a, err := registry.GetRegistry("redis://192.168.5.79:6379", logger.New("hydra"))
	if err != nil {
		return err
	}
	err = a.Update("/hydra_test/registry_redis/api/t/conf", `{"value":"{"address":":8070","status":"start"}","version":29660859}`)
	if err != nil {
		return err
	}
	return ""
}

var delete = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("registry-redis-delete")
	a, err := registry.GetRegistry("redis://192.168.5.79:6379", logger.New("hydra"))
	if err != nil {
		return err
	}
	err = a.Delete("hydra_test/node/persistent")
	if err != nil {
		return err
	}
	return ""
}

var exists = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("registry-redis-exists")
	a, err := registry.GetRegistry("redis://192.168.5.79:6379", logger.New("hydra"))
	if err != nil {
		return err
	}
	b, err := a.Exists("/hydra_test/registry_redis/api/t/conf")
	if err != nil {
		return err
	}
	b1, err := a.Exists("/hydra_test/registry_redis/api/t/conf1")
	if err != nil {
		return err
	}
	return map[string]bool{
		"conf":  b,
		"conf1": b1,
	}

}

var getvalue = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("registry-redis-getvalue")
	a, err := registry.GetRegistry("redis://192.168.5.79:6379", logger.New("hydra"))
	if err != nil {
		return err
	}
	v, ver, err := a.GetValue("/hydra_test/registry_redis/api/t/conf")
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
	a, err := registry.GetRegistry("redis://192.168.5.79:6379", logger.New("hydra"))
	if err != nil {
		return err
	}
	v, ver, err := a.GetChildren("/hydra_test/registry_redis/api/t/conf")
	if err != nil {
		return err
	}
	return map[string]interface{}{
		"version":  ver,
		"children": v,
	}
}
