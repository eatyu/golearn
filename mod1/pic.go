package main

import (
	"golang.org/x/tour/pic"
)

func main() {
	pic.Show(Pic)

	//data := Pic(256,256)
	//fmt.Println(data)
}

func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy)
	for j := 0; j < dy; j++ {
		s[j] = make([]uint8, dx)
		for i := 0; i < dx; i++ {
			s[j][i] = uint8((i + j) / 2)
		}
	}
	return s
}
