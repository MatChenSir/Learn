package test

import (
	"fmt"
	"runtime"
	"sync"
)

func TestPrint() {

	// 创建一个无缓冲的通道，用于同步两个协程
	syncCh := make(chan struct{})

	// 初始化syncCh，使第一个协程可以开始工作
	syncCh <- struct{}{}

	// 使用 WaitGroup 来等待两个协程完成
	var wg sync.WaitGroup
	wg.Add(2)

	// 启动打印字母的协程
	go printLetters(syncCh, &wg)

	// 启动打印数字的协程
	go printNumbers(syncCh, &wg)

	// 等待两个协程完成
	wg.Wait()
}

// printLetters 协程用于打印字母a到z
func printLetters(syncCh chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 'a'; i <= 'z'; i++ {
		<-syncCh // 等待信号
		fmt.Printf("%c ", i)
		runtime.Gosched()    // 让出CPU时间片，允许其他协程运行
		syncCh <- struct{}{} // 发送信号，允许另一个协程运行
	}
}

// printNumbers 协程用于打印数字0到9
func printNumbers(syncCh chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		<-syncCh // 等待信号
		fmt.Printf("%d ", i)
		runtime.Gosched()    // 让出CPU时间片，允许其他协程运行
		syncCh <- struct{}{} // 发送信号，允许另一个协程运行
	}
}
