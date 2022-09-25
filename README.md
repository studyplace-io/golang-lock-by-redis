# golang 基于redis实现简单分布式锁
## 本项目采用golang语言操作redis，实现简单分布式锁。
## 其中采用两种演示方式：1.gin请求 2.mysql写数据

### 项目目录
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

