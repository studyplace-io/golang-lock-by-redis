# golang 基于redis实现简单分布式锁
![](https://github.com/googs1025/golang-lock-by-redis/blob/main/image/%E6%B5%81%E7%A8%8B%E5%9B%BE%20(1).jpg?ram=true)
## 本项目采用golang语言操作redis，实现简单分布式锁。
## 其中采用两种演示方式：1.gin请求 2.mysql写数据
### 项目依赖
目前项目依赖Redis组件，需要先使用docker启动redis，并监听6379端口。
未来会支持ini文件配置，让使用者配置更便捷
```
拉取镜像
1. $ docker pull redis:latest
执行镜像
2. $ docker run -itd --name redis-test -p 6379:6379 redis
进入docker中测试镜像功能
3. $ docker exec -it redis-test /bin/bash
```
### 项目目录
```bigquery
.
├── README.md
├── go.mod
├── go.sum
├── job.go
├── jobtest
│   ├── db
│   │   ├── config.go
│   │   ├── config.ini
│   │   └── db_init.go
│   └── myjob.go
├── main.go
└── src
    ├── locker.go
    └── redisinit.go

```
**src**:放置redis锁实现的主要方法

**jobtest**: 放置连接mysql配置与执行操作的方法

**main.go**: gin请求事例的main入口

**job.go**: mysql请求的main入口

## 启动项目前操作
`go mod tidy`

### gin请求演示方式
`
默认开启localhost:8080端口
`

`在项目根目录执行  go run main.go`



### mysql写数据演示方式
`需要根据自己mysql情况，先配置mysql，在jobtest/db/db_init.go当中`

`需要注意库与表的建立`

`create table if not exists testjob ( id INT UNSIGNED , v INT UNSIGNED);`

`insert into testjob (id, v) values (100, 1);`

`启动测试main程序 go run job.go，默认启动两个goroutine同时写表中的值。`

