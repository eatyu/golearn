package main

import "fmt"

var c, java, python bool

func main() {
	var i = 10
	fmt.Println(i, c, java, python)
	kk()
}

func kk() {
	var b, v int = 1, 2
	i, j, k := "sda", "da", false
	fmt.Println(b, v)
	fmt.Println(i, j, k)

}

//	变量声明可以包含初始值，每个变量对应一个。
//
//	如果初始化值已存在，则可以省略类型；变量会从初始值中获得类型。

//	在函数中，简洁赋值语句 := 可在类型明确的地方代替 var 声明。
//
//	函数外的每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用。
