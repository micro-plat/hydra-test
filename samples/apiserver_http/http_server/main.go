package main

import (
	"time"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("http_server"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

func init() {
	hydra.Conf.API(":8099")
	app.API("/api", func(ctx hydra.IContext) (r interface{}) {
		ctx.Log().Info("log_session_id:", ctx.Log().GetSessionID())
		ctx.Log().Info("api_user_id:", ctx.User().GetRequestID())

		bytes, err := ctx.Request().GetBody()
		if err != nil {
			return err
		}
		ctx.Log().Debug("GetBody:", string(bytes), err)

		m := ctx.Request().GetMap()
		for k, v := range m {
			ctx.Log().Debugf("Map: %s=%s", k, v)
		}
		ctx.Log().Debug("api_body_map:", m)
		ctx.Log().Debug("api_method:", ctx.Request().Path().GetMethod())
		ctx.Log().Debug("api_encoding:", ctx.Request().Path().GetEncoding())
		ctx.Log().Debug("api_headers:", ctx.Request().Headers())
		return
	})

	app.API("/timeout", func(ctx hydra.IContext) (r interface{}) {
		time.Sleep(time.Second * 20)
		return
	})

}

//启动服务  ./http_server run
func main() {
	app.Start()
}
