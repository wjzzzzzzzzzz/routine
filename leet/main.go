package main

import (
	"log"
	"routine/leet/linklist"
	"reflect"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

}
func main(){
	a:="sss1"
	u := a[0]
	log.Println(reflect.TypeOf(u))

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

func Teststring(nums []int, target int)  {
	ints := make([]*linklist.ListNode, 0)
	s:=&ints
	*s = append(*s, &linklist.ListNode{})
}

