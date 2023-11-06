package main

import (
	"fmt"
	"myproject/simplecalc"
)

const Hello = "Hello World"

// function name always start in capital EX: PrintLn, Add
// variable declaration  a :=-"string", a, b := 12, 3, var a int = 20

func main() {
	a, b := 12, 3

	//fmt.Println(simplecalc.Add(a, b))
	fmt.Println(simplecalc.Sub(a, b))
	fmt.Println(Hello)

	fmt.Println("Hello")
}
