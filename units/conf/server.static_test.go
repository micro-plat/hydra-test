package conf

import (
	"net/http"
	"testing"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra-test/units/mocks"
	"github.com/micro-plat/hydra/conf"
	"github.com/micro-plat/hydra/conf/server/static"
	"github.com/micro-plat/lib4go/assert"
)

func TestStaticNew(t *testing.T) {
	defaultObj := &static.Static{}
	enObj := &static.Static{
		HomePage:    "index1.html",
		AutoRewrite: true,
		Excludes:    []string{"/views/", ".exe", ".so", ".zip"},
		Disable:     true,
	}
	tests := []struct {
		name string
		opts []static.Option
		want *static.Static
	}{
		{name: "1. Conf-StaticNew-初始化nil对象", opts: nil, want: defaultObj},
		{name: "2. Conf-StaticNew-初始化空对象", opts: []static.Option{}, want: defaultObj},
		{name: "3. Conf-StaticNew-初始化image对象", opts: []static.Option{}, want: defaultObj},
		{name: "4. Conf-StaticNew-初始化设置全量对象", opts: []static.Option{static.WithAssetsPath("test"), static.WithHomePage("index1.html"),
			static.WithAutoRewrite(), static.WithDisable(), static.WithExclude("/views/", ".exe", ".so", ".zip")},
			want: enObj},
	}
	for _, tt := range tests {
		got := static.New("web", tt.opts...)
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestStatic_AllowRequest(t *testing.T) {
	tests := []struct {
		name   string
		fields *static.Static
		args   string
		want   bool
	}{
		{name: "1. Conf-StaticAllowRequest-Get支持的方法", fields: static.New("web"), args: http.MethodGet, want: true},
		{name: "2. Conf-StaticAllowRequest-Head支持的方法", fields: static.New("web"), args: http.MethodHead, want: true},
		{name: "3. Conf-StaticAllowRequest-Post不支持的方法", fields: static.New("web"), args: http.MethodPost, want: false},
		{name: "4. Conf-StaticAllowRequest-PUT不支持的方法", fields: static.New("web"), args: http.MethodPut, want: false},
		{name: "5. Conf-StaticAllowRequest-PATCH不支持的方法", fields: static.New("web"), args: http.MethodPatch, want: false},
		{name: "6. Conf-StaticAllowRequest-DELETE不支持的方法", fields: static.New("web"), args: http.MethodDelete, want: false},
		{name: "7. Conf-StaticAllowRequest-CONNECT不支持的方法", fields: static.New("web"), args: http.MethodConnect, want: false},
		{name: "8. Conf-StaticAllowRequest-OPTIONS不支持的方法", fields: static.New("web"), args: http.MethodOptions, want: false},
		{name: "9. Conf-StaticAllowRequest-TRACE不支持的方法", fields: static.New("web"), args: http.MethodTrace, want: false},
		{name: "10. Conf-StaticAllowRequest-other不支持的方法", fields: static.New("web"), args: "other", want: false},
	}
	for _, tt := range tests {
		got := tt.fields.AllowRequest(tt.args)
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestStaticGetConf(t *testing.T) {
	type test struct {
		name    string
		cnf     conf.IServerConf
		want    *static.Static
		wantErr bool
	}

	confO := mocks.NewConfBy("hydra", "graytest")
	confB := confO.API("8090")
	hydra.G.SysName = "apiserver"

	test1 := test{name: "static节点不存在", cnf: confO.GetAPIConf().GetServerConf(),
		want: &static.Static{
			Path:     "./static",
			Excludes: []string{"/view/", "/views/", "/web/", ".exe", ".so"},
			HomePage: "index.html",
			Disable:  true,
		}, wantErr: false}
	staticObj, err := static.GetConf(test1.cnf)
	assert.Equal(t, test1.wantErr, (err != nil), test1.name+",err")
	assert.Equal(t, test1.want, staticObj, test1.name+",obj")

	confB.Static(static.WithAssetsPath("dddd"))
	test3 := test{name: "static节点存在,数据正确", cnf: confO.GetAPIConf().GetServerConf(), want: static.New("web", static.WithAssetsPath("dddd")), wantErr: false}
	staticObj, err = static.GetConf(test3.cnf)
	assert.Equal(t, test3.wantErr, (err != nil), test3.name+",err")
	assert.Equal(t, test3.want, staticObj, test3.name+",obj")

	//处理归档文件
}

func TestStatic_IsExclude(t *testing.T) {
	tests := []struct {
		name   string
		fields *static.Static
		rPath  string
		want   bool
	}{
		{name: "1. Conf-StaticIsExclude-空Exclude对象", fields: static.New("web"), rPath: "/test", want: false},
		{name: "2. Conf-StaticIsExclude-Exclude对象,路径匹配成功", fields: static.New("web", static.WithExclude("/test")), rPath: "/test", want: true},
		{name: "3. Conf-StaticIsExclude-Exclude对象，扩展名匹配成功", fields: static.New("web", static.WithExclude(".so")), rPath: "/test1.so", want: true},
		{name: "4. Conf-StaticIsExclude-Exclude对象，路径匹配失败", fields: static.New("web", static.WithExclude("/test11")), rPath: "/test1", want: false},
		{name: "5. Conf-StaticIsExclude-Exclude对象，扩展名匹配失败", fields: static.New("web", static.WithExclude(".so")), rPath: "/test11.txt", want: false},
	}
	for _, tt := range tests {
		got := tt.fields.IsExclude(tt.rPath)
		assert.Equal(t, tt.want, got, tt.name)
	}
}
