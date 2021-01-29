package context

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra-test/units/mocks"
	"github.com/micro-plat/hydra/conf"
	"github.com/micro-plat/hydra/context/ctx"
	"github.com/micro-plat/hydra/hydra/servers/pkg/middleware"
	"github.com/micro-plat/hydra/mock"
	"github.com/micro-plat/lib4go/assert"
	"github.com/micro-plat/lib4go/types"
)

type xmlResult struct {
	Xml result `m2s:"xml"`
}
type result struct {
	Key    string `json:"key" valid:"required" m2s:"key"`
	Number int    `json:"number" valid:"required" m2s:"number"`
	// Number int32         `json:"number" valid:"required" m2s:"number"`
	// Number int64         `json:"number" valid:"required" m2s:"number"`
	// Number float32       `json:"number" valid:"required" m2s:"number"`
	// Number float64       `json:"number" valid:"required" m2s:"number"`
	//Number types.Decimal `json:"number" valid:"required" m2s:"number"`
}

func Test_request_Bind(t *testing.T) {
	var res string
	num := 4585455141568261664
	tests := []struct {
		name        string
		contentType string
		encoding    string
		ctype       string
		queryRaw    string
		out         interface{}
		isMap       bool
		wantErrStr  string
		want        interface{}
	}{
		{name: "1.1 参数非指针", out: map[string]string{}, wantErrStr: "对象转换有误 输出对象必须为struct:map"},
		{name: "1.2 参数类型非struct,map", out: &res, wantErrStr: "对象转换有误 输出对象必须为struct:string"},

		//	{name: "2.1 内容为xml,绑定MAP", contentType: "application/xml",encoding: "utf-8", ctype: "xml",  isMap: true,   out: &map[string]interface{}{}, want: &map[string]interface{}{"xml": map[string]interface{}{"key": value}}},
		{name: "2.2 内容为xml,绑定Struct", contentType: "application/xml", encoding: "utf-8", ctype: "xml", out: &result{}, want: &result{Key: value, Number: num}},
		//		{name: "3.1 内容为json,绑定MAP", contentType: "application/json", encoding: "utf-8", ctype: "xml",   out: &map[string]interface{}{}, want: &map[string]interface{}{"key": value}},
		{name: "3.2 内容为json,绑定Struct", contentType: "application/json", encoding: "utf-8", ctype: "json", out: &result{}, want: &result{Key: value, Number: num}},
		//	{name: "4.1 内容为yaml,绑定MAP", contentType: "application/x-yaml",encoding: "utf-8", ctype: "xml",   out: &map[string]interface{}{}, want: &map[string]interface{}{"key": value}},
		{name: "4.2 内容为yaml,绑定Struct", contentType: "application/x-yaml", encoding: "utf-8", ctype: "yaml", out: &result{}, want: &result{Key: value, Number: num}},
		//	{name: "5.1 内容为form,绑定MAP", contentType: "application/x-www-form-urlencoded",encoding: "utf-8", ctype: "xml",   out: &map[string]interface{}{}, want: &map[string]interface{}{"key": value}},
		{name: "5.2 内容为form,绑定Struct", contentType: "application/x-www-form-urlencoded", encoding: "utf-8", ctype: "form", out: &result{}, want: &result{Key: value, Number: num}},
	}

	confObj := mocks.NewConf() //构建对象
	confObj.API("8080")        //初始化参数
	//serverConf := confObj.GetAPIConf() //获取配置
	//c, _ := gin.CreateTestContext(httptest.NewRecorder())
	for _, tt := range tests {

		//	err = req.Bind(tt.out)
		body := getTestBindBody(value, num, tt.encoding, tt.ctype)
		fmt.Println("body:", tt.ctype, string(body))
		mc := mock.NewContext(string(body), mock.WithRHeaders(types.XMap{
			"Content-Type": fmt.Sprintf("%s;chaset=%s", tt.contentType, tt.encoding),
		}))
		// b, q, _ := mc.Request().GetFullRaw()
		// fmt.Println("b:", string(b), "q:", q)

		err := mc.Request().Bind(tt.out)
		if tt.wantErrStr != "" {
			assert.Equal(t, tt.wantErrStr, err.Error(), tt.name)
			continue
		}

		assert.Equal(t, tt.want, tt.out, tt.name)
	}
}

