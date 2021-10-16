package main

import "fmt"

func generate() func() int {
	var i = 0

	return func() int {
		i++
		return i
	}
}

func main() {
	keepStateGenerate := generate()
	for i := 0; i < 5; i++ {
		fmt.Println(keepStateGenerate())
	}
}
