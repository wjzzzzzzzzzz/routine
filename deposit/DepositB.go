package concurrent

import (
	"log"
	"time"
)

func TransferB(acc *Account, moneyC chan int) {
	for money := range moneyC {
		acc.money = acc.money + money
	}
}
func DepositB() {
	B := &Account{
		money: 0,
	}
	moneyC := make(chan int, 5)
	for i := 0; i < 10000; i++ {
		go func(moneyC chan<- int) {
			moneyC <- 100
		}(moneyC)
	}
	go TransferB(B, moneyC)
	for {
		select {
		case <-time.After(time.Second * 1):
			log.Println(B.money, "money")
		}
	}
}