func Test_request_Check(t *testing.T) {
	tests := []struct {
		name        string
		queryRaw    string
		contentType string
		fields      []string
		body        string
		wantErr     string
	}{

		{name: "1.1 内容为xml,参数值为空", contentType: "application/xml", fields: []string{"key"}, body: "<key></key>", wantErr: "输入参数:key值不能为空"},
		{name: "1.2 内容为xml,参数不存在", contentType: "application/xml", fields: []string{"key"}, body: "<xml></xml>", wantErr: "输入参数:key值不能为空"},
		{name: "1.3 内容为xml,参数值不为空", contentType: "application/xml", fields: []string{"key"}, body: "<key>value</key>"},

		{name: "2.1 内容为json,参数值为空", contentType: "application/json", fields: []string{"key"}, body: `{"key":""}`, wantErr: "输入参数:key值不能为空"},
		{name: "2.2 内容为json,参数不存在", contentType: "application/json", fields: []string{"key"}, body: `{"key2":""}`, wantErr: "输入参数:key值不能为空"},
		{name: "2.3 内容为json,参数值不为空", contentType: "application/json", fields: []string{"key"}, body: `{"key":"value"}`},

		{name: "3.1 内容为yaml,参数值为空", contentType: "application/x-yaml", fields: []string{"key"}, body: `key: ""`, wantErr: "输入参数:key值不能为空"},
		{name: "3.2 内容为yaml,参数不存在", contentType: "application/x-yaml", fields: []string{"key"}, body: `kye2:`, wantErr: "输入参数:key值不能为空"},
		{name: "3.3 内容为yaml,参数值不为空", contentType: "application/x-yaml", fields: []string{"key"}, body: `key: value`},

		{name: "4.1 内容为form,参数值为空", contentType: "application/x-www-form-urlencoded", fields: []string{"key"}, body: `key=`, wantErr: "输入参数:key值不能为空"},
		{name: "4.2 内容为form,参数不存在", contentType: "application/x-www-form-urlencoded", fields: []string{"key"}, body: `key2=`, wantErr: "输入参数:key值不能为空"},
		{name: "4.3 内容为form,参数值不为空", contentType: "application/x-www-form-urlencoded", fields: []string{"key"}, body: `key=value`},
	}

	confObj := mocks.NewConfBy("contexttest", "cluster") //构建对象
	confObj.API("8080")
	hydra.G.SysName = "apiserver" //初始化参数

	serverConf := confObj.GetAPIConf() //获取配置
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	for _, tt := range tests {

		//构建请求
		r, err := http.NewRequest("POST", "http://localhost:8080/url?"+tt.queryRaw, bytes.NewReader([]byte(tt.body)))
		assert.Equal(t, nil, err, "构建请求")

		//设置content-type
		r.Header.Set("Content-Type", fmt.Sprintf("%s; charset=UTF-8", tt.contentType))

		c.Request = r
		req := ctx.NewRequest(middleware.NewGinCtx(c), serverConf, conf.NewMeta())

		err = req.Check(tt.fields...)
		if tt.wantErr != "" {
			assert.Equal(t, tt.wantErr, err.Error(), tt.name)
			continue
		}
		assert.Equal(t, nil, err, tt.name)
	}
}

