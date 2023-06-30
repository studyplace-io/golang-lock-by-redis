package pkg

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	redis2 "redis-practice/pkg/cache"
)

// Locker 锁对象
type Locker struct {
	key string
	// 过期时间
	expiration time.Duration
	unlock     bool
	incrScript *redis.Script
}

const incrLua = `
if cache.call('get', KEYS[1]) == ARGV[1] then
  return cache.call('expire', KEYS[1],ARGV[2]) 				
 else
   return '0' 					
end`

// NewLocker 返回一个分布式锁
func NewLocker(key string) *Locker {
	return &Locker{
		key:        key,
		expiration: time.Second * 30,
		incrScript: redis.NewScript(incrLua),
	}
}

// NewLockerWithTTL 设置过期时间的分布式锁
func NewLockerWithTTL(key string, expiration time.Duration) *Locker {
	if expiration.Seconds() <= 0 {
		panic("expiration err!")
	}

	return &Locker{
		key:        key,
		expiration: expiration,
		incrScript: redis.NewScript(incrLua),
	}
}

// Lock 从redis中使用SetNX获取锁
func (l *Locker) Lock() *Locker {
	ctx := context.Background()
	// 如果expiration设置为0，代表不过期，会导致死锁。
	boolRes := redis2.RedisClient.SetNX(ctx, l.key, "1", 0)
	if ok, err := boolRes.Result(); err != nil || !ok {
		panic(fmt.Sprintf("lock error with key: %s", l.key))
	}
	l.expandLockTime()
	// 返回锁对象，可以链式调用
	return l
}

// UnLock 释放redis锁
func (l *Locker) UnLock() {
	l.unlock = true
	redis2.RedisClient.Del(context.Background(), l.key)
}

// resetExpiration 重新设置过期时间
func (l *Locker) resetExpiration() {
	//ctx := context.Background()
	//cmd := redisClient.Expire(ctx, l.key, l.expiration)
	//fmt.Println("续期时间为：", cmd)
	cmd := l.incrScript.Run(context.Background(), redis2.RedisClient, []string{l.key}, 1, l.expiration.Seconds())
	v, err := cmd.Result()
	log.Printf("key=%s ,续期结果:%v,%v\n", l.key, err, v)

}

func (l *Locker) expandLockTime() {
	// 推荐的续期时间，经过2/3后 再次续期
	sleepTime := l.expiration.Seconds() * 2 / 3
	go func() {
		for {
			time.Sleep(time.Duration(sleepTime) * time.Second)
			if l.unlock {
				break
			}
			l.resetExpiration()
		}
	}()

}
