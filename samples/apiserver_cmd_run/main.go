package main

import (
	"os"

	"github.com/micro-plat/hydra"
)

var app = hydra.NewApp(
	hydra.WithRunFlag("flag", "-测试添加的run命令扩展参数"), // 添加run命令扩展参数
)

//go run 查看help打印
func main() {
	os.Args = []string{"apiserver_cmd_run", "run", "-h"}
	app.Start()
}
