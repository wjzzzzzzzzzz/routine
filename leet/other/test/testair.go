package main

import (
	"encoding/hex"
	"log"
	"net"
	"time"
)

func main() {
	var mm [][]byte
	m1, _ := hex.DecodeString("7E313030313636453630303030464439370D")
	m2, _ := hex.DecodeString("7E313030313636343230303030464441430D")
	m3, _ := hex.DecodeString("7E313030313636343330303030464441420D")
	m4, _ := hex.DecodeString("7E313030313636343430303030464441410D")
	m5, _ := hex.DecodeString("7e313030313636343730303030464441370d")
	mm = append(mm, m1, m2, m3, m4, m5)
	lns := []int{34, 50, 24, 32, 46}
	conn, err := net.Dial("tcp", "127.0.0.1:30003")
	if err != nil {
		log.Panic(err)
	}
	for {
		for i, m := range mm {
			n, err := conn.Write(m)
			if err != nil {
				log.Panic(err)
			}
			log.Println("发送", hex.EncodeToString(m), n)
			time.Sleep(time.Millisecond * 200)
			conn.SetReadDeadline(time.Now().Add(time.Second * 2))
			ln := 0
			var h = make([]byte, 500)
			for {
				n, err := conn.Read(h)
				if err != nil {
					log.Panic(err)
				}
				ln += n
				if ln >= lns[i] {
					break
				}
			}
		}
	}
}
