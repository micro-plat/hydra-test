// +build prod

package main

import "github.com/micro-plat/hydra"

func init() {
	hydra.Conf.API("8001").NFS("/home/yanglei/work/bin/static02")
}
