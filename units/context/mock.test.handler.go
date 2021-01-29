package context

import "github.com/micro-plat/hydra/context"

type Order struct {
	Result interface{}
}

func (o *Order) RequestHandle(ctx *context.IContext) interface{} {
	return o.Result
}
