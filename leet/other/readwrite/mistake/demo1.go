package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	getABCD1(3)
	for {
		time.Sleep(time.Second)
	}
}

type S1 struct {
	sl     []int
	time   int
	lock   sync.Mutex
	ch     []chan int
	result chan int
}

func getABCD1(a int) {
	s := new(S1)
	s.result = make(chan int, a*4)
	down := make(chan bool)
	for i := 0; i < 4; i++ {
		s.ch = append(s.ch, make(chan int))
		go s.get(s.ch[i], down)
	}
	for i := 0; i < a; i++ {
		for j := 0; j < 4; j++ {
			s.ch[j] <- 'A' + j
			<-s.result

		}
	}
	log.Println(s.sl, len(s.sl))

}
func (s *S1) get(a chan int, down chan bool) {
	for {
		select {
		case v := <-a:
			s.sl = append(s.sl, v)
			s.result <- 1

		}
	}
}
