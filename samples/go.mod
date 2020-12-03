module github.com/micro-plat/hydra-test/samples

go 1.15

replace github.com/micro-plat/lib4go => ../../../../github.com/micro-plat/lib4go

replace github.com/micro-plat/hydra => ../../../../github.com/micro-plat/hydra

require (
	github.com/fsnotify/fsnotify v1.4.9
	github.com/gin-gonic/gin v1.6.2
	github.com/go-sql-driver/mysql v1.5.0
	github.com/mattn/go-oci8 v0.1.0
	github.com/fsnotify/fsnotify v1.4.9
	github.com/gin-gonic/gin v1.6.2
	github.com/micro-plat/hydra v0.0.0-00010101000000-000000000000
	github.com/micro-plat/lib4go v1.0.2
	github.com/urfave/cli v1.22.4 // indirect
	gopkg.in/yaml.v2 v2.3.0
)
