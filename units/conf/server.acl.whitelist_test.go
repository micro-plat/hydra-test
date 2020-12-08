/*
author:taoshouyin
time:2020-10-15
*/

package conf

import (
	"strings"
	"testing"

	"github.com/micro-plat/hydra/conf/server/acl/whitelist"
	"github.com/micro-plat/hydra/test/assert"
	"github.com/micro-plat/hydra/test/mocks"
)

func TestWhitelistNew(t *testing.T) {
	tests := []struct {
		name      string
		opts      []whitelist.Option
		want      *whitelist.WhiteList
		allowIP   string
		allowReq  string
		wantallow bool
	}{
		{name: "1. Conf-WhitelistNew-初始化默认白名单配置", opts: []whitelist.Option{}, want: &whitelist.WhiteList{IPS: make([]*whitelist.IPList, 0, 1)}, allowIP: "192.168.0.101", allowReq: "/t1/t2/t3", wantallow: true},
		{name: "2. Conf-WhitelistNew-初始化Disable白名单配置", opts: []whitelist.Option{whitelist.WithDisable()}, want: &whitelist.WhiteList{IPS: make([]*whitelist.IPList, 0, 1), Disable: true}, allowIP: "192.168.0.101", allowReq: "/t1/t2/t3", wantallow: true},
		{name: "3. Conf-WhitelistNew-初始化Enable白名单配置", opts: []whitelist.Option{whitelist.WithEnable()}, want: &whitelist.WhiteList{IPS: make([]*whitelist.IPList, 0, 1), Disable: false}, allowIP: "192.168.0.101", allowReq: "/t1/t2/t3", wantallow: true},
		{name: "4. Conf-WhitelistNew-初始化自定义ip白名单配置", opts: []whitelist.Option{whitelist.WithIPList(whitelist.NewIPList([]string{"/t1/t2/*"}, []string{"192.168.0.101"}...))}, want: &whitelist.WhiteList{IPS: []*whitelist.IPList{&whitelist.IPList{Requests: []string{"/t1/t2/*"}, IPS: []string{"192.168.0.101"}}}}, allowIP: "192.168.0.101", allowReq: "/t1/t2/t3", wantallow: true},
	}
	for _, tt := range tests {
		got := whitelist.New(tt.opts...)
		assert.Equal(t, tt.want.Disable, got.Disable, tt.name+",disable")

		//比对白名单对象长度
		assert.Equal(t, len(tt.want.IPS), len(got.IPS), tt.name+",ips len")

		for i, item := range got.IPS {
			assert.Equal(t, tt.want.IPS[i].IPS, item.IPS, tt.name+",ips ips")
			assert.Equal(t, tt.want.IPS[i].Requests, item.Requests, tt.name+",ips request")
		}

		//测试私有匹配参数是否成功赋值
		allowgot := got.IsAllow(tt.allowReq, tt.allowIP)
		assert.Equal(t, tt.wantallow, allowgot, tt.name+",allow")
	}
}

//白名单匹配暂时不要测试  匹配方案没有确定
func TestWhiteList_IsAllow(t *testing.T) {

	tests := []struct {
		name    string
		opts    []whitelist.Option
		argPath string
		argIP   string
		want    bool
	}{}
	for _, tt := range tests {
		w := whitelist.New(tt.opts...)
		got := w.IsAllow(tt.argPath, tt.argIP)
		if got != tt.want {
			t.Errorf("WhiteList.IsAllow() = %v, want %v", got, tt.want)
		}
	}
}

func TestWhiteListGetConf(t *testing.T) {
	tests := []struct {
		name string
		opts []whitelist.Option
		want *whitelist.WhiteList
	}{
		{name: "1. Conf-WhitelistGetConf-节点不存在,获取默认对象", opts: []whitelist.Option{}, want: &whitelist.WhiteList{Disable: true}},
		{name: "2. Conf-WhitelistGetConf-节点为空,获取默认对象", opts: []whitelist.Option{}, want: whitelist.New()},
		{name: "3. Conf-WhitelistGetConf-正常对象获取", opts: []whitelist.Option{whitelist.WithIPList(whitelist.NewIPList([]string{"/t1/t2/*"}, []string{"192.168.0.101"}...))}, want: whitelist.New(whitelist.WithIPList(whitelist.NewIPList([]string{"/t1/t2/*"}, []string{"192.168.0.101"}...)))},
	}

	//初始化服务conf配置对象
	conf := mocks.NewConfBy("hydraconf_whitelist_test", "whitelist")
	confB := conf.API(":8081")
	for _, tt := range tests {
		if !strings.EqualFold(tt.name, "节点不存在,获取默认对象") {
			confB.WhiteList(tt.opts...)
		}
		obj, err := whitelist.GetConf(conf.GetAPIConf().GetServerConf())
		assert.Equal(t, nil, err, tt.name+",err")
		assert.Equal(t, len(tt.want.IPS), len(obj.IPS), tt.name)

	}
}
