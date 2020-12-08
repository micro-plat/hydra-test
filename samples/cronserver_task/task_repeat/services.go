package main

import (
	"encoding/json"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/task"
	"github.com/micro-plat/hydra/registry"
	"github.com/micro-plat/lib4go/logger"
)

var getCron = func(ctx hydra.IContext) (r interface{}) {
	reg, err := registry.GetRegistry("lm://./", logger.New("hydra"))
	if err != nil {
		return err
	}
	value, _, err := reg.GetValue("/hydra_test/task_repeat/cron/t/conf/task")
	if err != nil {
		return err
	}
	tasks := &task.Tasks{}
	err = json.Unmarshal(value, tasks)
	if err != nil {
		return err
	}
	for k, v := range tasks.Tasks {
		ctx.Log().Infof("%d:任务名[%s],服务[%s]", k, v.Cron, v.Service)
	}
	return
}
