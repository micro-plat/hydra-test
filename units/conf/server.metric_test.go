package conf

import (
	"testing"

	"github.com/micro-plat/hydra/conf"
	"github.com/micro-plat/hydra/test/assert"
	"github.com/micro-plat/hydra/test/mocks"

	"github.com/micro-plat/hydra/conf/server/metric"
)

func TestMetricNew(t *testing.T) {
	type args struct {
		host string
		db   string
		cron string
		opts []metric.Option
	}
	tests := []struct {
		name string
		args args
		want *metric.Metric
	}{
		{name: "1. Conf-MetricNew-设置空对象", args: args{host: "", db: "", cron: ""}, want: &metric.Metric{DataBase: "", Cron: ""}},
		{name: "2. Conf-MetricNew-只设置ip对象", args: args{host: "192.168.0.101", db: "", cron: ""}, want: &metric.Metric{Host: "192.168.0.101", DataBase: "", Cron: ""}},
		{name: "3. Conf-MetricNew-设置ip+prot对象", args: args{host: "192.168.0.101:8090", db: "", cron: ""}, want: &metric.Metric{Host: "192.168.0.101:8090", DataBase: "", Cron: ""}},
		{name: "4. Conf-MetricNew-设置disable对象", args: args{host: "192.168.0.101:8090", db: "", cron: "", opts: []metric.Option{metric.WithDisable()}}, want: &metric.Metric{Host: "192.168.0.101:8090", DataBase: "", Cron: "", Disable: true}},
		{name: "5. Conf-MetricNew-设置enable对象", args: args{host: "192.168.0.101:8090", db: "", cron: "", opts: []metric.Option{metric.WithEnable()}}, want: &metric.Metric{Host: "192.168.0.101:8090", DataBase: "", Cron: "", Disable: false}},
		{name: "6. Conf-MetricNew-设置全量对象", args: args{host: "192.168.0.101:8090", db: "1", cron: "cron", opts: []metric.Option{metric.WithEnable(), metric.WithUPName("upnem", "1223456")}},
			want: &metric.Metric{Host: "192.168.0.101:8090", DataBase: "1", Cron: "cron", Disable: false, UserName: "upnem", Password: "1223456"}},
	}
	for _, tt := range tests {
		got := metric.New(tt.args.host, tt.args.db, tt.args.cron, tt.args.opts...)
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestMetricGetConf(t *testing.T) {
	type test struct {
		name       string
		cnf        conf.IServerConf
		want       *metric.Metric
		wantErr    bool
		wantErrStr string
	}

	conf := mocks.NewConfBy("hydra", "graytest")
	confB := conf.API(":8090")
	test1 := test{name: "1. Conf-MetricGetConf-metric节点不存在", cnf: conf.GetAPIConf().GetServerConf(), want: &metric.Metric{Disable: true}, wantErr: false}
	limiterObj, err := metric.GetConf(test1.cnf)
	assert.Equal(t, test1.wantErr, (err != nil), test1.name)
	assert.Equal(t, test1.want, limiterObj, test1.name)

	confB.Metric("", "", "", metric.WithDisable())
	test2 := test{name: "2. Conf-MetricGetConf-metric节点存在,数据错误", cnf: conf.GetAPIConf().GetServerConf(), want: nil, wantErr: true, wantErrStr: "metric配置数据有误"}
	limiterObj, err = metric.GetConf(test2.cnf)
	assert.Equal(t, test2.wantErr, (err != nil), test2.name+",err")
	assert.Equal(t, test2.wantErrStr, err.Error()[:len(test2.wantErrStr)], test2.name+",err1")
	assert.Equal(t, test2.want, limiterObj, test2.name+",obj")

	confB.Metric("http://192.168.0.101", "1", "cron", metric.WithDisable(), metric.WithUPName("upnem", "1223456"))
	test3 := test{name: "3. Conf-MetricGetConf-metric节点存在,正确节点", cnf: conf.GetAPIConf().GetServerConf(),
		want: metric.New("http://192.168.0.101", "1", "cron", metric.WithDisable(), metric.WithUPName("upnem", "1223456")), wantErr: false}
	limiterObj, err = metric.GetConf(test3.cnf)
	assert.Equal(t, test3.wantErr, (err != nil), test3.name+",err")
	assert.Equal(t, test3.want, limiterObj, test3.name+",obj")

}
