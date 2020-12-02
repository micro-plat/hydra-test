package main

import (
	"fmt"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/global"
	_ "github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp()

func init() {
	global.OnReady(func() {
		fmt.Println("ServerTypes:", global.Def.ServerTypes)
		fmt.Println("PlatName:", global.Def.PlatName)
		fmt.Println("SysName:", global.Def.SysName)
		fmt.Println("ClusterName:", global.Def.ClusterName)
		fmt.Println("RegistryAddr:", global.Def.RegistryAddr)
	})
}

// go build
// 设置参数系统 ./apiserver_conf_dm run -r "lm://./" -p platname -S api -c t
// 查看onReady日志打印,判断系统参数是否绑定完成
func main() {
	app.Start()
}
