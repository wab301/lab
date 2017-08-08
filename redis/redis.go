package redis

import (
	"fmt"
	"time"
)

// var Redis *redis.Client

// func init() {
// 	Redis = redis.NewClient(&redis.Options{
// 		Addr:     "192.168.2.3:6379",
// 		Password: "",
// 	})
// 	Redis.Ping().Err()
// }

func test() {
	t1 := time.Now().UnixNano()
	for i := 0; i < 1000000; i++ {
		Redis.Set(fmt.Sprintf("test%d", i), i, time.Hour*24)
	}
	t2 := time.Now().UnixNano()

	result := Redis.Keys("test*")
	t3 := time.Now().UnixNano()
	fmt.Println("=======", len(result.Val()))
	// for _, v := range result.Val() {
	// 	fmt.Println("====", v)
	// }
	Redis.Del(result.Val()...)
	t4 := time.Now().UnixNano()
	fmt.Println("=========333", t2-t1, t3-t2, t4-t3)
}
