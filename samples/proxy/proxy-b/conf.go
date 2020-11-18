package main

import "github.com/micro-plat/hydra"

func init() {
	hydra.OnReady(func() {
		hydra.Conf.API(":8091").Proxy(`	
        request := import("request")
        app := import("app")
        text := import("text")
        types :=import("types")
        fmt := import("fmt")

        getUpCluster := func(){
            ip := request.getClientIP()
            current:= app.getCurrentClusterName()
            if text.has_prefix(ip,"192.168."){
                return app.getClusterNameBy("a")
            }
            return current
        }
        upcluster := getUpCluster()
		`)
	})
}
