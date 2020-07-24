package work

import (
	"db-interaction/common/api"
	"db-interaction/common/setting"
	"db-interaction/models"
	"db-interaction/repository"
	"fmt"
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

	// 获取key
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
		fmt.Println("获取用户数据错误:",err)
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

	// 获取 key
	key := api.GetKey(w.Config.WxAddr +"/interactive/key").(string)
	// 加密 key
	enkey := api.Encryption(key)

	// 构建url
	url := w.Config.WxAddr + "/interactive/visitor?EnKey=" + enkey
	// 获取数据
	visitors,err  := api.HttpGet(url)
	if err != nil {
		fmt.Println("拉取访客数据错误!")
		return err
	}
	if len(visitors) <= 0 {
		fmt.Println("暂无数据!")
		return nil
	}
	err = w.Repository.BatchSave(visitors)
	if err != nil {
		fmt.Println("保存访客数据错误:",err)
		return err
	}

	return nil
}

