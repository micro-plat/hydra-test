package context

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra-test/units/mocks"
	"github.com/micro-plat/hydra/conf"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/context/ctx"
	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/hydra/mock"
	"github.com/micro-plat/lib4go/assert"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/logger"
	"github.com/micro-plat/lib4go/types"
)

func Test_response_Write_ERR(t *testing.T) {
	tests := []struct {
		name    string
		status  int
		content interface{}
		wantRs  int
		wantRc  string
		wantCt  string
		header  http.Header
	}{
		{name: "1.1.内容为err,状态码未设置", status: 0, content: errs.NewError(999, "错误"), wantRs: 999, wantRc: "错误", wantCt: "text/plain; charset=utf-8"},
		{name: "1.2.内容为err,状态码300", status: 300, content: errors.New("err"), wantRs: 400, wantRc: "err", wantCt: "text/plain; charset=utf-8"},
		{name: "1.3.内容为err,状态码在200", status: 200, content: errors.New("err"), wantRs: 400, wantRc: "err", wantCt: "text/plain; charset=utf-8"},
		{name: "1.4.内容为err,状态码400", status: 400, content: errors.New("err"), wantRs: 400, wantRc: "err", wantCt: "text/plain; charset=utf-8"},
		{name: "1.5.内容为err,状态码900", status: 900, content: errors.New("err"), wantRs: 900, wantRc: "err", wantCt: "text/plain; charset=utf-8"},
	}

	confObj := mocks.NewConfBy("responsetest", "Writecluster") //构建对象
	confObj.API("8080")                                        //初始化参数
	hydra.G.SysName = "apiserver"
	serverConf := confObj.GetAPIConf() //获取配置

	global.IsDebug = true
	for _, tt := range tests {
		//构建response对象
		contx := &mocks.TestContxt{HttpHeader: tt.header}
		c := ctx.NewResponse(contx, serverConf, logger.New("hydra"), conf.NewMeta())
		err := c.Write(tt.status, tt.content)
		assert.Equal(t, nil, err, tt.name)

		//测试reponse状态码和内容
		rs, rc, cp := c.GetFinalResponse()
		assert.Equal(t, tt.wantRs, rs, tt.name)
		assert.Equal(t, tt.wantRc, rc, tt.name)
		assert.Equal(t, tt.wantCt, cp, tt.name)

	}
}

func Test_response_Write_Nil(t *testing.T) {
	tests := []struct {
		name    string
		status  int
		content interface{}
		wantRs  int
		wantRc  string
		wantCt  string
		header  string
	}{

		{name: "1.1.内容为空,未设置状态码,content-type未设置", status: 0, content: nil, wantRs: 200, wantRc: "", wantCt: "text/plain; charset=utf-8"},
		{name: "1.4.内容为空,未设置状态码,content-type为plain", status: 0, content: nil, wantRs: 200, header: context.UTF8PLAIN, wantRc: "", wantCt: "text/plain; charset=utf-8"},
		{name: "1.5.内容为空,未设置状态码,content-type为json", status: 0, content: nil, wantRs: 200, header: context.UTF8JSON, wantRc: "", wantCt: context.UTF8JSON},
		{name: "1.6.内容为空,未设置状态码,content-type为xml", status: 0, content: nil, wantRs: 200, header: context.UTF8XML, wantRc: "", wantCt: context.UTF8XML},
		{name: "1.6.内容为空,未设置状态码,content-type为yaml", status: 0, content: nil, wantRs: 200, header: context.UTF8YAML, wantRc: "", wantCt: context.UTF8YAML},

		{name: "2.1.内容为空,状态码为成功,content-type未设置", status: 200, content: nil, wantRs: 200, wantRc: "", wantCt: "text/plain; charset=utf-8"},
		{name: "2.2.内容为空,状态码为成功,content-type为plain", status: 200, content: nil, wantRs: 200, header: context.UTF8PLAIN, wantRc: "", wantCt: "text/plain; charset=utf-8"},
		{name: "2.3.内容为空,状态码为成功,content-type为json", status: 200, content: nil, wantRs: 200, header: context.UTF8JSON, wantRc: "", wantCt: context.UTF8JSON},
		{name: "2.4.内容为空,状态码为成功,content-type为xml", status: 200, content: nil, wantRs: 200, header: context.UTF8XML, wantRc: "", wantCt: context.UTF8XML},
		{name: "2.5.内容为空,状态码为成功,content-type为yaml", status: 200, content: nil, wantRs: 200, header: context.UTF8YAML, wantRc: "", wantCt: context.UTF8YAML},

		{name: "3.1.内容为空,状态码为600,content-type未设置", status: 600, content: nil, wantRs: 600, wantRc: "", wantCt: "text/plain; charset=utf-8"},
		{name: "3.2.内容为空,状态码为600,content-type为plain", status: 600, content: nil, wantRs: 600, header: context.UTF8PLAIN, wantRc: "", wantCt: "text/plain; charset=utf-8"},
		{name: "3.3.内容为空,状态码为600,content-type为json", status: 600, content: nil, wantRs: 600, header: context.UTF8JSON, wantRc: "", wantCt: context.UTF8JSON},
		{name: "3.4.内容为空,状态码为600,content-type为xml", status: 600, content: nil, wantRs: 600, header: context.UTF8XML, wantRc: "", wantCt: context.UTF8XML},
		{name: "3.5.内容为空,状态码为600,content-type为yaml", status: 600, content: nil, wantRs: 600, header: context.UTF8YAML, wantRc: "", wantCt: context.UTF8YAML},

		{name: "4.1.内容为空,状态码为400,content-type未设置", status: 400, content: nil, wantRs: 400, wantRc: "", wantCt: "text/plain; charset=utf-8"},
		{name: "4.2.内容为空,状态码为400,content-type为plain", status: 400, content: nil, wantRs: 400, header: context.UTF8PLAIN, wantRc: "", wantCt: "text/plain; charset=utf-8"},
		{name: "4.3.内容为空,状态码为400,content-type为json", status: 400, content: nil, wantRs: 400, header: context.UTF8JSON, wantRc: "", wantCt: context.UTF8JSON},
		{name: "4.4.内容为空,状态码为400,content-type为xml", status: 400, content: nil, wantRs: 400, header: context.UTF8XML, wantRc: "", wantCt: context.UTF8XML},
		{name: "4.5.内容为空,状态码为400,content-type为yaml", status: 400, content: nil, wantRs: 400, header: context.UTF8YAML, wantRc: "", wantCt: context.UTF8YAML},
	}

	for _, tt := range tests {

		contx := mock.NewContext("", mock.WithRHeaders(types.XMap{
			"Content-Type": tt.header,
		}))

		c := contx.Response()
		//写入响应流
		c.Data(tt.status, tt.header, tt.content)

		//测试reponse状态码和内容
		rs, rc, cp := c.GetFinalResponse()
		assert.Equal(t, tt.wantRs, rs, tt.name)
		assert.Equal(t, tt.wantRc, rc, tt.name)
		assert.Equal(t, tt.wantCt, cp, tt.name)

	}
}

