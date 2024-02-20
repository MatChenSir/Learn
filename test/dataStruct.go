package test

import (
	"fmt"
	"reflect"
)

func DataStruct() {
	defer RecoverPainc() //一定得放第一行，否则任意一行panic就不会执行deffer了
	// SiglePointer()
	// ArraryAndSlices()
	// GetRunne()
	GetMap()
	fmt.Printf("哈哈哈哈,我来了")
}

func SiglePointer() {
	fmt.Println("开始学习")
	// %p 获取地址，%v获取value 值，
	// %d: 十进制整数
	// %f: 浮点数
	// %s: 字符串
	// %t: 布尔值
	// %T: 值类型
	// %v: 通用格式 所有占位

	//指针 int
	fmt.Println("<-----------------------------指针 int start-------------------------------->")
	var a *int
	b := 10
	a = &b //该部分操作只是获取对应的地址，拿到这个指针对象，而非直接获取值
	fmt.Printf("指针a的地址为 %p\n", a)
	fmt.Printf("指针a的类型为 %T\n", a)
	//如果直接打印a，将会得到一个地址，而b则是值----> a为： 0xc0000160b0 b 为： 10
	fmt.Printf("指针a的值为 %d\n", a) //打印地址
	fmt.Println("a为：", a, "b 为：", b)

	//该处同样会拿到地址，因为即使你没有显式声明 c 是一个指针，&c 表达式创建了一个指向 c 的指针
	c := 10
	d := &c
	fmt.Printf("%+v\n", d)
	fmt.Println("d的值为", d)
	fmt.Println("<-----------------------------指针 int end-------------------------------->")

	//指针类型
	fmt.Println("<-----------------------------指针 type start-------------------------------->")
	type Student struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		Sex  string `json:"sex"`
	}
	fmt.Printf("%p\n", &Student{Name: "", Age: 2, Sex: ""})
	user := Student{Name: "", Age: 2, Sex: ""}
	user2 := &user
	user2.Age = 4
	fmt.Printf("%v\n", user.Age)
	fmt.Printf("user 指向的地址为%p,user2指向的地址为%p,俩个地址是一样的\n", &user, user2)

	//var user3 *Student
	user3 := new(Student)
	user3.Age = 5
	user3.Sex = "male"
	user3.Name = "tt"
	fmt.Printf("当前user3指向的地址为: %p, 当是如果取他的指针对象，其地址为: %p\n", user3, &user3)
	//user := &user{}

	fmt.Println("<-----------------------------指针 type end-------------------------------->")
}

//数组和切片  最大区别为 数组[5]int是有长度，切片是[]int动态长度！！！！！
func ArraryAndSlices() {
	//切片是数组的一个引用，因此切片是引用类型。但自身是结构体，值拷贝传递。
	//切片比数组更灵活，可以方便地进行动态调整、追加和截取等操作

	//定义一个数组
	fmt.Println("<-----------------------------定义一个数组 start-------------------------------->")

	var array [5]int //有默认值，且为0 长度为5
	fmt.Printf("第一个数组为%v,类型为:%T\n", array, array)

	array2 := [...]int{0, 1, 2, 3, 4, 5} //...表示自己计算长度不定义他,但仅对于后面数组，无法array2[6] =3这样
	fmt.Printf("第二个数组为%v,类型为:%T\n", array2, array2)

	//array3 := [4]int{0: 1, 1: 2}
	array3 := [4]string{"test", "test", "test"}
	if array3[3] == "" {
		fmt.Printf("arrar3[3]是空字符串")
	}
	fmt.Printf("第三个数组为%v,类型为:%T\n", array3, array3)

	//多维数组
	array4 := [4][2]string{{"test", "test"}, {"test1", "test2"}}
	fmt.Printf("第四个数组为%v,类型为:%T\n", array4, array4)

	fmt.Println("<-----------------------------定义一个数组 end-------------------------------->")

	//定义一个切片
	fmt.Println("<-----------------------------定义一个切片 start-------------------------------->")
	slice := make([]int, 5)
	//这里的输出[]int就是表明slice的类型是切片（slice）
	fmt.Printf("第一个切片为类型为:%T\n", slice)
	fmt.Printf("第一个切片为%v,类型为:%v\n", slice, getType(slice))

	slice2 := []int{2} //记住！！！！！！不指定长度就是切片
	fmt.Printf("第二个切片为%v,类型为:%v\n", slice2, getType(slice2))

	slice3 := array2[2:3] //输出2 包前不包后，可通过切数组或切片来定义新切片
	fmt.Printf("第三个切片为%v,类型为:%v\n", slice3, getType(slice3))

	slice4 := append([]int{}, slice3...) //通过定义新空或现有追加来定义新切片
	fmt.Printf("第四个切片为%v,类型为:%v\n", slice4, getType(slice4))

	fmt.Println("<-----------------------------定义一个切片 end-------------------------------->")

	//两种切片常见问题
	var slice5 []int
	fmt.Printf("第五个切片为%v,类型为:%v\n", slice5, getType(slice5))

	arr := [1026]int{}
	slice6 := make([]int, 100) //创建长度为0 ，容量为3
	slice6 = append(slice6, arr[1:1026]...)
	fmt.Printf("切片的容量为：%d\n", cap(slice6)) //扩容 当容量小于1024，则俩倍扩容，否则1.25倍扩容
	fmt.Printf("第五个切片为%v,类型为:%v\n", slice6, getType(slice6))

	// 局部：
	// arr2 := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	// slice5 := arr[start:end]
	// slice6 := arr[:end]
	// slice7 := arr[start:]
	// slice8 := arr[:]
	// slice9 := arr[:len(arr)-1] //去掉切片的最后一个元素

}

func getType(data interface{}) string {
	value := reflect.ValueOf(data)
	//这里的输出[]int就是表明slice的类型是切片（slice）
	return value.Kind().String()
}

//rune 实际上是 int32 的别名，用于表示 Unicode 码点
//Unicode即非字母外的符号，如中文  Unicode 字符可以用多种编码方式来表示，其中最常见的是 UTF-8、UTF-16 和 UTF-32。
func GgetRunne() {
	//简而言之， string底层三byte(ASCII 字符或一个非 ASCII 字符的字节), Unicode表示更占位置，用rune更好

	str := "Hello, 世界"
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
	fmt.Printf("(字符串长度为 %d  个 byte 字符,", len(str))
	fmt.Printf("字符串长度为 %d 个 Unicode 字符)\n", len([]rune(str)))
}

func GetMap() {

	//make 一个map
	map1 := make(map[string]int)
	map1["number1"] = 1
	map1["number2"] = 2
	fmt.Printf("map1的值为:%v, type为:%v\n", map1, getType(map1))

	//直接赋值
	map2 := map[string]string{"test1": "2", "test2": "2"}
	fmt.Printf("map1的值为:%v, type为:%v\n", map2, getType(map2))

	//var声明
	var map3 map[string]int //空map

	if map3 == nil {
		fmt.Println("map3是nil")
	}
	map4 := map3
	panic("报错了map3是nil")
	for k, v := range map4 {
		fmt.Println("key:", k, "value:", v)
	}
	fmt.Printf("map1的值为:%v, type为:%v\n", map3, getType(map3))

}

//报错捕获 保护性机制，防止程序崩溃
func RecoverPainc() {
	if r := recover(); r != nil {
		fmt.Println("捕获报错信息", r)
	}
}
