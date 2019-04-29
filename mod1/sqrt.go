package main

import "fmt"

func main() {
	fmt.Println(sqrt(4))
}

func sqrt(in float64) (out float64) {

	i := float64(1)
	for in != i*i {
		i -= (i*i - in) / (2 * i)
		fmt.Println(i)
	}
	return i
}
