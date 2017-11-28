package main

import (
	"log"
)

func main() {
	type m struct {
		k int
		e chan int
	}
	m1 := make(chan m)
	c := m{
		1,
		make(chan int),
	}

	go func() {
		select {
		case a := <-m1:
			log.Println("m1", a.k)
			a.e <- 100
		}
	}()

	m1 <- c
	log.Println("c", <-c.e)
    log.Print(123)
}
