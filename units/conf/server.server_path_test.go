package conf

import (
	"strings"
	"testing"

	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/hydra/test/assert"

	"github.com/micro-plat/hydra/conf/server"
)

func TestPub_GetServerPath(t *testing.T) {
	type fields struct {
		platName    string
		sysName     string
		serverType  string
		clusterName string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "1. Conf-PubGetServerPath-获取服务主路径1", fields: fields{platName: "p1", sysName: "sys1", serverType: "st1", clusterName: "cn1"}, want: "/p1/sys1/st1/cn1/conf"},
		{name: "2. Conf-PubGetServerPath-获取服务主路径2", fields: fields{platName: "p2", sysName: "sys2", serverType: "st2", clusterName: "cn2"}, want: "/p2/sys2/st2/cn2/conf"},
	}
	for _, tt := range tests {
		c := server.NewServerPub(tt.fields.platName, tt.fields.sysName, tt.fields.serverType, tt.fields.clusterName)
		got := c.GetServerPath()
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestPub_GetSubConfPath(t *testing.T) {
	type fields struct {
		platName    string
		sysName     string
		serverType  string
		clusterName string
	}
	tests := []struct {
		name   string
		fields fields
		args   []string
		want   string
	}{
		{name: "1. Conf-PubGetSubConfPath-获取主子节点下面的子路经,nil", fields: fields{platName: "p1", sysName: "sys1", serverType: "st1", clusterName: "cn1"}, args: nil, want: "/p1/sys1/st1/cn1/conf"},
		{name: "2. Conf-PubGetSubConfPath-获取主子节点下面的子路经,空", fields: fields{platName: "p1", sysName: "sys1", serverType: "st1", clusterName: "cn1"}, args: []string{}, want: "/p1/sys1/st1/cn1/conf"},
		{name: "3. Conf-PubGetSubConfPath-获取主子节点下面的子路经,一段", fields: fields{platName: "p1", sysName: "sys1", serverType: "st1", clusterName: "cn1"}, args: []string{"sub1"}, want: "/p1/sys1/st1/cn1/conf/sub1"},
		{name: "4. Conf-PubGetSubConfPath-获取主子节点下面的子路经,三段", fields: fields{platName: "p1", sysName: "sys1", serverType: "st1", clusterName: "cn1"}, args: []string{"sub1", "sub2", "sub3"}, want: "/p1/sys1/st1/cn1/conf/sub1/sub2/sub3"},
	}
	for _, tt := range tests {
		c := server.NewServerPub(tt.fields.platName, tt.fields.sysName, tt.fields.serverType, tt.fields.clusterName)
		got := c.GetSubConfPath(tt.args...)
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestPub_GetRPCServicePubPath(t *testing.T) {
	type fields struct {
		platName    string
		sysName     string
		serverType  string
		clusterName string
	}
	tests := []struct {
		name   string
		fields fields
		svName string
		want   string
	}{
		{name: "1. Conf-PubGetRPCServicePubPath-获取RPC服务发布跟路径,空", fields: fields{platName: "p1", sysName: "sys1", serverType: "st1", clusterName: "cn1"}, svName: "", want: "/p1/services/st1/providers"},
		{name: "2. Conf-PubGetRPCServicePubPath-获取RPC服务发布跟路径1", fields: fields{platName: "p1", sysName: "sys1", serverType: "st1", clusterName: "cn1"}, svName: "name1", want: "/p1/services/st1/name1/providers"},
		{name: "3. Conf-PubGetRPCServicePubPath-获取RPC服务发布跟路径2", fields: fields{platName: "p1", sysName: "sys1", serverType: "st1", clusterName: "cn1"}, svName: "name2", want: "/p1/services/st1/name2/providers"},
	}
	for _, tt := range tests {
		c := server.NewServerPub(tt.fields.platName, tt.fields.sysName, tt.fields.serverType, tt.fields.clusterName)
		got := c.GetRPCServicePubPath(tt.svName)
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestPub_GetServicePubPath(t *testing.T) {
	type fields struct {
		platName    string
		sysName     string
		serverType  string
		clusterName string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "1. Conf-PubGetServicePubPath-获取服务发布跟路径1", fields: fields{platName: "p1", sysName: "sys1", serverType: "st1", clusterName: "cn1"}, want: "/p1/services/st1/providers"},
		{name: "2. Conf-PubGetServicePubPath-获取服务发布跟路径2", fields: fields{platName: "p2", sysName: "sys1", serverType: "st1", clusterName: "cn1"}, want: "/p2/services/st1/providers"},
		{name: "3. Conf-PubGetServicePubPath-获取服务发布跟路径3", fields: fields{platName: "p3", sysName: "sys1", serverType: "st1", clusterName: "cn1"}, want: "/p3/services/st1/providers"},
	}
	for _, tt := range tests {
		c := server.NewServerPub(tt.fields.platName, tt.fields.sysName, tt.fields.serverType, tt.fields.clusterName)
		got := c.GetServicePubPath()
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestPub_GetDNSPubPath(t *testing.T) {
	type fields struct {
		platName    string
		sysName     string
		serverType  string
		clusterName string
	}
	tests := []struct {
		name   string
		fields fields
		svName string
		want   string
	}{
		{name: "1. Conf-PubGetDNSPubPath-获取DNS服务路径,空", fields: fields{platName: "p1", sysName: "sys1", serverType: "st1", clusterName: "cn1"}, svName: "", want: "/dns"},
		{name: "2. Conf-PubGetDNSPubPath-获取DNS服务路径1", fields: fields{platName: "p2", sysName: "sys2", serverType: "st2", clusterName: "cn2"}, svName: "name1", want: "/dns/name1"},
		{name: "3. Conf-PubGetDNSPubPath-获取DNS服务路径2", fields: fields{platName: "p3", sysName: "sys3", serverType: "st3", clusterName: "cn3"}, svName: "name2", want: "/dns/name2"},
	}
	for _, tt := range tests {
		c := server.NewServerPub(tt.fields.platName, tt.fields.sysName, tt.fields.serverType, tt.fields.clusterName)
		got := c.GetDNSPubPath(tt.svName)
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestPub_GetServerPubPath(t *testing.T) {
	type fields struct {
		platName    string
		sysName     string
		serverType  string
		clusterName string
	}
	tests := []struct {
		name      string
		fields    fields
		clustName []string
		want      string
	}{
		{name: "1. Conf-PubGetServerPubPath-获取服务器运行节点路径,nil", fields: fields{platName: "p1", sysName: "sys1", serverType: "st1", clusterName: "cn1"}, clustName: nil, want: "/p1/sys1/st1/cn1/servers"},
		{name: "2. Conf-PubGetServerPubPath-获取服务器运行节点路径,空", fields: fields{platName: "p1", sysName: "sys1", serverType: "st1", clusterName: "cn1"}, clustName: []string{}, want: "/p1/sys1/st1/cn1/servers"},
		{name: "3. Conf-PubGetServerPubPath-获取服务器运行节点路径,一段", fields: fields{platName: "p1", sysName: "sys1", serverType: "st1", clusterName: "cn1"}, clustName: []string{"xx"}, want: "/p1/sys1/st1/xx/servers"},
		{name: "4. Conf-PubGetServerPubPath-获取服务器运行节点路径,二段", fields: fields{platName: "p1", sysName: "sys1", serverType: "st1", clusterName: "cn1"}, clustName: []string{"cc", "xx"}, want: "/p1/sys1/st1/cc/servers"},
		{name: "5. Conf-PubGetServerPubPath-获取服务器运行节点路径,三段", fields: fields{platName: "p1", sysName: "sys1", serverType: "st1", clusterName: "cn1"}, clustName: []string{"aa", "cc", "xx"}, want: "/p1/sys1/st1/aa/servers"},
	}
	for _, tt := range tests {
		c := server.NewServerPub(tt.fields.platName, tt.fields.sysName, tt.fields.serverType, tt.fields.clusterName)
		got := c.GetServerPubPath(tt.clustName...)
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestPub_GetServerName(t *testing.T) {
	type fields struct {
		platName    string
		sysName     string
		serverType  string
		clusterName string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "1. Conf-PubGetServerName-获取服务器名称1", fields: fields{platName: "p1", sysName: "sys1", serverType: "st1", clusterName: "cn1"}, want: "sys1(cn1)"},
		{name: "2. Conf-PubGetServerName-获取服务器名称2", fields: fields{platName: "p2", sysName: "sys2", serverType: "st2", clusterName: "cn2"}, want: "sys2(cn2)"},
		{name: "3. Conf-PubGetServerName-获取服务器名称3", fields: fields{platName: "p3", sysName: "sys3", serverType: "st3", clusterName: "cn3"}, want: "sys3(cn3)"},
	}
	for _, tt := range tests {
		c := server.NewServerPub(tt.fields.platName, tt.fields.sysName, tt.fields.serverType, tt.fields.clusterName)
		got := c.GetServerName()
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestPub_AllowGray(t *testing.T) {
	type fields struct {
		platName    string
		sysName     string
		serverType  string
		clusterName string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "1. Conf-PubAllowGray-获取服务器名称API", fields: fields{platName: "p1", sysName: "sys1", serverType: global.API, clusterName: "cn1"}, want: true},
		{name: "2. Conf-PubAllowGray-获取服务器名称Web", fields: fields{platName: "p2", sysName: "sys2", serverType: global.Web, clusterName: "cn2"}, want: true},
		{name: "3. Conf-PubAllowGray-获取服务器名称CRON", fields: fields{platName: "p3", sysName: "sys3", serverType: global.CRON, clusterName: "cn3"}, want: false},
		{name: "4. Conf-PubAllowGray-获取服务器名称MQC", fields: fields{platName: "p3", sysName: "sys3", serverType: global.MQC, clusterName: "cn3"}, want: false},
		{name: "5. Conf-PubAllowGray-获取服务器名称RPC", fields: fields{platName: "p3", sysName: "sys3", serverType: global.RPC, clusterName: "cn3"}, want: false},
		{name: "6. Conf-PubAllowGray-获取服务器名称WS", fields: fields{platName: "p3", sysName: "sys3", serverType: global.WS, clusterName: "cn3"}, want: false},
		{name: "7. Conf-PubAllowGray-获取服务器名称other", fields: fields{platName: "p3", sysName: "sys3", serverType: "other", clusterName: "cn3"}, want: false},
	}
	for _, tt := range tests {
		c := server.NewServerPub(tt.fields.platName, tt.fields.sysName, tt.fields.serverType, tt.fields.clusterName)
		got := c.AllowGray()
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestPub_GetFunc(t *testing.T) {
	type fields struct {
		platName    string
		sysName     string
		serverType  string
		clusterName string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{name: "1. Conf-PubGetFunc-pub对象属性获取方法测试1", fields: fields{platName: "p1", sysName: "sys1", serverType: global.API, clusterName: "cn1"}},
		{name: "2. Conf-PubGetFunc-pub对象属性获取方法测试2", fields: fields{platName: "p2", sysName: "sys2", serverType: global.Web, clusterName: "cn2"}},
		{name: "3. Conf-PubGetFunc-pub对象属性获取方法测试3", fields: fields{platName: "p3", sysName: "sys3", serverType: global.CRON, clusterName: "cn3"}},
		{name: "4. Conf-PubGetFunc-pub对象属性获取方法测试4", fields: fields{platName: "p3", sysName: "sys3", serverType: global.MQC, clusterName: "cn3"}},
	}
	for _, tt := range tests {
		c := server.NewServerPub(tt.fields.platName, tt.fields.sysName, tt.fields.serverType, tt.fields.clusterName)
		got := c.GetPlatName()
		assert.Equal(t, tt.fields.platName, got, tt.name)
		got = c.GetSysName()
		assert.Equal(t, tt.fields.sysName, got, tt.name)
		got = c.GetServerType()
		assert.Equal(t, tt.fields.serverType, got, tt.name)
		got = c.GetClusterName()
		assert.Equal(t, tt.fields.clusterName, got, tt.name)
		got = c.GetServerID()
		assert.Equal(t, true, strings.HasPrefix(got, global.GetMatchineCode()), tt.name)
	}
}
