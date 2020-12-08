package mqc

import (
	"testing"
	"time"

	"github.com/micro-plat/hydra/conf/server/queue"
	"github.com/micro-plat/hydra/hydra/servers/mqc"
	"github.com/micro-plat/hydra/test/assert"
)

func TestNewProcessor(t *testing.T) {
	tests := []struct {
		name    string
		proto   string
		confRaw string
		wantP   *mqc.Processor
		wantErr string
	}{
		{name: "1. mqc-NewProcessor-协议错误", proto: "proto", confRaw: "{}", wantErr: "构建mqc服务失败(proto:proto,raw:{}) mqc: 未知的协议类型 proto"},
		{name: "2. mqc-NewProcessor-协议配置正确", proto: "redis", confRaw: `{"proto":"redis","addrs":["192.168.5.79:6379"]}`},
	}
	for _, tt := range tests {
		gotP, err := mqc.NewProcessor(tt.proto, tt.confRaw)
		if tt.wantErr != "" {
			assert.Equal(t, tt.wantErr, err.Error(), tt.name)
			continue
		}
		//4 : 中间件的个数
		assert.Equal(t, true, len(gotP.Engine.Handlers) >= 4, tt.name)
	}

}

func TestProcessor_Add(t *testing.T) {
	tests := []struct {
		name       string
		queueNames []string
		queues     []*queue.Queue
		wantErr    string
	}{
		{name: "1. mqc-ProcessorAdd-添加消息队列", queues: []*queue.Queue{queue.NewQueue("queue1", "services1"), queue.NewQueue("queue2", "services2")}},
		{name: "2. mqc-ProcessorAdd-再次添加消息队列", queues: []*queue.Queue{queue.NewQueue("queue1", "services1"), queue.NewQueue("queue3", "services3")}},
	}
	s, _ := mqc.NewProcessor("redis", `{"proto":"redis","addrs":["192.168.5.79:6379"]}`)
	for _, tt := range tests {
		err := s.Add(tt.queues...)
		if tt.wantErr != "" {
			assert.Equal(t, tt.wantErr, err.Error(), tt.name)
			continue
		}
		assert.Equal(t, nil, err, tt.name)
		for _, v := range tt.queues {
			assert.Equal(t, s.QueueItems()[v.Queue], v, tt.name)
		}
	}
}

func TestProcessor_Remove(t *testing.T) {
	s, _ := mqc.NewProcessor("redis", `{"proto":"redis","addrs":["192.168.5.79:6379"]}`)
	//添加消息队列
	queues := []*queue.Queue{queue.NewQueue("queue1", "services1"), queue.NewQueue("queue2", "services2")}
	err := s.Add(queues...)
	assert.Equal(t, nil, err, "Add")
	//移除消息队列
	l := len(queues)
	for _, v := range queues {
		err := s.Remove(v)
		assert.Equal(t, nil, err, "Remove")
		l--
		assert.Equal(t, len(s.QueueItems()), l, "Remove")
	}

}

func TestProcessor_Resume(t *testing.T) {

	s, _ := mqc.NewProcessor("redis", `{"proto":"redis","addrs":["192.168.5.79:6379"]}`)
	//添加消息队列
	queues := []*queue.Queue{queue.NewQueue("queue1", "services1"), queue.NewQueue("queue2", "services2")}
	err := s.Add(queues...)
	assert.Equal(t, nil, err, "Add")

	//暂停
	got, err := s.Pause()
	assert.Equal(t, nil, err, "Pause")
	assert.Equal(t, true, got, "Pause")

	//暂停
	got, err = s.Pause()
	assert.Equal(t, nil, err, "Pause2")
	assert.Equal(t, false, got, "Pause2")

	//重启
	got, err = s.Resume()
	assert.Equal(t, nil, err, "Resume")
	assert.Equal(t, true, got, "Resume")
	time.Sleep(time.Second)
	got, err = s.Resume()
	assert.Equal(t, nil, err, "Resume2")
	assert.Equal(t, false, got, "Resume2")

	assert.Equal(t, len(queues), len(s.QueueItems()), "Resume2")

	//再次添加
	addQueues := []*queue.Queue{queue.NewQueue("queue3", "services3"), queue.NewQueue("queue4", "services4")}
	err = s.Add(addQueues...)
	assert.Equal(t, nil, err, "Add")
	assert.Equal(t, len(queues)+len(addQueues), len(s.QueueItems()), "Resume2")

	//暂停
	got, err = s.Pause()
	assert.Equal(t, nil, err, "Pause3")
	assert.Equal(t, true, got, "Pause3")

}

func TestProcessor_Close(t *testing.T) {
	s, _ := mqc.NewProcessor("redis", `{"proto":"redis","addrs":["192.168.5.79:6379"]}`)
	//添加消息队列
	queues := []*queue.Queue{queue.NewQueue("queue1", "services1"), queue.NewQueue("queue2", "services2")}
	err := s.Add(queues...)
	assert.Equal(t, nil, err, "Add")
	s.Close()
	assert.Equal(t, 0, len(s.QueueItems()), "Close")
	assert.Equal(t, true, s.Done(), "Close")
}
