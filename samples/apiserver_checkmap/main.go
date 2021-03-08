package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

//服务器各种返回结果
func main() {
	app := hydra.NewApp(
		hydra.WithPlatName("hydra"),
		hydra.WithServerTypes(http.API),
		hydra.WithDebug(),
	)
	hydra.Conf.API("50001")
	app.API("/checkmap", checkmap, api.WithEncoding("gbk"))
	app.Start()
}
func checkmap(ctx hydra.IContext) interface{} {
	ctx.Log().Info(ctx.Request().GetBody())
	rules := map[string]interface{}{
		"uid": "required",
		"ids": "required",
	}
	if err := ctx.Request().CheckMap(rules); err != nil {
		ctx.Log().Error("1.", err)
		return err
	}
	var c = map[string]interface{}{}
	if err := ctx.Request().Bind(&c); err != nil {
		ctx.Log().Error("2.", err)
		return err
	}
	ctx.Log().Infof("%+v-%s", c, ctx.Request().Path().GetEncoding())
	return c
}
