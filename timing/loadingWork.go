package timing

import (
	"db-interaction/work"
	"fmt"
	cron "github.com/robfig/cron/v3"
	)

func RunWork() {

	crontab := cron.New()
	task := func() {
		fmt.Println("hello world")
	}
	// 添加定时任务, * * * * * 是 crontab,表示每分钟执行一次
	crontab.AddFunc("1 * * * *", task)
	works := work.CreateWork()
	works.Upload()
	works.LoadVisitor()
	// 启动定时器
	crontab.Start()
}