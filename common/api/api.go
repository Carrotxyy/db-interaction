package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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
	// 判断是否同步成功!
	if string(result) == "error" {
		return false
	}
	return true
}