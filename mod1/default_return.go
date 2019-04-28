package main

import "fmt"

func main() {
	a, b := fec(12, 34)
	fmt.Println(a, b)
}

func fec(i int, i2 int) (x, y int) {
	x = i + i2
	y = i - i2
	return
}

//	没有参数的 return 语句返回已命名的返回值。也就是 直接 返回。
