package other

import (
	"testing"
	"log"
)

func Test_map(t *testing.T) {
	ints := make(map[int]int)
	ints[1] = 1
	ints[2] = 2
	for key, value := range ints {
		log.Println(key,value)
	}
}
func Test_point(t *testing.T) {
	//指针为nil
	s:=append([]int{}, 1)
	log.Println(s)
}
