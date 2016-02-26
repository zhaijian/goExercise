package main

import "fmt"

type item struct {
	key   interface{}
	value interface{}
}

type q []*item

type i interface {
	show()
}

func (q1 *q) get() {
	fmt.Println("..............get..........",q1)
}

func (q1 q) delete() {
	fmt.Println("...........delete..........",q1)
}

func (q1 q) show() {
	fmt.Println("...........show..........")
}

type S struct{ i int }

func (p *S) Get() int  { return p.i }
func (p *S) Put(v int) { p.i = v }

type I interface {
	Get() int
	Put(int)
}

func f(p I) {
	fmt.Println(p.Get())
	p.Put(1)
}

func do(item i) {
	item.show()
}

func main() {
	s := S{}
	f(&s)

	a := "abc"
	b := a
	c := &a
	fmt.Println("a", a, "b", b, "c", *c)
	a = "efg"
	fmt.Println("a", a, "b", b, "c", *c)

	q1 := q{}
	item := &item{}
	item.key = "a"
	item.value = "b"
	q1 = append(q1, item)

	fmt.Println("q key", q1[0].key.(string), "q value", q1[0].value.(string))

	q2 := &q{}
	q1.get()
	q2.get()
	q2.delete()
	do(q1)
}
