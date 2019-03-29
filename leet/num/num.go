package num

func ReverseSlice(s []interface{}) {
	var res []interface{}
	for i := len(s) - 1; i >= 0; i-- {
		res = append(res, s[i])
	}
}
