//+build embedzip

package main

import (
	_ "embed"
	"github.com/micro-plat/hydra/conf/server/static"
)

//go:embed static.zip
var zipfs []byte

func init() {
	opts = []static.Option{
		static.WithEmbedBytes("static.zip", zipfs),
		static.WithAutoRewrite(),
	}
}
