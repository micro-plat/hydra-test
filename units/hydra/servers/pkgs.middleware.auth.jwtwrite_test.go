package servers

import (
	"fmt"
	"testing"

	"github.com/micro-plat/hydra-test/units/mocks"
	"github.com/micro-plat/hydra/conf"
	"github.com/micro-plat/hydra/conf/server/auth/jwt"
	octx "github.com/micro-plat/hydra/context/ctx"
	"github.com/micro-plat/hydra/hydra/servers/pkg/middleware"
	"github.com/micro-plat/lib4go/assert"
	wjwt "github.com/micro-plat/lib4go/security/jwt"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/lib4go/utility"
)

//author:taoshouyin
//time:2020-11-12
//desc:测试jwt设置中间件逻辑
func TestJWTWriter(t *testing.T) {
	secert := utility.GetGUID()
	requestPath := "/jwtwrtier/test"
	data := map[string]interface{}{"sdsd": "sdfd", "3ddfs": "gggggg"}
	rawData, _ := wjwt.Encrypt(secert, jwt.ModeHS512, data, 86400)
	type testCase struct {
		name       string
		jwtOpts    []jwt.Option
		authData   map[string]interface{}
		isSource   string //cookie/header
		isSet      bool
		isSucc     bool
		domain     string
		wantStatus int
		wanttoken  string
	}

	tests := []*testCase{
		{name: "1.1 jwtwrite-配置不存在", isSet: false, authData: data, domain: "", wantStatus: 200, jwtOpts: []jwt.Option{}},

		{name: "2.1 jwtwrite-配置存在-未启动-无数据", isSet: true, authData: data, domain: "", wantStatus: 200, jwtOpts: []jwt.Option{jwt.WithDisable()}},
		{name: "2.2 jwtwrite-配置存在-未启动-正常默认配置", isSet: true, authData: data, domain: "", wantStatus: 200, jwtOpts: []jwt.Option{jwt.WithDisable(), jwt.WithName("testjwtname"), jwt.WithExcludes("/jwt/test")}},
		{name: "2.3 jwtwrite-配置存在-未启动-authdata不存在,head", isSet: true, domain: "", isSource: "header", authData: nil, wanttoken: "", wantStatus: 200, jwtOpts: []jwt.Option{jwt.WithDisable()}},
		{name: "2.4 jwtwrite-配置存在-未启动-authdata不存在,cookie", isSet: true, domain: "", isSource: "cookie", authData: nil, wanttoken: "", wantStatus: 200, jwtOpts: []jwt.Option{jwt.WithDisable()}},
		{name: "2.5 jwtwrite-配置存在-未启动-jwt设置header无domain", isSet: true, domain: "", authData: data, isSource: "header", wanttoken: rawData, wantStatus: 200, jwtOpts: []jwt.Option{jwt.WithDisable(), jwt.WithHeader(), jwt.WithSecret(secert), jwt.WithExcludes("/jwt/test1")}},
		{name: "2.6 jwtwrite-配置存在-未启动-jwt设置header有domain", isSet: true, domain: "www.baidu.com", authData: data, isSource: "header", wanttoken: rawData, wantStatus: 200, jwtOpts: []jwt.Option{jwt.WithDisable(), jwt.WithHeader(), jwt.WithSecret(secert), jwt.WithDomain("www.baidu.com"), jwt.WithExcludes("/jwt/test1")}},
		{name: "2.7 jwtwrite-配置存在-未启动-jwt设置cookie无domain", isSet: true, domain: "", authData: data, isSource: "cookie", wanttoken: rawData, wantStatus: 200, jwtOpts: []jwt.Option{jwt.WithDisable(), jwt.WithCookie(), jwt.WithSecret(secert), jwt.WithExcludes("/jwt/test1")}},
		{name: "2.8 jwtwrite-配置存在-未启动-jwt设置cookie有domain", isSet: true, domain: "www.baidu.com", authData: data, isSource: "cookie", wanttoken: rawData, wantStatus: 200, jwtOpts: []jwt.Option{jwt.WithDisable(), jwt.WithCookie(), jwt.WithSecret(secert), jwt.WithDomain("www.baidu.com"), jwt.WithExcludes("/jwt/test1")}},

		{name: "3.1 jwtwrite-配置存在-启动-无数据", isSet: true, authData: data, domain: "", wantStatus: 200, jwtOpts: []jwt.Option{}},
		{name: "3.2 jwtwrite-配置存在-启动-正常默认配置", isSet: true, authData: data, domain: "", wantStatus: 200, jwtOpts: []jwt.Option{jwt.WithName("testjwtname"), jwt.WithExcludes("/jwt/test")}},
		{name: "3.3 jwtwrite-配置存在-启动-authdata不存在,head", isSet: true, domain: "", isSource: "header", authData: nil, wanttoken: "", wantStatus: 200, jwtOpts: []jwt.Option{}},
		{name: "3.4 jwtwrite-配置存在-启动-authdata不存在,cookie", isSet: true, domain: "", isSource: "cookie", authData: nil, wanttoken: "", wantStatus: 200, jwtOpts: []jwt.Option{}},
		{name: "3.5 jwtwrite-配置存在-启动-jwt设置header无domain", isSucc: true, isSet: true, domain: "", authData: data, isSource: "header", wanttoken: rawData, wantStatus: 200, jwtOpts: []jwt.Option{jwt.WithHeader(), jwt.WithSecret(secert), jwt.WithExcludes("/jwt/test1")}},
		{name: "3.6 jwtwrite-配置存在-启动-jwt设置header有domain", isSucc: true, isSet: true, domain: "www.baidu.com", authData: data, isSource: "header", wanttoken: rawData, wantStatus: 200, jwtOpts: []jwt.Option{jwt.WithHeader(), jwt.WithSecret(secert), jwt.WithDomain("www.baidu.com"), jwt.WithExcludes("/jwt/test1")}},
		{name: "3.7 jwtwrite-配置存在-启动-jwt设置cookie无domain", isSucc: true, isSet: true, domain: "", authData: data, isSource: "cookie", wanttoken: rawData, wantStatus: 200, jwtOpts: []jwt.Option{jwt.WithCookie(), jwt.WithSecret(secert), jwt.WithExcludes("/jwt/test1")}},
		{name: "3.8 jwtwrite-配置存在-启动-jwt设置cookie有domain", isSucc: true, isSet: true, domain: "www.baidu.com", authData: data, isSource: "cookie", wanttoken: rawData, wantStatus: 200, jwtOpts: []jwt.Option{jwt.WithCookie(), jwt.WithSecret(secert), jwt.WithDomain("www.baidu.com"), jwt.WithExcludes("/jwt/test1")}},
	}

	for _, tt := range tests {
		fmt.Println(tt.name)
		mockConf := mocks.NewConfBy("middleware_jwtwrite_test", "jwtwrite")
		//初始化测试用例参数
		confB := mockConf.GetAPI()
		if tt.isSet {
			confB.Jwt(tt.jwtOpts...)
		}
		serverConf := mockConf.GetAPIConf()
		userAuth := &octx.Auth{}
		userAuth.Response(tt.authData)
		ctx := &mocks.MiddleContext{
			MockMeta:     conf.NewMeta(),
			MockUser:     &mocks.MockUser{MockClientIP: "192.168.0.1", MockAuth: userAuth},
			MockResponse: &mocks.MockResponse{MockStatus: 200, MockHeader: types.XMap{}},
			MockRequest: &mocks.MockRequest{
				MockPath: &mocks.MockPath{
					MockRequestPath: requestPath,
				},
			},
			MockAPPConf: serverConf,
		}
		//ctx := mock.NewContext("")
		//midCtx := middleware.NewMiddleContext(ctx, &mock.Middle{})

		//获取中间件
		handler := middleware.JwtWriter()
		//调用中间件
		handler(ctx)
		//断言结果
		gotStatus, _, _ := ctx.Response().GetFinalResponse()
		assert.Equalf(t, tt.wantStatus, gotStatus, tt.name, tt.wantStatus, gotStatus)
		headers := ctx.Response().GetHeaders()
		if tt.isSucc {
			if tt.isSource == "header" {
				assert.Equalf(t, jwt.TokenBearerPrefix+tt.wanttoken, headers[jwt.AuthorizationHeader], tt.name, tt.wanttoken, headers[jwt.AuthorizationHeader])
			} else {
				cookies := headers.GetString("Set-Cookie")

				assert.Equal(t, true, len(cookies) > 0, tt.name, tt.wanttoken)

			}
		}
	}
}
