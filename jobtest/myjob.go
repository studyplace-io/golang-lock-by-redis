package jobtest

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	db2 "redis-practice/jobtest/db"
	"redis-practice/src"

	//"redis-practice/src"
	"time"
)

func Run() {
	job("job1")
	job("job2")
}

func job(jobname string) {
	c := cron.New(cron.WithSeconds())

	id, err := c.AddFunc("0/5 * * * * *", func() {
		defer func() {
			if e := recover(); e != nil {
				log.Println(jobname,"取锁失败", e)
			}
		}()

		lock := src.NewLockerWithTTL("job11", time.Second*5).Lock()
		defer lock.UnLock()
		time.Sleep(time.Second*2)
		db:= db2.DB.Exec("update testjob set v=v+1 where id=100")
		if db.Error!=nil{
			log.Println(db.Error)
		}else{
			log.Println(jobname,"任务执行完毕")
		}
	})

	if err!=nil{
		log.Fatal(err)
	}
	fmt.Printf("%s任务ID是:%d 启动\n",jobname,id)
	c.Start()

}
