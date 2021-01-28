package conf

import (
	"testing"

	"github.com/micro-plat/hydra/conf/vars/queue"
	"github.com/micro-plat/hydra/conf/vars/queue/queueredis"

	"github.com/micro-plat/lib4go/assert"
)

func TestQueueRedisNew(t *testing.T) {

	tests := []struct {
		name    string
		address string
		opts    []queueredis.Option
		want    *queueredis.Redis
	}{
		{name: "1. Conf-QueueRedisNew-测试新增-无option", opts: []queueredis.Option{queueredis.WithAddrs("192.168.5.79:6379")},
			want: &queueredis.Redis{Queue: &queue.Queue{Proto: "redis"}, Addrs: []string{"192.168.5.79:6379"}, DbIndex: 0, DialTimeout: 10, ReadTimeout: 10, WriteTimeout: 10, PoolSize: 10}},
		{name: "2. Conf-QueueRedisNew-测试新增-WithDbIndex", opts: []queueredis.Option{queueredis.WithAddrs("192.168.5.79:6379"), queueredis.WithDbIndex(2)},
			want: &queueredis.Redis{Queue: &queue.Queue{Proto: "redis"}, Addrs: []string{"192.168.5.79:6379"}, DbIndex: 2, DialTimeout: 10, ReadTimeout: 10, WriteTimeout: 10, PoolSize: 10}},
		{name: "3. Conf-QueueRedisNew-测试新增-WithTimeout", opts: []queueredis.Option{queueredis.WithAddrs("192.168.5.79:6379"), queueredis.WithDbIndex(2), queueredis.WithTimeout(11, 22, 33)},
			want: &queueredis.Redis{Queue: &queue.Queue{Proto: "redis"}, Addrs: []string{"192.168.5.79:6379"}, DbIndex: 2, DialTimeout: 11, ReadTimeout: 22, WriteTimeout: 33, PoolSize: 10}},
		{name: "4. Conf-QueueRedisNew-测试新增-WithTimeout", opts: []queueredis.Option{queueredis.WithAddrs("192.168.5.79:6379"), queueredis.WithDbIndex(2), queueredis.WithTimeout(11, 22, 33), queueredis.WithPoolSize(40)},
			want: &queueredis.Redis{Queue: &queue.Queue{Proto: "redis"}, Addrs: []string{"192.168.5.79:6379"}, DbIndex: 2, DialTimeout: 11, ReadTimeout: 22, WriteTimeout: 33, PoolSize: 40}},
	}
	for _, tt := range tests {
		got := queueredis.New("", tt.opts...)
		assert.Equal(t, "redis", got.Queue.Proto, tt.name+",Proto")
		assert.Equal(t, tt.want.Addrs, got.Addrs, tt.name+",Addrs")
		assert.Equal(t, tt.want.Password, got.Password, tt.name+",Password")
		assert.Equal(t, tt.want.DbIndex, got.DbIndex, tt.name+",DbIndex")
		assert.Equal(t, tt.want.DialTimeout, got.DialTimeout, tt.name+",DialTimeout")
		assert.Equal(t, tt.want.ReadTimeout, got.ReadTimeout, tt.name+",ReadTimeout")
		assert.Equal(t, tt.want.WriteTimeout, got.WriteTimeout, tt.name+",WriteTimeout")
		assert.Equal(t, tt.want.PoolSize, got.PoolSize, tt.name+",PoolSize")
	}
}
