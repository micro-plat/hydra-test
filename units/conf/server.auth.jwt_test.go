/*
author:taoshouyin
time:2020-10-16
*/

package conf

import (
	"strings"
	"testing"

	"github.com/micro-plat/hydra/conf"
	"github.com/micro-plat/hydra/conf/server/auth/jwt"
	"github.com/micro-plat/hydra/test/assert"
	"github.com/micro-plat/hydra/test/mocks"
)

func TestNewJWT(t *testing.T) {
	tests := []struct {
		name string
		opts []jwt.Option
		want *jwt.JWTAuth
	}{
		{name: "1. Conf-NewJWT-设置secert", opts: []jwt.Option{jwt.WithSecret("12345678")}, want: &jwt.JWTAuth{Name: "Authorization-Jwt", Mode: "HS512", Secret: "12345678", ExpireAt: 86400, Source: "COOKIE", PathMatch: conf.NewPathMatch()}},
		{name: "2. Conf-NewJWT-设置disable", opts: []jwt.Option{jwt.WithSecret("12345678"), jwt.WithDisable()}, want: &jwt.JWTAuth{Name: "Authorization-Jwt", Mode: "HS512", Secret: "12345678", Disable: true, ExpireAt: 86400, Source: "COOKIE", PathMatch: conf.NewPathMatch()}},
		{name: "3. Conf-NewJWT-设置Enable", opts: []jwt.Option{jwt.WithSecret("12345678"), jwt.WithEnable()}, want: &jwt.JWTAuth{Name: "Authorization-Jwt", Mode: "HS512", Secret: "12345678", Disable: false, ExpireAt: 86400, Source: "COOKIE", PathMatch: conf.NewPathMatch()}},
		{name: "4. Conf-NewJWT-设置自定义对象", opts: []jwt.Option{jwt.WithSecret("12345678"), jwt.WithHeader(), jwt.WithExcludes("/t1/**"), jwt.WithExpireAt(1000), jwt.WithMode("ES256"), jwt.WithName("test"), jwt.WithAuthURL("1111")}, want: &jwt.JWTAuth{Name: "test", AuthURL: "1111", Mode: "ES256", Secret: "12345678", ExpireAt: 1000, Source: "HEADER", Excludes: []string{"/t1/**"}, PathMatch: conf.NewPathMatch("/t1/**")}},
	}
	for _, tt := range tests {
		got := jwt.NewJWT(tt.opts...)
		assert.Equal(t, tt.want, got, tt.name)
	}
}

//还有panic的异常情况没有测试  等修改后测试
func TestJWTGetConf(t *testing.T) {

	tests := []struct {
		name string
		opts []jwt.Option
		want *jwt.JWTAuth
	}{
		{name: "1. Conf-JWTGetConf-未设置jwt节点", opts: []jwt.Option{}, want: &jwt.JWTAuth{Disable: true}},
		{name: "2. Conf-JWTGetConf-配置参数正确", opts: []jwt.Option{jwt.WithExpireAt(123), jwt.WithSecret("11111")}, want: jwt.NewJWT(jwt.WithExpireAt(123), jwt.WithSecret("11111"))},
	}

	conf := mocks.NewConfBy("hydraconf_jwt_test2", "jwttest")
	confB := conf.API(":8081")
	for _, tt := range tests {
		if !strings.EqualFold(tt.name, "1. Conf-JWTGetConf-未设置jwt节点") {
			confB.Jwt(tt.opts...)
		}
		got, err := jwt.GetConf(conf.GetAPIConf().GetServerConf())
		assert.Equal(t, nil, err, tt.name+",err")
		assert.Equal(t, got, tt.want, tt.name)
	}
}
