package main
import (
	"log"
	"time"
)

func main() {
	type m struct {
		k int
		e chan int
	}
	c := make(chan m)
	go func() {
		select {
		case a := <-c.k:
			log.Println("a",a)
			m.e<-1
		}
	}()


	gg := m{}
	gg.k=2
	c<-gg
	e:= <-gg.e
	log.Println("e",e)
	time.Sleep()
}
