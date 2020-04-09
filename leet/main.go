package main

import (
	"log"
	"routine/leet/linklist"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

}
func main() {
	for i := 0; i < 3; i++ {
		go func(i int) {
			log.Println(i)
		}(i)
	}
	for {
		time.Sleep(time.Second)
	}
}
func searchInsert1(nums []int, target int) {
	log.Printf("%p %v", target, target)
	nums = append(nums, target)
	log.Printf("%p %v", nums, nums)
}
func searchInsert2(nums []int, target int) {
	nums[0] = 100
	log.Printf("%p %v", nums, nums)

}
func test1(nums []int, target int) {
	ints := make([]*linklist.ListNode, 0)
	s := &ints
	*s = append(*s, &linklist.ListNode{})
}

func Teststring(nums []int, target int) {
	ints := make([]*linklist.ListNode, 0)
	s := &ints
	*s = append(*s, &linklist.ListNode{})
}
