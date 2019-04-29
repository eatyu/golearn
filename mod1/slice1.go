package main

import "fmt"

func main() {
	b := make([]int, 5, 10)
	c := b[4:5]

	printSlice("c", c)
	fmt.Println("%%%%\n\n ")

	language := []string{"java", "c", "c++", "c#", "go", "python"}

	for i, v := range language {
		language[i] = "golang"
		fmt.Println(i, v)
	}

	fmt.Println(language)

}

// range 返回的两个参数，一个是下标，一个是副本

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
