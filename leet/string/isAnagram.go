package string

//func isAnagram(s string, t string) bool {
//	if len(s) != len(t) {
//		return false
//	}
//	ss := strings.Split(s, "")
//	tt := strings.Split(t, "")
//	sort.Strings(ss)
//	sort.Strings(tt)
//	return strings.Join(ss, "") == strings.Join(tt, "")
//
//}
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	//一个map 去加减
	b := make([]int, 26)
	for _, r := range s {
		b[r-'a']++
	}
	for _, r := range t {
		b[r-'a']--
		if b[r-'a'] < 0 {
			return false
		}
	}
	for _, v := range b {
		if v != 0 {
			return false
		}
	}
	return true
}