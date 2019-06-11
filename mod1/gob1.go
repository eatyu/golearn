package main

import "fmt"

func mm() func(int) int {
	sum := 0
	return func(i int) int {
		sum = i + sum
		return sum
	}
}

func main() {
	a := mm()
	//b := mm()
	for i := 0; i < 10; i++ {
		fmt.Println(a(i))
		//fmt.Println(b(i))
	}

}
