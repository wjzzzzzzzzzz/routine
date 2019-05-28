package concurrent

import (
	"log"
	"sync"
	"time"
)

type Account struct {
	money int
	sync.RWMutex
}

func TransferA(acc *Account, money int, valueC chan int) {
	//acc.Lock()
	//defer  acc.Unlock()
	acc.money = acc.money + money
	valueC <- acc.money
}
func DepositA() {
	A := &Account{
		money: 0,
	}
	valueC := make(chan int)
	for i := 0; i < 1000; i++ {
		go TransferA(A, 100, valueC)
	}
	for {
		select {
		case value := <-valueC:
			log.Println(value, "value")
		case <-time.After(time.Second * 2):
			log.Println(A.money, "money")
		}
	}
}
