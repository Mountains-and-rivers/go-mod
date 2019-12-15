package main //必须

import (
	"mypkg/calc"
	"mypkg/jkl"
	"fmt"
)

func init() {
	fmt.Println("this is main init")
}

func main() {
	fmt.Println("----------------calc----------------")
	a := jkl.Add(10, 20)
	fmt.Println("a = ", a)

	fmt.Println("r = ", jkl.Minus(10, 5))
	fmt.Println("----------------jkl----------------")
	b := calc.Add(10, 20)
	fmt.Println("a = ", b)

	fmt.Println("r = ", calc.Minus(10, 5))
}