func Test_response_Write_String(t *testing.T) {
	tests := []struct {
		name    string
		status  int
		content interface{}
		wantRs  int
		wantRc  string
		wantCt  string
		header  string
	}{

		{name: "1.1.内容字符串,未设置状态码,content-type未设置", status: 0, content: "success", wantRs: 200, wantRc: "success", wantCt: "text/plain; charset=utf-8"},
		{name: "1.2.内容字符串,未设置状态码,content-type为plain", status: 0, content: "success", wantRs: 200, header: context.UTF8PLAIN, wantRc: "success", wantCt: "text/plain; charset=utf-8"},
		{name: "1.3.内容字符串,未设置状态码,content-type为json", status: 0, content: "success", wantRs: 200, header: context.UTF8JSON, wantRc: `success`, wantCt: "application/json; charset=utf-8"},
		{name: "1.4.内容字符串,未设置状态码,content-type为xml", status: 0, content: "success", wantRs: 200, header: context.UTF8XML, wantRc: "success", wantCt: "application/xml; charset=utf-8"},
		{name: "1.5.内容字符串,未设置状态码,content-type为yaml", status: 0, content: "success", wantRs: 200, header: context.UTF8YAML, wantRc: "success", wantCt: "text/yaml; charset=utf-8"},

		{name: "2.1.内容字符串,状态码为成功,content-type未设置", status: 200, content: "success", wantRs: 200, wantRc: "success", wantCt: "text/plain; charset=utf-8"},
		{name: "2.2.内容字符串,状态码为成功,content-type为plain", status: 200, content: "success", wantRs: 200, header: context.UTF8PLAIN, wantRc: "success", wantCt: "text/plain; charset=utf-8"},
		{name: "2.3.内容字符串,状态码为成功,content-type为json", status: 200, content: "success", wantRs: 200, header: context.UTF8JSON, wantRc: "success", wantCt: "application/json; charset=utf-8"},
		{name: "2.4.内容字符串,状态码为成功,content-type为xml", status: 200, content: "success", wantRs: 200, header: context.UTF8XML, wantRc: "success", wantCt: "application/xml; charset=utf-8"},
		{name: "2.5.内容字符串,状态码为成功,content-type为yaml", status: 200, content: "success", wantRs: 200, header: context.UTF8YAML, wantRc: "success", wantCt: "text/yaml; charset=utf-8"},

		{name: "3.1.内容字符串,状态码为600,content-type未设置", status: 600, content: "success", wantRs: 600, wantRc: "success", wantCt: "text/plain; charset=utf-8"},
		{name: "3.2.内容字符串,状态码为600,content-type为plain", status: 600, content: "success", wantRs: 600, header: context.UTF8PLAIN, wantRc: "success", wantCt: "text/plain; charset=utf-8"},
		{name: "3.3.内容字符串,状态码为600,content-type为json", status: 600, content: "success", wantRs: 600, header: context.UTF8JSON, wantRc: "success", wantCt: "application/json; charset=utf-8"},
		{name: "3.4.内容字符串,状态码为600,content-type为xml", status: 600, content: "success", wantRs: 600, header: context.UTF8XML, wantRc: "success", wantCt: "application/xml; charset=utf-8"},
		{name: "3.5.内容字符串,状态码为600,content-type为yaml", status: 600, content: "success", wantRs: 600, header: context.UTF8YAML, wantRc: "success", wantCt: "text/yaml; charset=utf-8"},

		{name: "4.1.内容字符串,状态码为400,content-type未设置", status: 400, content: "success", wantRs: 400, wantRc: "success", wantCt: "text/plain; charset=utf-8"},
		{name: "4.2.内容字符串,状态码为400,content-type为plain", status: 400, content: "success", wantRs: 400, header: context.UTF8PLAIN, wantRc: "success", wantCt: "text/plain; charset=utf-8"},
		{name: "4.3.内容字符串,状态码为400,content-type为json", status: 400, content: "success", wantRs: 400, header: context.UTF8JSON, wantRc: "success", wantCt: "application/json; charset=utf-8"},
		{name: "4.4.内容字符串,状态码为400,content-type为xml", status: 400, content: "success", wantRs: 400, header: context.UTF8XML, wantRc: "success", wantCt: "application/xml; charset=utf-8"},
		{name: "4.5.内容字符串,状态码为400,content-type为yaml", status: 400, content: "success", wantRs: 400, header: context.UTF8YAML, wantRc: "success", wantCt: "text/yaml; charset=utf-8"},
	}

	confObj := mocks.NewConfBy("context_response_test1", "response1") //构建对象
	confObj.API("8080")                                               //初始化参数
	global.IsDebug = true

	for _, tt := range tests {

		contx := mock.NewContext("", mock.WithRHeaders(types.XMap{
			"Content-Type": tt.header,
		}))

		c := contx.Response()
		//写入响应流
		c.Data(tt.status, tt.header, tt.content)
		c.Flush()

		//测试reponse状态码和内容
		rs, rc, cp := c.GetFinalResponse()
		assert.Equal(t, tt.wantRs, rs, tt.name)
		assert.Equal(t, tt.wantRc, rc, tt.name)
		assert.Equal(t, tt.wantCt, cp, tt.name)

	}
}

