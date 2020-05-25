package main

import (
	"concurrency/error_handling/code"
	"fmt"
)

func main() {
	fmt.Println("bad")
	code.BadHandling()
	fmt.Println("good")
	code.GoodHandling()

}
