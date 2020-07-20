package work

import (
	"db-interaction/common/api"
	"db-interaction/common/setting"
	"db-interaction/models"
	"db-interaction/repository"
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
	url := w.Config.WxAddr + "/interactive/saveperson"
	contentType := "application/json"
	bool := api.HttpPost(url,persons,contentType)
	if bool {
		log.Panic("同步接受预约人员错误: ",err)
	}
	return nil
}