package test

import "fmt"

func DataStruct() {
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
