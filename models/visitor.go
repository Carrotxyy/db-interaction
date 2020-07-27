package models

type Visitor struct {
	Name      string `json:"name"`      // 业主姓名
	Number    string `json:"number"`    // 业主电话
	UName     string `json:"uname"`     // 访客姓名
	UNumber   string `json:"unumber"`   // 访客电话
	IdNum     string `json:"idnum"`     // 访客身份证
	IdType    string `json:"idtype"`    // 访客 证件类型
	StartTime string `json:"starttime"` // 访问开始时间
	EndTime   string `json:"endtime"`   // 访问结束时间
	Message   string `json:"message"`   // 留言
	IsAcc     bool   `json:"isAcc"`     // 业主是否操作
	State     bool   `json:"state"`     // 预约结果
	FileName  string `json:"filename"`  // 头像文件名
	Per_ID    int    `json:"Per_ID"`    // 业主id
}
