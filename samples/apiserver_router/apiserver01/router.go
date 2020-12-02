package main

import "github.com/micro-plat/hydra/context"

type GetStruct struct{}

func (s *GetStruct) GetHandle(ctx context.IContext) interface{} {
	return "get.handle"
}

type PostStruct struct{}

func (s *PostStruct) PostHandle(ctx context.IContext) interface{} {
	return "post.handle"
}

type PutStruct struct{}

func (s *PutStruct) PutHandle(ctx context.IContext) interface{} {
	return "put.handle"
}

type DeleteStruct struct{}

func (s *DeleteStruct) DeleteHandle(ctx context.IContext) interface{} {
	return "delete.handle"
}

type AddStruct struct{}

func (s *AddStruct) Handle(ctx context.IContext) interface{} {
	return "add.handle"
}
