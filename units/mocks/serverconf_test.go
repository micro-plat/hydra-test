package mocks

import (
	"testing"

	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/test/assert"
)

func TestGetConf(t *testing.T) {
	conf := NewConfBy("hydra_serverconf_test", "serverconf") //构建对象

	conf.API(":8081", api.WithHeaderReadTimeout(30), api.WithTimeout(31, 32))

	server := conf.GetAPIConf() //获取配置
	assert.Equal(t, server.GetServerConf().GetMainConf().GetString("address"), ":8081", "端口一致性检查")
	assert.Equal(t, server.GetServerConf().GetMainConf().GetInt("rhTimeout"), 30, "rhTimeout")
	assert.Equal(t, server.GetServerConf().GetMainConf().GetInt("rTimeout"), 31, "rTimeout")
	assert.Equal(t, server.GetServerConf().GetMainConf().GetInt("wTimeout"), 32, "wTimeout")
	assert.Equal(t, server.GetServerConf().GetServerPath(), "/hydra_serverconf_test/apiserver/api/serverconf/conf", "地址")
}
func TestRouters(t *testing.T) {
	conf := NewConfBy("test", "testrouters") //构建对象

	// conf.API(":8082")

	conf.Service.API.Add("/abc", "/abc", []string{"GET"})

	server := conf.GetAPIConf()

	rconf, err := server.GetRouterConf()

	assert.Equal(t, nil, err, "TestRouters")
	assert.Equal(t, 1, len(rconf.Routers), "TestRouters")

}
