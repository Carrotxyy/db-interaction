package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
func HttpPost(url string, data interface{}, out interface{}, contentType string) error {

	jsonStr, _ := json.Marshal(data)
	resp, err := http.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("响应数据:", string(result))
	// 解析数据
	err = json.Unmarshal(result, out)
	if err != nil {
		fmt.Println("解析响应数据错误!", err)
		return err
	}

	return nil
}

// 请求数据
func HttpGet(path string, out interface{}) error {
	resp, err := http.Get(path)
	if err != nil {
		fmt.Println("请求url错误:", err)
		return err
	}

	result, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(result))
	// json 解析
	err = json.Unmarshal(result, out)
	if err != nil {
		fmt.Println("解析响应数据错误！", err)
		return err
	}
	return nil
}

// 获取key
func GetKey(path string) interface{} {
	res, err := http.Get(path)
	if err != nil {
		fmt.Println("获取key:", err)
		return ""
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}
