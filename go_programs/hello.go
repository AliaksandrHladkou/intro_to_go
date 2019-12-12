//Author: Aliaksandr Hladkou
//Date: 10/05/2019
//Go basics

package main

import (
	"fmt"
)

func main() {
	text := "Hello, world"
	num1, num2 := 5, 10
	fmt.Println(text)
	fmt.Println("The sum is: ", add(num1, num2))
}

func add(a, b int) int {
	return a + b
}