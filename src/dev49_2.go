package main

import (
	"fmt"
	"time"
)

func goroutine(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(130 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go goroutine("world!")
	normal("hello")
}
