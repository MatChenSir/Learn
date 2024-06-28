package test

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
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

	//wg.Add() 和 wg.Wait() 区别
	//是 sync.WaitGroup 结构体提供的两个方法，用于并发控制
	//wg.Add() 方法用于向 WaitGroup 中添加等待的计数器，而wg.Wait() 方法时，主 goroutine 会被阻塞，直到所有的等待计数器都被逐个减为零！！！！！

}

//并法和并行  ---->>>>两者的区别在于对cpu核的利有，并行往往多核处理，并发是单核或处理不足情况下一起执行的并由调度器处理分配使用cpu的一种情况
func ConcurrencyAndParallelism() {
	Concurrency() //并发执行，并由调度器分配，如果个人想优先执行某个goroutines,可通过sync.Mutex或者channel来辅助
	//Parallelism() //并行执行
}

//并发
func Concurrency() {
	fmt.Println("开发并发执行了")
	//并发
	runtime.GOMAXPROCS(1)
	wg.Add(2)
	go func() {
		defer wg.Done()
		// mu.Lock() //加入这两行将会使代码不并发，并单个goroutines的执行    ！！！！！！虽然会单个执行，但因为是goroutines，执行顺序会由调度器决定
		// defer mu.Unlock()
		for i := 0; i < 200; i++ {
			fmt.Printf("开始执行test1了,已经执行到了第%v 个了;\n", i)
			time.Sleep(time.Microsecond * 500) //模拟消耗的资源来测到 因为调度器设计来尽可能公平地分配 CPU 时间给所有 goroutines，如果资源太少可能看不到调度的过程
		}

	}()
	go func() {
		defer wg.Done()
		// mu.Lock() //加入这两行将会使代码不并发，并单个goroutines的执行   ！！！！！！虽然会单个执行，但因为是goroutines，执行顺序会由调度器决定
		// defer mu.Unlock()
		for i := 0; i < 200; i++ {
			fmt.Printf("开始执行test2了,已经执行到了第%v 个了;\n", i)
			time.Sleep(time.Microsecond * 500) //模拟消耗的资源来测到 因为调度器设计来尽可能公平地分配 CPU 时间给所有 goroutines，如果资源太少可能看不到调度的过程
		}

	}()
	wg.Wait()

}

//并行
func Parallelism() {
	//并发执行可以让他并行，并行一定是并发

	fmt.Println("开发并发执行了")
	//并发
	runtime.GOMAXPROCS(2)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 200; i++ {
			fmt.Printf("开始执行并行1了,已经执行到了第%v 个了;\n", i)
		}

	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 200; i++ {
			fmt.Printf("开始执行并行2了,已经执行到了第%v 个了;\n", i)
		}

	}()
	wg.Wait()
}

// func channel() {
// 	//通道（channel）是一种用于在goroutine之间进行通信和同步的机制, 但也可以通过make来声明，故我放在了数据结构里面来展示

// 	//非缓冲通道
// 	ch1 := make(chan int)
// 	// 	带缓冲的通道允许在通道中存储多个值，直到缓冲区被填满。
// 	// 发送数据到带缓冲通道不会阻塞，除非缓冲区已满。
// 	// 从带缓冲通道接收数据不会阻塞，除非缓冲区为空。
// 	// 当通道中缓冲区已满时，发送操作会阻塞；当通道中缓冲区为空时，接收操作会阻塞  !!!!!!!!!

// 	//缓冲通道
// 	ch2 := make(chan int, 10)
// 	// 	不带缓冲的通道在发送数据时会立即阻塞等待接收方接收，接收方在接收数据之前也会阻塞等待发送方发送。
// 	// 通道的发送和接收操作是同步的，数据直接从发送方传递到接收方。
// 	// 当通道中没有接收方时，发送操作会一直阻塞；当通道中没有发送方时，接收操作会一直阻塞。
// }
