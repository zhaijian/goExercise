package main

import (
	"log"
	"net"
	"sync"
	"time"
	"github.com/quexer/utee"
	//"fmt"
	_ "net/http/pprof"
	"net/http"
)

type tcpConn struct {
	sync.RWMutex
	m map[net.Conn]int64
}

func (c *tcpConn) add(conn net.Conn) {
	c.Lock()
	defer c.Unlock()
	if c.m == nil {
		c.m = map[net.Conn]int64{}
	}
	c.m[conn] = utee.TickSec()

}

func (c *tcpConn) remove(conn net.Conn) {
	c.Lock()
	defer c.Unlock()
	if c.m == nil {
		return
	}
	delete(c.m, conn)
}

func (c *tcpConn) len() int {
	c.RLock()
	defer c.RUnlock()
	return len(c.m)
}

var (
	tc = tcpConn{}

	syncMap = utee.SyncMap{}
)

func main() {
	l, err := net.Listen("tcp", "localhost:6666")
	utee.Chk(err)
	go sendLoop(l)
	go printTotal()
	log.Fatal(http.ListenAndServe("localhost:8888", nil))

}

func printTotal() {
	ticker := time.NewTicker(time.Second * 5)
	for _ = range ticker.C {
		log.Println("now online user is ", tc.len())
		//log.Println("now online user is ", syncMap.Len())
	}
}

func sendLoop(l net.Listener) {
	i := 0
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("accept: err", err)
			return
		}
		//syncMap.Put(conn,fmt.Sprintf("%d",utee.TickSec()))
		tc.add(conn)
		i++
		go read(conn)
		go write(conn)

	}
}

func read(conn net.Conn) {
	for {
		b := []byte{1}
		_, err := conn.Read(b)
		if err != nil {
			log.Println("read error ", err)
			//syncMap.Remove(conn)
			tc.remove(conn)
			conn.Close()
			return
		}
		//log.Println("read from client ", string(b))
	}

}

func write(conn net.Conn) {
	b := "hi,s"
	ticker := time.NewTicker(time.Second * 5)
	for _ = range ticker.C {
		_, err := conn.Write([]byte(b))
		if err != nil {
			//syncMap.Remove(conn)
			tc.remove(conn)
			log.Println("write error ", err, "i ")
			conn.Close()
			ticker.Stop()
			return
		}
	}
}
