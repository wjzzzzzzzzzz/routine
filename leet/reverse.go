package main

import (
	"log"
	"strconv"
)

//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转
func reverse(x int) int {
	if x > 0 {
		s := []byte(strconv.Itoa(x))
		log.Println(s)
		for i, v := range s {
			log.Println(i, v)
			s[len(s)-i-1] = v

		}
		i, _ := strconv.Atoi(string(s))

		return i
	}
	return 0

}
func main() {
	i := reverse(123)
	log.Println(i)
}
