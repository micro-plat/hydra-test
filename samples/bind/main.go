package main

import (
	"github.com/micro-plat/hydra"

	"github.com/micro-plat/hydra/hydra/servers/http"
)

//服务器各种返回结果
func main() {
	app := hydra.NewApp(
		hydra.WithPlatName("hydra"),
		hydra.WithServerTypes(http.API),
	)
	app.API("/request", request)
	app.Start()
}
func request(ctx hydra.IContext) interface{} {
	var c = CouponNotifyInfo{}
	if err := ctx.Request().Bind(&c); err != nil {
		return err
	}
	return c
}

//CouponNotifyInfo 卡券通知参数
type CouponNotifyInfo struct {
	Type           string  `json:"tp" form:"tp" m2s:"type" valid:"required"`                     //通知类型 CONSUMED 券码核销通知 RECHARGE 渠道充值通知
	Time           string  `json:"time" form:"time" m2s:"time"`                                  //核销时间
	Code           string  `json:"code" form:"code" m2s:"code"`                                  //type=CONSUMED 券码核销通知时时券码编号
	Sitecode       string  `json:"sitecode" form:"sitecode" m2s:"sitecode"`                      //type=CONSUMED 核销站点
	RechargeAmount float64 `json:"recharge_amount" form:"recharge_amount" m2s:"recharge_amount"` //type= RECHARGE 渠道充值金额
	Sign           string  `json:"sign" form:"sign" m2s:"sign" valid:"required"`
}
