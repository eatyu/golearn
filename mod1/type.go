package main

import "fmt"

func main() {
	v := 0.867 + 0.5i // 修改这里！
	fmt.Printf("v is of type : %T\n", v)
}

//	类型推导
//	在声明一个变量而不指定其类型时（即使用不带类型的 := 语法或 var = 表达式语法），变量的类型由右值推导得出。
