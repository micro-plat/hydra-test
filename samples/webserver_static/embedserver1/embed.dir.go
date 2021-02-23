//+build embeddir

package main

import "embed"

//go:embed static
var fs embed.FS

func init() {
	opts = []static.Option{
		static.WithEmbed("static", fs),
		static.WithAutoRewrite(),
	}
}
