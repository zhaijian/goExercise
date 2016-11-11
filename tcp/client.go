package main

import (
	"log"
	"net"
	"strconv"
	"time"
)

type tcpClient struct {
	conn net.Conn
	name string
}

func (c tcpClient) read() {
	for {
		fixedBuf := make([]byte, 4)
		_, err := c.conn.Read(fixedBuf)
		if err != nil {
			log.Println("read error ", err)
			c.conn.Close()
			return
		}
		log.Println(c.name, "read from server:", string(fixedBuf))
	}
}

func (c tcpClient) write() {
	ticker := time.NewTicker(time.Second * 1)
	for _ = range ticker.C {
		_, err := c.conn.Write([]byte(c.name))
		if err != nil {
			log.Println("write err ", err)
			c.conn.Close()
			ticker.Stop()
			return
		}
	}
}

func main() {
	mockTcpClient("localhost:6666")
	select {}
}

func mockTcpClient(addr string) {
	i := 0
	for {
		c, err := net.Dial("tcp", addr)
		if err != nil {
			log.Println("dial tcp err ", err, i)
			continue
		}
		client := tcpClient{
			conn: c,
			name: strconv.Itoa(i),
		}
		go client.read()
		go client.write()
		i++
		if i%100 == 0 {
			log.Printf("mock tcp client %d total", i)
			time.Sleep(time.Second * 1)
		}
	}
}