func Test_response_Write_JSON(t *testing.T) {
	tests := []struct {
		name    string
		status  int
		content interface{}
		wantRs  int
		wantRc  string
		wantCt  string
		header  string
	}{
		{name: "1.1.内容JSON字符串,未设置状态码,content-type未设置", status: 0, content: `{"key":"value"}`, wantRs: 200, wantRc: `{"key":"value"}`, wantCt: "application/json; charset=utf-8"},
		{name: "1.2.内容JSON字符串,未设置状态码,content-type为plain", status: 0, content: `{"key":"value"}`, wantRs: 200, header: context.UTF8PLAIN, wantRc: `{"key":"value"}`, wantCt: "text/plain; charset=utf-8"},
		{name: "1.3.内容JSON字符串,未设置状态码,content-type为json", status: 0, content: `{"key":"value"}`, wantRs: 200, header: context.UTF8JSON, wantRc: `{"key":"value"}`, wantCt: "application/json; charset=utf-8"},
		{name: "1.4.内容JSON字符串,未设置状态码,content-type为xml", status: 0, content: `{"key":"value"}`, wantRs: 200, header: context.UTF8XML, wantRc: `{"key":"value"}`, wantCt: "application/xml; charset=utf-8"},
		{name: "1.5.内容JSON字符串,未设置状态码,content-type为yaml", status: 0, content: `{"key":"value"}`, wantRs: 200, header: context.UTF8YAML, wantRc: `{"key":"value"}`, wantCt: "text/yaml; charset=utf-8"},

		{name: "2.1.内容JSON字符串,状态码为成功,content-type未设置", status: 200, content: `{"key":"value"}`, wantRs: 200, wantRc: `{"key":"value"}`, wantCt: "application/json; charset=utf-8"},
		{name: "2.2.内容JSON字符串,状态码为成功,content-type为plain", status: 200, content: `{"key":"value"}`, wantRs: 200, header: context.UTF8PLAIN, wantRc: `{"key":"value"}`, wantCt: "text/plain; charset=utf-8"},
		{name: "2.3.内容JSON字符串,状态码为成功,content-type为json", status: 200, content: `{"key":"value"}`, wantRs: 200, header: context.UTF8JSON, wantRc: `{"key":"value"}`, wantCt: "application/json; charset=utf-8"},
		{name: "2.4.内容JSON字符串,状态码为成功,content-type为xml", status: 200, content: `{"key":"value"}`, wantRs: 200, header: context.UTF8XML, wantRc: `{"key":"value"}`, wantCt: "application/xml; charset=utf-8"},
		{name: "2.5.内容JSON字符串,状态码为成功,content-type为yaml", status: 200, content: `{"key":"value"}`, wantRs: 200, header: context.UTF8YAML, wantRc: `{"key":"value"}`, wantCt: "text/yaml; charset=utf-8"},

		{name: "3.1.内容JSON字符串,状态码为600,content-type未设置", status: 600, content: `{"key":"value"}`, wantRs: 600, wantRc: `{"key":"value"}`, wantCt: "application/json; charset=utf-8"},
		{name: "3.2.内容JSON字符串,状态码为600,content-type为plain", status: 600, content: `{"key":"value"}`, wantRs: 600, header: context.UTF8PLAIN, wantRc: `{"key":"value"}`, wantCt: "text/plain; charset=utf-8"},
		{name: "3.3.内容JSON字符串,状态码为600,content-type为json", status: 600, content: `{"key":"value"}`, wantRs: 600, header: context.UTF8JSON, wantRc: `{"key":"value"}`, wantCt: "application/json; charset=utf-8"},
		{name: "3.4.内容JSON字符串,状态码为600,content-type为xml", status: 600, content: `{"key":"value"}`, wantRs: 600, header: context.UTF8XML, wantRc: `{"key":"value"}`, wantCt: "application/xml; charset=utf-8"},
		{name: "3.5.内容JSON字符串,状态码为600,content-type为yaml", status: 600, content: `{"key":"value"}`, wantRs: 600, header: context.UTF8YAML, wantRc: `{"key":"value"}`, wantCt: "text/yaml; charset=utf-8"},

		{name: "4.1.内容JSON字符串,状态码为400,content-type未设置", status: 400, content: `{"key":"value"}`, wantRs: 400, wantRc: `{"key":"value"}`, wantCt: "application/json; charset=utf-8"},
		{name: "4.2.内容JSON字符串,状态码为400,content-type为plain", status: 400, content: `{"key":"value"}`, wantRs: 400, header: context.UTF8PLAIN, wantRc: `{"key":"value"}`, wantCt: "text/plain; charset=utf-8"},
		{name: "4.3.内容JSON字符串,状态码为400,content-type为json", status: 400, content: `{"key":"value"}`, wantRs: 400, header: context.UTF8JSON, wantRc: `{"key":"value"}`, wantCt: "application/json; charset=utf-8"},
		{name: "4.4.内容JSON字符串,状态码为400,content-type为xml", status: 400, content: `{"key":"value"}`, wantRs: 400, header: context.UTF8XML, wantRc: `{"key":"value"}`, wantCt: "application/xml; charset=utf-8"},
		{name: "4.5.内容JSON字符串,状态码为400,content-type为yaml", status: 400, content: `{"key":"value"}`, wantRs: 400, header: context.UTF8YAML, wantRc: `{"key":"value"}`, wantCt: "text/yaml; charset=utf-8"},
	}

	confObj := mocks.NewConfBy("context_response_test1", "response1") //构建对象
	confObj.API("8080")                                               //获取配置

	global.IsDebug = true
	for _, tt := range tests {

		contx := mock.NewContext("", mock.WithRHeaders(types.XMap{
			"Content-Type": tt.header,
		}))

		c := contx.Response()
		//写入响应流
		c.Data(tt.status, tt.header, tt.content)
		c.Flush()

		//测试reponse状态码和内容
		rs, rc, cp := c.GetFinalResponse()
		assert.Equal(t, tt.wantRs, rs, tt.name)
		assert.Equal(t, tt.wantRc, rc, tt.name)
		assert.Equal(t, tt.wantCt, cp, tt.name)

	}
}

