package setting

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	WxAddr      string `json:"wxaddr"`       // 微信服务器域名
	DbType      string `json:"dbtype"`       // 数据库类型
	DbUser      string `json:"dbuser"`       // 数据库账号
	DbPassword  string `json:"dbpassword"`   // 数据库密码
	DbIP        string `json:"dbip"`         // 数据库ip
	DbName      string `json:"dbname"`       // 数据库名
	TablePrefix string `json:"table-prefix"` // 数据库表扩展
	SyncTime    string `json:"sync-time"`    // 数据同步时间
}

// 加载配置文件
func LoadingConf() *Config {

	file, err := os.Open("./conf/conf.json")
	if err != nil {
		log.Panic("Open file failed:", err.Error())
		os.Exit(1)
	}
	// 关闭文件流
	defer file.Close()

	var config Config

	// 创建json解码器
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Panic("Decoder failed:", err.Error())
		os.Exit(1)
	} else {
		fmt.Println("Decoder success")
	}
	// 返回配置文件
	return &config
}
