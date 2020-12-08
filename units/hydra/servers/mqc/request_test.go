package mqc

import (
	"testing"

	"github.com/micro-plat/hydra/components/queues/mq/redis"
	"github.com/micro-plat/hydra/conf/server/queue"
	"github.com/micro-plat/hydra/hydra/servers/mqc"
	"github.com/micro-plat/hydra/test/assert"
)

func TestNewRequest(t *testing.T) {
	tests := []struct {
		name      string
		queueName string
		service   string
		message   string
		hasData   bool
		errStr    string
	}{
		{name: "1. mqc-NewRequest-队列数据格式错误", queueName: "queue", service: "service", message: `message`, hasData: true, errStr: "队列queue中存放的数据不是有效的json:message invalid character 'm' looking for beginning of value"},
		{name: "2. mqc-NewRequest-添加队列数据", queueName: "queue", service: "service", message: `{"key":"value","__header__":{"Content-Type":"application/json"}}`, hasData: true},
	}
	for _, tt := range tests {
		gotR, err := mqc.NewRequest(queue.NewQueue(tt.queueName, tt.service), &redis.RedisMessage{Message: tt.message, HasData: tt.hasData})
		if tt.errStr != "" {
			assert.Equal(t, tt.errStr, err.Error(), tt.name)
			continue
		}
		assert.Equal(t, tt.queueName, gotR.GetName(), tt.name)
		assert.Equal(t, tt.service, gotR.GetService(), tt.name)
		assert.Equal(t, mqc.DefMethod, gotR.GetMethod(), tt.name)

		assert.Equal(t, tt.message, gotR.GetForm()["__body__"], tt.name)
	}
}