func Test_response_Write_XML(t *testing.T) {
	tests := []struct {
		name    string
		status  int
		content interface{}
		wantRs  int
		wantRc  string
		wantCt  string
		header  string
	}{
		{name: "1.1.内容XML字符串,未设置状态码,content-type未设置", status: 0, content: "<xml><key>value</key></xml>", wantRs: 200, wantRc: "<xml><key>value</key></xml>", wantCt: "application/xml; charset=utf-8"},
		{name: "1.2.内容XML字符串,未设置状态码,content-type为plain", status: 0, content: "<xml><key>value</key></xml>", wantRs: 200, header: context.UTF8PLAIN, wantRc: "<xml><key>value</key></xml>", wantCt: "text/plain; charset=utf-8"},
		{name: "1.3.内容XML字符串,未设置状态码,content-type为json", status: 0, content: "<xml><key>value</key></xml>", wantRs: 200, header: context.UTF8JSON, wantRc: "<xml><key>value</key></xml>", wantCt: "application/json; charset=utf-8"},
		{name: "1.4.内容XML字符串,未设置状态码,content-type为xml", status: 0, content: "<xml><key>value</key></xml>", wantRs: 200, header: context.UTF8XML, wantRc: "<xml><key>value</key></xml>", wantCt: "application/xml; charset=utf-8"},
		{name: "1.5.内容XML字符串,未设置状态码,content-type为yaml", status: 0, content: "<xml><key>value</key></xml>", wantRs: 200, header: context.UTF8YAML, wantRc: "<xml><key>value</key></xml>", wantCt: "text/yaml; charset=utf-8"},

		{name: "2.1.内容XML字符串,状态码为成功,content-type未设置", status: 200, content: "<xml><key>value</key></xml>", wantRs: 200, wantRc: "<xml><key>value</key></xml>", wantCt: "application/xml; charset=utf-8"},
		{name: "2.2.内容XML字符串,状态码为成功,content-type为plain", status: 200, content: "<xml><key>value</key></xml>", wantRs: 200, header: context.UTF8PLAIN, wantRc: "<xml><key>value</key></xml>", wantCt: "text/plain; charset=utf-8"},
		{name: "2.3.内容XML字符串,状态码为成功,content-type为json", status: 200, content: "<xml><key>value</key></xml>", wantRs: 200, header: context.UTF8JSON, wantRc: "<xml><key>value</key></xml>", wantCt: "application/json; charset=utf-8"},
		{name: "2.4.内容XML字符串,状态码为成功,content-type为xml", status: 200, content: "<xml><key>value</key></xml>", wantRs: 200, header: context.UTF8XML, wantRc: "<xml><key>value</key></xml>", wantCt: "application/xml; charset=utf-8"},
		{name: "2.5.内容XML字符串,状态码为成功,content-type为yaml", status: 200, content: "<xml><key>value</key></xml>", wantRs: 200, header: context.UTF8YAML, wantRc: "<xml><key>value</key></xml>", wantCt: "text/yaml; charset=utf-8"},

		{name: "3.1.内容XML字符串,状态码为600,content-type未设置", status: 600, content: "<xml><key>value</key></xml>", wantRs: 600, wantRc: "<xml><key>value</key></xml>", wantCt: "application/xml; charset=utf-8"},
		{name: "3.2.内容XML字符串,状态码为600,content-type为plain", status: 600, content: "<xml><key>value</key></xml>", wantRs: 600, header: context.UTF8PLAIN, wantRc: "<xml><key>value</key></xml>", wantCt: "text/plain; charset=utf-8"},
		{name: "3.3.内容XML字符串,状态码为600,content-type为json", status: 600, content: "<xml><key>value</key></xml>", wantRs: 600, header: context.UTF8JSON, wantRc: "<xml><key>value</key></xml>", wantCt: "application/json; charset=utf-8"},
		{name: "3.4.内容XML字符串,状态码为600,content-type为xml", status: 600, content: "<xml><key>value</key></xml>", wantRs: 600, header: context.UTF8XML, wantRc: "<xml><key>value</key></xml>", wantCt: "application/xml; charset=utf-8"},
		{name: "3.5.内容XML字符串,状态码为600,content-type为yaml", status: 600, content: "<xml><key>value</key></xml>", wantRs: 600, header: context.UTF8YAML, wantRc: "<xml><key>value</key></xml>", wantCt: "text/yaml; charset=utf-8"},

		{name: "4.1.内容XML字符串,状态码为400,content-type未设置", status: 400, content: "<xml><key>value</key></xml>", wantRs: 400, wantRc: "<xml><key>value</key></xml>", wantCt: "application/xml; charset=utf-8"},
		{name: "4.2.内容XML字符串,状态码为400,content-type为plain", status: 400, content: "<xml><key>value</key></xml>", wantRs: 400, header: context.UTF8PLAIN, wantRc: "<xml><key>value</key></xml>", wantCt: "text/plain; charset=utf-8"},
		{name: "4.3.内容XML字符串,状态码为400,content-type为json", status: 400, content: "<xml><key>value</key></xml>", wantRs: 400, header: context.UTF8JSON, wantRc: "<xml><key>value</key></xml>", wantCt: "application/json; charset=utf-8"},
		{name: "4.4.内容XML字符串,状态码为400,content-type为xml", status: 400, content: "<xml><key>value</key></xml>", wantRs: 400, header: context.UTF8XML, wantRc: "<xml><key>value</key></xml>", wantCt: "application/xml; charset=utf-8"},
		{name: "4.5.内容XML字符串,状态码为400,content-type为yaml", status: 400, content: "<xml><key>value</key></xml>", wantRs: 400, header: context.UTF8YAML, wantRc: "<xml><key>value</key></xml>", wantCt: "text/yaml; charset=utf-8"},
	}

	confObj := mocks.NewConfBy("context_response_test1", "response1") //构建对象
	confObj.API("8080")                                               //初始化参数
	global.IsDebug = true

	for _, tt := range tests {
		contx := mock.NewContext("", mock.WithRHeaders(types.XMap{
			"Content-Type": tt.header,
		}))

		c := contx.Response()
		//写入响应流
		c.Data(tt.status, tt.header, tt.content)
		c.Flush()

		//测试reponse状态码和内容
		rs, rc, cp := c.GetFinalResponse()
		assert.Equal(t, tt.wantRs, rs, tt.name)
		assert.Equal(t, tt.wantRc, rc, tt.name)
		assert.Equal(t, tt.wantCt, cp, tt.name)

	}
}

