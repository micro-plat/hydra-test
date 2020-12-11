package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/hydra/conf/vars/queue/queueredis"
	"github.com/micro-plat/hydra/hydra/servers/http"
	"github.com/micro-plat/hydra/hydra/servers/mqc"
)

var app = hydra.NewApp(
	hydra.WithDebug(),
	hydra.WithServerTypes(http.API, mqc.MQC),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("mqcserver"),
	hydra.WithClusterName("taosytest"),
	hydra.WithRegistry("lm://."),
)

func init() {
	hydra.Conf.API(":8070")
	hydra.Conf.MQC("redis://redis")
	hydra.Conf.Vars().Redis("redis", "192.168.5.79:6379")
	hydra.Conf.Vars().Queue().Redis("redis", "", queueredis.WithConfigName("redis"))
	app.API("/hydratest/mqcserver/apijson", funcAPI)
	app.API("/hydratest/mqcserver/apijson1", funcAPI1)
	app.MQC("/hydratest/mqcserver/jsonqueue", funcMQC1, "mqcserver::json:queue")
}

// mqcserver-json json内容在mqcserver中解析demo
//1.1 安装程序 ./mqcserver_json conf install -cover
//1.2 使用 ./mqcserver_json run

//1.3 调用返回结果接口：http://localhost:8070/hydratest/mqcserver/apijson 观察日志是否与api中初始化数据相同
//1.4 调用返回结果接口：http://localhost:8070/hydratest/mqcserver/apijson1 数据类型不同数据解析
func main() {
	app.Start()
}

type DateTime time.Time

func (d DateTime) MarshalJSON() (bytes []byte, err error) {
	fmt.Println("marshal")
	val := time.Time(d)
	tmpV := fmt.Sprintf("\"%s\"", val.Format("2006-01-02 15:04:05"))
	fmt.Println("marshal:", string(bytes))
	return []byte(tmpV), nil
}
func (d DateTime) UnmarshalJSON(bytes []byte) error {
	fmt.Println("unmarshal:", string(bytes))
	val, err := time.Parse("\"2006-01-02 15:04:05\"", string(bytes))
	d = DateTime(val)
	return err
}

type Param struct {
	Param1 string                 `json:"param1"`
	Param2 bool                   `json:"param2"`
	Param3 int32                  `json:"param3"`
	Param4 float32                `json:"param4"`
	Param5 []string               `json:"param5"`
	Param6 DateTime               `json:"param6" time_format:"2006-01-02 15:04:05"`
	Param7 map[string]interface{} `json:"param7"`
}

var funcAPI = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("mqcserver-json-api json内容在mqcserver中解析demo")
	queue := "mqcserver::json:queue"
	input := Param{
		Param1: "34ddf#$*@大!@#$%^&*()_+~锅饭都是",
		Param2: true,
		Param3: 1024,
		Param4: 10.24,
		Param5: []string{"1", "2"},
		Param6: DateTime(time.Now()),
		Param7: map[string]interface{}{"t1": 123, "t2": "sdfs@@###", "t3": 12.2},
	}
	ver, err := json.Marshal(input)
	value := string(ver)
	fmt.Println("value:", value, err)
	queueObj := components.Def.Queue().GetRegularQueue("redis")
	if err := queueObj.Send(queue, value); err != nil {
		ctx.Log().Errorf("发送消息队列异常：%s", queue)
		return
	}
	return
}

var funcAPI1 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("mqcserver-json-api1 json内容在mqcserver中解析demo")
	queue := "mqcserver::json:queue"
	value := fmt.Sprint(`{"param1":"@#$%^&*()_+~锅饭都是","param2":"true","param3":"1024","param4":"10.24","param5":"1,2","param6":"2020/11/12 11:12:59","param7":"{}"}`)
	ctx.Log().Info("-------------:", value)
	queueObj := components.Def.Queue().GetRegularQueue("redis")
	if err := queueObj.Send(queue, value); err != nil {
		ctx.Log().Errorf("发送消息队列异常：%s", queue)
		return
	}
	return
}

var funcMQC1 = func(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("mqcserver-json-mqc json内容在mqcserver中解析demo")
	if err := ctx.Request().Check(); err != nil {
		ctx.Log().Errorf("ctx.Request().Check()异常：%s", err)
		return
	}
	param := &Param{}
	if err := ctx.Request().Bind(param); err != nil {
		ctx.Log().Errorf("ctx.Request().Bind()异常：%s", err)
	}
	ctx.Log().Info("----bind data:", param, param.Param6)

	body, err := ctx.Request().GetBody()
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetBody()异常：%s", err)
	}
	ctx.Log().Info("----body data:", string(body))

	bodyBytes, raw, err := ctx.Request().GetFullRaw()
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetFullRaw()异常：%s", err)
	}
	ctx.Log().Info("----GetFullRaw data:", string(bodyBytes), raw)

	jsonD, err := ctx.Request().GetJSON("param7")
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetJSON()异常：%s", err)
	}
	ctx.Log().Info("----GetJSON data:", string(jsonD))

	xmap, err := ctx.Request().GetMap()
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetMap()异常：%s", err)
	}
	ctx.Log().Info("----GetMap()：", xmap)

	tm, err := ctx.Request().GetDatetime("param6", "2006-01-02 15:04:05")
	if err != nil {
		ctx.Log().Errorf("ctx.Request().GetDatetime()异常：%s", err)
	}
	ctx.Log().Info("----GetDatetime()：", tm)
	ctx.Log().Info("----Keys data:", ctx.Request().Keys())
	ctx.Log().Info("----Len data:", ctx.Request().Len())
	ctx.Log().Info("----GetArray data:", ctx.Request().GetArray("param5"))
	ctx.Log().Info("----GetInt32 data:", ctx.Request().GetInt32("param3"))
	ctx.Log().Info("----GetFloat32 data:", ctx.Request().GetFloat32("param4"))
	ctx.Log().Info("----GetBool data:", ctx.Request().GetBool("param2"))
	return
}
