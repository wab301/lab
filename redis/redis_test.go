package redis

import (
	"fmt"
	"testing"

	"gopkg.in/redis.v5"
)

var Redis *redis.Client

func init() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "192.168.2.3:6379",
		Password: "",
	})
	Redis.Ping().Err()

}

func BenchmarkRedis(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Redis.Keys(fmt.Sprintf("test*%d*", i))
	}
}