func Test_response_Write_HTML(t *testing.T) {
	tests := []struct {
		name    string
		status  int
		content interface{}
		wantRs  int
		wantRc  string
		wantCt  string
		header  string
	}{
		{name: "1.1.内容XML字符串,未设置状态码,content-type未设置", status: 0, content: "<!DOCTYPE html><html></html>", wantRs: 200, wantRc: "<!DOCTYPE html><html></html>", wantCt: "text/html; charset=utf-8"},
		{name: "1.2.内容XML字符串,未设置状态码,content-type为plain", status: 0, content: "<!DOCTYPE html><html></html>", wantRs: 200, header: context.UTF8PLAIN, wantRc: "<!DOCTYPE html><html></html>", wantCt: "text/plain; charset=utf-8"},
		{name: "1.3.内容XML字符串,未设置状态码,content-type为json", status: 0, content: "<!DOCTYPE html><html></html>", wantRs: 200, header: context.UTF8JSON, wantRc: "<!DOCTYPE html><html></html>", wantCt: "application/json; charset=utf-8"},
		{name: "1.4.内容XML字符串,未设置状态码,content-type为xml", status: 0, content: "<!DOCTYPE html><html></html>", wantRs: 200, header: context.UTF8XML, wantRc: "<!DOCTYPE html><html></html>", wantCt: "application/xml; charset=utf-8"},
		{name: "1.5.内容XML字符串,未设置状态码,content-type为yaml", status: 0, content: "<!DOCTYPE html><html></html>", wantRs: 200, header: context.UTF8YAML, wantRc: "<!DOCTYPE html><html></html>", wantCt: "text/yaml; charset=utf-8"},

		{name: "2.1.内容XML字符串,状态码为成功,content-type未设置", status: 200, content: "<!DOCTYPE html><html></html>", wantRs: 200, wantRc: "<!DOCTYPE html><html></html>", wantCt: "text/html; charset=utf-8"},
		{name: "2.2.内容XML字符串,状态码为成功,content-type为plain", status: 200, content: "<!DOCTYPE html><html></html>", wantRs: 200, header: context.UTF8PLAIN, wantRc: "<!DOCTYPE html><html></html>", wantCt: "text/plain; charset=utf-8"},
		{name: "2.3.内容XML字符串,状态码为成功,content-type为json", status: 200, content: "<!DOCTYPE html><html></html>", wantRs: 200, header: context.UTF8JSON, wantRc: "<!DOCTYPE html><html></html>", wantCt: "application/json; charset=utf-8"},
		{name: "2.4.内容XML字符串,状态码为成功,content-type为xml", status: 200, content: "<!DOCTYPE html><html></html>", wantRs: 200, header: context.UTF8XML, wantRc: "<!DOCTYPE html><html></html>", wantCt: "application/xml; charset=utf-8"},
		{name: "2.5.内容XML字符串,状态码为成功,content-type为yaml", status: 200, content: "<!DOCTYPE html><html></html>", wantRs: 200, header: context.UTF8YAML, wantRc: "<!DOCTYPE html><html></html>", wantCt: "text/yaml; charset=utf-8"},

		{name: "3.1.内容XML字符串,状态码为600,content-type未设置", status: 600, content: "<!DOCTYPE html><html></html>", wantRs: 600, wantRc: "<!DOCTYPE html><html></html>", wantCt: "text/html; charset=utf-8"},
		{name: "3.2.内容XML字符串,状态码为600,content-type为plain", status: 600, content: "<!DOCTYPE html><html></html>", wantRs: 600, header: context.UTF8PLAIN, wantRc: "<!DOCTYPE html><html></html>", wantCt: "text/plain; charset=utf-8"},
		{name: "3.3.内容XML字符串,状态码为600,content-type为json", status: 600, content: "<!DOCTYPE html><html></html>", wantRs: 600, header: context.UTF8JSON, wantRc: "<!DOCTYPE html><html></html>", wantCt: "application/json; charset=utf-8"},
		{name: "3.4.内容XML字符串,状态码为600,content-type为xml", status: 600, content: "<!DOCTYPE html><html></html>", wantRs: 600, header: context.UTF8XML, wantRc: "<!DOCTYPE html><html></html>", wantCt: "application/xml; charset=utf-8"},
		{name: "3.5.内容XML字符串,状态码为600,content-type为yaml", status: 600, content: "<!DOCTYPE html><html></html>", wantRs: 600, header: context.UTF8YAML, wantRc: "<!DOCTYPE html><html></html>", wantCt: "text/yaml; charset=utf-8"},

		{name: "4.1.内容XML字符串,状态码为400,content-type未设置", status: 400, content: "<!DOCTYPE html><html></html>", wantRs: 400, wantRc: "<!DOCTYPE html><html></html>", wantCt: "text/html; charset=utf-8"},
		{name: "4.2.内容XML字符串,状态码为400,content-type为plain", status: 400, content: "<!DOCTYPE html><html></html>", wantRs: 400, header: context.UTF8PLAIN, wantRc: "<!DOCTYPE html><html></html>", wantCt: "text/plain; charset=utf-8"},
		{name: "4.3.内容XML字符串,状态码为400,content-type为json", status: 400, content: "<!DOCTYPE html><html></html>", wantRs: 400, header: context.UTF8JSON, wantRc: "<!DOCTYPE html><html></html>", wantCt: "application/json; charset=utf-8"},
		{name: "4.4.内容XML字符串,状态码为400,content-type为xml", status: 400, content: "<!DOCTYPE html><html></html>", wantRs: 400, header: context.UTF8XML, wantRc: "<!DOCTYPE html><html></html>", wantCt: "application/xml; charset=utf-8"},
		{name: "4.5.内容XML字符串,状态码为400,content-type为yaml", status: 400, content: "<!DOCTYPE html><html></html>", wantRs: 400, header: context.UTF8YAML, wantRc: "<!DOCTYPE html><html></html>", wantCt: "text/yaml; charset=utf-8"},
	}

	confObj := mocks.NewConfBy("context_response_test1", "response1") //构建对象
	confObj.API("8080")                                               //初始化参数
	global.IsDebug = true
	for _, tt := range tests {
		contx := mock.NewContext("", mock.WithRHeaders(types.XMap{
			"Content-Type": tt.header,
		}))

		c := contx.Response()
		//写入响应流
		c.Data(tt.status, tt.header, tt.content)
		c.Flush()

		//测试reponse状态码和内容
		rs, rc, cp := c.GetFinalResponse()
		assert.Equal(t, tt.wantRs, rs, tt.name)
		assert.Equal(t, tt.wantRc, rc, tt.name)
		assert.Equal(t, tt.wantCt, cp, tt.name)

	}
}

