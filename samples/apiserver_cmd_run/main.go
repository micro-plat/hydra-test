package main

import (
	"os"

	"github.com/micro-plat/hydra"
	_ "github.com/micro-plat/hydra/components/caches/cache/redis"
)

var app = hydra.NewApp(
	hydra.WithRunFlag("flag", "-测试添加的run命令扩展参数"), // 添加run命令扩展参数
)

func main() {
	os.Args = []string{"apiserver_cmd_run", "run", "-h"}
	app.Start()
}
