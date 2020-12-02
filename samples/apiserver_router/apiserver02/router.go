package main

import "github.com/micro-plat/hydra"

type SrvStruct struct{}

func (s *SrvStruct) GetHandle(ctx hydra.IContext) interface{} {
	return "get.handle"
}

func (s *SrvStruct) PostHandle(ctx hydra.IContext) interface{} {
	return "post.handle"
}

//Handle 注释该方法，可正常运行
func (s *SrvStruct) Handle(ctx hydra.IContext) interface{} {
	return ".handle"
}