func Test_response_Write_MAP(t *testing.T) {
	tests := []struct {
		name    string
		status  int
		content interface{}
		wantRs  int
		wantRc  string
		wantCt  string
		header  string
	}{

		{name: "1.1.内容Map,未设置状态码,content-type未设置", status: 0, content: map[string]interface{}{"id": 100}, wantRs: 200, wantRc: `{"id":100}`, wantCt: "application/json; charset=utf-8"},
		{name: "1.2.内容Map,未设置状态码,content-type为plain", status: 0, content: map[string]interface{}{"id": 100}, wantRs: 200, header: context.UTF8PLAIN, wantRc: "map[id:100]", wantCt: "text/plain; charset=utf-8"},
		{name: "1.3.内容Map,未设置状态码,content-type为json", status: 0, content: map[string]interface{}{"id": 100}, wantRs: 200, header: context.UTF8JSON, wantRc: `{"id":100}`, wantCt: "application/json; charset=utf-8"},
		{name: "1.4.内容Map,未设置状态码,content-type为xml", status: 0, content: map[string]interface{}{"id": 100}, wantRs: 200, header: context.UTF8XML, wantRc: "<xml><id>100</id></xml>", wantCt: "application/xml; charset=utf-8"},
		{name: "1.5.内容Map,未设置状态码,content-type为yaml", status: 0, content: map[string]interface{}{"id": 100}, wantRs: 200, header: context.UTF8YAML, wantRc: "id: 100\n", wantCt: "text/yaml; charset=utf-8"},

		{name: "2.1.内容Map,状态码为成功,content-type未设置", status: 200, content: map[string]interface{}{"id": 100}, wantRs: 200, wantRc: `{"id":100}`, wantCt: "application/json; charset=utf-8"},
		{name: "2.2.内容Map,状态码为成功,content-type为plain", status: 200, content: map[string]interface{}{"id": 100}, wantRs: 200, header: context.UTF8PLAIN, wantRc: "map[id:100]", wantCt: "text/plain; charset=utf-8"},
		{name: "2.3.内容Map,状态码为成功,content-type为json", status: 200, content: map[string]interface{}{"id": 100}, wantRs: 200, header: context.UTF8JSON, wantRc: `{"id":100}`, wantCt: "application/json; charset=utf-8"},
		{name: "2.4.内容Map,状态码为成功,content-type为xml", status: 200, content: map[string]interface{}{"id": 100}, wantRs: 200, header: context.UTF8XML, wantRc: "<xml><id>100</id></xml>", wantCt: "application/xml; charset=utf-8"},
		{name: "2.5.内容Map,状态码为成功,content-type为yaml", status: 200, content: map[string]interface{}{"id": 100}, wantRs: 200, header: context.UTF8YAML, wantRc: "id: 100\n", wantCt: "text/yaml; charset=utf-8"},

		{name: "3.1.内容Map,状态码为600,content-type未设置", status: 600, content: map[string]interface{}{"id": 100}, wantRs: 600, wantRc: `{"id":100}`, wantCt: "application/json; charset=utf-8"},
		{name: "3.2.内容Map,状态码为600,content-type为plain", status: 600, content: map[string]interface{}{"id": 100}, wantRs: 600, header: context.UTF8PLAIN, wantRc: "map[id:100]", wantCt: "text/plain; charset=utf-8"},
		{name: "3.3.内容Map,状态码为600,content-type为json", status: 600, content: map[string]interface{}{"id": 100}, wantRs: 600, header: context.UTF8JSON, wantRc: `{"id":100}`, wantCt: "application/json; charset=utf-8"},
		{name: "3.4.内容Map,状态码为600,content-type为xml", status: 600, content: map[string]interface{}{"id": 100}, wantRs: 600, header: context.UTF8XML, wantRc: "<xml><id>100</id></xml>", wantCt: "application/xml; charset=utf-8"},
		{name: "3.5.内容Map,状态码为600,content-type为yaml", status: 600, content: map[string]interface{}{"id": 100}, wantRs: 600, header: context.UTF8YAML, wantRc: "id: 100\n", wantCt: "text/yaml; charset=utf-8"},

		{name: "4.1.内容Map,状态码为400,content-type未设置", status: 400, content: map[string]interface{}{"id": 100}, wantRs: 400, wantRc: `{"id":100}`, wantCt: "application/json; charset=utf-8"},
		{name: "4.2.内容Map,状态码为400,content-type为plain", status: 400, content: map[string]interface{}{"id": 100}, wantRs: 400, header: context.UTF8PLAIN, wantRc: "map[id:100]", wantCt: "text/plain; charset=utf-8"},
		{name: "4.3.内容Map,状态码为400,content-type为json", status: 400, content: map[string]interface{}{"id": 100}, wantRs: 400, header: context.UTF8JSON, wantRc: `{"id":100}`, wantCt: "application/json; charset=utf-8"},
		{name: "4.4.内容Map,状态码为400,content-type为xml", status: 400, content: map[string]interface{}{"id": 100}, wantRs: 400, header: context.UTF8XML, wantRc: "<xml><id>100</id></xml>", wantCt: "application/xml; charset=utf-8"},
		{name: "4.5.内容Map,状态码为400,content-type为yaml", status: 400, content: map[string]interface{}{"id": 100}, wantRs: 400, header: context.UTF8YAML, wantRc: "id: 100\n", wantCt: "text/yaml; charset=utf-8"},
	}

	confObj := mocks.NewConf() //构建对象
	confObj.API("8080")        //初始化参数
	global.IsDebug = true
	for _, tt := range tests {
		contx := mock.NewContext("", mock.WithRHeaders(types.XMap{
			"Content-Type": tt.header,
		}))

		c := contx.Response()
		//写入响应流
		c.Data(tt.status, tt.header, tt.content)
		c.Flush()

		//测试reponse状态码和内容
		rs, rc, cp := c.GetFinalResponse()
		assert.Equal(t, tt.wantRs, rs, tt.name)
		assert.Equal(t, tt.wantRc, rc, tt.name)
		assert.Equal(t, tt.wantCt, cp, tt.name)

	}
}

