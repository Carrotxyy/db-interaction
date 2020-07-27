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
	url := w.Config.WxAddr + "/interactive/saveperson?EnKey=" + enkey
	// 数据类型
	contentType := "application/json"

	// 发送请求
	bool := api.HttpPost(url,persons,contentType)
	if !bool {
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


// 同步 商汤机构
func (w *Work)Sense_Orginfo()error{
	var orginfos []*models.Orginfo
	where := models.Orginfo{Org_SenseMark: "0"}
	err,count := w.Repository.Get(where,&orginfos,"")
	if err != nil {
		fmt.Println("商汤同步：获取机构错误:",err)
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

	return nil
}