func Test_request_GetKeys(t *testing.T) {
	tests := []struct {
		name        string
		queryRaw    string
		contentType string
		want        []string
		body        string
	}{

		{name: "1.1 内容为xml，无参数", contentType: "application/xml", want: []string{}, body: ""},
		{name: "1.2 内容为xml，有多个参数", contentType: "application/xml", want: []string{"xml"}, body: "<xml></xml>"},

		{name: "2.1 内容为json，无参数", contentType: "application/json", want: []string{}, body: ""},
		{name: "2.2 内容为json，有多个参数", contentType: "application/json", want: []string{"key1", "key2"}, body: `{"key1":"","key2":""}`},

		{name: "3.1 内容为yaml,无参数", contentType: "application/x-yaml", want: []string{}, body: ``},
		{name: "3.2 内容为yaml,有多个参数", contentType: "application/x-yaml", want: []string{"key1", "key2"}, body: "key1: value1\nkey2: value2"},

		{name: "4.1 内容为form,无参数", contentType: "application/x-www-form-urlencoded", want: []string{}, body: ``},
		{name: "4.2 内容为form,有多个参数", contentType: "application/x-www-form-urlencoded", want: []string{"key1", "key2"}, body: `key1=value1&key2=value2`},
	}

	confObj := mocks.NewConf() //构建对象
	confObj.API("8080")        //初始化参数
	hydra.G.SysName = "apiserver"

	serverConf := confObj.GetAPIConf() //获取配置
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	for _, tt := range tests {

		//构建请求
		r, err := http.NewRequest("POST", "http://localhost:8080/url?"+tt.queryRaw, bytes.NewReader([]byte(tt.body)))
		assert.Equal(t, nil, err, "构建请求")

		//设置content-type
		r.Header.Set("Content-Type", fmt.Sprintf("%s; charset=UTF-8", tt.contentType))

		c.Request = r
		req := ctx.NewRequest(middleware.NewGinCtx(c), serverConf, conf.NewMeta())

		got := req.Keys()
		assert.Equal(t, len(tt.want), len(got), tt.name)
		keyMap := map[string]bool{}
		for _, v := range tt.want {
			keyMap[v] = true
		}
		for _, v := range got {
			if _, ok := keyMap[v]; !ok {
				t.Errorf("%s:GeKeys错误", tt.name)
			}
		}
	}
}

func Test_request_GetCookies(t *testing.T) {
	tests := []struct {
		name        string
		contentType string
		cookie      http.Cookie
		want        types.XMap
	}{
		//net/http: invalid byte 'Ö' in Cookie.Value; dropping invalid bytes
		{name: "1. cookie内容为中文GBK", contentType: "application/json;charset=gbk", cookie: http.Cookie{Name: "cname", Value: Utf8ToGbk("中文")}, want: types.XMap{"cname": ""}},
		//net/http: invalid byte 'ä' in Cookie.Value; dropping invalid bytes
		{name: "2. cookie内容为中文UTF-8", contentType: "application/json;charset=utf-8", cookie: http.Cookie{Name: "cname", Value: "中文"}, want: types.XMap{"cname": ""}},
		{name: "3. cookie内容不存在中文", contentType: "application/json;charset=utf-8", cookie: http.Cookie{Name: "cname", Value: "value!@#$%^&*()_+="}, want: types.XMap{"cname": "value!@#$%^\u0026*()_+="}},
		{name: "4. cookie内容为中文UTF-8-Escape", contentType: "application/json;charset=utf-8", cookie: http.Cookie{Name: "cname", Value: url.QueryEscape("中文")}, want: types.XMap{"cname": url.QueryEscape("中文")}},
	}

	confObj := mocks.NewConf() //构建对象
	confObj.API("8080")        //初始化参数
	hydra.G.SysName = "apiserver"
	serverConf := confObj.GetAPIConf() //获取配置
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	for _, tt := range tests {
		r, err := http.NewRequest("POST", "http://localhost:9091/getcookies/encoding", strings.NewReader(""))
		assert.Equal(t, nil, err, "构建请求")

		//设置content-type
		r.Header.Set("Content-Type", tt.contentType)

		//添加cookie
		r.AddCookie(&tt.cookie)
		c.Request = r

		req := ctx.NewRequest(middleware.NewGinCtx(c), serverConf, conf.NewMeta())
		got := req.Cookies()
		assert.Equal(t, tt.want, got, tt.name)
	}
}
