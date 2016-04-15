package main

import (
	"flag"
	"fmt"
	"log"
	"sync"
)

var (
	msgs = make(chan string, 5)
	N    int
)

func main() {
	flag.IntVar(&N, "n", 10, "total of produce num")
	flag.Parse()
	var wg sync.WaitGroup
	wg.Add(1)
	go produce(N, &wg)
	wg.Add(1)
	go consume(N, &wg)
	wg.Wait()
}

func consume(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		log.Printf("consume %s", <-msgs)
	}
}

func produce(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		msg := fmt.Sprint("msg-%n", i)
		msgs <- msg
		log.Printf("produce %s", msg)
	}
}
