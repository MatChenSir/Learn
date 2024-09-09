package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

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
