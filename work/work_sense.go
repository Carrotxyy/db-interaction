package work

import (
	"db-interaction/common/api"
	"db-interaction/models"
	"fmt"
)

/**
	商汤系统同步
 */

// 同步商汤机构
func (w *Work)Sense_Orginfo()error{
	var orginfos []*models.Orginfo
	// 获取商汤标志位不为 0的数据
	where := models.Orginfo{Org_SenseMark: "0"}

	err,count := w.Repository.Get(where,&orginfos,"")
	if err != nil {
		fmt.Println("商汤同步：本地获取机构错误:",err)
		return err
	}
	// 判断是否有新数据
	if count == 0 {
		fmt.Println("暂无新数据")
		return nil
	}

	for _, v := range orginfos {
		fmt.Println("机构数据：",v)
	}

	// 构建请求 路径  携带 加密钥匙
	url := w.SplicUrl("/cloudSync/sense/orginfos")
	// 数据类型
	contentType := "application/json"
	// 响应数据
	var res struct{
		Status bool `json:"status"`
	}
	// 发送请求
	err = api.HttpPost(url,orginfos,&res,contentType)
	if err != nil || !res.Status {
		fmt.Println("上传数据错误")
		return err
	}
	// 恢复标志位


	return nil
}