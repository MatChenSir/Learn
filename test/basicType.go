package test

import (
	"fmt"
	"strconv"
	"unsafe"
)

func BasicType() {
	//int
	var int1 int
	int1 = 9223372036854775807
	fmt.Printf("int1的值为:%d,类型为:%v\n", int1, getType(int1))
	/*int8:  有符号 8 位整型 (-128 到 127)
	int16:  有符号 16 位整型 (-32768 到 32767)
	int32:  有符号 32 位整型 (-2147483648 到 2147483647)
	int64:  有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)

	uint8:  无符号 8 位整型 (0 到 255)
	uint16:  无符号 16 位整型 (0 到 65535)
	uint32:  无符号 32 位整型 (0 到 4294967295)
	uint64:  无符号 64 位整型 (0 到 18446744073709551615)
	*/

	//string
	var string1 string
	string1 = strconv.FormatInt(int64(int1), 10) //  转换为十进制字符串
	fmt.Printf("string1的值为:%s,类型为:%v\n", string1, getType(string1))

	//浮点型
	var float1 float64
	float1, _ = strconv.ParseFloat(string1, 64) //字符串转float
	float1 = float64(int1)                      //int转float 可直接转
	fmt.Printf("float1的值为:%f,类型为:%v\n", float1, getType(float1))

	string2 := strconv.FormatFloat(float1, 'f', 2, 64) // // 将浮点数转换为字符串，保留两位小数
	fmt.Printf("string2的值为:%s,类型为:%v\n", string2, getType(string2))
	// float32:  IEEE-754 32位浮点型数
	// float64:  IEEE-754 64位浮点型数
	// complex64:  32 位实数和虚数
	// complex128:  64 位实数和虚数
	complex()

	//其他数字类型
	// byte:  uint8的类型别名
	str := "hello"
	byteSlice := []byte(str)

	// 修改字节切片
	byteSlice[0] = 'H'

	// 将修改后的字节切片转换回字符串
	modifiedStr := string(byteSlice)
	fmt.Printf("byte的值为:%v,string为:%v\n", byteSlice, modifiedStr)
	//结果byte的值为:[72 101 108 108 111],string为:Hello

	// rune:  int32的类型别名   多用于用于表示 Unicode 码点，可见方法
	GgetRunne()
	// uint:  32 或 64 位
	// int:  32 或 64 位
	// uintptr:  无符号整型，用于存放一个指针
	uintptr1()

}

func uintptr1() {
	// 通过指针获取变量地址
	var num int = 42
	ptr := &num

	// 将指针转换为 uintptr 类型
	ptrInt := uintptr(unsafe.Pointer(ptr))

	// 输出 uintptr 类型的整数值
	fmt.Printf("指针的整数值为: %d\n", ptrInt)

	// 将 uintptr 转换回指针
	newPtr := (*int)(unsafe.Pointer(ptrInt))

	// 输出新的指针指向的值
	fmt.Printf("通过 uintptr 转换回指针，值为: %d\n", *newPtr)
}

//复数 用于信号处理，控制系统，数学计算，图像处理
func complex() {
	// 定义一个 complex64 类型的复数变量
	var z1 complex64 = 3 + 4i
	fmt.Println(z1) // 输出: (3+4i)

	// 复数运算示例
	var z2 complex64 = 2 + 3i
	sum := z1 + z2
	fmt.Println(sum) // 输出: (5+7i)

	// 通过内置的 real() 和 imag() 函数获取复数的实部和虚部
	fmt.Println(real(sum)) // 输出: 5
	fmt.Println(imag(sum)) // 输出: 7
}
