package test

import (
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"
)

//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//请注意 ，必须在不复制数组的情况下原地对数组进行操作。
func MoveZeros(numbers []int) {
	//计算数组长度
	length := len(numbers)
	j := 0

	for i := 0; i < length; i++ {
		//建立一个循环，
		if numbers[i] != 0 { //首先找到不为0的数据
			//交换位置给到numbers[j],相当于一次洗牌
			numbers[j] = numbers[i]

			if j != i {
				//i完成了交接使命，可以替换数据了
				numbers[i] = 0
			}
			j++

		}

	}
	fmt.Println(numbers)

}

//给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
//字母异位词 是由重新排列源单词的所有字母得到的一个新单词。
func GroupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)

	for _, str := range strs {
		// 将字符串转换为字符数组并排序
		characters := strings.Split(str, "")
		sort.Strings(characters)
		sortedStr := strings.Join(characters, "") //思路：所有都排序组合好了，肯定就能对上了，管他什么顺序

		// 将排序后的字符串作为键，原始字符串作为值加入哈希表
		if _, ok := anagrams[sortedStr]; !ok {
			anagrams[sortedStr] = []string{str}
		} else {
			anagrams[sortedStr] = append(anagrams[sortedStr], str)
		}
	}

	// 构建结果列表
	result := make([][]string, 0, len(anagrams))
	for _, v := range anagrams {
		result = append(result, v)
	}

	return result

}

// 1. 使用sync.waitGroup实现如下要求
// 	  1) 开启三个协程，1 2 3
// 	  2) 三个协程分别等待1 2 3s，结束后打印出本协程的次序
// 	  3) 三个协程完成之后，主线程结束
func Worker() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Second * 3)
			fmt.Printf("协程第%v开始\n", i)

		}(i)

	}
	wg.Wait()
}

// 2. 使用channel实现以下结果，
// 	  1) 开启2个协程 1 2
//     2) 2个协程分别等待1 2s，等待结束后，分别向channel中塞入整数1 2
//     3) 主协程一直等待，直到从channel获取两个协程的结束信号，并打印协程传递过来的数字
//     4) 主流程结束

func Worker2() {
	ch := make(chan int, 2) // 创建一个带缓冲区大小为2的通道

	// 开启协程1
	go func(id int, duration time.Duration) {
		time.Sleep(duration * time.Second)
		fmt.Printf("协程%d等待结束\n", id)
		ch <- id
	}(1, 1)
	// 开启协程2
	go func(id int, duration time.Duration) {
		time.Sleep(duration * time.Second)
		fmt.Printf("协程%d等待结束\n", id)
		ch <- id
	}(2, 2)

	// 接收两个协程的结束信号
	count := 0
	for count < 2 {
		result := <-ch
		fmt.Printf("主协程收到协程%d的结束信号: %d\n", result, result)
		count++
	}

	close(ch) // 关闭通道

	fmt.Println("主协程结束")
}

/*4. 有一个接口interface a, a 有一个方法SetA(),请写出一个interface a的实现
 */

// 定义接口 a
type a interface {
	SetA(value int)
}

// 定义实现接口 a 的结构体
type myStruct struct {
	aValue int
}

// 实现接口方法 SetA
func (m *myStruct) SetA(value int) {
	m.aValue = value
}

func Woker4() {
	// 创建结构体实例
	obj := &myStruct{}

	// 调用实现的方法
	obj.SetA(42)

	// 输出结构体的值
	fmt.Println("Value set:", obj.aValue)
}
