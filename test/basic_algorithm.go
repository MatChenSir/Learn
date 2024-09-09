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

//二分发查询一组数组的指定值
func GetTargetNums(target int) {
	numbers := []int{3, 7, 2, 8, 34, 221, 85, 94, 3, 9, 90}

	//首先给数组排序转化成有序数组
	sort.Ints(numbers)

	//取得最大值和最小值区间
	low, high := 0, len(numbers)
	i := 0
	for low < high {
		i++
		fmt.Println("寻找下标中...", i)
		if target == numbers[high-1] {
			fmt.Println("目标值下标是%v", high-1)
			break
		}

		if high-low-1 <= 0 {
			fmt.Println("找不到目标值")
			break
		}
		half := (high + low) / 2 //low本身为0 ，如果二分发生偏移，说明需要找到中间定位需要往上加
		if target < numbers[half] {
			high = half
		} else if target > numbers[half] {
			low = half
		} else {
			fmt.Printf("目标值 %v 的下标是 %v\n", target, half) //即中间值刚好为目标值
			break
		}

		// if i == 10 {
		// 	break
		// }
	}
}

/*
package main

import "fmt"

// 已知数组 A, B, 如果 A 中元素在 B 数组存在，打印出这个元素的下标，B 数组是不重复的.
// Input: [5, 3, 1, 5, 4] [5, 3]
// Output: [0, 1, 3]

func main() {
  a := []int{5, 3, 1, 5, 4}
  b := []int{5, 3}
  fmt.Println(a)
  fmt.Println(b)
}*/

func FindIndex() {
	a := []int{5, 3, 1, 5, 4}
	b := []int{5, 3}

	map1 := make(map[int]int)

	for index, value := range b {
		map1[value] = index
	}

	for in, second := range a {
		if _, ok := map1[second]; ok {
			fmt.Println(in)
		}
	}

	fmt.Println(a)
	fmt.Println(b)

}

/*以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

示例 1：
输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2：

输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。*/
func MergeArrary(array [][]int) {
	sort.Slice(array, func(i, j int) bool { return array[i][0] < array[j][0] })
	var out [][]int
	var index []int
	for e, arr := range array {
		exists := false
		for _, v2 := range index {
			if e == v2 {
				exists = true
			}
		}
		if exists {
			continue
		}
		for t, v := range array {
			if v[0] == arr[0] && v[1] == arr[1] {
				index = append(index, t)
				continue

			}
			//if v[0]+1 <= arr[len(arr)-1] && v[len(arr)-1] >= arr[0]+1 {
			fmt.Printf("the arr is%+v\n", arr)
			fmt.Printf("the v is%+v\n", v)
			//[[0,4],[3,5]]
			if arr[0] <= v[0] && arr[len(arr)-1] >= v[0] { //第一个数组第一位小于第二个数组第一位，第一个数组第二位小于第二个数组第二位，第一个数组第二位小于等于第二个数组第一位
				//if arr[0]+1 <= v[len(v)-1] && arr[len(arr)-1] >= v[0]+1 {
				exists = true
				//the data is [[1 3] [2 3] [1 6] [2 6] [8 10] [15 18]]
				if v[0] < arr[0] {
					if arr[len(arr)-1] <= v[len(v)-1] {
						out = append(out, []int{v[0], v[len(v)-1]})
					} else {
						out = append(out, []int{v[0], arr[len(arr)-1]})
					}
				} else {
					//out = append(out, []int{arr[0], v[len(v)-1]})
					if arr[len(arr)-1] <= v[len(v)-1] {
						out = append(out, []int{arr[0], v[len(v)-1]})
					} else {
						out = append(out, []int{arr[0], arr[len(arr)-1]})
					}
				}

				index = append(index, t)
			}
		}

		if !exists {
			out = append(out, arr)
		}

	}
	fmt.Printf("the data is %+v\n", out)
}
