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
	hydra.WithSystemName("cmd_conf"),
	hydra.WithClusterName("t"),
	hydra.WithConfFlag("flag", "-测试添加的conf命令扩展参数"), // 添加run命令扩展参数
)

// 通过代码为run命令模式下指定参数，并在程序中获得cli输入的值
// go build
// ./apiserver_cmd_conf conf show -flag flagvalue 查看打印的flag对应的值
func main() {
	app.Cli.Conf.OnStarting(func(cli global.ICli) error {
		fmt.Println("IsSet:", cli.IsSet("flag"))
		fmt.Println("String:", cli.String("flag"))
		return nil
	})
	app.Start()
}
