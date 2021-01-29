package conf

import (
	"testing"

	"github.com/micro-plat/hydra/conf/vars/cache"
	"github.com/micro-plat/hydra/conf/vars/cache/cacheredis"
	"github.com/micro-plat/lib4go/assert"
)

func TestCacheRedisNew(t *testing.T) {

	tests := []struct {
		name    string
		address string
		opts    []cacheredis.Option
		want    *cacheredis.Redis
	}{
		{name: "1. Conf-CacheRedisNew-测试新增-无option", opts: []cacheredis.Option{cacheredis.WithAddrs("192.168.5.79:6379")},
			want: &cacheredis.Redis{Cache: &cache.Cache{Proto: "redis"}, Addrs: []string{"192.168.5.79:6379"}, DbIndex: 0, DialTimeout: 0, ReadTimeout: 0, WriteTimeout: 0, PoolSize: 0}},
		{name: "2. Conf-CacheRedisNew-测试新增-WithDbIndex", opts: []cacheredis.Option{cacheredis.WithAddrs("192.168.5.79:6379"), cacheredis.WithDbIndex(2)},
			want: &cacheredis.Redis{Cache: &cache.Cache{Proto: "redis"}, Addrs: []string{"192.168.5.79:6379"}, DbIndex: 2, DialTimeout: 0, ReadTimeout: 0, WriteTimeout: 0, PoolSize: 0}},
		{name: "3. Conf-CacheRedisNew-测试新增-WithTimeout", opts: []cacheredis.Option{cacheredis.WithAddrs("192.168.5.79:6379"), cacheredis.WithDbIndex(2), cacheredis.WithTimeout(11, 22, 33)},
			want: &cacheredis.Redis{Cache: &cache.Cache{Proto: "redis"}, Addrs: []string{"192.168.5.79:6379"}, DbIndex: 2, DialTimeout: 11, ReadTimeout: 22, WriteTimeout: 33, PoolSize: 0}},
		{name: "4. Conf-CacheRedisNew-测试新增-WithPoolSize", opts: []cacheredis.Option{cacheredis.WithAddrs("192.168.5.79:6379"), cacheredis.WithDbIndex(2), cacheredis.WithTimeout(11, 22, 33), cacheredis.WithPoolSize(40)},
			want: &cacheredis.Redis{Cache: &cache.Cache{Proto: "redis"}, Addrs: []string{"192.168.5.79:6379"}, DbIndex: 2, DialTimeout: 11, ReadTimeout: 22, WriteTimeout: 33, PoolSize: 40}},
	}

	for _, tt := range tests {
		got := cacheredis.New("", tt.opts...)
		assert.Equal(t, "redis", got.Cache.Proto, tt.name+",Proto")
		assert.Equal(t, tt.want.Addrs, got.Addrs, tt.name+",Addrs")
		assert.Equal(t, tt.want.Password, got.Password, tt.name+",Password")
		assert.Equal(t, tt.want.DbIndex, got.DbIndex, tt.name+",DbIndex")
		assert.Equal(t, tt.want.DialTimeout, got.DialTimeout, tt.name+",DialTimeout")
		assert.Equal(t, tt.want.ReadTimeout, got.ReadTimeout, tt.name+",ReadTimeout")
		assert.Equal(t, tt.want.WriteTimeout, got.WriteTimeout, tt.name+",WriteTimeout")
		assert.Equal(t, tt.want.PoolSize, got.PoolSize, tt.name+",PoolSize")
	}
}
