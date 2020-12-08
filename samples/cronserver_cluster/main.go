package main

import (
	"encoding/json"
	"time"

	"github.com/micro-plat/hydra"
	confcron "github.com/micro-plat/hydra/conf/server/cron"
	"github.com/micro-plat/hydra/hydra/servers/cron"
	"github.com/micro-plat/lib4go/concurrent/cmap"
	"github.com/micro-plat/lib4go/types"
)

var app = hydra.NewApp(
	hydra.WithDebug(),
	hydra.WithServerTypes(cron.CRON),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("cronservercluster"),
	hydra.WithClusterName("taosytest"),
	hydra.WithRegistry("zk://192.168.0.101"),
)

func init() {
	hydra.Conf.CRON(confcron.WithMasterSlave())
	app.CRON("/hydratest/cronservercluster/cron1", funcCycle1, "@now")
	app.CRON("/hydratest/cronservercluster/cron2", funcCycle2, "@once")
	app.CRON("/hydratest/cronservercluster/cron3", funcCycle3, "@every 10s")
}

// cronserver_cluster 集群模式：对等、主从、分片，变更后自动切换模式测试demo
//因为是集群情况测试，所以需要运行多台服务器，分别复制cronserver_cluster1和cronserver_cluster2两个执行文件
//1.1 安装程序：./cronserver_cluster conf install -cover
//1.2 运行cronserver_cluster： ./cronserver_cluster run
//1.3 运行cronserver_cluster1： ./cronserver_cluster1 run
//1.4 运行cronserver_cluster2： ./cronserver_cluster2 run
//1.5 直接观察服务器运行的日志了解运行情况是否符合预期

/*
1. 初始化为主备用模式,预期结果:（一台master，两台slave）
	1.1.三台服务器启动分别为一台master和两台slave服务器;
	1.2 master服务器按照配置执行任务，slave服务器不执行;
	1.3 关闭master服务器,其中一台slave服务器自动转为master服务器，并且按照配置要求执行任务;

2. 主从模式修改为对等模式，预期结果：
	2.1 两台slave服务器会重启服务器变为master服务器;原来的master服务器保持原有运行状态;（三台master，无slave）

3. 对等模式修为分片模式，预期结果：(因为集群只有三台服务器，所以直接分为2片)
	3.1 其中一台master服务器会变更为slave服务器，另外两台master服务器保持原有运行状态;（两台master，一台slave）

4.分片模式修改为主从模式，预期结果：
	4.1 两台master中的其中一台会变更为slave服务器暂停运行。剩下的一台master服务器保持原有运行状态;（一台master，两台slave）

5. 主从模式修改为分片模式，预期结果：(因为集群只有三台服务器，所以直接分为2片)
	5.1 两台slave服务器的其中一台会重启服务器变为master服务器，按照配置要求执行任务;原来的master服务器保持原有运行状态;（两台master，一台slave）

6. 分片模式修改为对等模式，预期结果：
	6.1 slave服务器会重启服务器变为master服务器;原来的master服务器保持原有运行状态;（三台master，无slave）

7. 对等模式修改为主从模式，预期结果：
	7.1 三台master中的其中两台会变为slave服务器暂停执行服务;剩下的一台master服务器保持原有运行状态;（一台master，两台slave）
*/

func main() {
	uuidMap.Set("startTime", time.Now().Format("2006-01-02 15:04:05"))
	app.Start()
}

var uuidMap = cmap.New(1)

var funcCycle1 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("cronserver_cluster now单次执行次数")
	c, e := uuidMap.Get("now")
	if !e {
		uuidMap.Set("now", 1)
	} else {
		count := types.GetInt(c) + 1
		uuidMap.Set("now", count)
	}
	res := uuidMap.Items()
	b, _ := json.Marshal(res)
	ctx.Log().Info("--------程序执行情况汇总：", string(b))
	return
}

var funcCycle2 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("cronserver_cluster once单次执行次数")
	c, e := uuidMap.Get("once")
	if !e {
		uuidMap.Set("once", 1)
	} else {
		count := types.GetInt(c) + 1
		uuidMap.Set("once", count)
	}

	res := uuidMap.Items()
	b, _ := json.Marshal(res)
	ctx.Log().Info("--------程序执行情况汇总：", string(b))
	return
}

var funcCycle3 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("cronserver_cluster 每10s执行一次")
	c, e := uuidMap.Get("10s")
	if !e {
		uuidMap.Set("10s", 1)
	} else {
		count := types.GetInt(c) + 1
		uuidMap.Set("10s", count)
	}
	res := uuidMap.Items()
	b, _ := json.Marshal(res)
	ctx.Log().Info("--------程序执行情况汇总：", string(b))
	return
}
