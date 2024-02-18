package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendGetRequestWithParams() {
	baseURL := "https://is-gateway-test.corp.kuaishou.com/token/get"
	param1 := "d5b9f475-c9fc-4cde-9cab-c3b36c3035c0"
	param2 := "1189947"
	// 将参数绑定到 URL
	url := fmt.Sprintf("%s?appKey=%s", baseURL, param1)
	url2 := fmt.Sprintf("https://is-gateway-test.corp.kuaishou.com/kuberec/v1/open-api/environment/query-progress/%s", param2)

	// 创建 GET 请求
	req, err := http.NewRequest("GET", url, nil)
	req2, err2 := http.NewRequest("GET", url2, nil)
	// // 创建 POST 请求
	// requestBody := []byte(`{"key": "value"}`)
	// req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil && err2 != nil {
		fmt.Println("创建 GET 请求失败:", err, err2)
		return
	}
	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送 GET 请求失败:", err)
		return
	}
	// 处理响应
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	// 解码JSON数据到结构体
	var responseData JSONData
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		fmt.Println("解码JSON失败:", err)
		return
	}
	// 设置请求头
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("Authorization", "aB2XvR5wL9yOzQ8")
	req2.Header.Set("OpenAccessToken", responseData.Result.AccessToken)
	resp2, err2 := client.Do(req2)
	if err2 != nil {
		fmt.Println("发送 GET 请求失败:", err)
		return
	}
	body2, _ := ioutil.ReadAll(resp2.Body)
	defer resp2.Body.Close()
	fmt.Println("GET 请求响应:", string(body))
	fmt.Println("GET2 请求响应:", string(body2))

}

func MakeHttps() {
	//构建一个https请求
	req, err := http.NewRequest("GET", "https://www.baidu.com", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// req.Header("Accept", "application/json")
	// req.Header("Authorization", "ww")
	//发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
