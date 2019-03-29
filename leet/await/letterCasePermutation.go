package await

//给定一个字符串S，通过将字符串S中的每个字母转变大小写，我们可以获得一个新的字符串。返回所有可能得到的字符串集合。
//返回所有的字串[]string
//返回字串的大小写
func letterCasePermutation(S string) []string {
	l := 1 //窗口大小
	//if len(S) == 1 {
	//
	//	return
	//}
	var res []string
	for l = 1; l <= len(S); l++ {
		for i := 0; i <=len(S)-l; i++ {
			res = append(res, S[i:i+l])
		}
	}
	return res
}
