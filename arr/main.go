package main

import (
	"log"
	"github.com/coocood/freecache"
)

func main() {
	c := freecache.NewCache(512 * 1024)
	c.Set([]byte("a"),[]byte("b"),1024)
	v,_:=c.Get([]byte("a"))
	log.Println(string(v))
}
