package main

import (
	"log"
	"math"
	"strconv"
)

//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转
func reverse(x int) int {
	var i int
	s := []byte(strconv.Itoa(x))
	if x >= 0 {
		l := len(s) - 1
		r := make([]byte, len(s))
		for i, v := range s {
			r[l-i] = v
		}
		i, _ = strconv.Atoi(string(r))
	} else {
		s = s[1:]
		l := len(s) - 1
		r := make([]byte, len(s))
		for i, v := range s {
			r[l-i] = v
		}
		i, _ = strconv.Atoi("-" + string(r))

	}
	pow := int(math.Pow(float64(2), float64(31)))
	if i < pow && (i+pow) >= 0 {
		return i
	}
	return 0
}
func main() {
	i := reverse(1534236469)
	log.Println(i)
}
