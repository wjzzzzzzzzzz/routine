package main

import (
	"log"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

type S2 struct {
	sl     []int
	ch     []chan int
	result chan int
}
func main() {
	getABCD2(4)
	for {
		time.Sleep(time.Second)
	}
}

func getABCD2(a int) {
	s := new(S2)
	s.result = make(chan int)
	for i := 0; i < 4; i++ {
		s.ch = append(s.ch, make(chan int))
		go s.run(i)
	}
	for i := 0; i < a; i++ {
		s.ch[0] <- 'A'
		<-s.result
	}
	log.Println(s.sl, len(s.sl))
}
func (s *S2) run(i int) {
	for {
		select {
		case a := <-s.ch[i]:
			s.sl = append(s.sl, a)
			if i < 3 {
				s.ch[i+1] <- i + 1 + 'A'
			} else {
				s.result <- 1
			}
		}
	}
}

//func (c *Connect) Getconns(port string) chan net.Conn {
//	c.conns = make(chan *net.Conn, 10)
//	ln, err := net.Listen("tcp", ":"+port)
//	if err != nil {
//		log.Println(err)
//	}
//	for {
//		conn, err := ln.Accept()
//		log.Println("重连")
//		if err != nil {
//			log.Println(err)
//			return nil
//		}
//		c.conns <- &conn
//	}
//}

//func (c *Connect) Dealconn() {
//	for conn := range c.conns {
//		go read(conn)
//	}
//}
//func read(conn net.Conn) {
//	var buf [20]byte
//	for {
//		n, err := conn.Read(buf[:])
//		//bytes, err := ioutil.ReadAll(conn)
//		if err != nil {
//			log.Println(err)
//			return
//		}
//		log.Println("上报<-------", buf[:n])
//	}
//}
//
//func work(conn net.Conn) {
//
//	go write(conn)
//
//}
//
//func write(conn net.Conn) {
//	a := numconv.HexToBytes("1007FBFB0709FF")
//	for {
//		_, err1 := conn.Write(a)
//
//		log.Println("下发------->", a)
//		time.Sleep(time.Second * 600)
//		if err1 != nil {
//			log.Println(err1)
//			return
//		}
//	}
//
//}
//
////func getABCD(a int) {
////	var  m []int
////	s := &S{
////		sl: m  ,
////		time: a,
////	}
////	for i := 0; i < 4; i++ {
////		ch := make(chan int)
////		s.ch = append(s.ch, ch)
////		go get(i, s)
////	}
////	for i := 0; i < a; i++ {
////		for j := 0; j < 4; j++ {
////			s.ch[j] <- 1
////		}
////	}
////	log.Println(s.sl,len(s.sl))
////}
////func get(i int, s *S) {
////	for {
////		<-s.ch[i]
////		s.sl = append(s.sl, i+'A')
////	}
////
////}
