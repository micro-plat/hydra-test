package main

import (
	"encoding/json"
	"fmt"

	"github.com/micro-plat/hydra/conf/app"
	"github.com/micro-plat/hydra/hydra/servers"
	"github.com/micro-plat/lib4go/logger"
	"github.com/micro-plat/lib4go/types"
)

//Responsive 响应式服务器
type Responsive struct {
	*CusServer
	conf app.IAPPConf
	log  logger.ILogger
}

//NewResponsive 创建响应式服务器
func NewResponsive(cnf app.IAPPConf) (h *Responsive, err error) {
	h = &Responsive{
		conf: cnf,
		log:  logger.New(cnf.GetServerConf().GetServerName()),
	}
	app.Cache.Save(cnf)
	h.CusServer, err = h.getServer(cnf)
	return h, err
}

//Start 启用服务
func (w *Responsive) Start() (err error) {

	if !w.conf.GetServerConf().IsStarted() {
		w.log.Warnf("%s被禁用，未启动", w.conf.GetServerConf().GetServerType())
		return
	}

	if err = w.CusServer.Start(); err != nil {
		err = fmt.Errorf("%s启动失败 %w", w.conf.GetServerConf().GetServerType(), err)
		return
	}

	w.log.Infof("启动成功(%s,%s,[%d])", w.conf.GetServerConf().GetServerType(), w.CusServer.GetAddress(), w.serverNum())
	return nil
}

//Notify 服务器配置变更通知
func (w *Responsive) Notify(c app.IAPPConf) (change bool, err error) {

	return true, nil
}

//Shutdown 关闭服务器
func (w *Responsive) Shutdown() {
	w.log.Infof("关闭[%s]服务...", w.conf.GetServerConf().GetServerType())
	w.CusServer.Shutdown()
	return
}

//serverNum 获取服务数量
func (w *Responsive) serverNum() int {
	routers := w.CusServer.engine.Routes()
	serverMap := map[string]string{}
	for _, item := range routers {
		if _, ok := serverMap[item.Path]; !ok {
			serverMap[item.Path] = item.Path
		}
	}
	return len(serverMap)
}

//根据main.conf创建服务嚣
func (w *Responsive) getServer(cnf app.IAPPConf) (*CusServer, error) {
	tp := cnf.GetServerConf().GetServerType()
	srvCnf := cnf.GetServerConf()
	data := types.XMap{}
	srvCnf.GetMainObject(&data)

	rawConf, err := srvCnf.GetSubConf("router")
	rawBytes := rawConf.GetRaw()
	fmt.Println("rawBytes:", string(rawBytes))
	routers := &RouterList{}
	json.Unmarshal(rawBytes, routers)

	if err != nil {
		return nil, err
	}
	return NewServer(tp,
		data.GetString("address"),
		routers,
	)

}

func init() {
	fn := func(c app.IAPPConf) (servers.IResponsiveServer, error) {
		return NewResponsive(c)
	}
	servers.Register(CusServerName, fn)
}
