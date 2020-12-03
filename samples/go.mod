module github.com/micro-plat/hydra-test/samples

go 1.15

replace github.com/micro-plat/lib4go => ../../../../github.com/micro-plat/lib4go

replace github.com/micro-plat/hydra => ../../../../github.com/micro-plat/hydra

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/mattn/go-oci8 v0.1.0
	github.com/micro-plat/hydra v0.0.0-00010101000000-000000000000
	github.com/micro-plat/lib4go v1.0.2
)
