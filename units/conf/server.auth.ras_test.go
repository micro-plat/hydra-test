/*
author:taoshouyin
time:2020-10-16
*/

package conf

import (
	"testing"

	"github.com/micro-plat/hydra/test/assert"
	"github.com/micro-plat/hydra/test/mocks"

	"github.com/micro-plat/hydra/conf"
	"github.com/micro-plat/hydra/conf/server/auth/ras"
)

//
func xTestNewAuth(t *testing.T) {
	tests := []struct {
		name    string
		service string
		opts    []ras.AuthOption
		want    *ras.Auth
	}{
		{name: "1. conf-NewAuth-设置默认对象", service: "", opts: []ras.AuthOption{}, want: &ras.Auth{Service: "", Requests: []string{"*"}, Connect: &ras.Connect{}, PathMatch: conf.NewPathMatch([]string{"*"}...),
			Params: make(map[string]interface{}), Required: make([]string, 0, 1), Alias: make(map[string]string), Decrypt: make([]string, 0, 1)}},
		{name: "2. conf-NewAuth-设置service对象", service: "test-tsy", opts: []ras.AuthOption{}, want: &ras.Auth{Service: "test-tsy", Requests: []string{"*"}, Connect: &ras.Connect{}, PathMatch: conf.NewPathMatch([]string{"*"}...),
			Params: make(map[string]interface{}), Required: make([]string, 0, 1), Alias: make(map[string]string), Decrypt: make([]string, 0, 1)}},
		{name: "3. conf-NewAuth-设置Enable对象", service: "", opts: []ras.AuthOption{ras.WithAuthEnable()}, want: &ras.Auth{Service: "", Requests: []string{"*"}, Connect: &ras.Connect{}, PathMatch: conf.NewPathMatch([]string{"*"}...),
			Params: make(map[string]interface{}), Required: make([]string, 0, 1), Alias: make(map[string]string), Decrypt: make([]string, 0, 1)}},
		{name: "4. conf-NewAuth-设置全量对象", service: "test-tsy",
			opts: []ras.AuthOption{ras.WithRequest("/t1/t2"), ras.WithRequired("taofield"), ras.WithUIDAlias("userID"), ras.WithTimestampAlias("timespan"), ras.WithSignAlias("signname"),
				ras.WithCheckTimestamp(false), ras.WithDecryptName("duser"), ras.WithParam("key1", "v1"), ras.WithParam("key2", "v2"), ras.WithAuthDisable()},
			want: &ras.Auth{Service: "test-tsy", Requests: []string{"/t1/t2"}, Connect: &ras.Connect{}, Params: map[string]interface{}{"key1": "v1", "key2": "v2"}, PathMatch: conf.NewPathMatch([]string{"/t1/t2"}...),
				Required: []string{"taofield"}, Alias: map[string]string{"euid": "userID", "timestamp": "timespan", "sign": "signname"}, Disable: true, CheckTS: false, Decrypt: []string{"duser"}}},
	}
	for _, tt := range tests {
		got := ras.New(tt.service, tt.opts...)
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestNewRASAuth(t *testing.T) {
	tests := []struct {
		name string
		opts []ras.Option
		want *ras.RASAuth
	}{
		{name: "1. Conf-NewRASAuth-设置默认对象", opts: []ras.Option{}, want: &ras.RASAuth{Disable: false, Auth: nil}},
		{name: "2. Conf-NewRASAuth-设置disable对象", opts: []ras.Option{ras.WithDisable()}, want: &ras.RASAuth{Disable: true, Auth: nil}},
		{name: "3. Conf-NewRASAuth-设置enable对象", opts: []ras.Option{ras.WithEnable()}, want: &ras.RASAuth{Disable: false, Auth: nil}},
		{name: "4. Conf-NewRASAuth-设置auth对象", opts: []ras.Option{ras.WithAuths(ras.New("tt", ras.WithRequest("/t1/t2")), ras.New("tt1", ras.WithRequired("taofield")))},
			want: &ras.RASAuth{Disable: false, Auth: []*ras.Auth{ras.New("tt", ras.WithRequest("/t1/t2")), ras.New("tt1", ras.WithRequired("taofield"))}}},
	}
	for _, tt := range tests {
		got := ras.NewRASAuth(tt.opts...)
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestRASAuth_Match(t *testing.T) {
	tests := []struct {
		name string
		auth *ras.RASAuth
		args string
		want bool
	}{
		{name: "1. Conf-RASAuthMatch-空对象匹配", auth: ras.NewRASAuth(), args: "/t1", want: false},
		{name: "2. Conf-RASAuthMatch-默认auth对象匹配", auth: ras.NewRASAuth(ras.WithAuths(ras.New(""))), args: "/t1", want: true},
		{name: "3. Conf-RASAuthMatch-自定义auth对象匹配失败", auth: ras.NewRASAuth(ras.WithAuths(ras.New("", ras.WithRequest("/t1/t2")))), args: "/t1", want: false},
		{name: "4. Conf-RASAuthMatch-自定义auth对象匹配成功", auth: ras.NewRASAuth(ras.WithAuths(ras.New("", ras.WithRequest("/t1/t2")))), args: "/t1/t2", want: true},
		{name: "5. Conf-RASAuthMatch-自定义auth对象模糊匹配失败", auth: ras.NewRASAuth(ras.WithAuths(ras.New("", ras.WithRequest("/t1/*")))), args: "/t1/t2/t3", want: false},
		{name: "6. Conf-RASAuthMatch-自定义auth对象模糊匹配成功", auth: ras.NewRASAuth(ras.WithAuths(ras.New("", ras.WithRequest("/t1/*")))), args: "/t1/tt", want: true},
		{name: "7. Conf-RASAuthMatch-自定义auth对象模糊匹配成功1", auth: ras.NewRASAuth(ras.WithAuths(ras.New("", ras.WithRequest("/t1/**")))), args: "/t1/t2/tt", want: true},
	}
	for _, tt := range tests {
		got, _ := tt.auth.Match(tt.args)
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestAuthRASGetConf(t *testing.T) {
	type test struct {
		name       string
		opts       []ras.Option
		wantAuths  *ras.RASAuth
		wantErr    bool
		wantErrStr string
	}

	conf := mocks.NewConfBy("hydraconf_ras_test2", "rastest")
	confB := conf.API(":8081")
	test1 := test{name: "1.1. Conf-AuthRASGetConf-未设置ras节点", opts: []ras.Option{}, wantAuths: &ras.RASAuth{Disable: true}, wantErr: false}
	gotAuths, err := ras.GetConf(conf.GetAPIConf().GetServerConf())
	assert.Equal(t, (err != nil), test1.wantErr, test1.name+",err")
	assert.Equal(t, gotAuths, test1.wantAuths, test1.name)

	tests := []test{
		{name: "2.1. Conf-AuthRASGetConf-设置ras数据格式错误节点", opts: []ras.Option{ras.WithAuths(ras.New(""))}, wantAuths: ras.NewRASAuth(ras.WithAuths(ras.New(""))), wantErr: true, wantErrStr: "RASAuth配置数据有误"},
		{name: "2.2. Conf-AuthRASGetConf-设置正确的配置数据", opts: []ras.Option{ras.WithAuths(ras.New("taosy", ras.WithRequest("/t1/t2")))},
			wantAuths: ras.NewRASAuth(ras.WithAuths(ras.New("taosy", ras.WithRequest("/t1/t2")))), wantErr: false, wantErrStr: ""},
	}
	for _, tt := range tests {
		confB.Ras(tt.opts...)
		gotAuths, err = ras.GetConf(conf.GetAPIConf().GetServerConf())
		assert.Equal(t, (err != nil), tt.wantErr, tt.name+",err")
		if !tt.wantErr {
			assert.Equal(t, gotAuths.Disable, tt.wantAuths.Disable, tt.name)
			assert.Equal(t, len(gotAuths.Auth), len(tt.wantAuths.Auth), tt.name)
		} else {
			assert.Equal(t, tt.wantErrStr, err.Error()[:len(tt.wantErrStr)], tt.name+",err")
		}
	}
}
