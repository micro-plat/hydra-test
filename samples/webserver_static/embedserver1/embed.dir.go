//+build embeddir

package main

import (
	"embed"
	"github.com/micro-plat/hydra/conf/server/static"
)

//go:embed static
var fs embed.FS

func init() {
	opts = []static.Option{
		static.WithEmbed("static", fs),
		static.WithAutoRewrite(),
	}
}
