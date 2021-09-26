// +build !prod

package main

import "github.com/micro-plat/hydra"

func init() {
	hydra.Conf.API("8000").NFS("/home/yanglei/work/bin/static01")
}
