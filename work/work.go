package work

import (
	"db-interaction/common/api"
	"db-interaction/common/setting"
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

// 创建访问路由
// @path : 具体的路由 例: http://xyz.szlimaiyun.cn/cloudSync/test   path = /cloudSync/test
func (w *Work)SplicUrl(path string)string{

	// 获取 key
	key := api.GetKey(w.Config.WxAddr +"/interactive/key").(string)
	// 加密 key
	enkey := api.Encryption(key)

	// 构建url
	url := fmt.Sprintf("%s%s?EnKey=%s",w.Config.WxAddr,path,enkey)
	return url
}