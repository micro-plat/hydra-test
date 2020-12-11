package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/hydra/registry"
	"github.com/micro-plat/lib4go/logger"
)

var addCron1 = func(ctx hydra.IContext) (r interface{}) {

	reg, err := registry.GetRegistry(global.Def.RegistryAddr, logger.New("hydra"))
	if err != nil {
		return err
	}
	value := `{"tasks":[{"cron":"@every 10s","service":"/cron"},{"cron":"@every 10s","service":"/cron2"}]}`
	err = reg.Update("/hydratest/task_conf/cron/t/conf/task", value)
	if err != nil {
		return err
	}
	return reg.Update("/hydratest/task_conf/cron/t/conf", `{"status":"start"}`)
}

var deleteCron1 = func(ctx hydra.IContext) (r interface{}) {
	reg, err := registry.GetRegistry(global.Def.RegistryAddr, logger.New("hydra"))
	if err != nil {
		return err
	}
	value := `{"tasks":[{"cron":"@every 10s","service":"/cron2"}]}`
	err = reg.Update("/hydratest/task_conf/cron/t/conf/task", value)
	if err != nil {
		return err
	}
	return reg.Update("/hydratest/task_conf/cron/t/conf", `{"status":"start"}`)
}

var updateCron2 = func(ctx hydra.IContext) (r interface{}) {
	reg, err := registry.GetRegistry(global.Def.RegistryAddr, logger.New("hydra"))
	if err != nil {
		return err
	}
	value := `{"tasks":[{"cron":"@every 10s","service":"/cron"},{"cron":"@every 1m","service":"/cron2"}]}`
	err = reg.Update("/hydratest/task_conf/cron/t/conf/task", value)
	if err != nil {
		return err
	}
	return reg.Update("/hydratest/task_conf/cron/t/conf", `{"status":"start"}`)
}
