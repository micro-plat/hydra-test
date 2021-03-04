module github.com/micro-plat/hydra-test/units

go 1.16

replace github.com/micro-plat/lib4go => ../../../../github.com/micro-plat/lib4go

replace github.com/micro-plat/hydra => ../../../../github.com/micro-plat/hydra

require (
	github.com/asaskevich/govalidator v0.0.0-20200907205600-7a23bdc65eef // indirect
	github.com/clbanning/mxj v1.8.4
	github.com/gin-gonic/gin v1.6.3
	github.com/micro-plat/hydra v0.0.0-00010101000000-000000000000
	github.com/micro-plat/lib4go v1.0.10
	github.com/shopspring/decimal v1.2.0
	github.com/urfave/cli v1.22.5
	golang.org/x/text v0.3.4
)
