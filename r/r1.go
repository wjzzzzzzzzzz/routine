package r

import (
	"log"
	"time"
)

//使用多协程打印穿插打印数据多个元素

func main() {
	arra := []int64{1, 2, 3, 4}
	arrb := []string{"a", "b", "c", "d"}
	chanm := make(chan int)
	go func(v []string, chanm chan int) {
		for _, n := range arra {
			log.Println(n)
			<-chanm
			chanm <- 1
		}
	}(arrb, chanm)
	go func(v []int64, chanm chan int) {
		for _, n := range arrb {
			log.Println(n)
			chanm <- 1
			<-chanm
		}
	}(arra, chanm)
	for {
		time.Sleep(time.Second)
	}
}
