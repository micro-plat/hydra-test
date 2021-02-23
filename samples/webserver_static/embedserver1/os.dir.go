//+build osdir

package main

import (
	"github.com/micro-plat/hydra/conf/server/static"

)

func init() {
	opts = []static.Option{ 
		static.WithAssetsPath("static"), 
		static.WithAutoRewrite(),
	}
}
