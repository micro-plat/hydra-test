package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/static"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/hydra/servers/http"
	_ "github.com/micro-plat/hydra/hydra/servers/rpc"
)

var opts []static.Option

var app = hydra.NewApp(
	hydra.WithServerTypes(http.Web),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("webserverstatic"),
	hydra.WithClusterName("test"),
	hydra.WithRegistry("lm://."),
	hydra.WithDebug(),
)

//go build -mod=mod -tags=embeddir
//go build -mod=mod -tags=embedzip
//go build -mod=mod -tags=osdir
//go build -mod=mod -tags=oszip

func main() {
	hydra.Conf.Web("50005").Static(opts...)
	app.Micro("/struct", &GetStruct{})
	app.Micro("/func", singleFunc)
	app.Micro("/func.view", singleFunc)
	app.Start()
}

func singleFunc(ctx context.IContext) interface{} {
	return "func.handle"
}

type GetStruct struct{}

func (s *GetStruct) GetHandle(ctx context.IContext) interface{} {
	return "get.handle"
}

func (s *GetStruct) PostHandle(ctx context.IContext) interface{} {
	return "post.handle"
}

func (s *GetStruct) PutHandle(ctx context.IContext) interface{} {
	return "put.handle"
}
