package main

import "fmt"

func main() {
	fmt.Println("Start...")
	fmt.Println("calculate...")
	cal := Calc(23, 10)
	res := fmt.Sprintf("%d * %d = %d", 23, 10, cal)
	fmt.Println("done, result: ", res)
}

func Calc(a int, b int) int {
	return a * b
}
