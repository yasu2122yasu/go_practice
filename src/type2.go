package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Println("I don't know %T\n", v)
	}
}

func main() {
	do(10)
	do("Mike!")
	do(true)

	var i int = 10
	ii := float64(10)
	fmt.Println(i, ii)
}
