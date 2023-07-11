package jobtest

import (
	"fmt"
	"log"
	"time"

	db2 "github.com/practice/redis-practice/jobtest/db"
	"github.com/practice/redis-practice/pkg"
	"github.com/robfig/cron/v3"
)

func Run() {
	job("job1")
	job("job2")
}

func job(jobName string) {
	c := cron.New(cron.WithSeconds())

	id, err := c.AddFunc("0/5 * * * * *", func() {
		defer func() {
			if e := recover(); e != nil {
				log.Println(jobName, "取锁失败", e)
			}
		}()

		lock := pkg.NewLockerWithTTL("job11", time.Second*5).Lock()
		defer lock.UnLock()
		time.Sleep(time.Second * 2)
		db := db2.DB.Exec("update testjob set v=v+1 where id=100")
		if db.Error != nil {
			log.Println(db.Error)
		} else {
			log.Println(jobName, "任务执行完毕")
		}
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s任务ID是:%d 启动\n", jobName, id)
	c.Start()

}