func Test_response_Write_Struct(t *testing.T) {
	type content struct {
		Key string `json:"key" xml:"key"`
	}

	tests := []struct {
		name    string
		status  int
		content interface{}
		wantRs  int
		wantRc  string
		wantCt  string
		header  string
	}{
		{name: "1.1.内容Struct,未设置状态码,content-type未设置", status: 0, content: &content{Key: "value"}, wantRs: 200, wantRc: `{"key":"value"}`, wantCt: "application/json; charset=utf-8"},
		{name: "1.2.内容Struct,未设置状态码,content-type为plain", status: 0, content: &content{Key: "value"}, wantRs: 200, header: context.UTF8PLAIN, wantRc: "&{value}", wantCt: "text/plain; charset=utf-8"},
		{name: "1.3.内容Struct,未设置状态码,content-type为json", status: 0, content: &content{Key: "value"}, wantRs: 200, header: context.UTF8JSON, wantRc: `{"key":"value"}`, wantCt: "application/json; charset=utf-8"},
		{name: "1.4.内容Struct,未设置状态码,content-type为xml", status: 0, content: &content{Key: "value"}, wantRs: 200, header: context.UTF8XML, wantRc: "<xml><key>value</key></xml>", wantCt: "application/xml; charset=utf-8"},
		{name: "1.5.内容Struct,未设置状态码,content-type为yaml", status: 0, content: &content{Key: "value"}, wantRs: 200, header: context.UTF8YAML, wantRc: "key: value\n", wantCt: "text/yaml; charset=utf-8"},

		{name: "2.1.内容Struct,状态码为成功,content-type未设置", status: 200, content: &content{Key: "value"}, wantRs: 200, wantRc: `{"key":"value"}`, wantCt: "application/json; charset=utf-8"},
		{name: "2.2.内容Struct,状态码为成功,content-type为plain", status: 200, content: &content{Key: "value"}, wantRs: 200, header: context.UTF8PLAIN, wantRc: "&{value}", wantCt: "text/plain; charset=utf-8"},
		{name: "2.3.内容Struct,状态码为成功,content-type为json", status: 200, content: &content{Key: "value"}, wantRs: 200, header: context.UTF8JSON, wantRc: `{"key":"value"}`, wantCt: "application/json; charset=utf-8"},
		{name: "2.4.内容Struct,状态码为成功,content-type为xml", status: 200, content: &content{Key: "value"}, wantRs: 200, header: context.UTF8XML, wantRc: "<xml><key>value</key></xml>", wantCt: "application/xml; charset=utf-8"},
		{name: "2.5.内容Struct,状态码为成功,content-type为yaml", status: 200, content: &content{Key: "value"}, wantRs: 200, header: context.UTF8YAML, wantRc: "key: value\n", wantCt: "text/yaml; charset=utf-8"},

		{name: "3.1.内容Struct,状态码为600,content-type未设置", status: 600, content: &content{Key: "value"}, wantRs: 600, wantRc: `{"key":"value"}`, wantCt: "application/json; charset=utf-8"},
		{name: "3.2.内容Struct,状态码为600,content-type为plain", status: 600, content: &content{Key: "value"}, wantRs: 600, header: context.UTF8PLAIN, wantRc: "&{value}", wantCt: "text/plain; charset=utf-8"},
		{name: "3.3.内容Struct,状态码为600,content-type为json", status: 600, content: &content{Key: "value"}, wantRs: 600, header: context.UTF8JSON, wantRc: `{"key":"value"}`, wantCt: "application/json; charset=utf-8"},
		{name: "3.4.内容Struct,状态码为600,content-type为xml", status: 600, content: &content{Key: "value"}, wantRs: 600, header: context.UTF8XML, wantRc: "<xml><key>value</key></xml>", wantCt: "application/xml; charset=utf-8"},
		{name: "3.5.内容Struct,状态码为600,content-type为yaml", status: 600, content: &content{Key: "value"}, wantRs: 600, header: context.UTF8YAML, wantRc: "key: value\n", wantCt: "text/yaml; charset=utf-8"},

		{name: "4.1.内容Struct,状态码为400,content-type未设置", status: 400, content: &content{Key: "value"}, wantRs: 400, wantRc: `{"key":"value"}`, wantCt: "application/json; charset=utf-8"},
		{name: "4.2.内容Struct,状态码为400,content-type为plain", status: 400, content: &content{Key: "value"}, wantRs: 400, header: context.UTF8PLAIN, wantRc: "&{value}", wantCt: "text/plain; charset=utf-8"},
		{name: "4.3.内容Struct,状态码为400,content-type为json", status: 400, content: &content{Key: "value"}, wantRs: 400, header: context.UTF8JSON, wantRc: `{"key":"value"}`, wantCt: "application/json; charset=utf-8"},
		{name: "4.4.内容Struct,状态码为400,content-type为xml", status: 400, content: &content{Key: "value"}, wantRs: 400, header: context.UTF8XML, wantRc: "<xml><key>value</key></xml>", wantCt: "application/xml; charset=utf-8"},
		{name: "4.5.内容Struct,状态码为400,content-type为yaml", status: 400, content: &content{Key: "value"}, wantRs: 400, header: context.UTF8YAML, wantRc: "key: value\n", wantCt: "text/yaml; charset=utf-8"},
	}

	confObj := mocks.NewConf() //构建对象
	confObj.API("8080")        //初始化参数
	global.IsDebug = true
	for _, tt := range tests {
		contx := mock.NewContext("", mock.WithRHeaders(types.XMap{
			"Content-Type": tt.header,
		}))

		c := contx.Response()
		//写入响应流
		c.Data(tt.status, tt.header, tt.content)
		c.Flush()

		//测试reponse状态码和内容
		rs, rc, cp := c.GetFinalResponse()
		assert.Equal(t, tt.wantRs, rs, tt.name)
		assert.Equal(t, tt.wantRc, rc, tt.name)
		assert.Equal(t, tt.wantCt, cp, tt.name)

	}
}

