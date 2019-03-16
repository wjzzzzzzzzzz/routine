package await

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
