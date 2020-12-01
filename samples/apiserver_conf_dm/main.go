package main

import (
	"fmt"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydra_test"),
	hydra.WithSystemName("conf_dm"),
	hydra.WithClusterName("t"),
	hydra.WithRegistry("zk://192.168.0.101"),
)

func init() {
	global.OnReady(func() {
		fmt.Println("ServerTypes:", global.Def.ServerTypes)
		fmt.Println("PlatName:", global.Def.PlatName)
		fmt.Println("SysName:", global.Def.SysName)
		fmt.Println("ClusterName:", global.Def.ClusterName)
	})
}

//go build
//启动服务
//  ./apiserver_conf_dm run -p 
func main() {
	app.Start()
}
