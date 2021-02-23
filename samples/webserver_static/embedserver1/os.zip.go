//+build oszip

package main

import (
	"github.com/micro-plat/hydra/conf/server/static"
)

func init() {
	opts = []static.Option{
		static.WithAssetsPath("static.zip"),
		static.WithAutoRewrite(),
	}
}
