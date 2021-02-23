//+build embedzip

package main

import "embed"

//go:embed static.zip
var zipfs []byte

func init() {
	opts = []static.Option{ 
		static.WithEmbedBytes("static.zip",zipfs), 
		static.WithAutoRewrite(),
	}
}
