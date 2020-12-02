package main

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/micro-plat/hydra/conf/server/router"
	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/hydra/hydra/servers/pkg/middleware"
)

func (s *CusServer) addHttpRouters(routers ...*router.Router) {
	if !global.IsDebug {
		gin.SetMode(gin.ReleaseMode)
	}
	s.engine = gin.New()
	s.engine.Use(middleware.Logging().GinFunc(CusServerName)) //记录请求日志

	for _, router := range routers {
		for _, method := range router.Action {
			s.engine.Handle(strings.ToUpper(method), router.Path, middleware.ExecuteHandler(router.Service).GinFunc(CusServerName))
		}
	}
	s.server.Handler = s.engine
	return
}
