module github.com/micro-plat/hydra-test/samples

go 1.15

replace github.com/micro-plat/lib4go => ../../../../github.com/micro-plat/lib4go

replace github.com/micro-plat/hydra => ../../../../github.com/micro-plat/hydra

require (
	github.com/micro-plat/hydra v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20200707034311-ab3426394381
)
