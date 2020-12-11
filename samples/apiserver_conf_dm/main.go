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
		//打印命令行输入的参数
		fmt.Println("ServerTypes:", global.Def.ServerTypes)
		fmt.Println("PlatName:", global.Def.PlatName)
		fmt.Println("SysName:", global.Def.SysName)
		fmt.Println("ClusterName:", global.Def.ClusterName)
		fmt.Println("RegistryAddr:", global.Def.RegistryAddr)

		//以OnReady绑定参数
		global.Def.ServerTypes = []string{"ws"}
		global.Def.PlatName = "onReady"
		global.Def.SysName = "onReady"
		global.Def.ClusterName = "onReady"
		global.Def.RegistryAddr = "lm://./"
	})
}

// go build
// 设置参数系统 ./apiserver_conf_dm run -r "lm://./" -p platname -S api -c t
// 查看onReady日志打印,系统参数绑定正确
// 查看服务启动日志,以onReady方式的参数绑定正确
func main() {
	app.Start()
}
