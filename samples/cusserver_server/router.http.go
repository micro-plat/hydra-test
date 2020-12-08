package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/micro-plat/hydra/global"
)
var handlers map[string]gin.HandlerFunc

var onceLock sync.Once 

 

type RouterList struct {
	List []*Router `json:"routers"`
}

type Router struct {
	Action  []string `json:"action"`
	Service string   `json:"service"`
}


func (s *CusServer) addHttpRouters(routers *RouterList) {
	if !global.IsDebug {
		gin.SetMode(gin.ReleaseMode)
	}
	s.engine = gin.New()

	for _, router := range routers.List {
		if len(router.Action) == 0 {
			router.Action = []string{"get", "post"}
		}
		for _, action := range router.Action {
			fmt.Println("gin:", router.Service, action)
			s.engine.Handle(strings.ToUpper(action), router.Service, GetHandler(router.Service))
		}
	}
	s.server.Handler = s.engine
	return
}



func GetHandler(service string) gin.HandlerFunc {
	if handler, ok := handlers[service]; ok {
		return handler
	}
	return nil
}

func Registry(service string, handler gin.HandlerFunc) {
	onceLock.Do(func(){
		handlers = map[string]gin.HandlerFunc{}
	})	
	handlers[service] = handler
}
