package await

//func lengthOfLongestSubstring(s string) int {
//	ints := make(map[string]int)
//	for i := 0; i < len(ints); i++ {
//		ints[:]
//	}
//}


//获得
//两个以上所有的子串
//func getSub(s string) []string {
//	var res []string
//	for i := 2; i < len(s); i++ {
//		for j := 0; j <=len(s)-i; j++ {
//			res = append(res, s[j:j+i])
//		}
//	}
//	log.Println(res)
//	return res
//}

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	isall := make(map[string]bool)
	res := 1
	for i := 2; i <= len(s); i++ {
		for j := 0; j <= len(s)-i; j++ {
			temp := s[j : j+i]
			if _, ok := isall[temp]; !ok {
				isall[temp] = all(temp,isall)
			}
			if isall[temp] {
				if res < len(temp) {
					res = len(temp)
				}
			}
		}
	}
	return res
}
func all(s string,m map[string]bool ) bool {





	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}


//判断是否所有元素均唯一
//func all(s string) bool {
//	for i := 0; i < len(s); i++ {
//		for j := i + 1; j < len(s); j++ {
//			if s[i] == s[j] {
//				return false
//			}
//		}
//	}
//	return true
//}

func getSub(s string) []string {
	var res []string
	for i := 2; i <= len(s); i++ {
		for j := 0; j <= len(s)-i; j++ {
			res = append(res, s[j:j+i])
		}
	}
	return res
}
