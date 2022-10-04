package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gopkg.in/ini.v1"
	"os"
	"strconv"
	"strings"
)

//RedisClient Redis缓存客户端单例
var (
	RedisClient *redis.Client
	RedisAddr  			string
	RedisPw    			string
	RedisDbName    		string
)




func init() {
	// 取得配置文件路径
	workdir, _ := os.Getwd()
	var str = []string{workdir, "/redis_config.ini"}
	path := strings.Join(str, "")
	file, err := ini.Load(path)
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadRedis(file)
	RedisInit()
}

//初始化redis链接
func RedisInit() {
	db, _ := strconv.ParseUint(RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		Password: RedisPw,
		DB:       int(db),
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	RedisClient = client
}

func LoadRedis(file *ini.File) {
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}