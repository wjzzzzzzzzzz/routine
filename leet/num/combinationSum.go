package num

import (
	"sort"
	"log"
)

//给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
//candidates 中的数字可以无限制重复被选取。
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var s []int
	sort.Ints(candidates)
	find(res, s, candidates, target, 0)
	return res

}
func find(listAll [][]int, tmp []int, candidates []int, target int, num int) {
	log.Println(candidates)
	if target == 0 {
		listAll = append(listAll, tmp)
		return
	}
	if target <candidates[0] {
		return
	}

	for i := num; i < len(candidates) &&candidates[i]<=target;i++ {
		log.Println(i,target)
		tmp=append(tmp ,candidates[i])
		log.Println(tmp)
		//find(listAll,tmp,candidates,target-candidates[i],i)

	}

}
