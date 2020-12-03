module github.com/micro-plat/hydra-test/samples

go 1.15

replace github.com/micro-plat/lib4go => ../../../../github.com/micro-plat/lib4go

replace github.com/micro-plat/hydra => ../../../../github.com/micro-plat/hydra

require (
	github.com/fsnotify/fsnotify v1.4.9
	github.com/gin-gonic/gin v1.6.2
	github.com/micro-plat/hydra v0.0.0-00010101000000-000000000000
	github.com/micro-plat/lib4go v1.0.2
	github.com/urfave/cli v1.22.4 // indirect
)
