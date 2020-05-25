package main

import (
	"concurrency/pipeline/code"
	"fmt"
)

func main() {
	values := []int{1, 2, 3, 4, 5}
	multipliedValues := code.Multiply(values, 3)
	fmt.Println(multipliedValues)
}
