package timing

import (
	"db-interaction/common/setting"
	"db-interaction/work"
	"fmt"
	cron "github.com/robfig/cron/v3"
	"strings"
)

func RunWork() {


	// 创建一个秒级的定时器
	crontab := cron.New(cron.WithSeconds())
	works := work.CreateWork()
	task := func() {

		fmt.Println("同步--上传业主数据 开始")

		err := works.Upload()
		if err != nil {
			fmt.Println("同步--上传业主出现错误:",err)
		}

		fmt.Println("同步--上传业主数据 结束")


		fmt.Println("同步--下载访客数据 开始")
		err = works.LoadVisitor()
		if err != nil {
			fmt.Println("同步--下载访客出现错误:",err)
		}
		fmt.Println("同步--下载访客数据 结束")
	}
	// 获取执行时间
	spec := getSpec()
	// 添加定时任务
	crontab.AddFunc(spec, task)
	// 启动定时器
	crontab.Start()



}

// 获取运行时间格式
func getSpec()string{
	var spec string
	// 获取配置
	config := setting.LoadingConf()
	// 获取同步时间 秒
	s := config.SyncTime

	arr := strings.Split(s,"-")

	switch arr[1] {
	case "s":
		spec = "*/" + arr[0] + " * * * * *"
	case "m":
		spec = "0 */" + arr[0] + " * * * *"
	case "h":
		spec = "0 0 */" + arr[0] + " * * *"
	default:
		// 默认15s一次同步
		spec = "*/15 * * * * *"
	}

	return spec
}