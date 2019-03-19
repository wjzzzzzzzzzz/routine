package main

import (
	"context"
	"time"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx)
	time.Sleep(time.Second * 10)
	cancel()
	for {
		time.Sleep(time.Second * 1)
	}

}
func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println("退出", ctx.Err())
			return
		default:
			time.Sleep(time.Second * 2)
			log.Println("监控")

		}
	}

}
