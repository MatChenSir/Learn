package test

import (
	"fmt"
	"strconv"
)

func FunctionAsParam() {
	Test2(c)
}

func c(num int) {
	num = 4 + num
	fmt.Println("num:" + strconv.Itoa(num))
}

func Test2(test func(in int)) {
	fmt.Println("test")
	test(3)
}
