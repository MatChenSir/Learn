package main

import (
	"fmt"
	"main/test"
	"net"
)

func main() {
	defer test.RecoverPainc() //一定得放第一行，否则任意一行panic就不会执行deffer了

	//路径相关
	//test.Test()

	//https请求相关
	//test.MakeHttps()

	//数据结构定义相关
	//test.DataStruct()

	//数据类型
	//test.BasicType()

	//interface转string
	//test.ConvertInterfaceToString()

	//将函数作为参数传递
	//test.FunctionAsParam()

	//闭包
	test.GetClosePackage()

	// http.HandleFunc("/", handler.Handler)
	// getLocahost()
	// err := http.ListenAndServe(":8080", nil)

	// if err != nil {
	// 	log.Fatal(err) // 打印错误信息并退出程序
	// 	return
	// }
	//test.SendGetRequestWithParams()

	//bindData()

	//getLocahost()

	//记录日志
	//test.CachedLog()

	//佛祖保佑
	//fmt.Println(test.StrConsole)
	//test.ReadTxt()

	fmt.Println(".............................................\n" +
		"     佛祖镇楼                  BUG辟易")
}

// 获得本地的IP地址
func getLocahost() {
	host := ""
	addrs, _ := net.InterfaceAddrs()
	// 获得本地的IP地址
	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			host = ipNet.IP.String()

		}
	}
	port := "8080"
	url := fmt.Sprintf("http://%s:%s", host, port)
	fmt.Println(url)
}

//绑定函数调用
func bindData() {
	//绑定函数调用
	model := &test.UserRepository{}
	model.Name = "test"

	//绑定函数调用一 实例一个函数进行调用
	user1, _ := test.IUserRepository.InsertAdUserInfoWithDefaults(model, "user")
	//或者
	//user1, _ := model.InsertAdUserInfoWithDefaults("user")

	//绑定函数调用二 通过内在的方法已经声明的进行调用
	user2, _ := test.GetUserRepository().InsertAdUserInfoWithDefaults("user")

	////绑定函数调用三 通过内在的方法实例化，和实例函数一样
	test.SingletonUserRepository = &test.UserRepository{Name: "bbt"}
	user3, _ := test.SingletonUserRepository.InsertAdUserInfoWithDefaults("user")
	println("user:" + user1.Username + "  user2:" + user2.Username + "  user3:" + user3.Username)
}
