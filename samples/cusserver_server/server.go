package main

import (
	"context"
	"fmt"
	xnet "net"
	x "net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/lib4go/types"
)

const CusServerName = "customserver"

//CusServer api服务器
type CusServer struct {
	server  *x.Server
	engine  *gin.Engine
	running bool
	ip      string
	proto   string
	host    string
	port    string
}

//NewServer 创建http api服务嚣
func NewServer(name string, addr string, routers *RouterList) (t *CusServer, err error) {
	t, err = new(name, addr)
	if err != nil {
		return
	}
	t.addHttpRouters(routers)
	return
}

//new 创建http api服务嚣
func new(name string, addr string) (t *CusServer, err error) {
	t = &CusServer{
		proto: "http",
		ip:    global.LocalIP(), // net.GetLocalIPAddress(),

	}

	t.host, t.port, err = global.GetHostPort(addr)
	if err != nil {
		return nil, err
	}

	t.server = &x.Server{
		Addr:              xnet.JoinHostPort(t.host, t.port),
		ReadHeaderTimeout: time.Second * 6,
		ReadTimeout:       time.Second * 6,
		WriteTimeout:      time.Second * 6,
		MaxHeaderBytes:    1 << 20,
	}
	return
}

// Start the http server
func (s *CusServer) Start() error {
	s.running = true
	errChan := make(chan error, 1)

	go func(ch chan error) {
		if err := s.server.ListenAndServe(); err != nil {
			ch <- err
		}
	}(errChan)

	select {
	case <-time.After(time.Millisecond * 500):
		return nil
	case err := <-errChan:
		s.running = false
		return err
	}
}

//Shutdown 关闭服务器
func (s *CusServer) Shutdown() error {
	if s.server != nil && s.running {
		s.running = false
		ctx, cannel := context.WithTimeout(context.Background(), time.Second*10)
		defer cannel()
		if err := s.server.Shutdown(ctx); err != nil {
			if err == x.ErrServerClosed {
				return nil
			}
			return fmt.Errorf("关闭出现错误:%w", err)
		}
	}
	return nil
}

//GetAddress 获取当前服务地址
func (s *CusServer) GetAddress(h ...string) string {
	if len(h) > 0 && h[0] != "" {
		return fmt.Sprintf("%s://%s:%s", s.proto, h[0], s.port)
	}
	if s.host == "0.0.0.0" {
		return fmt.Sprintf("%s://%s:%s", s.proto, s.ip, s.port)
	}
	return fmt.Sprintf("%s://%s:%s", s.proto, s.host, s.port)
}

//GetStatus 获取当前服务器状态
func (s *CusServer) GetStatus() string {
	return types.DecodeString(s.running, true, "运行中", "停止")
}
