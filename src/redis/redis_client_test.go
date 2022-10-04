package redis

import (
	"context"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRedis(t *testing.T) {
	Convey("test redis client", t, func() {
		ctx := context.Background()
		pong, err := RedisClient.Ping(ctx).Result()
		fmt.Println(pong, err)
		So(err,ShouldBeNil)
	})
}

func TestRedisSet(t *testing.T) {
	Convey("test redis set get", t, func() {
		ctx := context.Background()
		err := RedisClient.Set(ctx, "key", "value", 0).Err()
		if err != nil {
			panic(err)
		}

		val, err := RedisClient.Get(ctx, "key").Result()
		if err != nil {
			panic(err)
		}
		fmt.Println("key", val)
		So(err, ShouldBeNil)

		//val2, err := RedisClient.Get(ctx, "missing_key").Result()
		//if err == redis.Nil {
		//	fmt.Println("missing_key does not exist")
		//} else if err != nil {
		//	panic(err)
		//} else {
		//	fmt.Println("missing_key", val2)
		//}

	})
}
