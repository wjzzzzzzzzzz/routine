package main

import (
	"github.com/tuotoo/qrcode"
	"log"
	"os"
	"time"
)

var logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

func main() {
	startAt := time.Now()
	fi, err := os.Open("learnTool/qr/qrcode.jpg")
	if !check(err) {
		return
	}
	defer fi.Close()
	qrMatrix, err := qrcode.Decode(fi)
	if !check(err) {
		return
	}
	logger.Println(qrMatrix.Content)
	logger.Println(time.Since(startAt))
}

func check(err error) bool {
	if err != nil {
		logger.Println(err)
	}
	return err == nil
}
