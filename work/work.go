package work

import (
	"db-interaction/common/api"
	"db-interaction/common/setting"
	"db-interaction/models"
	"db-interaction/repository"
	"fmt"
	"log"
)

type Work struct {
	Repository *repository.BaseRepository
	Config *setting.Config
}

// 创建业务
func CreateWork()*Work{
	r := repository.NewRepository()
	config := setting.LoadingConf()
	return &Work{Repository: r,Config: config}
}
// 上传 可接受预约 的业主数据
func (w *Work)Upload()error{

	// 获取 key
	key := api.GetKey(w.Config.WxAddr +"/interactive/key").(string)
	// 加密 key
	enkey := api.Encryption(key)

	var persons []*models.Personinfo
	// 查找拥有 接受预约 权限的人员
	where := models.Personinfo{
		Per_Status:"0",
		Per_AllowVisit:"0",
	}
	// 获取数据
	err := w.Repository.Get(where,&persons,"Per_Name,Per_ContactTel")
	if err != nil {
		log.Panic("获取用户数据错误错误:",err)
		return err
	}
	// 构建请求 路径  携带 加密钥匙
	url := w.Config.WxAddr + "/interactive/saveperson?EnKey=" + enkey
	// 数据类型
	contentType := "application/json"

	// 发送请求
	bool := api.HttpPost(url,persons,contentType)
	if !bool {
		fmt.Println("同步接受预约人员错误: ",err)
	}
	return nil
}

// 下载 所有访客数据
func (w *Work)LoadVisitor()error{

	err := w.Repository.TruncateTable("go_visitor")
	if err != nil {
		fmt.Println("删除表数据错误:",err)
	}
	// 获取 key
	key := api.GetKey(w.Config.WxAddr +"/interactive/key").(string)
	// 加密 key
	enkey := api.Encryption(key)
	cursor := "0"
	for {

		// 构建url
		url := w.Config.WxAddr + "/interactive/visitor?EnKey=" + enkey + "&cursor=" + cursor
		// 获取数据
		visitors , cursor := api.HttpGet(url)
		err := w.Repository.BatchSave(visitors)
		if err != nil {
			fmt.Println("保存访客数据错误:",err)
		}
		// 判断访客数据是否已经下载完毕
		if cursor == "0" {
			break
		}
	}
	fmt.Println("下载完成！")
	return nil
}

