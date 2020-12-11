package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/acl/blacklist"
	"github.com/micro-plat/hydra/conf/server/acl/limiter"
	"github.com/micro-plat/hydra/conf/server/acl/whitelist"
	"github.com/micro-plat/hydra/conf/server/auth/basic"
	"github.com/micro-plat/hydra/conf/server/auth/jwt"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.WS),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("ws_establish"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("zk://192.168.0.101"),
)

func init() {
	h := hydra.Conf.WS(":8080")
	h.Jwt(jwt.WithEnable()).Basic(basic.WithEnable())                                                              //配置不生效的中间件
	h.BlackList(blacklist.WithEnable(), blacklist.WithIP("192.168.5.106"))                                         //配置黑名单中间件
	h.WhiteList(whitelist.WithEnable(), whitelist.WithIPList(whitelist.NewIPList([]string{"/"}, "192.168.5.115"))) //配置白名单的中间件
	h.Limit(limiter.WithEnable(), limiter.WithRuleList(limiter.NewRule("/", 0, limiter.WithMaxWait(1), limiter.WithFallback(), limiter.WithReponse(200, "success"))))
}

//测试wsserver[连接建立]的中间件[BlackList/WhiteList]

//启动服务 /middleware_establish run
//查看 /conf几点下配置  [/jwt/basic/blacklist/whitelist中间件配置均为启用]

//使用 机器192.168.5.106建立连接ws://localhost:8080/ [jwt/basic中间件未生效] [192.168.5.106不允许访问,黑名单限制生效]
//使用 机器192.168.5.115建立连接ws://localhost:8080/ [jwt/basic中间件未生效] [连接建立成功,白名单限制生效]
func main() {
	app.Start()
}
