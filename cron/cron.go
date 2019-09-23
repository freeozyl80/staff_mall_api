package cron

import (
	"fmt"
	"log"
	"staff-mall-center/models/dao"
	"time"

	"github.com/robfig/cron"
)

func CronRun() {
	log.Println("Starting...")

	c := cron.New()
	c.AddFunc("0 0 0 *", func() {
		log.Println("Run models.ClearVaildOrder...")
		err := dao.ClearVaildOrder()
		fmt.Println(err)
	})

	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
