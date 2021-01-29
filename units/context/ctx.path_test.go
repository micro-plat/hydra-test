package context

import (
	"testing"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra-test/units/mocks"
	"github.com/micro-plat/hydra/conf"
	"github.com/micro-plat/hydra/conf/app"
	"github.com/micro-plat/hydra/conf/server/api"
	c "github.com/micro-plat/hydra/conf/server/cron"
	"github.com/micro-plat/hydra/conf/server/queue"
	"github.com/micro-plat/hydra/conf/vars/queue/queueredis"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/context/ctx"
	"github.com/micro-plat/hydra/services"
	"github.com/micro-plat/lib4go/assert"
)

func Test_rpath_GetEncoding(t *testing.T) {
	confObj := mocks.NewConf() //构建对象
	confObj.API("8080")
	hydra.G.SysName = "apiserver"

	confObj.Vars().Redis("5.79", "192.168.5.79:6379")
	confObj.Vars().Queue().Redis("xxx", "", queueredis.WithConfigName("5.79"))
	confObj.MQC("redis://xxx").Queue(queue.NewQueue("queue1", "/service1")).Queue(queue.NewQueue("queue2", "/service2"))
	confObj.CRON(c.WithMasterSlave(), c.WithTrace())

	services.API.Add("/api", "/api", []string{"GET"}, api.WithEncoding("utf-8"))
	services.API.Add("/api2", "/api2", []string{"GET"}, api.WithEncoding("gbk"))
	services.API.Add("/api3", "/api3", []string{"GET"})
	services.WEB.Add("/web", "/web", []string{"GET"}, api.WithEncoding("utf-8"))
	services.WEB.Add("/web2", "/web2", []string{"GET"}, api.WithEncoding("gbk"))
	services.WEB.Add("/web3", "/web3", []string{"GET"})
	services.WS.Add("/ws", "/ws", []string{"GET"}, api.WithEncoding("utf-8"))
	services.WS.Add("/ws2", "/ws2", []string{"GET"}, api.WithEncoding("gbk"))
	services.WS.Add("/ws3", "/ws3", []string{"GET"})
	apiConf := confObj.GetAPIConf() //获取配置

	hydra.G.SysName = "webserver"
	webConf := confObj.GetWebConf() //获取配置

	hydra.G.SysName = "wsserver"
	wsConf := confObj.GetWSConf() //获取配置

	hydra.G.SysName = "mqcserver"
	mqcConf := confObj.GetMQCConf() //获取配置

	hydra.G.SysName = "cronserver"
	cronConf := confObj.GetCronConf() //获取配置

	tests := []struct {
		name       string
		ctx        context.IInnerContext
		serverConf app.IAPPConf
		meta       conf.IMeta
		want       string
	}{
		{name: "1.1 api类型,注册时设置encoding为utf-8", ctx: &mocks.TestContxt{Routerpath: "/api", Method: "GET"}, serverConf: apiConf, meta: conf.NewMeta(), want: "utf-8"},
		{name: "1.2 api类型,注册时设置encoding为gbk", ctx: &mocks.TestContxt{Routerpath: "/api2", Method: "GET"}, serverConf: apiConf, meta: conf.NewMeta(), want: "gbk"},
		{name: "1.3 api类型,注册时未设置encoding", ctx: &mocks.TestContxt{Routerpath: "/api3", Method: "GET"}, serverConf: apiConf, meta: conf.NewMeta(), want: "utf-8"},
		{name: "2.1 web类型,注册时设置encoding为utf-8", ctx: &mocks.TestContxt{Routerpath: "/web", Method: "GET"}, serverConf: webConf, meta: conf.NewMeta(), want: "utf-8"},
		{name: "2.2 web类型,注册时设置encoding为gbk", ctx: &mocks.TestContxt{Routerpath: "/web2", Method: "GET"}, serverConf: webConf, meta: conf.NewMeta(), want: "gbk"},
		{name: "2.3 web类型,注册时未设置encoding", ctx: &mocks.TestContxt{Routerpath: "/web3", Method: "GET"}, serverConf: webConf, meta: conf.NewMeta(), want: "utf-8"},
		{name: "3.1 ws类型,注册时设置encoding为utf-8", ctx: &mocks.TestContxt{Routerpath: "/ws", Method: "GET"}, serverConf: wsConf, meta: conf.NewMeta(), want: "utf-8"},
		{name: "3.2 ws类型,注册时设置encoding为gbk", ctx: &mocks.TestContxt{Routerpath: "/ws2", Method: "GET"}, serverConf: wsConf, meta: conf.NewMeta(), want: "gbk"},
		{name: "3.3 ws类型,注册时未设置encoding", ctx: &mocks.TestContxt{Routerpath: "/ws3", Method: "GET"}, serverConf: wsConf, meta: conf.NewMeta(), want: "utf-8"},
		{name: "4 cron类型获取默认值", ctx: &mocks.TestContxt{Routerpath: "/mqc", Method: "GET"}, serverConf: mqcConf, meta: conf.NewMeta(), want: "utf-8"},
		{name: "5 mqc类型获取默认值", ctx: &mocks.TestContxt{Routerpath: "/cron", Method: "GET"}, serverConf: cronConf, meta: conf.NewMeta(), want: "utf-8"},
	}
	for _, tt := range tests {
		c := ctx.NewRpath(tt.ctx, tt.serverConf, tt.meta)
		got := c.GetEncoding()
		assert.Equal(t, tt.want, got, tt.name)

		//再次获取
		got2 := c.GetEncoding()
		assert.Equal(t, got, got2, tt.name)

	}
}
