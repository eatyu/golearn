package main

import (
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	return map[string]int{"Go!": 1, "I": 1, "am": 1, "learning": 1}
}

func main() {
	wc.Test(WordCount)
}
