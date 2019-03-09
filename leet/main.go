package main

import (
	"log"
	"routine/leet/linklist"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

}
func main(){
	var s1 = []int{0, 0, 0, 0, 0, 0, 1, 2, 2, 2}
	var s2 = []int{0, 0, 0, 0, 0, 0, 1, 2, 2, 2}
	//log.Printf("%p %v",s1,s1)
	searchInsert1(s1,3)
	//log.Printf("%p %v",s1,s1)
	searchInsert2(s2,3)
	//log.Printf("%p %v",s2,s2)
}
func searchInsert1(nums []int, target int)  {
	log.Printf("%p %v",target,target)
	nums = append(nums, target)
	log.Printf("%p %v",nums,nums)
}
func searchInsert2(nums []int, target int)  {
	nums[0]=100
	log.Printf("%p %v",nums,nums)

}
func test1(nums []int, target int)  {
	ints := make([]*linklist.ListNode, 0)
	s:=&ints
	*s = append(*s, &linklist.ListNode{})
}
