package main

import (
	"fmt"

	"github.com/micro-plat/hydra/registry"
)

const registrycustom = "custom"

//customFactory 基于文件的注册中心
type customFactory struct {
	opts *registry.Options
}

//Resolve 根据配置生成zookeeper客户端
func (z *customFactory) Create(opts ...registry.Option) (registry.IRegistry, error) {
	for i := range opts {
		opts[i](z.opts)
	}

	if len(z.opts.Addrs) <= 0 {
		return nil, fmt.Errorf("FS注册中心，需要指定一个本地目录地址：%v", z.opts.Addrs)
	}
	if len(z.opts.Addrs) > 1 {
		return nil, fmt.Errorf("FS注册中心，只允许传递一个本地目录地址：%v", z.opts.Addrs)
	}
	rootDir := z.opts.Addrs[0]
	client, err := NewFileSystem(rootDir)
	if err != nil {
		return nil, err
	}
	client.Start()
	return client, nil
}
