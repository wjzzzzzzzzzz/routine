package string

//给定两个字符串 s 和 t，判断它们是否是同构的。
//
//如果 s 中的字符可以被替换得到 t ，那么这两个字符串是同构的。
//
//所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。两个字符不能映射到同一个字符上，但字符可以映射自己本身。
func isIsomorphic(s string, t string) bool {
	ints := make(map[uint8]uint8)
	for i := 0; i < len(s); i++ {
		v, ok := ints[s[i]]
		if !ok {
			ints[s[i]] = s[i] - t[i]
			if i>=1&&t[i]==t[i-1]{
				return false
			}
		}else{
			if s[i]-t[i]!=v{
				return false
			}
		}
	}
	return true
}
