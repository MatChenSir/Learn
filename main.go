package main

import (
	"context"
	"fmt"
	"main/test"
	"net"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func main() {
	//defer test.RecoverPainc() //一定得放第一行，否则任意一行panic就不会执行deffer了

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
	//test.GetClosePackage()

	//并发和并行
	//test.ConcurrencyAndParallelism()

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

	//timeout()
	//test.TestGrountine()

	// test.CheckStruct()
	// time.Sleep(time.Second * 3)

	//算法题目相关

	// nums := []int{7, 8, 9, 0, 2, 0, 5, 11}
	// test.MoveZeros(nums)

	//strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	// strs := []string{"eat", "tea"}
	// tt := test.GroupAnagrams(strs)
	// fmt.Println(tt)

	//test.Worker()
	//test.Worker2()

	//	numbers := []int{3, 7, 2, 8, 34, 221, 85, 94, 3, 9, 90}
	test.GetTargetNums(221)

	fmt.Println(".............................................\n" +
		"     佛祖镇楼                  BUG辟易")

	i := 7
	b := &i
	i = 8
	fmt.Println(*b)
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

func TestWiths(t *testing.T) {

}

func worker(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker")
		//time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
	//wg.Done()
}

func timeout() {
	ctx, cancel := context.WithCancel(context.Background())
	//wg.Add(1)
	go worker(ctx)

	cancel() // 通知子goroutine结束
	//wg.Wait()
	fmt.Println("over")
}
