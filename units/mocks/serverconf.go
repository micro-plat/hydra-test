package mocks

import (
	"fmt"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/app"
	"github.com/micro-plat/hydra/creator"
	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/hydra/hydra/servers/cron"
	"github.com/micro-plat/hydra/hydra/servers/http"
	"github.com/micro-plat/hydra/hydra/servers/mqc"
	"github.com/micro-plat/hydra/registry"
	_ "github.com/micro-plat/hydra/registry/registry/localmemory"
	"github.com/micro-plat/hydra/services"
)

type service struct {
	API *services.ORouter
	Web *services.ORouter
	WS  *services.ORouter
	RPC *services.ORouter
}

//SConf 服务器配置
type SConf struct {
	creator.IConf
	PlatName     string
	ClusterName  string
	Service      *service
	registryAddr string
	Registry     registry.IRegistry
}

//NewConf 构建配置信息
func NewConf() *SConf {
	return NewConfBy("hydra", "test")
}

//NewConfBy 构建配置信息
func NewConfBy(platName, clusterName string, addr ...string) *SConf {
	c := &SConf{
		PlatName:    platName,
		ClusterName: clusterName,
		Service:     &service{},
		//registryAddr: "zk://192.168.0.101",
		registryAddr: "lm://.",
	}
	if len(addr) > 0 {
		c.registryAddr = addr[0]
	}
	hydra.G.PlatName = platName
	hydra.G.ClusterName = clusterName
	//API  路由信息
	c.Service.API = services.NewORouter()

	//WEB web服务的路由信息
	c.Service.Web = services.NewORouter()

	//WS web socket路由信息
	c.Service.WS = services.NewORouter()

	//RPC rpc服务的路由信息
	c.Service.RPC = services.NewORouter()

	c.IConf = creator.Conf // creator.NewByLoader(c.getRouter)
	var err error
	c.Registry, err = registry.GetRegistry(c.registryAddr, global.Def.Log())
	if err != nil {
		panic(err)
	}

	//处理iconf.load中，服务检查问题
	global.Def.ServerTypes = []string{http.API, http.Web, http.WS, cron.CRON}
	return c
}

//Conf 配置
func (s *SConf) Conf() creator.IConf {
	return s.IConf
}

//GetAPIConf 获取API服务器配置
func (s *SConf) GetAPIConf() app.IAPPConf {
	hydra.G.PlatName = s.PlatName
	hydra.G.SysName = "apiserver"
	hydra.G.ClusterName = s.ClusterName
	conf := s.GetConf(s.PlatName, "apiserver", "api", s.ClusterName)
	return conf
}

//GetWebConf 获取web服务器配置
func (s *SConf) GetWebConf() app.IAPPConf {
	hydra.G.PlatName = s.PlatName
	hydra.G.SysName = "webserver"
	hydra.G.ClusterName = s.ClusterName
	conf := s.GetConf(s.PlatName, "webserver", "web", s.ClusterName)
	return conf
}

//GetWSConf 获取API服务器配置
func (s *SConf) GetWSConf() app.IAPPConf {
	hydra.G.PlatName = s.PlatName
	hydra.G.SysName = "wsserver"
	hydra.G.ClusterName = s.ClusterName
	conf := s.GetConf(s.PlatName, "wsserver", "ws", s.ClusterName)
	return conf
}

//GetCronConf 获取cron服务器配置
func (s *SConf) GetCronConf() app.IAPPConf {
	hydra.G.PlatName = s.PlatName
	hydra.G.SysName = "cronserver"
	hydra.G.ClusterName = s.ClusterName
	conf := s.GetConf(s.PlatName, "cronserver", "cron", s.ClusterName)
	return conf
}

//GetMQCConf 获取mqc服务器配置
func (s *SConf) GetMQCConf() app.IAPPConf {
	hydra.G.PlatName = s.PlatName
	hydra.G.SysName = "mqcserver"
	hydra.G.ClusterName = s.ClusterName
	global.Def.ServerTypes = []string{http.API, http.Web, http.WS, cron.CRON, mqc.MQC}
	return s.GetConf(s.PlatName, "mqcserver", "mqc", s.ClusterName)
}

//GetRPCConf 获取rpc服务器配置
func (s *SConf) GetRPCConf() app.IAPPConf {
	hydra.G.PlatName = s.PlatName
	hydra.G.SysName = "rpcserver"
	hydra.G.ClusterName = s.ClusterName
	return s.GetConf(s.PlatName, "rpcserver", "rpc", s.ClusterName)
}

//GetConf 获取配置信息
func (s *SConf) GetConf(platName string, systemName string, serverType string, clusterName string) app.IAPPConf {

	if err := s.IConf.Pub(platName, systemName, clusterName, s.registryAddr, true); err != nil {
		panic(err)
	}

	path := registry.Join(platName, systemName, serverType, clusterName, "conf")
	conf, err := app.NewAPPConf(path, s.Registry)
	if err != nil {
		panic(err)
	}
	return conf
}

//GetRouter 获取服务器的路由配置
func (s *SConf) getRouter(tp string) *services.ORouter {
	switch tp {
	case global.API:
		return s.Service.API
	case global.Web:
		return s.Service.Web
	case global.WS:
		return s.Service.WS
	case global.RPC:
		return s.Service.RPC
	default:
		panic(fmt.Sprintf("无法获取服务%s的路由配置", tp))
	}
}
