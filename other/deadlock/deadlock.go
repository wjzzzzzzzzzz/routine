package main

import (
	"sync"
	"errors"
	"log"
	"time"
)

func main() {
	ac := &account{
		id:    "111",
		money: 0,
	}
	bi := &account{
		id:    "222",
		money: 100000,
	}
	for i:=0;i<10000;i++ {
		go tranfer(ac,bi,50)
		log.Println(ac.money,bi.money)
	}
	time.Sleep(time.Second)
	log.Println(ac.money,bi.money)
}

type account struct {
	sync.Mutex
	id    string
	money int
}

func (this *account) deposit(add int) {
	this.Lock()
	defer this.Unlock()
	this.money += add
}
func (this *account) out(out int) error {
	this.Lock()
	defer this.Unlock()
	if this.money <= out {
		return errors.New("www")
	} else {
		this.money -= out
	}
	return nil
}
func tranfer(a,b *account, money int)(){
	a.deposit(money)
	b.out(money)

}
