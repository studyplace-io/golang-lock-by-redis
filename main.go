package main

import (
	"github.com/gin-gonic/gin"
	"redis-practice/src"
	"strconv"
	"time"
)

const TTL = time.Second * 10

func main() {
	r := gin.New()

	r.Use(func(c *gin.Context) {
		defer func() {
			if e := recover(); e != nil {
				c.AbortWithStatusJSON(400, gin.H{"message:":e})
			}
		}()
		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		a := 1
		//lock := src.NewLocker("tryLock1111")
		// 自定义过期时间，可能会发生一个情况，就是TTL设置3妙，但是请求需要5秒的情况，所以需要自动续期的功能。
		lock := src.NewLockerWithTTL("trylock1111", TTL)
		lock.Lock()
		defer lock.UnLock()
		if c.Query("t") != "" {
			// 范例一
			// 请求http://127.0.0.1:8080/?t=1 会报panic
			//panic("t")
			// 范例二
			time.Sleep(time.Second * 5)
			// 当路由到这里时，需要等待五秒才能再次获得锁
			// 如果不设置过期时间，服务当机后，会迟迟无法释放锁，导致之后重启时，也不能正常运行。
		}
		a++
		c.JSON(200, gin.H{"message:":"ok"+ strconv.Itoa(a)})
	})

	defer func() {
		r.Run(":8080")
	}()

}
