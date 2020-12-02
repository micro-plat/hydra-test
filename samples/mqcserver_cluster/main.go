package main

import (
	"fmt"
	"time"

	"github.com/micro-plat/hydra"
	mqcconf "github.com/micro-plat/hydra/conf/server/mqc"
	"github.com/micro-plat/hydra/conf/vars/queue/queueredis"
	"github.com/micro-plat/hydra/conf/vars/redis"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/hydra/servers/mqc"
	"github.com/micro-plat/hydra/registry"
	"github.com/micro-plat/lib4go/logger"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(mqc.MQC),
	hydra.WithPlatName("hydra_test"),
	hydra.WithSystemName("mqc_cluster_master"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

func init() {
	hydra.Conf.Vars().Redis("5.79", redis.New([]string{"192.168.5.79:6379"}))
	hydra.Conf.Vars().Queue().Redis("xxx", queueredis.New(queueredis.WithConfigName("5.79")))
	hydra.Conf.MQC("redis://xxx", mqcconf.WithMasterSlave()) //设置为主从模式
	app.MQC("/mqc", func(ctx context.IContext) (r interface{}) {
		ctx.Log().Info("mqc")
		return
	}, "mqcName")
}

//消息队列服务器异常关闭后正常启动，服务是否自动恢复
//启动服务 ./cluster_master run 启动3个mqc服务
//查询服务器在集群模式下相互切换时是否正常
func main() {
	go updateServerStatus()
	app.Start()
}

func updateServerStatus() {
	for k := 0; k < 54; k++ {
		time.Sleep(time.Second * 5)
		sharding := 0
		if k%3 == 0 {
			sharding = 0
		}
		if (k+1)%3 == 0 {
			sharding = 1
		}
		if (k+2)%3 == 0 {
			sharding = 2
		}
		reg, _ := registry.NewRegistry("lm://./", logger.New("hydra"))
		path := "/hydra_test/mqc_cluster_master/mqc/t/conf"
		err := reg.Update(path, fmt.Sprintf(`{"status":"start","sharding":%d,"addr":"redis://xxx"}`, sharding))
		fmt.Println(err)
	}

}
