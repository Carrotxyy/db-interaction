package work

import (
	"db-interaction/common/api"
	"db-interaction/models"
	"fmt"
)

/**
	微信访客同步任务
 */

// 上传 可接受预约 的业主数据
func (w *Work)VisitorUpload()error{

	var persons []*models.Personinfo
	// 查找修改过数据的业主,并且这些数据没有被同步到微信系统中
	where := models.Personinfo{
		Per_WXMark: "0",
	}
	// 获取数据
	err,count := w.Repository.Get(where,&persons,"Per_ID,Per_Name,Per_ContactTel,Per_Status,Per_AllowVisit")
	if err != nil {
		fmt.Println("获取用户数据错误:",err)
		return err
	}
	fmt.Println("数据数量:",count)
	if count == 0{
		fmt.Println("暂无新数据")
		return nil
	}
	// 构建请求 路径  携带 加密钥匙
	url := w.SplicUrl("/interactive/saveperson")
	// 数据类型
	contentType := "application/json"
	var res struct{
		Status bool `json:"status"`
	}
	// 发送请求
	err = api.HttpPost(url,persons,&res,contentType)
	if err != nil || !res.Status {
		fmt.Println("发送同步请求错误!")
	}else {
		// 上传同步成功后，需要将所有同步过数据的标志位设置为1，表示数据已经同步
		err = w.recovery(persons)
	}

	return err
}

// 恢复标志位
func (w *Work)recovery(persons []*models.Personinfo)error{
	var ids []int
	for _, v := range persons {
		// 获取Per_ID
		ids = append(ids,v.Per_ID)
	}
	return w.Repository.BatchUpdate(ids)
}


// 下载 所有访客数据
func (w *Work)VisitorLoad()error{

	// 构建url
	url := w.SplicUrl("/interactive/visitor")
	obj := struct {
		Data   []*models.Visitor `json:"data"`
		Status bool              `json:"status"`
	}{}
	// 获取数据
	err  := api.HttpGet(url,&obj)
	if err != nil {
		fmt.Println("拉取访客数据错误!")
		return err
	}
	if len(obj.Data) <= 0 {
		fmt.Println("暂无数据!")
		return nil
	}
	err = w.Repository.BatchSave(obj.Data)
	if err != nil {
		fmt.Println("保存访客数据错误:",err)
		return err
	}

	return nil
}

