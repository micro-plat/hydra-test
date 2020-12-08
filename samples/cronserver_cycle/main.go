package main

import (
	"time"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/task"
	"github.com/micro-plat/hydra/hydra/servers/cron"
	"github.com/micro-plat/hydra/hydra/servers/http"
	"github.com/micro-plat/lib4go/concurrent/cmap"
	"github.com/micro-plat/lib4go/types"
)

var app = hydra.NewApp(
	hydra.WithDebug(),
	hydra.WithServerTypes(http.API, cron.CRON),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("cronserverCycle"),
	hydra.WithClusterName("taosytest"),
	hydra.WithRegistry("zk://192.168.0.101"),
)

func init() {
	hydra.Conf.API(":8070")
	hydra.Conf.CRON().Task(
		task.NewTask("@now", "/hydratest/cronserverCycle/cron2"),
		task.NewTask("@every 5m", "/hydratest/cronserverCycle/cron4"),
		task.NewTask("@every 24h", "/hydratest/cronserverCycle/cron6"),
		task.NewTask("30 6 * * 5", "/hydratest/cronserverCycle/cron8"),
	)
	app.API("/hydratest/cronserverCycle/show", funcAPI)
	app.CRON("/hydratest/cronserverCycle/cron2", funcCycle2)
	app.CRON("/hydratest/cronserverCycle/cron3", funcCycle3, "@every 10s")
	app.CRON("/hydratest/cronserverCycle/cron4", funcCycle4)
	app.CRON("/hydratest/cronserverCycle/cron5", funcCycle5, "@every 2h")
	app.CRON("/hydratest/cronserverCycle/cron6", funcCycle6)
	app.CRON("/hydratest/cronserverCycle/cron7", funcCycle7, "30 6 * * *")
	app.CRON("/hydratest/cronserverCycle/cron8", funcCycle8)
	app.CRON("/hydratest/cronserverCycle/cron9", funcCycle9, "30 6 * 5,6 *")
}

// cronserver_cycle 对于不同cron配置循环执行次数测试demo
//1.1 安装程序 ./cronserver_cycle conf install -cover
//1.2 使用 ./cronserver_cycle run

//1.3 调用接口：http://localhost:8070/hydratest/cronserverCycle/show  查看当前的执行次数情况
func main() {
	uuidMap.Set("startTime", time.Now().Format("2006-01-02 15:04:05"))
	app.Start()
}

var uuidMap = cmap.New(1)

var funcAPI = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("cronserver_cycle 查询所有任务的执行次数情况")
	uuidMap.Set("nowTime", time.Now().Format("2006-01-02 15:04:05"))
	res := uuidMap.Items()
	return res
}

var funcCycle2 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("cronserver_cycle now单次执行次数")
	c, e := uuidMap.Get("now")
	if !e {
		uuidMap.Set("now", 1)
	} else {
		count := types.GetInt(c) + 1
		uuidMap.Set("now", count)
	}
	return
}

var funcCycle3 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("cronserver_cycle 每10s执行一次")
	c, e := uuidMap.Get("10s")
	if !e {
		uuidMap.Set("10s", 1)
	} else {
		count := types.GetInt(c) + 1
		uuidMap.Set("10s", count)
	}
	return
}

var funcCycle4 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("cronserver_cycle 每5m执行一次")
	c, e := uuidMap.Get("5m")
	if !e {
		uuidMap.Set("5m", 1)
	} else {
		count := types.GetInt(c) + 1
		uuidMap.Set("5m", count)
	}
	return
}

var funcCycle5 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("cronserver_cycle 每2h执行一次")
	c, e := uuidMap.Get("2h")
	if !e {
		uuidMap.Set("2h", 1)
	} else {
		count := types.GetInt(c) + 1
		uuidMap.Set("2h", count)
	}
	return
}

var funcCycle6 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("cronserver_cycle 启动时间开始每天执行一次")
	c, e := uuidMap.Get("everyday")
	if !e {
		uuidMap.Set("everyday", 1)
	} else {
		count := types.GetInt(c) + 1
		uuidMap.Set("everyday", count)
	}
	return
}

var funcCycle7 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("cronserver_cycle 每天上午6:30执行一次")
	c, e := uuidMap.Get("EveryDay6-30")
	if !e {
		uuidMap.Set("EveryDay6-30", 1)
	} else {
		count := types.GetInt(c) + 1
		uuidMap.Set("EveryDay6-30", count)
	}
	return
}

var funcCycle8 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("cronserver_cycle 每周星期五上午6:30执行一次")
	c, e := uuidMap.Get("Week6-30")
	if !e {
		uuidMap.Set("Week6-30", 1)
	} else {
		count := types.GetInt(c) + 1
		uuidMap.Set("Week6-30", count)
	}
	return
}

var funcCycle9 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("cronserver_cycle 每月5.6号上午6:30执行一次")
	c, e := uuidMap.Get("Month6-30")
	if !e {
		uuidMap.Set("Month6-30", 1)
	} else {
		count := types.GetInt(c) + 1
		uuidMap.Set("Month6-30", count)
	}
	return
}
