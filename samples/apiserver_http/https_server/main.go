package main

import (
	"time"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("https_server"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("lm://./"),
)

func init() {
	hydra.Conf.API(":8098")
	app.API("/api", func(ctx hydra.IContext) (r interface{}) {
		ctx.Log().Info("log_session_id:", ctx.Log().GetSessionID())
		ctx.Log().Info("api_user_id:", ctx.User().GetRequestID())
		m, err := ctx.Request().GetMap()
		if err != nil {
			return err
		}
		ctx.Log().Info("api_body_map:", m)
		ctx.Log().Info("api_method:", ctx.Request().Path().GetMethod())
		ctx.Log().Info("api_encoding:", ctx.Request().Path().GetEncoding())
		ctx.Log().Info("api_headers:", ctx.Request().Headers())
		return
	})

	app.API("/timeout", func(ctx hydra.IContext) (r interface{}) {
		time.Sleep(time.Second * 20)
		return
	})

}

//启动服务  ./https_server run
func main() {
	app.Start()
}
