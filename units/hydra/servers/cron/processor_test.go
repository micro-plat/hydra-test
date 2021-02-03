package cron

import (
	"testing"
	"time"

	//	"github.com/micro-plat/hydra/global"

	"github.com/micro-plat/hydra-test/units/mocks"
	"github.com/micro-plat/hydra/components"
	_ "github.com/micro-plat/hydra/components/caches/cache/redis"
	"github.com/micro-plat/hydra/conf/app"
	"github.com/micro-plat/hydra/conf/server/task"
	"github.com/micro-plat/hydra/conf/vars/cache/cacheredis"
	"github.com/micro-plat/hydra/conf/vars/queue/queueredis"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/hydra/servers/cron"
	"github.com/micro-plat/hydra/services"
	"github.com/micro-plat/lib4go/assert"
)

func TestProcessor_Start(t *testing.T) {
	confObj := mocks.NewConfBy("cronserver_resserivece_testx", "testcronsdfx")
	confObj.CRON()
	confObj.Vars().Redis("redis", "192.168.5.79:6379")
	confObj.Vars().Cache().Redis("redis", "", cacheredis.WithConfigName("redis"))
	confObj.Vars().Queue().Redis("redis", "", queueredis.WithConfigName("redis"))
	app.Cache.Save(confObj.GetCronConf())
	services.Def.CRON("/taosy/services1", func(ctx context.IContext) (r interface{}) {
		queueObj := components.Def.Queue().GetRegularQueue("redis")
		if err := queueObj.Send("key1", `{"services1":"queue1"}`); err != nil {
			ctx.Log().Errorf("发送queue1队列消息异常, err:%v", err)
		}
		return
	})

	services.Def.CRON("/taosy/services2", func(ctx context.IContext) (r interface{}) {
		queueObj := components.Def.Queue().GetRegularQueue("redis")
		if err := queueObj.Send("key2", `{"services2":"queue2"}`); err != nil {
			ctx.Log().Errorf("发送queue1队列消息异常, err:%v", err)
		}
		return
	})

	services.Def.CRON("/taosy/services3", func(ctx context.IContext) (r interface{}) {
		queueObj := components.Def.Queue().GetRegularQueue("redis")
		if err := queueObj.Send("key3", `{"services3":"queue3"}`); err != nil {
			ctx.Log().Errorf("发送queue1队列消息异常, err:%v", err)
		}
		return
	})

	services.Def.CRON("/taosy/services4", func(ctx context.IContext) (r interface{}) {
		queueObj := components.Def.Queue().GetRegularQueue("redis")
		if err := queueObj.Send("key4", `{"services4":"queue4"}`); err != nil {
			ctx.Log().Errorf("发送queue1队列消息异常, err:%v", err)
		}
		return
	})

	s := cron.NewProcessor()
	test1 := task.NewTask("@every 1s", "/taosy/services1")
	test2 := task.NewTask("@every 5s", "/taosy/services2")
	test3 := task.NewTask("@every 10s", "/taosy/services3")
	test4 := task.NewTask("@every 40s", "/taosy/services4")
	err := s.Add(test1, test2, test3, test4)
	assert.Equalf(t, true, err == nil, ",err")
	s.Resume()
	go s.Start()
	time.Sleep(1 * time.Second)
	s.Close()

	cacheObj := components.Def.Cache().GetRegularCache("redis")
	cacheObj.Delete("taosytest:services1:queue1")
	cacheObj.Delete("taosytest:services2:queue2")
	cacheObj.Delete("taosytest:services3:queue3")
	cacheObj.Delete("taosytest:services4:queue4")
}
