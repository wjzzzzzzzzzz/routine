package num

import "log"
//给定两个数组，编写一个函数来计算它们的交集。
func intersection1(nums1 []int, nums2 []int) []int {
	map1 := make(map[int]int)
	map2 := make(map[int]int)
	for _, v := range nums1 {
		if _, ok := map1[v]; !ok {
			map1[v] = v
		} else {
			continue
		}
	}
	for _, v := range nums2 {
		if _, ok := map2[v]; !ok {
			map2[v] = v
		} else {
			continue
		}
	}
	return getCommKey(map1, map2)
}
func getCommKey(map1 map[int]int, map2 map[int]int) []int {
	if len(map1) > len(map2) {
		return getCommKey(map2, map1)
	}
	var a []int
	for key, _ := range map1 {
		if _, ok := map2[key]; ok {
			a = append(a, key)
		}
	}
	return a
}
func intersection(nums1 []int, nums2 []int) []int {
	if nums1 == nil || len(nums1) == 0{
		return nil
	}else if nums2 == nil || len(nums2) == 0{
		return nil
	}
	dict := make(map[int]bool)
	for _,v := range nums1{
		log.Println(dict[v])
		if dict[v] == false{
			dict[v] = true
		}
	}

	var res []int
	for _,v := range nums2{
		if dict[v]{
			res = append(res,v)
			dict[v] = false
		}
	}
	return res
}