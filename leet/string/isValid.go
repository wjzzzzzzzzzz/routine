package string



//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//有效字符串需满足：
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//注意空字符串可被认为是有效字符串。
//ok

//用切片
//func isValid1(s string) bool {
//	stack := make([]string, len(s))
//	left := "{[("
//	right := "}])"
//	ms := strings.Split(s, "")
//	top := 0
//	if len(s)%2 != 0 {
//		return false
//	}
//	for _, v := range ms {
//		stack[top] = v
//		if top >= 1 {
//			func(a, b string) {
//				if strings.Index(left, a) != -1 && strings.Index(right, b) != -1 {
//					if a == "{" && b == "}" || a == "[" && b == "]" || a == "(" && b == ")" {
//						stack = append(stack[:top-1], stack[top+1:]...)
//						top = top - 2
//					}
//				}
//			}(stack[top-1], v)
//		}
//		top++
//	}
//	if top == 0 {
//		return true
//	} else {
//		return false
//	}
//}

//map
//func isValid2(s string) bool {
//	if len(s)%2 != 0 {
//		return false
//	}
//	stack := make([]string, len(s))
//	m := map[string]string{
//		"{": "}",
//		"[": "]",
//		"(": ")",
//	}
//	ms := strings.Split(s, "")
//	top := 0
//	for _, v := range ms {
//		stack[top] = v
//		if top >= 1 {
//			func(a, b string) {
//				if m[a] == b {
//					stack = append(stack[:top-1], stack[top+1:]...)
//					top = top - 2
//				}
//			}(stack[top-1], v)
//		}
//		top++
//	}
//	if top == 0 {
//		return true
//	} else {
//		return false
//	}
//}
type Statck struct {
	Link []string
}

func (s *Statck) Pop() string {
	if len(s.Link) == 0 {
		return ""
	}
	v := s.Link[len(s.Link)-1]
	s.Link = s.Link[:len(s.Link)-1]
	return v
}
//在左边就入栈，不在左边 就弹栈对比
func isValid(s string) bool {
	f := Statck{}
	hash_map := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	for _, b := range s {
		if _, ok := hash_map[string(b)]; !ok {
			f.Link = append(f.Link, string(b))
		} else {
			temp := f.Pop()
			if temp != hash_map[string(b)] {
				return false
			}
		}
	}
	if len(f.Link) == 0 {
		return true
	}
	return false
}