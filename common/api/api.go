package api

import (
	"bytes"
	"db-interaction/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
func HttpPost(url string,data interface{},contentType string)bool{

	jsonStr, _ := json.Marshal(data)
	resp, err :=  http.Post(url, contentType,bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	var r struct{
		Status bool `json:"status"`
	}
	// 解析数据
	err = json.Unmarshal(result,&r)
	if err != nil {
		fmt.Println("解析响应数据错误!",err)
		return false
	}
	// 判断是否同步成功!
	if !r.Status  {
		fmt.Println("错误了")
		return false
	}
	return true
}

// 请求访客数据
func HttpGet(path string)([]*models.Visitor,string){
	var visitors []*models.Visitor
	var cursor string
	resp,err := http.Get(path)
	if err != nil {
		fmt.Println("请求访客数据错误:",err)
		return visitors,cursor
	}

	result, _ := ioutil.ReadAll(resp.Body)

	obj := struct {
		Data []*models.Visitor `json:"data"`	// 此处解析错误！！！！不知道为什么
		Cursor string `json:"cursor"`
	}{}
	fmt.Println(string(result))
	err = json.Unmarshal(result,&obj)
	if err != nil {
		fmt.Println("解析访客数据错误！",err)
	}
	visitors = obj.Data
	cursor = obj.Cursor
	return visitors,cursor
}

// 获取key
func GetKey(path string)interface{}{
	res,err :=http.Get(path)
	if err != nil {
		log.Panic("获取key:",err)
		return ""
	}
	defer res.Body.Close()
	body,_ := ioutil.ReadAll(res.Body)

	return string(body)
}
