package await
//给定一个正整数，返回它在 Excel 表中相对应的列名称。
//
//例如，
//168
//1 -> A
//2 -> B
//3 -> C
//26 -> Z
//27 -> AA
//28 -> AB
import "log"

func convertToTitle(n int) string {
	var s string
	A:="A"
	for n >26 {
		h:=n/26
		n = n % 26
		log.Println(h,n)
		s = s + string("A"[0]+uint8(h)-2)
	}
	re:=string(A[0]+uint8(n)-1)

	s=s+re
	return s

}
