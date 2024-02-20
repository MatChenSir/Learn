package test

import (
	"fmt"
	"strconv"
	"sync"
)

func FunctionAsParam() {
	//方法作为参数
	Test2(c)
}

func c(num int) {
	num = 4 + num
	fmt.Println("num:" + strconv.Itoa(num))
}

func Test2(test func(in int)) {
	fmt.Println("test")
	test(3) //相当于一个闭包，但如果这里不调用将不会走到c方法
}

var mu sync.Mutex
var wg sync.WaitGroup

//闭包
func closePackage() func() int {
	//闭包（Closure）是指一个函数值捕获并绑定了其周围的环境变量的函数
	count := 0
	return func() int {
		mu.Lock()
		defer mu.Unlock() //加入资源锁
		count++
		fmt.Printf("count: %v\n", count)
		return count
	}
}

//调用闭包
func GetClosePackage() {
	closePa := closePackage()
	//在 函数中，我们先调用 outerFunc 函数获取一个闭包函数 increment，然后多次调用 increment 函数，每次调用递增的计数值都会被保留并累加。
	for i := 0; i < 10; i++ {
		//要等待 1 个 goroutine 的完成。然后通过一个循环启动了 3 个 goroutine，并在每个 goroutine 中调用 wg.Done() 来表示完成了一次任务
		wg.Add(1)   //这里的数字就是决定等待几个同时执行，比如公交车可以坐10个人，但我决定只要到了3个就开车不等了,这里如果是3，下面就会一直陷入循环等待中，但不影响主进程
		go func() { // 可以解读为go test(), 因为没有取名直接用func替换，也是闭包的一种
			defer wg.Done()
			fmt.Printf("getClosePackage: %v\n", closePa()) //多次调用被绑定函数，简单理解就是这个函数永远在执行，但又不影响外在的方法，静态存在
		}()

	}
	//wg.Wait() //等待上一个wg done后再执行下一次，这里的意识就是，先候着
	wg.Wait() // 等待所有的 goroutine 完成，如果没有，可能就跳过上面的goroutines直接执行下面的了

	fmt.Println("All goroutines completed.")
}
