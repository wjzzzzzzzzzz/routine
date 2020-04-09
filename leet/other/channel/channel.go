package channel

import (
	"fmt"
	"time"
)

func main() {

	orders := make(chan int, 200)
	timeOut := make(chan int, 200)

	go accept(orders, timeOut)
	orders <- 1

	for {
		select {
		case x := <-timeOut:
			fmt.Println("订单", x, "过期")
		}
	}
}

//接受订单 id 返回过期id
func accept(orders chan int, timeout chan int) {
	for {
		select {
		case <-orders:
			map[int]int
		}
	}
}

type order struct {
	id    int
	score time.Time
}
