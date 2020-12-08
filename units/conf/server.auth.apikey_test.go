/*
author:taoshouyin
time:2020-10-16
*/

package conf

import (
	"strings"
	"testing"

	"github.com/micro-plat/hydra/conf"
	"github.com/micro-plat/hydra/conf/server/auth/apikey"
	"github.com/micro-plat/hydra/test/assert"
	"github.com/micro-plat/hydra/test/mocks"
	"github.com/micro-plat/lib4go/security/md5"
	"github.com/micro-plat/lib4go/security/sha1"
	"github.com/micro-plat/lib4go/security/sha256"
)

var md5secret = "12345678"
var sha1secret = "1234567812345678"
var sha256secret = "9876543210222222222"
var rawData = "taosy hydra test"
var md5Sign = md5.Encrypt(rawData + md5secret)
var sha1Sign = sha1.Encrypt(rawData + sha1secret)
var sha256Sign = sha256.Encrypt(rawData + sha256secret)

func TestAPIKeyNew(t *testing.T) {
	tests := []struct {
		name   string
		secret string
		opts   []apikey.Option
		want   *apikey.APIKeyAuth
	}{
		{name: "1. Conf-APIKeyNew-初始化默认对象", secret: "", opts: []apikey.Option{}, want: &apikey.APIKeyAuth{Mode: "MD5", PathMatch: conf.NewPathMatch()}},
		{name: "2. Conf-APIKeyNew-设置密钥和路径", secret: "1111", opts: []apikey.Option{apikey.WithSecret("123456"), apikey.WithExcludes("/t/tw", "/t1/t2")}, want: &apikey.APIKeyAuth{Secret: "123456", Excludes: []string{"/t/tw", "/t1/t2"}, Mode: "MD5", PathMatch: conf.NewPathMatch("/t/tw", "/t1/t2")}},
		{name: "3. Conf-APIKeyNew-设置md5", secret: "1111", opts: []apikey.Option{apikey.WithMD5Mode()}, want: &apikey.APIKeyAuth{Secret: "1111", Mode: "MD5", Disable: false, PathMatch: conf.NewPathMatch()}},
		{name: "4. Conf-APIKeyNew-设置sha1", secret: "1111", opts: []apikey.Option{apikey.WithSHA1Mode(), apikey.WithDisable()}, want: &apikey.APIKeyAuth{Secret: "1111", Mode: "SHA1", Disable: true, PathMatch: conf.NewPathMatch()}},
		{name: "5. Conf-APIKeyNew-设置sha256", secret: "", opts: []apikey.Option{apikey.WithSHA256Mode(), apikey.WithEnable()}, want: &apikey.APIKeyAuth{Mode: "SHA256", Disable: false, PathMatch: conf.NewPathMatch()}},
	}
	for _, tt := range tests {
		got := apikey.New(tt.secret, tt.opts...)
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestAPIKeyAuth_Verify(t *testing.T) {
	type args struct {
		raw  string
		sign string
	}
	tests := []struct {
		name       string
		mode       string
		secret     string
		args       args
		wantErr    bool
		wantErrStr string
	}{
		{name: "1. Conf-APIKeyVerify-不支持的签名方式", mode: "md4", secret: md5secret, args: args{raw: rawData, sign: md5Sign}, wantErr: true, wantErrStr: "不支持的签名验证方式:md4"},
		{name: "2. Conf-APIKeyVerify-签名方式不正确", mode: "md5", secret: sha1secret, args: args{raw: rawData, sign: sha1Sign}, wantErr: true, wantErrStr: "签名错误:raw:"},
		{name: "3. Conf-APIKeyVerify-签名数据错误", mode: "md5", secret: md5secret, args: args{raw: "rawData", sign: md5Sign}, wantErr: true, wantErrStr: "签名错误:raw:"},
		{name: "4. Conf-APIKeyVerify-密钥错误", mode: "md5", secret: "md5secret", args: args{raw: rawData, sign: md5Sign}, wantErr: true, wantErrStr: "签名错误:raw:"},
		{name: "5. Conf-APIKeyVerify-md5签名成功", mode: "md5", secret: md5secret, args: args{raw: rawData, sign: md5Sign}, wantErr: false, wantErrStr: ""},
		{name: "6. Conf-APIKeyVerify-sha1签名成功", mode: "sha1", secret: sha1secret, args: args{raw: rawData, sign: sha1Sign}, wantErr: false, wantErrStr: ""},
		{name: "7. Conf-APIKeyVerify-sha256签名成功", mode: "sha256", secret: sha256secret, args: args{raw: rawData, sign: sha256Sign}, wantErr: false, wantErrStr: ""},
	}
	for _, tt := range tests {
		a := &apikey.APIKeyAuth{Mode: tt.mode, Secret: tt.secret}
		err := a.Verify(tt.args.raw, tt.args.sign)
		assert.Equal(t, tt.wantErr, (err != nil), tt.name)
		if tt.wantErr {
			assert.Equal(t, tt.wantErrStr, err.Error()[:len(tt.wantErrStr)], tt.name)
		}
	}
}

func TestApikeyGetConf(t *testing.T) {
	type test struct {
		name string
		opts []apikey.Option
		want *apikey.APIKeyAuth
	}

	apiConf := mocks.NewConfBy("hydraconf_apikey_test", "apikey")
	confB := apiConf.API(":8081")
	test1 := test{name: "1.1 Conf-APIKeyGetConf-未设置apikey节点", want: &apikey.APIKeyAuth{Disable: true, PathMatch: conf.NewPathMatch()}}
	got, err := apikey.GetConf(apiConf.GetAPIConf().GetServerConf())
	assert.Equal(t, nil, err, test1.name+",err")
	assert.Equal(t, got, test1.want, test1.name)

	test2 := test{name: "2.1 Conf-APIKeyGetConf-配置参数正确", opts: []apikey.Option{apikey.WithMD5Mode(), apikey.WithDisable(), apikey.WithExcludes("/t1/t2"), apikey.WithSecret("123456")},
		want: apikey.New("123456", apikey.WithMD5Mode(), apikey.WithDisable(), apikey.WithExcludes("/t1/t2"))}
	confB.APIKEY("", test2.opts...)
	got, err = apikey.GetConf(apiConf.GetAPIConf().GetServerConf())
	assert.Equal(t, nil, err, test2.name+",err")
	assert.Equal(t, got, test2.want, test2.name)
}

func TestApikeyGetConf1(t *testing.T) {
	type test struct {
		name string
		opts []apikey.Option
		want *apikey.APIKeyAuth
	}

	apiConf := mocks.NewConfBy("hydraconf_apikey_test1", "apikey1")
	confB := apiConf.API(":8081")
	test1 := test{name: "3.1 Conf-APIKeyGetConf-节点密钥不存在,验证异常", opts: []apikey.Option{apikey.WithMD5Mode()}, want: nil}
	confB.APIKEY("", test1.opts...)
	got, err := apikey.GetConf(apiConf.GetAPIConf().GetServerConf())
	assert.Equal(t, true, strings.Contains(err.Error(), "apikey配置数据有误"), test1.name)
	assert.Equal(t, test1.want, got, test1.name)
}
func TestApikeyGetConf2(t *testing.T) {
	type test struct {
		name string
		opts []apikey.Option
		want *apikey.APIKeyAuth
	}

	apiConf := mocks.NewConfBy("hydraconf_apikey_test2", "apikey2")
	confB := apiConf.API(":8081")
	test1 := test{name: "4.1 Conf-APIKeyGetConf-apikey修改为错误json串", opts: []apikey.Option{apikey.WithMD5Mode(), apikey.WithDisable(), apikey.WithExcludes("/t1/t2"), apikey.WithSecret("123456")},
		want: apikey.New("123456", apikey.WithMD5Mode(), apikey.WithDisable(), apikey.WithExcludes("/t1/t2"))}
	confB.APIKEY("", test1.opts...)
	// 修改json数据不合法
	path := apiConf.GetAPIConf().GetServerConf().GetSubConfPath("auth", "apikey")
	apiConf.Registry.Update(path, "错误的json字符串")
}
