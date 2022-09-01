package main

import (
	"fmt"
	"sync"
)

func producer(ch chan int, i int) {
	ch <- i * 2
}

//closeで閉じられているからrange chが終わってDoneしている。
func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		fmt.Println("process", i*1000)
		wg.Done()
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	//Producer
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, i)
	}

	//Consumer
	go consumer(ch, &wg)
	wg.Wait()
	close(ch)

}

/*
wgについての説明をGoogleドキュメントに加えておく。
*/
