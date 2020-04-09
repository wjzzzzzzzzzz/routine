package other

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"testing"
)

// 格式 一，二,三，日 |13:00-15:00
func Test_map(t *testing.T) {
	b, e := checkShareTimes("一，二,三，日 |13:00-13:00")
	fmt.Println(b, e)
}

func deal(conn net.Conn) {
	var buf []byte
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(buf[:n])
	}

}
func checkShareTimes(sharetimes string) (bool, error) {
	ss := strings.Split(sharetimes, "|")
	if len(ss) < 2 {
		return false, errors.New("时间格式错误")
	}
	ss1 := strings.Split(ss[1], "-")
	if len(ss1) < 2 {
		return false, errors.New("时间格式错误")
	}
	start := strings.Split(ss1[0], ":")
	end := strings.Split(ss1[1], ":")
	if len(start) < 2 || len(end) < 2 {
		return false, errors.New("时间格式错误")
	}
	startH, _ := strconv.Atoi(start[0])
	startM, _ := strconv.Atoi(start[1])
	endH, _ := strconv.Atoi(end[0])
	endM, _ := strconv.Atoi(end[1])
	if endH > startH {
		return true, nil
	}
	if endH == startH {
		if endM > startM {
			return true, nil
		}
	}
	return false, nil
}

//func lengthOfLongestSubstring(s string) int {
//	split := strings.Split("s", "")
//	tmp := split[1:]
//	res:=1
//	for key, value := range split {
//		for k, v := range tmp {
//			if value == v{
//				tmp=append(tmp[:k],tmp[k+1:]...)
//			}
//			res+=1
//		}
//	}
//
//}
func replaceSpaces(S string, length int) string {
	strings.Split(S, "")
	S = S[:length]
	return strings.Replace(S, " ", "%20", -1)
}
func Test_point1(t *testing.T) {
	fmt.Println(replaceSpaces("                ", 5))
}