func Test_response_Write_Int(t *testing.T) {

	tests := []struct {
		name    string
		status  int
		content interface{}
		wantRs  int
		wantRc  string
		wantCt  string
		header  string
	}{
		{name: "1.1.内容Int,未设置状态码,content-type未设置", status: 0, content: 1, wantRs: 200, wantRc: `1`, wantCt: "text/plain; charset=utf-8"},
		{name: "1.2.内容Int,未设置状态码,content-type为plain", status: 0, content: 1, wantRs: 200, header: context.UTF8PLAIN, wantRc: "1", wantCt: "text/plain; charset=utf-8"},
		{name: "1.3.内容Int,未设置状态码,content-type为json", status: 0, content: 1, wantRs: 200, header: context.UTF8JSON, wantRc: `1`, wantCt: "application/json; charset=utf-8"},
		{name: "1.4.内容Int,未设置状态码,content-type为xml", status: 0, content: 1, wantRs: 200, header: context.UTF8XML, wantRc: "1", wantCt: "application/xml; charset=utf-8"},
		{name: "1.5.内容Int,未设置状态码,content-type为yaml", status: 0, content: 1, wantRs: 200, header: context.UTF8YAML, wantRc: "1", wantCt: "text/yaml; charset=utf-8"},

		{name: "2.1.内容Int,状态码为成功,content-type未设置", status: 200, content: 1, wantRs: 200, wantRc: `1`, wantCt: "text/plain; charset=utf-8"},
		{name: "2.2.内容Int,状态码为成功,content-type为plain", status: 200, content: 1, wantRs: 200, header: context.UTF8PLAIN, wantRc: "1", wantCt: "text/plain; charset=utf-8"},
		{name: "2.3.内容Int,状态码为成功,content-type为json", status: 200, content: 1, wantRs: 200, header: context.UTF8JSON, wantRc: `1`, wantCt: "application/json; charset=utf-8"},
		{name: "2.4.内容Int,状态码为成功,content-type为xml", status: 200, content: 1, wantRs: 200, header: context.UTF8XML, wantRc: "1", wantCt: "application/xml; charset=utf-8"},
		{name: "2.5.内容Int,状态码为成功,content-type为yaml", status: 200, content: 1, wantRs: 200, header: context.UTF8YAML, wantRc: "1", wantCt: "text/yaml; charset=utf-8"},

		{name: "3.1.内容Int,状态码为600,content-type未设置", status: 600, content: 1, wantRs: 600, wantRc: `1`, wantCt: "text/plain; charset=utf-8"},
		{name: "3.2.内容Int,状态码为600,content-type为plain", status: 600, content: 1, wantRs: 600, header: context.UTF8PLAIN, wantRc: "1", wantCt: "text/plain; charset=utf-8"},
		{name: "3.3.内容Int,状态码为600,content-type为json", status: 600, content: 1, wantRs: 600, header: context.UTF8JSON, wantRc: `1`, wantCt: "application/json; charset=utf-8"},
		{name: "3.4.内容Int,状态码为600,content-type为xml", status: 600, content: 1, wantRs: 600, header: context.UTF8XML, wantRc: "1", wantCt: "application/xml; charset=utf-8"},
		{name: "3.5.内容Int,状态码为600,content-type为yaml", status: 600, content: 1, wantRs: 600, header: context.UTF8YAML, wantRc: "1", wantCt: "text/yaml; charset=utf-8"},

		{name: "4.1.内容Int,状态码为400,content-type未设置", status: 400, content: 1, wantRs: 400, wantRc: `1`, wantCt: "text/plain; charset=utf-8"},
		{name: "4.2.内容Int,状态码为400,content-type为plain", status: 400, content: 1, wantRs: 400, header: context.UTF8PLAIN, wantRc: "1", wantCt: "text/plain; charset=utf-8"},
		{name: "4.3.内容Int,状态码为400,content-type为json", status: 400, content: 1, wantRs: 400, header: context.UTF8JSON, wantRc: `1`, wantCt: "application/json; charset=utf-8"},
		{name: "4.4.内容Int,状态码为400,content-type为xml", status: 400, content: 1, wantRs: 400, header: context.UTF8XML, wantRc: "1", wantCt: "application/xml; charset=utf-8"},
		{name: "4.5.内容Int,状态码为400,content-type为yaml", status: 400, content: 1, wantRs: 400, header: context.UTF8YAML, wantRc: "1", wantCt: "text/yaml; charset=utf-8"},
	}

	confObj := mocks.NewConf() //构建对象
	confObj.API("8080")        //初始化参数
	global.IsDebug = true
	for _, tt := range tests {
		contx := mock.NewContext("", mock.WithRHeaders(types.XMap{
			"Content-Type": tt.header,
		}))

		c := contx.Response()
		//写入响应流
		c.Data(tt.status, tt.header, tt.content)
		c.Flush()

		//测试reponse状态码和内容
		rs, rc, cp := c.GetFinalResponse()
		assert.Equal(t, tt.wantRs, rs, tt.name)
		assert.Equal(t, tt.wantRc, rc, tt.name)
		assert.Equal(t, tt.wantCt, cp, tt.name)

	}
}

func Test_response_Header(t *testing.T) {

	rc := mock.NewContext("", mock.WithRHeaders(types.XMap{}))
	c := rc.Response()

	c.Header("header1", "value1")
	assert.Equal(t, types.XMap{"header1": "value1"}, c.GetHeaders(), "设置header")

	c.Header("header1", "value1-1")
	assert.Equal(t, types.XMap{"header1": "value1-1"}, c.GetHeaders(), "更新已设置的header")

	c.Header("header2", "value2")
	assert.Equal(t, types.XMap{"header1": "value1-1", "header2": "value2"}, c.GetHeaders(), "设置不存在的header")
}

func Test_response_Abort(t *testing.T) {

	rc := mock.NewContext("", mock.WithRHeaders(types.XMap{}))
	c := rc.Response()
	global.IsDebug = false
	//测试Abort
	c.Abort(200, fmt.Errorf("终止"))
	rs, content, cp := c.GetFinalResponse()

	assert.Equal(t, 400, rs, "验证上下文中的状态码")
	assert.Equal(t, "Bad Request", content, "验证返回内容")
	assert.Equal(t, context.UTF8PLAIN, cp, "验证上下文中的content-type")

}

func Test_response_File(t *testing.T) {

	rc := mock.NewContext("", mock.WithRHeaders(types.XMap{}))
	c := rc.Response()

	//测试File
	c.File("upload.test.txt")
	c.Flush()
	rs, content, cp := c.GetFinalResponse()

	assert.Equal(t, 200, rs, "验证上下文中的状态码")
	assert.Equal(t, "", content, "验证返回内容")
	assert.Equal(t, context.UTF8PLAIN, cp, "验证上下文中的content-type")
}

func Test_response_Special(t *testing.T) {
	rc := mock.NewContext("", mock.WithRHeaders(types.XMap{}))
	c := rc.Response()

	//添加响应的特殊字符
	c.AddSpecial("proxy")
	c.AddSpecial("logging")
	c.AddSpecial("render")

	//获取
	assert.Equal(t, "proxy|logging|render", c.GetSpecials(), "获取响应的特殊字符")
}

func Test_response_GetRaw(t *testing.T) {

	rc := mock.NewContext("获取content", mock.WithRHeaders(types.XMap{}))
	c := rc.Response()

	//获取
	assert.Equal(t, nil, c.GetRaw(), "获取content")

	//写入
	c.Write(200, "content")

	s, content, ctp := c.GetRawResponse()
	assert.Equal(t, 200, s, "获取status")
	assert.Equal(t, "content", content, "获取content")
	assert.Equal(t, context.UTF8PLAIN, ctp, "获取content-type")
}
